package instagram_api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

type instagram struct {
	userName          string
	password          string
	token             string
	isLoggedIn        bool
	uuid              string
	deviceId          string
	userNameId        int64
	rankToken         string
	cookies           []*http.Cookie
}

func NewInstagram(userName, password string) *instagram {
	return &instagram{
		userName:          userName,
		password:          password,
		token:             "",
		isLoggedIn:        false,
		uuid:              generateUUID(true),
		deviceId:          generateDeviceId(),
		userNameId:        0,
		rankToken:         "",
		cookies:           nil,
	}
}

// Login to Instagram.
func (this *instagram) Login() error {

	fetch := API_URL + "/si/fetch_headers/?challenge_type=signup&guid=" + generateUUID(false)

	resp, err := this.request("GET", fetch, nil)
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

	resp, err = this.request("POST", API_URL+"/accounts/login/?", bytes.NewReader([]byte(signature)))
	if err != nil {
		return err
	}

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
func (this *instagram) GetMediaLikers(mediaId string) (*MediaLikers, error) {

	endpoint := API_URL + "/media/" + mediaId + "/likers/?"

	resp, err := this.request("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var object *MediaLikers
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&object)
	if err != nil {
		return nil, err
	}

	if object.Status == "fail" {
		return nil, errors.New(object.Message)
	}

	return object, nil
}

// Get media comments.
func (this *instagram) GetMediaComments(mediaId string) (*MediaComments, error) {

	endpoint := API_URL + "/media/" + mediaId + "/comments/?"

	resp, err := this.request("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var object *MediaComments
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&object)
	if err != nil {
		return nil, err
	}

	if object.Status == "fail" {
		return nil, errors.New(object.Message)
	}

	return object, nil
}

// Get recent activity.
func (this *instagram) GetRecentActivity() (*RecentActivity, error) {

	endpoint := API_URL + "/news/inbox/?"

	resp, err := this.request("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var object *RecentActivity
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&object)
	if err != nil {
		return nil, err
	}

	if object.Status == "fail" {
		return nil, errors.New(object.Message)
	}

	return object, nil
}

// Search users.
func (this *instagram) SearchUsers(query string) (*SearchUsers, error) {

	endpoint := API_URL + "/users/search/?ig_sig_key_version=" + SIG_KEY_VERSION +
		"&is_typeahead=true&query=" + query + "&rank_token=" + this.rankToken

	resp, err := this.request("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var object *SearchUsers
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&object)
	if err != nil {
		return nil, err
	}

	if object.Status == "fail" {
		return nil, errors.New(object.Message)
	}

	return object, nil
}

// Get username info.
func (this *instagram) GetUserNameInfo(userNameId int64) (*UserNameInfo, error) {

	endpoint := API_URL + "/users/" + strconv.FormatInt(userNameId, 10) + "/info/?"

	resp, err := this.request("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var object *UserNameInfo
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&object)
	if err != nil {
		return nil, err
	}

	if object.Status == "fail" {
		return nil, errors.New(object.Message)
	}

	return object, nil
}

// Get user tags.
func (this *instagram) GetUserTags(userNameId int64) (*UserTags, error) {

	endpoint := API_URL + "/usertags/" + strconv.FormatInt(userNameId, 10) + "/feed/?rank_token=" +
		this.rankToken + "&ranked_content=false"

	resp, err := this.request("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var object *UserTags
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&object)
	if err != nil {
		return nil, err
	}

	if object.Status == "fail" {
		return nil, errors.New(object.Message)
	}

	return object, nil
}

// Search tags.
func (this *instagram) SearchTags(query string) (*SearchTags, error) {

	endpoint := API_URL + "/tags/search/?is_typeahead=true&q=" + query + "&rank_token=" + this.rankToken

	resp, err := this.request("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var object *SearchTags
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&object)
	if err != nil {
		return nil, err
	}

	if object.Status == "fail" {
		return nil, errors.New(object.Message)
	}

	return object, nil
}

// Get tagged media.
func (this *instagram) TagFeed(tag, maxId string) (*TagFeed, error) {

	endpoint := API_URL + "/feed/tag/" + tag + "/?rank_token=" + this.rankToken + "&ranked_content=false&max_id=" + maxId

	resp, err := this.request("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var object *TagFeed
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&object)
	if err != nil {
		return nil, err
	}

	if object.Status == "fail" {
		return nil, errors.New(object.Message)
	}

	return object, nil
}

func (this *instagram) request(method, endpoint string, body io.Reader) (*http.Response, error) {
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
