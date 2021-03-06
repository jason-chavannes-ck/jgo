package web

import (
	"net/http"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"strconv"
)

type Request struct {
	HttpRequest http.Request
}

func (r *Request) GetCsrfToken() (string, error) {
	if r.HttpRequest.Method != "POST" {
		return "", errors.New("Not a POST request.")
	}
	csrfToken := r.HttpRequest.Header.Get("X-CSRF-Token")
	if csrfToken == "" {
		return "", errors.New("Header empty or not set.")
	}
	return csrfToken, nil
}

func (r *Request) GetFormValue(key string) string {
	r.HttpRequest.ParseForm()
	return r.HttpRequest.Form.Get(key)
}

func (r *Request) GetFormValueInt(key string) int {
	r.HttpRequest.ParseForm()
	valString := r.HttpRequest.Form.Get(key)
	i, _ := strconv.Atoi(valString)
	return i
}

func (r *Request) GetUrlNamedQueryVariable(key string) string {
	vars := mux.Vars(&r.HttpRequest)
	return vars[key]
}

func (r *Request) GetUrlParameter(key string) string {
	return r.HttpRequest.URL.Query().Get(key)
}

func (r *Request) GetHeader(key string) string {
	return r.HttpRequest.Header.Get(key)
}

func (r *Request) GetCookie(key string) string {
	cookie, _ := r.HttpRequest.Cookie(key)
	if cookie == nil {
		return ""
	}
	return cookie.Value
}

func (r *Request) GetURI() string {
	return r.HttpRequest.RequestURI
}

func (r *Request) GetBody() []byte {
	body, _ := ioutil.ReadAll(r.HttpRequest.Body)
	return body
}

func (r *Request) GetPotentialFilename() string {
	return r.HttpRequest.RequestURI[1:len(r.HttpRequest.RequestURI)]
}
