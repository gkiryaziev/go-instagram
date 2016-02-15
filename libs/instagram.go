package libs

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/pborman/uuid"
)

const (
	API_URL         = "https://i.instagram.com/api/v1"
	USER_AGENT      = "Instagram 7.13.1 Android (23/6.0.1; 515dpi; 1440x2416; huawei/google; Nexus 6P; angler; angler; en_US)"
	IG_SIG_KEY      = "8b46309eb680f272cc770d214b7dbe5f0c5d26b6cb82b0b740257360b43618f0"
	SIG_KEY_VERSION = "4"
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
		return nil, errors.New("Not logged in")
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
		this.isLoggedIn = false
		return nil, errors.New(likes.Message)
	}

	return likes, nil
}

// Get media comments.
func (this *instagram) GetMediaComments(mediaId string) (*MediaComments, error) {

	if !this.isLoggedIn {
		return nil, errors.New("Not logged in")
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
		this.isLoggedIn = false
		return nil, errors.New(mediaComments.Message)
	}

	return mediaComments, nil
}

// Get recent activity.
func (this *instagram) GetRecentActivity() (*RecentActivity, error) {

	if !this.isLoggedIn {
		return nil, errors.New("Not logged in")
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
		this.isLoggedIn = false
		return nil, errors.New(recentActivity.Message)
	}

	return recentActivity, nil
}

// Search users.
func (this *instagram) SearchUsers(query string) (*SearchUsers, error) {

	if !this.isLoggedIn {
		return nil, errors.New("Not logged in")
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
		this.isLoggedIn = false
		return nil, errors.New(searchUsers.Message)
	}

	return searchUsers, nil
}

// Get username info.
func (this *instagram) GetUserNameInfo(userNameId int64) (*UserNameInfo, error) {

	if !this.isLoggedIn {
		return nil, errors.New("Not logged in")
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
		this.isLoggedIn = false
		return nil, errors.New(userNameInfo.Message)
	}

	return userNameInfo, nil
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

func generateUUID(t bool) string {
	u := uuid.New()
	if !t {
		return strings.Replace(u, "-", "", -1)
	}
	return u
}

func generateSignature(data []byte) string {
	h := hmac.New(sha256.New, []byte(IG_SIG_KEY))
	h.Write(data)
	hash := hex.EncodeToString(h.Sum(nil))
	return "ig_sig_key_version=" + SIG_KEY_VERSION + "&signed_body=" + hash + "." + url.QueryEscape(string(data))
}

func generateDeviceId() string {
	buffer := make([]byte, 32)
	rand.Read(buffer)
	hash := md5.New()
	hash.Write(buffer)
	return "android-" + hex.EncodeToString(hash.Sum(nil))[:16]
}
