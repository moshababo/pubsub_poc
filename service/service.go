package service

import (
	"context"
	"fmt"
	"pubsub_poc/common"
	"pubsub_poc/config"
	"pubsub_poc/consumer"
	"pubsub_poc/wire"
	"strings"
)

type Service struct {
	cfg    *config.Config
	om     *orderedMap[string, string]
	logger *common.Logger
}

func New(cfg *config.Config, logger *common.Logger) *Service {
	return &Service{
		cfg:    cfg,
		om:     newOrderedMap[string, string](),
		logger: logger.WithTag("SERVICE"),
	}
}

func (s *Service) Start(ctx context.Context) error {
	conn, err := common.NewConnection(s.cfg.Url, s.logger)
	if err != nil {
		return err
	}

	msgChan, errChan, err := consumer.New(conn, s.logger).Consume(ctx)
	if err != nil {
		return err
	}

	for {
		select {
		case err, ok := <-errChan:
			if !ok {
				return nil
			}
			s.logger.Error("consuming error: %v", err)
		case msg, ok := <-msgChan:
			if !ok {
				return nil
			}
			s.logger.Info("incoming msg: %v", msg)
			s.handleMsg(msg)
		}
	}
}

func (s *Service) handleMsg(msg interface{}) {
	switch t := msg.(type) {
	case *wire.MsgAddItem:
		s.addItem(t.Key, t.Val)
	case *wire.MsgRemoveItem:
		s.removeItem(t.Key)
	case *wire.MsgGetItem:
		s.getItem(t.Key)
	case *wire.MsgGetAllItems:
		s.getAllItems()
	default:
		panic(fmt.Sprintf("invalid type: %v", t))
	}
}

func (s *Service) addItem(key string, val string) {
	s.logger.Info("adding [key: %v, val: %v]", key, val)
	s.om.Add(key, val)
}

func (s *Service) removeItem(key string) {
	s.logger.Info("removing [key: %v]", key)
	s.om.Remove(key)
}

func (s *Service) getItem(key string) {
	val := s.om.Get(key)
	if val == nil {
		s.logger.Info("[key: %v] doesn't exist", key)
		return
	}

	s.logger.Info(val.String())
}

func (s *Service) getAllItems() {
	pairs := s.om.GetAll()

	var sb strings.Builder
	for i, kv := range pairs {
		sb.WriteString(kv.String())
		if i < len(pairs)-1 {
			sb.WriteString(", ")
		}
	}

	s.logger.Info(sb.String())
}
