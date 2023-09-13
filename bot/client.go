// miss voeg ik wel later proxied versie, had ff geen zin

package bot

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

var (
	Timeout time.Duration = time.Second * 13
)

type (
	ClientI interface {
		CalculateDistance(origin, destination string) (*DestinationData, error)
	}

	BotClient struct {
		ClientI
		Token   string
		Timeout time.Duration
	}

	DestinationData struct {
		Distance float64
		Duration float64
	}

	location struct {
		Lat float64
		Lon float64
	}

	coordinates struct {
		Features []struct {
			Geo struct {
				Cord []float64 `json:"coordinates"`
			} `json:"geometry"`
		}
	}

	calculateCtx struct {
		Durations [][]float64 `json:"durations"`
		Distances [][]float64 `json:"distances"`
	}

	payload struct {
		Location [][]float64 `json:"locations"`
		// Destination []float64   `json:"destinations"`
		Metrics []string `json:"metrics"`
		Unit    string   `json:"units"`
	}
)

func (bot *BotClient) CalculateDistance(origin, destination string, done chan<- int) (DestinationData, error) {
	defer func() {
		done <- 1
	}()

	fmt.Println(bot.Token)
	var origin_ location
	var dest_ location

	client := new_http_client(bot.Timeout)

	task := map[int]string{
		0: origin,
		1: destination,
	}

	for i, v := range task {
		loc, err := geocode(client, v, bot.Token)
		if err != nil {
			return DestinationData{}, err
		}

		if i == 0 {
			origin_ = loc
			continue
		}

		dest_ = loc
	}

	ctx, err := matrix(client, origin_, dest_, bot.Token)
	if err != nil {
		return DestinationData{}, err
	}

	return DestinationData{
		Distance: ctx.Distances[0][1],
		Duration: ctx.Durations[0][1],
	}, nil
}

func new_http_client(timeout time.Duration) *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: timeout,
	}

	return client
}

func geocode(client *http.Client, input, token string) (location, error) {
	req := &http.Request{
		Method: http.MethodGet,
		Header: map[string][]string{
			//"authorization": {bot.Token},
			"content-type": {"application/json; charset=utf-8"},
			"user-agent":   {"Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1"},
			"accept":       {"application/json, application/geo+json, application/gpx+xml, img/png; charset=utf-8"},
		},
		Body: nil,
	}

	url, err := url.ParseRequestURI(fmt.Sprintf("https://api.openrouteservice.org/geocode/search?api_key=%s&text=%s", token, url.PathEscape(input)))
	if err != nil {
		return location{}, err
	}

	req.URL = url

	resp, err := client.Do(req)
	if err != nil {
		return location{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return location{}, fmt.Errorf("%s [%d], free quota reached, provide a token as querystring named 'token'. Get a key  from 'openrouteservice.org'", resp.Status, resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return location{}, err
	}

	cord := &coordinates{}
	if err := json.Unmarshal(data, cord); err != nil {
		return location{}, err
	}

	return location{
		Lat: cord.Features[0].Geo.Cord[1],
		Lon: cord.Features[0].Geo.Cord[0],
	}, nil
}

func matrix(client *http.Client, loc, dest location, token string) (calculateCtx, error) {
	pl := &payload{
		Location: [][]float64{{loc.Lon, loc.Lat}, {dest.Lon, dest.Lat}},
		// Destination: []float64{dest.Lat, dest.Lon},
		Metrics: []string{"distance", "duration"},
		Unit:    "km",
	}

	data, err := json.Marshal(pl)
	if err != nil {
		return calculateCtx{}, err
	}

	body := bytes.NewReader(data)
	req, err := http.NewRequest("POST", "https://api.openrouteservice.org/v2/matrix/driving-car", body)
	if err != nil {
		return calculateCtx{}, err
	}

	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return calculateCtx{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return calculateCtx{}, fmt.Errorf("%s [%d], free quota reached, provide a token as querystring from 'openrouteservice.org'", resp.Status, resp.StatusCode)
	}

	ctx := &calculateCtx{}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return calculateCtx{}, err
	}

	if err := json.Unmarshal(data, ctx); err != nil {
		return calculateCtx{}, err
	}

	return *ctx, nil
}

// func XEST() {
// 	client := &BotClient{
// 		Token:   os.Getenv("token"),
// 		Timeout: time.Second * 15,
// 	}

// 	ctx, err := client.CalculateDistance("Velperweg Nederland", "Grutoplein 18 Velp, Nederland")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(ctx.Distance) // KM
// 	fmt.Println(ctx.Duration) // SECONDS
// }
