package statx

import "errors"

// GroupsService handles calls to /groups
type GroupsService struct {
	client *Client
}

// GroupList holds the list of groups from GET /groups
type GroupList struct {
	Data []Group `json:"data"`
}

// Group holds the group info
type Group struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// List gets the lists of StatX Groups
func (s *GroupsService) List() (*GroupList, *Response, error) {
	if s.client.Credentials == nil {
		return nil, nil, errors.New("Must be authenticated")
	}

	u := "groups"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(GroupList)
	resp, err := s.client.Do(req, uResp)
	return uResp, resp, err
}

// Get - gets information about a single group
func (s *GroupsService) Get(groupID string) (*Group, *Response, error) {
	if s.client.Credentials == nil {
		return nil, nil, errors.New("Must be authenticated")
	}

	u := "groups/" + groupID
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(Group)
	resp, err := s.client.Do(req, uResp)
	return uResp, resp, err
}
