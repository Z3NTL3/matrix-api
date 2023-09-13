package bot

import (
	"context"
	"errors"
	"os"
	"time"
)

func RunBot(origin, destination string) (DestinationData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	ctx.Done()
	client := &BotClient{
		Token:   os.Getenv("token"),
		Timeout: time.Second * 15,
	}

	done := make(chan int, 1)

	result := &DestinationData{}
	var err error = nil

	go func() {
		ctx, err_ := client.CalculateDistance(origin, destination, done)
		if err != nil {
			err = err_
		}

		*result = ctx
	}()

	select {
	case <-done:
		return *result, nil

	case <-ctx.Done():
		return DestinationData{}, errors.New("got no signal back bru")
	}
}
