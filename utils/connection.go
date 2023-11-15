package utils

import (
	"net/http"

	"github.com/ODDInvictus/aether/logger"
	"github.com/spf13/viper"
)

type ConnectionStatus struct {
	spotify bool
	mp3 bool
}

func InitConnectionStatus() *ConnectionStatus {
	var s ConnectionStatus

	logger.Log("Testing spotify connection")
	s.spotify = s.TestSpotifyConnection()
	logger.Log("Testing MP3 connection")
	s.mp3 = s.TestMP3Connection()

	return &s
}

func (s *ConnectionStatus) TestSpotifyConnection() bool {

	url := viper.GetString("spotify.url")

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		logger.Err("Could not do request", err)
		return false
	}

	_, err = http.DefaultClient.Do(req)

	if err != nil {
		logger.Err("Spotify down?", err)
		s.spotify = false
	}

	return true
}

func (s *ConnectionStatus) TestMP3Connection() bool {
	return true
}
