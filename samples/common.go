//
// Auto-generated with frappe-doctype-to-go
//
package erpnext_api

import (
	"net/http/cookiejar"
)

// Filter defines a filter type
type Filter {
  Key       string
  Operation string
  Value     string
}

// Credentials defines the credentials
type Credentials struct {
  Username string
  Password string
}

// Session defines the session within ERPNext
type Session struct {
  CookieJar   *cookiejar.Jar
  Credentials Credentials
  ApiURL      string
}

// AuthRoundTripper used for http request middleware
type AuthRoundTripper struct {
	Proxied http.RoundTripper
  Session Session
}

// httpRequest used for internal CRUD request
type httpRequest {
  Method      string,
  URLEndpoint string,
  URLData     *url.Values,
  Body        interface{},
  Session     *Session
}

func (s *Session) Retry() (err error) {
  urlData := url.Values{}
	urlData.Set("usr", s.Credentials.User)
	urlData.Set("pwd", s.Credentials.Password)
  authURL := fmt.Sprintf("%s/method/login", s.ApiURL)

  // Post form
	r, err := http.PostForm(authURL, urlData)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	s.CookieJar, err = cookiejar.New(nil)
	if err != nil {
		return err
	}

	url, err := url.Parse(authURL)
	if err != nil {
		return err
	}
	s.CookieJar.SetCookies(s.ApiURL, r.Cookies())

	return nil
}

// RoundTrip middleware to handle authentication retries
func (art AuthRoundTripper) RoundTrip(req *http.Request) (res *http.Response, e error) {
	do := func() (res *http.Response, e error) {
		return art.Proxied.RoundTrip(req)
	}

	res, e = do()
	if res.StatusCode == http.StatusForbidden {
		if e = art.Session.Retry(); e == nil {
			u, _ := url.Parse(s.ApiURL)
			newCookies := art.Session.CookieJar.Cookies(u)
			req.Header.Del("Cookie")
			for _, c := range newCookies {
				req.AddCookie(c)
			}
			return do()
		}
	}

	return res, e
}

func makeRequest(r *httpRequest) (res *http.Response, err error) {
	var reqBody io.Reader
	if r.Body != nil {
		b, err := json.Marshal(r.Body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewReader(b)
	}

	req, err := http.NewRequest(r.Method, r.URLEndpoint, reqBody)
	if err != nil {
		return
	}

	req.Header.Add("Accept", "application/json, */*")

	if r.URLData != nil {
		req.URL.RawQuery = r.URLData.Encode()
	}

	client := &http.Client{
		Transport: AuthRoundTripper{http.DefaultTransport, r.Session},
		Jar:       r.Session.CookieJar,
	}
	res, err = client.Do(req)

	if err != nil {
		return
	}

	return res, checkForHTTPErrors(res)
}

func checkForHTTPErrors(res *http.Response) (err error) {
	if res.StatusCode >= http.StatusBadRequest {
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		return fmt.Errorf("%d - %s\n%s", res.StatusCode, res.Status, buf.String())
	}
	return
}