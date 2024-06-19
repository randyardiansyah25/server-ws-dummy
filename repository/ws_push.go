package repository

import (
	"context"
	"errors"
	"server-ws-dummy/entities/request"
	"server-ws-dummy/entities/ws"
	"time"

	"github.com/randyardiansyah25/wsbase-handler"
)

func PushMessage(message wsbase.Message) (response string, er error) {

	deadline := time.Now().Add(5 * time.Second)
	c, cancelC := context.WithDeadline(context.Background(), deadline)
	defer func() {
		cancelC()
	}()

	resCh := make(chan string)
	defer func() {
		close(resCh)
	}()

	request.Pool[message.SenderId] = resCh

	defer func() {
		delete(request.Pool, message.SenderId)
	}()

	ws.Hub.PushMessage(message)
	
	select {
	case response = <-resCh:
		return response, nil
	case <-c.Done():
		//Timeout exceeded..
		return response, errors.New("timeout exceeded")
	}
}
