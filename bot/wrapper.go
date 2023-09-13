package bot

import (
	"context"
	"errors"
	"os"
	"time"
)

func RunBot(origin, destination, token string) (DestinationData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	done := make(chan int, 1)

	result := &DestinationData{}
	var err error = nil

	go func() {
		client := &BotClient{
			Token:   os.Getenv("token"),
			Timeout: time.Second * 15,
		}

		if token != "" {
			client.Token = token
		}

		ctx, err_ := client.CalculateDistance(origin, destination, done)
		if err_ != nil {
			err = err_
		}

		*result = ctx
	}()

	select {
	case <-done:
		if err != nil {
			return DestinationData{}, err
		}
		return *result, nil

	case <-ctx.Done():
		if err != nil {
			return DestinationData{}, err
		}
		return DestinationData{}, errors.New("got no signal back bru")
	}
}
