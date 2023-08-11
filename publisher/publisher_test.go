package publisher

import (
	"github.com/stretchr/testify/require"
	"pubsub_poc/common"
	"pubsub_poc/config"
	"pubsub_poc/wire"
	"testing"
)

var (
	cfg    = config.DefaultConfig()
	logger = common.NewLogger(common.LogLevelInfo, "TEST")
)

func TestPublisher_Publish(t *testing.T) {
	r := require.New(t)

	conn, err := common.NewConnection(cfg.Url, logger)
	r.NoError(err)
	r.NotNil(conn)

	pub := New(conn, logger)
	msg := wire.MsgAddItem{}
	err = pub.Publish(msg.Encode())
	r.NoError(err)
}
