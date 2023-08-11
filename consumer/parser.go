package consumer

import (
	"errors"
	"fmt"
	"pubsub_poc/common"
	"pubsub_poc/wire"
	"reflect"
)

func parseMessage(data []byte) (interface{}, error) {
	if len(data) == 0 {
		return nil, errors.New("empty message")
	}

	msgType := wire.MsgType(data[0])
	t, ok := wire.TypeToStructType[msgType]
	if !ok {
		return nil, fmt.Errorf("invalid msg type: %v", msgType)
	}

	decoder, ok := reflect.New(t).Interface().(common.Decoder)
	if !ok {
		panic(fmt.Sprintf("implementation missing, msgType: %v", msgType))
	}

	decoder.Decode(data)
	return decoder, nil
}
