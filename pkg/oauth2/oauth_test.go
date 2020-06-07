package oauth2

import (
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestMakeOauthHandlers(t *testing.T) {
	fileOauthConfig := []byte(`oauth2:
  - provider: github
    clientId: 999
    clientSecret: 666
    authURL: "https://github.com/login/oauth/authorize"
    tokenURL: "https://github.com/login/oauth/access_token"
    urlLoginSuccess: "http://localhost:3000"
    scopes: 
      - user
      - repo
      - read:org`)

	r := mux.NewRouter()
	MakeOauthHandlers(r, fileOauthConfig)

	path, err := r.GetRoute("oauth").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/api/oauth", path)

	path, err = r.GetRoute("oauthRedirect").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/oauth/redirect", path)
}