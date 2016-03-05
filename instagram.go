package instagram_api

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

// Login to Instagram.
func (this *Instagram) Login() error {

	fetch := API_URL + "/si/fetch_headers/?challenge_type=signup&guid=" + generateUUID(false)

	resp, err := this.requestLogin("GET", fetch, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// get csrftoken
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "csrftoken" {
			this.token = cookie.Value
		}
	}

	// login
	login := &Login{
		DeviceId:          this.deviceId,
		Guid:              this.uuid,
		UserName:          this.userName,
		Password:          this.password,
		Csrftoken:         this.token,
		LoginAttemptCount: "0",
	}

	jsonData, err := json.Marshal(login)
	if err != nil {
		return err
	}

	signature := generateSignature(jsonData)

	resp, err = this.requestLogin("POST", API_URL+"/accounts/login/?", bytes.NewReader([]byte(signature)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// get new csrftoken
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "csrftoken" {
			this.token = cookie.Value
		}
	}
	this.cookies = resp.Cookies()

	var object *LoginResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&object)
	if err != nil {
		return err
	}

	if object.Status == "fail" {
		return errors.New(object.Message)
	}

	this.userNameId = object.LoggedInUser.Pk
	this.rankToken = strconv.FormatInt(this.userNameId, 10) + "_" + this.uuid

	return nil
}

// Get media likers.
func (this *Instagram) GetMediaLikers(mediaId string) (*MediaLikers, error) {

	endpoint := API_URL + "/media/" + mediaId + "/likers/?"

	resp, err := this.request("GET", endpoint, nil)
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

// Get media comments.
func (this *Instagram) GetMedia(mediaId string) (*Media, error) {

	endpoint := API_URL + "/media/" + mediaId + "/comments/?"

	resp, err := this.request("GET", endpoint, nil)
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

// Get recent activity.
func (this *Instagram) GetRecentActivity() (*RecentActivity, error) {

	endpoint := API_URL + "/news/inbox/?"

	resp, err := this.request("GET", endpoint, nil)
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

// Search users.
func (this *Instagram) SearchUsers(query string) (*SearchUsers, error) {

	endpoint := API_URL + "/users/search/?ig_sig_key_version=" + SIG_KEY_VERSION +
		"&is_typeahead=true&query=" + query + "&rank_token=" + this.rankToken

	resp, err := this.request("GET", endpoint, nil)
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

// Get username info.
func (this *Instagram) GetUserNameInfo(userNameId int64) (*UserNameInfo, error) {

	endpoint := API_URL + "/users/" + strconv.FormatInt(userNameId, 10) + "/info/?"

	resp, err := this.request("GET", endpoint, nil)
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

// Get user tags.
func (this *Instagram) GetUserTags(userNameId int64) (*UserTags, error) {

	endpoint := API_URL + "/usertags/" + strconv.FormatInt(userNameId, 10) + "/feed/?rank_token=" +
		this.rankToken + "&ranked_content=false"

	resp, err := this.request("GET", endpoint, nil)
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

// Search tags.
func (this *Instagram) SearchTags(query string) (*SearchTags, error) {

	endpoint := API_URL + "/tags/search/?is_typeahead=true&q=" + query + "&rank_token=" + this.rankToken

	resp, err := this.request("GET", endpoint, nil)
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

// Get tagged media.
func (this *Instagram) TagFeed(tag, maxId string) (*TagFeed, error) {

	endpoint := API_URL + "/feed/tag/" + tag + "/?rank_token=" + this.rankToken + "&ranked_content=false&max_id=" + maxId

	resp, err := this.request("GET", endpoint, nil)
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

// Request for Login method. Needs to get the authorization cookies.
func (this *Instagram) requestLogin(method, endpoint string, body io.Reader) (*http.Response, error) {
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

// Main request for all other methods. Reading the authorization cookies.
func (this *Instagram) requestMain(method, endpoint string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", USER_AGENT)
	for _, cookie := range this.cookies {
		req.AddCookie(cookie)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Request with five attempts re-login. Re-login if getting error 'login_required'.
func (this *Instagram) request(method, endpoint string, body io.Reader) ([]byte, error) {

	for attempt := 0; attempt < 5; attempt++ {

		resp, err := this.requestMain(method, endpoint, body)
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
			err = this.Login()
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
