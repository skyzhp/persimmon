package mycsrf

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"html/template"
	"io"
	"math"
	"net/url"

	"github.com/revel/revel"
)

// allowMethods are HTTP methods that do NOT require a token
var allowedMethods = map[string]bool{
	"GET":     true,
	"HEAD":    true,
	"OPTIONS": true,
	"TRACE":   true,
}

func RandomString(length int) (string, error) {
	buffer := make([]byte, int(math.Ceil(float64(length)/2)))
	if _, err := io.ReadFull(rand.Reader, buffer); err != nil {
		return "", err
	}
	str := hex.EncodeToString(buffer)
	return str[:length], nil
}

func RefreshToken(c *revel.Controller) {
	token, err := RandomString(64)
	if err != nil {
		panic(err)
	}
	c.Session["csrf_token"] = token
}

// CsrfFilter enables CSRF request token creation and verification.
//
// Usage:
//  1) Add `csrf.CsrfFilter` to the app's filters (it must come after the revel.SessionFilter).
//  2) Add CSRF fields to a form with the template tag `{{ csrftoken . }}`. The filter adds a function closure to the `ViewArgs` that can pull out the secret and make the token as-needed, caching the value in the request. Ajax support provided through the `X-CSRFToken` header.
func CsrfFilter(c *revel.Controller, fc []revel.Filter) {
	token, foundToken := c.Session["csrf_token"]

	if !foundToken {
		RefreshToken(c)
	}

	referer, refErr := url.Parse(c.Request.Referer())
	if refErr != nil {
		c.Result = c.Forbidden("REVEL CSRF: Unable to fetch referer")
		return
	}

	isSameOrigin := false
	if !allowedMethods[c.Request.Method] {
		isSameOrigin = sameOrigin(c.Request.Host, referer)
	} else {
		isSameOrigin = true
	}

	// If the Request method isn't in the white listed methods
	if !allowedMethods[c.Request.Method] && !IsExempt(c) {
		// Token wasn't present at all
		if !foundToken {
			c.Result = c.Forbidden("REVEL CSRF: Session token missing.")
			return
		}

		// Referer header is invalid
		if refErr != nil {
			c.Result = c.Forbidden("REVEL CSRF: HTTP Referer malformed.")
			return
		}

		// Same origin
		if !isSameOrigin {
			c.Result = c.Forbidden("REVEL CSRF: Same origin mismatch.")
			return
		}

		var requestToken string
		// First check for token in post data
		if c.Request.Method == "POST" {
			requestToken = c.Params.Get("csrftoken")
		}

		// Then check for token in custom headers, as with AJAX
		if requestToken == "" {
			requestToken = c.Request.GetHttpHeader("X-CSRFToken")
		}

		if requestToken == "" || !compareToken(requestToken, token) {
			c.Result = c.Forbidden("REVEL CSRF: Invalid token.")
			return
		}
	}

	fc[0](c, fc[1:])

	// Only add token to ViewArgs if the request is: not AJAX, not missing referer header, and is same origin.
	if c.Request.GetHttpHeader("X-CSRFToken") == "" && isSameOrigin {
		c.ViewArgs["_csrftoken"] = token
	}
}

func compareToken(requestToken, token string) bool {
	// ConstantTimeCompare will panic if the []byte aren't the same length
	if len(requestToken) != len(token) {
		return false
	}
	return subtle.ConstantTimeCompare([]byte(requestToken), []byte(token)) == 1
}

// Validates same origin policy
func sameOrigin(u1Host string, u2 *url.URL) bool {
	//revel.INFO.Printf("url:%s,u1.Scheme:%s, u2.Scheme:%s, u1.Host:%s, u2.Host:%s", u1.Path, u1.Scheme, u2.Scheme, u1.Host, u2.Host)
	//revel.INFO.Printf("u1Host:%s,u2Host:%s", u1Host, u2.Host)
	return u1Host == u2.Host
}

func init() {
	revel.TemplateFuncs["csrftoken"] = func(viewArgs map[string]interface{}) template.HTML {
		if tokenFunc, ok := viewArgs["_csrftoken"]; !ok {
			panic("REVEL CSRF: _csrftoken missing from ViewArgs.")
		} else {
			return template.HTML(tokenFunc.(string))
		}
	}
}
