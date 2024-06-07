package repository

import (
	"context"
	"errors"
	"server-ws-dummy/entities/request"
	"time"
)

func ListenResponse(requestId string) (response string, er error) {

	deadline := time.Now().Add(5 * time.Second)
	c, cancelC := context.WithDeadline(context.Background(), deadline)
	defer func() {
		cancelC()
	}()

	resCh := make(chan string)
	defer func() {
		close(resCh)
	}()

	request.Pool[requestId] = resCh

	defer func() {
		delete(request.Pool, requestId)
	}()

	select {
	case response = <-resCh:
		return response, nil
	case <-c.Done():
		//Timeout exceeded..
		return response, errors.New("timeout exceeded")
	}
}
