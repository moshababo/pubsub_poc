package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"pubsub_poc/common"
	"pubsub_poc/config"
	"pubsub_poc/publisher"
	"pubsub_poc/wire"
)

var logger = common.NewLogger(common.LogLevelInfo, "CLIENT")

func main() {
	app := cli.NewApp()
	app.Name = "cli client"
	app.Version = "0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "url",
			Usage: "AMQP endpoint",
			Value: config.DefaultConfig().Url,
		},
	}
	app.Commands = []cli.Command{
		{
			Name: "add",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "key", Required: true},
				cli.StringFlag{Name: "val", Required: true},
			},
			Action: func(ctx *cli.Context) error {
				url := ctx.GlobalString("url")
				key := ctx.String("key")
				val := ctx.String("val")

				conn, err := common.NewConnection(url, logger)
				if err != nil {
					return err
				}
				msg := &wire.MsgAddItem{Key: key, Val: val}
				err = publisher.New(conn, logger).Publish(msg.Encode())
				if err != nil {
					return err
				}

				return nil
			},
		},
		{
			Name: "remove",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "key", Required: true},
			},
			Action: func(ctx *cli.Context) error {
				url := ctx.GlobalString("url")
				key := ctx.String("key")

				conn, err := common.NewConnection(url, logger)
				if err != nil {
					return err
				}
				msg := &wire.MsgRemoveItem{Key: key}
				err = publisher.New(conn, logger).Publish(msg.Encode())
				if err != nil {
					return err
				}

				return nil
			},
		},
		{
			Name: "get",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "key", Required: true},
			},
			Action: func(ctx *cli.Context) error {
				url := ctx.GlobalString("url")
				key := ctx.String("key")

				conn, err := common.NewConnection(url, logger)
				if err != nil {
					return err
				}
				msg := &wire.MsgGetItem{Key: key}
				logger.Info("publishing msg: %v", msg)
				err = publisher.New(conn, logger).Publish(msg.Encode())
				if err != nil {
					return err
				}

				return nil
			},
		},
		{
			Name: "getall",
			Action: func(ctx *cli.Context) error {
				url := ctx.GlobalString("url")

				conn, err := common.NewConnection(url, logger)
				if err != nil {
					return err
				}
				msg := &wire.MsgGetAllItems{}
				err = publisher.New(conn, logger).Publish(msg.Encode())
				if err != nil {
					return err
				}

				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
