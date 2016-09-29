package statx

import (
	"errors"
	"fmt"
	"time"
)

// StatsService handles calls to /stats
type StatsService struct {
	client *Client
}

// StatList holds stats
type StatList struct {
	Data []Stat `json:"data"`
}

// Stat holds generic stat data
type Stat struct {
	ID                       string     `json:"id,omitempty"`
	Title                    string     `json:"title,omitempty"`
	VisualType               string     `json:"visualType,omitempty"`
	Notes                    string     `json:"notes,omitempty"`
	Value                    string     `json:"value,omitempty"`
	DueDateTime              *time.Time `json:"dueDateTime,omitempty"`
	LastUpdatedDateTime      *time.Time `json:"lastUpdatedDateTime,omitempty"`
	NotesLastUpdatedDateTime *time.Time `json:"notesLastUpdatedDateTime,omitempty"`
}

// UpdateTimestamps - Populate the timestamps for a stat before saving
func (stat *Stat) UpdateTimestamps() {
	now := time.Now()
	if len(stat.Value) > 0 {
		stat.LastUpdatedDateTime = &now
	}
	if len(stat.Notes) > 0 {
		stat.NotesLastUpdatedDateTime = &now
	}
}

// List gets the lists of StatX Groups
func (s *StatsService) List(groupID string) (*StatList, *Response, error) {
	if s.client.Credentials == nil {
		return nil, nil, errors.New("Must be authenticated")
	}

	u := fmt.Sprintf("groups/%s/stats", groupID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(StatList)
	resp, err := s.client.Do(req, uResp)
	return uResp, resp, err
}

// Get - gets information about a single group
func (s *StatsService) Get(groupID, statID string) (*Stat, *Response, error) {
	if s.client.Credentials == nil {
		return nil, nil, errors.New("Must be authenticated")
	}

	u := fmt.Sprintf("groups/%s/stats/%s", groupID, statID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(Stat)
	resp, err := s.client.Do(req, uResp)
	return uResp, resp, err
}

// Update - gets information about a single group
func (s *StatsService) Update(groupID, statID string, stat *Stat) (*Stat, *Response, error) {
	if s.client.Credentials == nil {
		return nil, nil, errors.New("Must be authenticated")
	}

	u := fmt.Sprintf("groups/%s/stats/%s", groupID, statID)
	stat.UpdateTimestamps()
	req, err := s.client.NewRequest("PUT", u, &stat)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(Stat)
	resp, err := s.client.Do(req, uResp)
	return uResp, resp, err
}
