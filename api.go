package vk

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

var (
	// Version of VK API
	Version = "5.42"
	// APIURL is a base to make API calls
	APIURL = "https://api.vk.com/method/"
	// HTTPS defines if use https instead of http. 1 - use https. 0 - use http
	HTTPS = 1
)

// API holds data to use for communication
type API struct {
	AppID           string
	Secret          string
	Scope           []string
	AccessToken     string
	Expiry          time.Time
	UserID          string
	UserEmail       string
	callbackURL     *url.URL
	requestTokenURL *url.URL
	accessTokenURL  *url.URL
}

// NewAPI creates instance of API
func NewAPI(appID, secret string, scope []string, callback string) (api *API, e error) {
	var (
		callbackURL *url.URL
		reqTokURL   *url.URL
		accTokURL   *url.URL
	)

	if appID == "" {
		e = fmt.Errorf("AppId is nil")
		return
	}
	if secret == "" {
		e = fmt.Errorf("Secret is nil")
		return
	}
	if callbackURL, e = url.Parse(callback); e != nil {
		//e = fmt.Errorf("CallbackURL is nil")
		return
	}
	reqTokURL, e = url.Parse("https://oauth.vk.com/authorize")
	if e != nil {
		return
	}
	accTokURL, e = url.Parse("https://oauth.vk.com/access_token")
	if e != nil {
		return
	}

	return &API{
		AppID:           appID,
		Secret:          secret,
		Scope:           scope,
		callbackURL:     callbackURL,
		requestTokenURL: reqTokURL,
		accessTokenURL:  accTokURL,
	}, nil
}

// getAPIURL prepares URL instance with defined method
func (api *API) getAPIURL(method string) *url.URL {
	q := url.Values{
		"v":            {Version},
		"https":        {strconv.Itoa(HTTPS)},
		"access_token": {api.AccessToken},
	}.Encode()
	apiURL, _ := url.Parse(APIURL + method + "?" + q)
	return apiURL
}
