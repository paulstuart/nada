package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getBody(fn http.HandlerFunc, path string) (string, int, error) {
	req := httptest.NewRequest("GET", path, nil)
	w := httptest.NewRecorder()
	fn(w, req)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), resp.StatusCode, err
}

func TestHey(t *testing.T) {
	path := "/foo"
	body, _, err := getBody(hey, path)
	if err != nil {
		t.Error(err)
		return
	}
	expects := fmt.Sprintf("you said: %s\n", path[1:])
	if body != expects {
		t.Errorf("Expected: '%s' -- Got: '%s'\n", expects, body)
	}
}

func TestBad(t *testing.T) {
	path := "/404"
	_, code, err := getBody(pageNotFound, path)
	if err != nil {
		t.Error(err)
		return
	}
	expects := 404
	if code != expects {
		t.Errorf("Expected: '%d' -- Got: '%d'\n", expects, code)
	}
}
