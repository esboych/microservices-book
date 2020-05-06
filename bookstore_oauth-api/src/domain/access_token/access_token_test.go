package access_token

import (
	"fmt"
	"testing"
	"time"
	//"github.com/stretchr/testify/assert"
)
//func TestAccessTokenConstants (t *testing.T) {
//	assert.EqualValues(t, 24, expirationTime, "exp time shoul be 24h")
//}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	fmt.Println(at)
	if at.isExpired() {
		t.Error("brand new access token should not be expired")
	}
	if at.AccessToken != "" {
		t.Error("new access token should not be empty")
	}
	if at.UserId != 0 {
		t.Error("new access token should not have associated UserId")
	}
}

func TestAccessTokenIsExpired (t *testing.T) {
	at := AccesToken{}
	fmt.Println(at)
	if !at.isExpired() {
		t.Error("empty access token should be expired by default")
	}
	at.Expires = time.Now().UTC().Add(3*time.Hour).Unix()
	if at.isExpired() {
		t.Error("access token created 3 hours ago shoul NOT be expired")
	}
}

