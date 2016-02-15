package libs

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	//"io/ioutil"
)

type instagram struct {
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

func NewInstagram(userName, password string) *instagram {
	return &instagram{
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
}

// Login to Instagram.
func (this *instagram) Login() error {

	fetch := fmt.Sprintf("%s/si/fetch_headers/?challenge_type=signup&guid=%s", API_URL, generateUUID(false))

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
	data := &Login{
		DeviceId:          this.deviceId,
		Guid:              this.uuid,
		UserName:          this.userName,
		Password:          this.password,
		Csrftoken:         this.token,
		LoginAttemptCount: "0",
	}

	jsonData, err := json.Marshal(data)
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

	var loginResponse *LoginResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&loginResponse)
	if err != nil {
		return err
	}

	if loginResponse.Status == "fail" {
		return errors.New(loginResponse.Message)
	}

	this.isLoggedIn = true
	this.userNameId = loginResponse.LoggedInUser.Pk
	this.rankToken = strconv.FormatInt(this.userNameId, 10) + "_" + this.uuid

	return nil
}

// Get media likers.
func (this *instagram) GetMediaLikers(mediaId string) (*MediaLikers, error) {

	if !this.isLoggedIn {
		return nil, errors.New("Not logged in.")
	}

	endpoint := fmt.Sprintf("%s/media/%s/likers/?", API_URL, mediaId)

	resp, err := this.request("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var likes *MediaLikers
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&likes)
	if err != nil {
		return nil, err
	}

	if likes.Status == "fail" {
		return nil, errors.New(likes.Message)
	}

	return likes, nil
}

// Get media comments.
func (this *instagram) GetMediaComments(mediaId string) (*MediaComments, error) {

	if !this.isLoggedIn {
		return nil, errors.New("Not logged in.")
	}

	endpoint := fmt.Sprintf("%s/media/%s/comments/?", API_URL, mediaId)

	resp, err := this.request("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var mediaComments *MediaComments
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&mediaComments)
	if err != nil {
		return nil, err
	}

	if mediaComments.Status == "fail" {
		return nil, errors.New(mediaComments.Message)
	}

	return mediaComments, nil
}

// Get recent activity.
func (this *instagram) GetRecentActivity() (*RecentActivity, error) {

	if !this.isLoggedIn {
		return nil, errors.New("Not logged in.")
	}

	endpoint := API_URL + "/news/inbox/?"

	resp, err := this.request("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var recentActivity *RecentActivity
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&recentActivity)
	if err != nil {
		return nil, err
	}

	if recentActivity.Status == "fail" {
		return nil, errors.New(recentActivity.Message)
	}

	return recentActivity, nil
}

// Search users.
func (this *instagram) SearchUsers(query string) (*SearchUsers, error) {

	if !this.isLoggedIn {
		return nil, errors.New("Not logged in.")
	}

	endpoint := API_URL + "/users/search/?ig_sig_key_version=" + SIG_KEY_VERSION +
		"&is_typeahead=true&query=" + query + "&rank_token=" + this.rankToken

	resp, err := this.request("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var searchUsers *SearchUsers
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&searchUsers)
	if err != nil {
		return nil, err
	}

	if searchUsers.Status == "fail" {
		return nil, errors.New(searchUsers.Message)
	}

	return searchUsers, nil
}

// Get username info.
func (this *instagram) GetUserNameInfo(userNameId int64) (*UserNameInfo, error) {

	if !this.isLoggedIn {
		return nil, errors.New("Not logged in.")
	}

	endpoint := fmt.Sprintf("%s/users/%d/info/?", API_URL, userNameId)

	resp, err := this.request("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userNameInfo *UserNameInfo
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&userNameInfo)
	if err != nil {
		return nil, err
	}

	if userNameInfo.Status == "fail" {
		return nil, errors.New(userNameInfo.Message)
	}

	return userNameInfo, nil
}

// Get user tags.
func (this *instagram) GetUserTags(userNameId int64) (*UserTags, error) {

	if !this.isLoggedIn {
		return nil, errors.New("Not logged in.")
	}

	endpoint := fmt.Sprintf("%s/usertags/%d/feed/?rank_token=%s&ranked_content=true&",
		API_URL, userNameId, this.rankToken)

	resp, err := this.request("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userTags *UserTags
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&userTags)
	if err != nil {
		return nil, err
	}

	if userTags.Status == "fail" {
		return nil, errors.New(userTags.Message)
	}

	return userTags, nil
}

// Search tags.
func (this *instagram) SearchTags(query string) (*SearchTags, error) {

	if !this.isLoggedIn {
		return nil, errors.New("Not logged in.")
	}

	endpoint := fmt.Sprintf("%s/tags/search/?is_typeahead=true&q=%s&rank_token=%s",
		API_URL, query, this.rankToken)

	resp, err := this.request("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var searchTags *SearchTags
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&searchTags)
	if err != nil {
		return nil, err
	}

	if searchTags.Status == "fail" {
		return nil, errors.New(searchTags.Message)
	}

	return searchTags, nil
}

//body, _ := ioutil.ReadAll(resp.Body)
//fmt.Println(string(body))

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