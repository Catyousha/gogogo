package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

func TestTimeHandler(t *testing.T) {
	// initiate request
	req, err := http.NewRequest("GET", "/time", nil)
	if err != nil {
		t.Fatal(err)
	}

	// recorder for testing
	rr := httptest.NewRecorder()

	// set handler
	handler := http.HandlerFunc(TimeHandler)

	// bind recorder to request
	handler.ServeHTTP(rr, req)

	// compare actual result from recorder to expected res
	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestMethodNotAllowed(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/time", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MethodNotAllowedHandler)
	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestLogin(t *testing.T) {
	UserPass := []byte(`{"Username": "admin", "Password": "admin"}`)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(UserPass))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LoginHandler)
	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		return
	}
}

func TestLogout(t *testing.T) {
	UserPass := []byte(`{"Username": "admin", "Password": "admin"}`)
	req, err := http.NewRequest("POST", "/logout", bytes.NewBuffer(UserPass))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LogoutHandler)
	handler.ServeHTTP(rr, req)

	// Check the HTTP status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		return
	}
}

func TestAdd(t *testing.T) {
	now := int(time.Now().Unix())
	username := "test_" + strconv.Itoa(now)
	users := `[{"Username": "admin", "Password": "admin"}, {"Username":"` + username + `", "Password": "myPass"}]`
	UserPass := []byte(users)
	req, err := http.NewRequest("POST", "/add", bytes.NewBuffer(UserPass))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddHandler)
	handler.ServeHTTP(rr, req)

	// Check the HTTP status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		return
	}
}

func TestGetUserDataHandler(t *testing.T) {
	UserPass := []byte(`{"Username": "admin", "Password": "admin"}`)
	req, err := http.NewRequest("GET", "/username/1", bytes.NewBuffer(UserPass))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// `gorilla/mux` provides the `SetURLVars` function for testing purposes
	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUserDataHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		return
	}

	// exercise 2: The code in testHTTP/server_test.go uses the same value for lastlogin in the expected variable. This is clearly a bug in restdb.go as the value of lastlogin should be updated. After correcting the bug, modify testHTTP/server_test.go to take into account the different values of the lastlogin field.
	expected := `{"id":1,"username":"admin","password":"admin","lastlogin":1620922454,"admin":1,"active":1}`
	serverResponse := rr.Body.String()

	// result := strings.Split(serverResponse, "lastlogin")
	// serverResponse = result[0]
	serverResponse = strings.TrimSpace(serverResponse)

	if serverResponse != expected {
		t.Errorf("handler returned unexpected body: got %v but wanted %v",
			serverResponse, expected)
	}
}
