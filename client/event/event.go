package event

import (
	"errors"

	"github.com/LucasLan666666/frp/pkg/msg"
)

var ErrPayloadType = errors.New("error payload type")

type Handler func(payload interface{}) error

type StartProxyPayload struct {
	NewProxyMsg *msg.NewProxy
}

type CloseProxyPayload struct {
	CloseProxyMsg *msg.CloseProxy
}
