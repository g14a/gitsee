package api

import (
	_ "github.com/joho/godotenv/autoload"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	req, err := http.NewRequest("GET", "/user/g14a", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUserInfo)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("HTTP Code expected %v, got %v", http.StatusOK, status)
	}

	expected := `{
	    "avatar_url": "https://avatars3.githubusercontent.com/u/17702388?u=9c7235b6f5909386ad20945df7414f88246fe581&v=4",
	    "created_at": "Joined Github 4 years ago",
	    "followers": 26,
	    "location": "Chennai, India",
	    "name": "Gowtham Munukutla",
	    "url": "https://github.com/g14a"
	}`

	if rr.Body.String() != expected {
		t.Errorf("Expected Response %v, got %v", expected, rr.Body.String())
	}
}
