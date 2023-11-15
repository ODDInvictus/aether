package spotify

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/ODDInvictus/aether/logger"
	"github.com/spf13/viper"
)

var apiUrl string

func Init(startPlaying bool) {
	logger.Log("Initializing spotify player")
	apiUrl = viper.GetString("spotify.url")

	if startPlaying {
		Load(viper.GetString("playlist.fallback"), true, true)
	}
}

/*
Load a track from a given URI uri, can specify to start playing with play and to shuffle with shuffle.
*/
func Load(uri string, startPlaying bool, shuffle bool) (bool, error) {
	url := fmt.Sprintf("/player/load?uri=%s&play=%t&shuffle=%t", uri, startPlaying, shuffle)

	Log(url)
	return emptyPost(url)
}

/*
Toggle play/pause status. Useful when using a remote.
*/
func PlayPause() (bool, error) {
	return emptyPost("/player/play-pause")
}

/*
Pause playback.
*/
func Pause() (bool, error) {
	return emptyPost("/player/pause")
}

/*
Resume playback.
*/
func Resume() (bool, error) {
	return emptyPost("/player/resume")
}

/*
Skip to next track.
*/
func Next() (bool, error) {
	return emptyPost("/player/next")
}

/*
Skip to previous track.
*/
func Prev() (bool, error) {
	return emptyPost("/player/prev")
}

/*
Seek to a given position in ms specified by pos.
*/
func Seek(pos int) (bool, error) {
	return emptyPost("/player/seek?pos=" + fmt.Sprint(pos))
}

/*
Set shuffle enabled or disabled accordingly to val.
*/
func Shuffle(shuffle bool) (bool, error) {
	return emptyPost("/player/shuffle?val=" + strconv.FormatBool(shuffle))
}

/*
Set repeating mode as specified by val (modes are none, track, context).
*/
func Repeat(val string) (bool, error) {
	if val != "none" && val != "track" && val != "context" {
		return false, errors.New("invalid context mode, possible options: none, track, context")
	}

	return emptyPost("/player/repeat?val=" + val)
}

/*
Either set volume to a given volume value (from 0 to 65536), or change it by a step count (positive or negative).

Will use step if volume is negative
*/
func SetVolume(volume int, step int) (bool, error) {
	if (volume < 0 && step == 0) {
		return false, errors.New("invalid parameters, volume is negative and step is not set")
	}

	if (volume < 0) {
		return emptyPost("/player/set-volume?step=" + fmt.Sprint(step))
	}

	if (volume < 0 || volume > 65536) {
		return false, errors.New("invalid parameters, volume should be between 0 and 65536")
	}

	return emptyPost("/player/set-volume?volume=" + fmt.Sprint(volume))
}

/*
Up the volume a little bit.
*/
func VolumeUp() (bool, error) {
	return emptyPost("/player/volume-up")
}

/*
Lower the volume a little bit.
*/
func VolumeDown() (bool, error) {
	return emptyPost("/player/volume-down")
}

/*
Retrieve information about the current track (metadata and time).
*/
func Current() (*PlaybackState, error) {
	url := "/player/current"
	var state PlaybackState

	_, err := postWithReturn(url, &state)

	return &state, err
}

/*
Retrieve all the tracks in the player state with metadata, you can specify withQueue.
*/
func Tracks(withQueue bool) (*TracksState, error) {
	url := "/player/tracks"

	if withQueue {
		url += "?withQueue=true"
	}

	var state TracksState

	_, err := postWithReturn(url, &state)

	return &state, err
}

/*
Add a track to the queue, specified by uri.
*/
func AddToQueue(uri string) (bool, error) {
	return emptyPost("/player/addToQueue?uri=" + uri)
}

/*
Remove a track from the queue, specified by uri.
*/
func RemoveFromQueue(uri string) (bool, error) {
	return emptyPost("/player/removeFromQueue?uri=" + uri)
}

/*
Retrieve metadata. metadataType can be one of episode, track, album, show, artist or playlist, uri is the standard Spotify uri.
*/
func Metadata(metadataType string, uri string) (bool, error) {
	return false, errors.New("metadata is not implemented")
}

/*
Retrieve metadata. uri is the standard Spotify uri, the type will be guessed based on the provided uri.
*/
func MetadataPerUri(uri string) (bool, error) {
	return false, errors.New("metadata per uri is not implemented")
}

/*
Make a search.
*/
func Search(query string) (*SearchResult, error) {
	url := "/search/" + query

	var state SearchResult

	_, err := postWithReturn(url, &state)

	return &state, err
}

/*
Request an access token for a specific scope (or a comma separated list of scopes).
*/
func Token(scope string) (bool, error) {
	return false, errors.New("token is not implemented")
}

/*
Retrieve a list of profiles that are followers of the specified user.
*/
func ProfileFollowers(uid string) (bool, error) {
	return false, errors.New("profile followers is not implemented")
}

/*
Retrieve a list of profiles that the specified user is following.
*/
func ProfileFollowing(uid string) (bool, error) {
	return false, errors.New("profile following is not implemented")
}

/*
Returns a json model that contains basic information about the current session.
*/
func Instance() (*InstanceData, error) {
	url := "/instance"

	var state InstanceData

	_, err := getWithReturn(url, &state)

	return &state, err
}

/*
Terminates the API server.
*/
func TerminateServer() (bool, error) {
	return emptyPost("/instance/terminate")
}

/*
Closes the current session (and player).
*/
func CloseSession() (bool, error) {
	return emptyPost("/instance/close")
}

/*
List all Spotify Connect devices on the network.
*/
func DiscoveryList() (bool, error) {
	return false, errors.New("discovery list is not implemented")
}

/*
Use any endpoint from the public Web API by appending it to /web-api/, 
the request will be made to the API with the correct Authorization header and the result will be returned. 
The method, body, and content type headers will pass through. 
Additionally, you can specify an X-Spotify-Scope header to override the requested scope, by default all will be requested.
*/
func WebApiPassthrough() (bool, error) {
	return false, errors.New("api passthrough is not implemented")
}


func Log(str string) {
	logger.Verbose("[Spotify] " + str)
}

func fail(str string) {
	logger.Warn("[Spotify] " + str)
}

func emptyPost(url string) (bool, error) {
	Log("Calling " + url)

	_, err := http.Post(apiUrl + url, "", bytes.NewBufferString(""))

	if err != nil {
		fail(fmt.Sprintf("Call to %s failed", url))
		return false, err
	}

	Log(fmt.Sprintf("Call to %s successful", url))

	return true, nil
}

func postWithReturn(url string, v any) (bool, error) {
	Log("Calling " + url)

	resp, err := http.Post(apiUrl + url, "", bytes.NewBufferString(""))

	if err != nil {
		fail(fmt.Sprintf("Call to %s failed", url))
		return false, err
	}

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		fail(fmt.Sprintf("Call to %s failed", url))
		return false, err
	}

	err = json.Unmarshal(respBody, &v)

	if err != nil {
		fail("Unmarschal failed for Current()")
		return false, err
	}

	Log(fmt.Sprintf("Call to %s successful", url))

	return false, nil
}

func getWithReturn(url string, v any) (bool, error) {
	Log("Calling " + url)

	resp, err := http.Get(apiUrl + url)

	if err != nil {
		fail(fmt.Sprintf("Call to %s failed", url))
		return false, err
	}

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		fail(fmt.Sprintf("Call to %s failed", url))
		return false, err
	}

	err = json.Unmarshal(respBody, &v)

	if err != nil {
		fail("Unmarschal failed for Current()")
		return false, err
	}

	Log(fmt.Sprintf("Call to %s successful", url))

	return false, nil
}