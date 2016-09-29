package statx

// AuthService handles calls to /groups
type AuthService struct {
	client *Client
}

// AuthResponse the reponse back from /auth/login
type AuthResponse struct {
	ClientID    *string `json:"clientId"`
	ClientName  *string `json:"clientName"`
	PhoneNumber *string `json:"phoneNumber"`
}

// Credentials is the response from /auth/verifyCode
type Credentials struct {
	APIKey    *string `json:"apiKey"`
	AuthToken *string `json:"authToken"`
}

// Login to a users account
func (s *AuthService) Login(phoneNumber, clientName string) (*AuthResponse, *Response, error) {
	u := "auth/login"
	loginInfo := map[string]interface{}{
		"phoneNumber": phoneNumber,
		"clientName":  clientName,
	}
	req, err := s.client.NewRequest("POST", u, loginInfo)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(AuthResponse)
	resp, err := s.client.Do(req, uResp)
	return uResp, resp, err
}

// Verify to a users account
func (s *AuthService) Verify(verificationCode string, auth *AuthResponse) (*Credentials, *Response, error) {
	u := "auth/verifyCode"
	verification := map[string]interface{}{
		"phoneNumber":      auth.PhoneNumber,
		"clientId":         auth.ClientID,
		"verificationCode": verificationCode,
	}
	req, err := s.client.NewRequest("POST", u, verification)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(Credentials)
	resp, err := s.client.Do(req, uResp)
	return uResp, resp, err
}
