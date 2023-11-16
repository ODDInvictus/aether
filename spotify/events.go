package spotify

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/KokopelliMusic/go-lib/logger"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
)

func ListenToEvents(state *SpotifyPlayer) {
	Log("Listening to player events")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: viper.GetString("spotify.ws"), Path: "/events"}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		logger.Err("failed to connect to websocket", err)
	}

	defer conn.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				panic("failed to read msg")
			}

			var res map[string]interface{}

			err = json.Unmarshal(message, &res)

			if err != nil {
				continue
			}

			Log(fmt.Sprint(res["event"]))

			switch res["event"] {
			case "contextChanged":
				state.contextUri = fmt.Sprint(res["uri"])
			case "trackChanged":
				state.uri = fmt.Sprint(res["uri"])
			case "playbackEnded":
				state.paused = true
				state.trackTime = 0
			case "playbackPaused":
				state.paused = true
				state.trackTime = int64(res["trackTime"].(float64))
			case "playbackResumed":
				state.paused = false
				state.trackTime = int64(res["trackTime"].(float64))
			case "playbackFailed":
				state.paused = true
				state.trackTime = 0
			case "trackSeeked":
				state.trackTime = int64(res["trackTime"].(float64))
			case "metadataAvailable":
				track := Track{}

				jsonString, _ := json.Marshal(res["track"])
				json.Unmarshal(jsonString, &track)

				state.metadata = track
			case "playbackHaltStateChanged":
				x, err := strconv.ParseBool(fmt.Sprint(res["halted"]))
				
				if err != nil {
					state.paused = true
				} else {
					state.paused = x
				}
				
				state.trackTime = int64(res["trackTime"].(float64))
			case "panic":
				Log("Spotify failed, restarting song")
				PlayPause()
				PlayPause()
				// Load(viper.GetString("fallback.playlist"), true, true)
			}
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := conn.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				logger.Err("write:", err)
				return
			}
		case <-interrupt:
			Log("closing connection...")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				logger.Err("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}