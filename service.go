package instagram

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"strings"

	"github.com/pborman/uuid"
)

const (
	// APIURL api url
	APIURL = "https://i.instagram.com/api/v1"
	// USERAGENT user agent
	USERAGENT = "Instagram 7.13.1 Android (23/6.0.1; 515dpi; 1440x2416; huawei/google; Nexus 6P; angler; angler; en_US)"
	// IGSIGKEY ig sig key
	IGSIGKEY = "8b46309eb680f272cc770d214b7dbe5f0c5d26b6cb82b0b740257360b43618f0"
	// SIGKEYVERSION sig key version
	SIGKEYVERSION = "4"
)

// generateUUID return uuid
func generateUUID(t bool) string {
	u := uuid.New()
	if !t {
		return strings.Replace(u, "-", "", -1)
	}
	return u
}

// generateSignature return signature
func generateSignature(data []byte) string {
	h := hmac.New(sha256.New, []byte(IGSIGKEY))
	h.Write(data)
	hash := hex.EncodeToString(h.Sum(nil))
	return "ig_sig_key_version=" + SIGKEYVERSION + "&signed_body=" + hash + "." + url.QueryEscape(string(data))
}

// generateDeviceId return deviceId
func generateDeviceID() string {
	buffer := make([]byte, 32)
	rand.Read(buffer)
	hash := md5.New()
	hash.Write(buffer)
	return "android-" + hex.EncodeToString(hash.Sum(nil))[:16]
}
