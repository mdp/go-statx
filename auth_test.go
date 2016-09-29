package statx

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAuthService(t *testing.T) {
	setup()
	defer teardown()

	mockResponse := `{"phoneNumber":"+15558675309", "clientId":"123", "clientName":"test"}`

	mux.HandleFunc("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, mockResponse)
	})

	authResponse, _, err := client.Auth.Login("+15558675309", "Test Service")
	if err != nil {
		t.Errorf("Auth.Login returned error: %v", err)
	}

	want := &AuthResponse{}
	json.Unmarshal([]byte(mockResponse), want)

	if !reflect.DeepEqual(authResponse, want) {
		t.Errorf("Auth.Login returned %+v, want %+v", authResponse, want)
	}
}

func TestAuthServiceVerification(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/auth/verifyCode", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"authToken":"123", "apiKey":"test"}`)
	})

	authResponse := &AuthResponse{}
	json.Unmarshal([]byte(`{"phoneNumber":"+15558675309", "clientId":"123", "clientName":"test"}`), authResponse)

	credentials, _, err := client.Auth.Verify("123456", authResponse)
	if err != nil {
		t.Errorf("Auth.Login returned error: %v", err)
	}

	want := &Credentials{}
	json.Unmarshal([]byte(`{"authToken":"123", "apiKey":"test"}`), want)

	if !reflect.DeepEqual(credentials, want) {
		t.Errorf("Auth.Verify returned %+v, want %+v", credentials, want)
	}
}
