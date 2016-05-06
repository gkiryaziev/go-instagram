package instagram

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Instagram struct {
	userName   string
	password   string
	token      string
	isLoggedIn bool
	uuid       string
	deviceId   string
	userNameId int64
	rankToken  string
	cookies    []*http.Cookie
}

// NewInstagram api
func NewInstagram(userName, password string) (*Instagram, error) {
	i := &Instagram{
		userName:   userName,
		password:   password,
		token:      "",
		isLoggedIn: false,
		uuid:       generateUUID(true),
		deviceId:   generateDeviceId(),
		userNameId: 0,
		rankToken:  "",
		cookies:    nil,
	}

	err := i.Login()
	if err != nil {
		return nil, err
	}

	return i, nil
}

// Login to instagram.
func (i *Instagram) Login() error {

	fetch := API_URL + "/si/fetch_headers/?challenge_type=signup&guid=" + generateUUID(false)

	resp, err := i.requestLogin("GET", fetch, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// get csrftoken
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "csrftoken" {
			i.token = cookie.Value
		}
	}

	// login
	login := &Login{
		DeviceId:          i.deviceId,
		Guid:              i.uuid,
		UserName:          i.userName,
		Password:          i.password,
		Csrftoken:         i.token,
		LoginAttemptCount: "0",
	}

	jsonData, err := json.Marshal(login)
	if err != nil {
		return err
	}

	signature := generateSignature(jsonData)

	resp, err = i.requestLogin("POST", API_URL+"/accounts/login/?", bytes.NewReader([]byte(signature)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// get new csrftoken
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "csrftoken" {
			i.token = cookie.Value
		}
	}
	i.cookies = resp.Cookies()

	var object *LoginResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&object)
	if err != nil {
		return err
	}

	if object.Status == "fail" {
		return errors.New(object.Message)
	}

	i.userNameId = object.LoggedInUser.Pk
	i.rankToken = strconv.FormatInt(i.userNameId, 10) + "_" + i.uuid

	return nil
}

// GetMediaLikers return media likers.
func (i *Instagram) GetMediaLikers(mediaId string) (*MediaLikers, error) {

	endpoint := API_URL + "/media/" + mediaId + "/likers/?"

	resp, err := i.request("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var object *MediaLikers
	err = json.Unmarshal(resp, &object)
	if err != nil {
		return nil, err
	}

	return object, nil
}

// GetMedia return media comments.
func (i *Instagram) GetMedia(mediaId string) (*Media, error) {

	endpoint := API_URL + "/media/" + mediaId + "/comments/?"

	resp, err := i.request("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var object *Media
	err = json.Unmarshal(resp, &object)
	if err != nil {
		return nil, err
	}

	return object, nil
}

// GetRecentActivity return recent activity.
func (i *Instagram) GetRecentActivity() (*RecentActivity, error) {

	endpoint := API_URL + "/news/inbox/?"

	resp, err := i.request("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var object *RecentActivity
	err = json.Unmarshal(resp, &object)
	if err != nil {
		return nil, err
	}

	return object, nil
}

// SearchUsers return users.
func (i *Instagram) SearchUsers(query string) (*SearchUsers, error) {

	endpoint := API_URL + "/users/search/?ig_sig_key_version=" + SIG_KEY_VERSION +
		"&is_typeahead=true&query=" + query + "&rank_token=" + i.rankToken

	resp, err := i.request("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var object *SearchUsers
	err = json.Unmarshal(resp, &object)
	if err != nil {
		return nil, err
	}

	return object, nil
}

// GetUserNameInfo return username info.
func (i *Instagram) GetUserNameInfo(userNameId int64) (*UserNameInfo, error) {

	endpoint := API_URL + "/users/" + strconv.FormatInt(userNameId, 10) + "/info/?"

	resp, err := i.request("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var object *UserNameInfo
	err = json.Unmarshal(resp, &object)
	if err != nil {
		return nil, err
	}

	return object, nil
}

// GetUserTags return user tags.
func (i *Instagram) GetUserTags(userNameId int64) (*UserTags, error) {

	endpoint := API_URL + "/usertags/" + strconv.FormatInt(userNameId, 10) + "/feed/?rank_token=" +
		i.rankToken + "&ranked_content=false"

	resp, err := i.request("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var object *UserTags
	err = json.Unmarshal(resp, &object)
	if err != nil {
		return nil, err
	}

	return object, nil
}

// SearchTags return tags.
func (i *Instagram) SearchTags(query string) (*SearchTags, error) {

	endpoint := API_URL + "/tags/search/?is_typeahead=true&q=" + query + "&rank_token=" + i.rankToken

	resp, err := i.request("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var object *SearchTags
	err = json.Unmarshal(resp, &object)
	if err != nil {
		return nil, err
	}

	return object, nil
}

// TagFeed return tagged media.
func (i *Instagram) TagFeed(tag, maxId string) (*TagFeed, error) {

	endpoint := API_URL + "/feed/tag/" + tag + "/?rank_token=" + i.rankToken + "&ranked_content=false&max_id=" + maxId

	resp, err := i.request("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var object *TagFeed
	err = json.Unmarshal(resp, &object)
	if err != nil {
		return nil, err
	}

	return object, nil
}

// requestLogin return http.Response. Needs to get the authorization cookies.
func (i *Instagram) requestLogin(method, endpoint string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", USER_AGENT)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// requestMain is main request for all other methods. Reading the authorization cookies.
func (i *Instagram) requestMain(method, endpoint string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", USER_AGENT)
	for _, cookie := range i.cookies {
		req.AddCookie(cookie)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// request with five attempts re-login. Re-login if getting error 'login_required'.
func (i *Instagram) request(method, endpoint string, body io.Reader) ([]byte, error) {

	for attempt := 0; attempt < 5; attempt++ {

		resp, err := i.requestMain(method, endpoint, body)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		jsonBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var message *Message
		err = json.Unmarshal(jsonBody, &message)
		if err != nil {
			return nil, err
		}

		if message.Status == "fail" {
			if message.Message != "login_required" {
				return nil, errors.New(message.Message)
			}
			// relogin
			err = i.Login()
			if err != nil {
				return nil, err
			}
			time.Sleep(time.Millisecond * 500)
		} else {
			return jsonBody, nil
		}
	}

	return nil, errors.New("max_attempts")
}
