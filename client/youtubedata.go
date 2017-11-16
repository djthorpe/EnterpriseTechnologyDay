package client

import (
	"net/http"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

////////////////////////////////////////////////////////////////////////////////

type YouTubeService struct {
	// Data API
	API *youtube.Service
}

////////////////////////////////////////////////////////////////////////////////

// NewYouTubeServiceFromServiceAccountJSON returns a service object given service account details
func NewYouTubeServiceFromDeveloperKey(key string, debug bool) (*YouTubeService, error) {

	client := &http.Client{
		Transport: &transport.APIKey{Key: key},
	}
	service, err := youtube.New(client)
	if err != nil {
		return nil, err
	} else {
		return &YouTubeService{service}, nil
	}
}
