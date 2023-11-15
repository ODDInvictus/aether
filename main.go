package main

import (
	"github.com/ODDInvictus/aether/http"
	"github.com/ODDInvictus/aether/logger"
	"github.com/ODDInvictus/aether/spotify"
	"github.com/ODDInvictus/aether/utils"
)

var spotifyState spotify.SpotifyPlayer

func main() {
	logger.Log("Starting the Aether")
	logger.Debug(true)
	utils.LoadConfig()
	utils.InitConnectionStatus()
	
	spotify.Init(false)
	go spotify.ListenToEvents(&spotifyState)

	router := http.Init()
	router.Run()
}
