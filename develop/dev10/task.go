package main

import (
	"develop/dev10/task10/mytelnet"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		UseShortOptionHandling: true,
		Commands: []*cli.Command{
			{
				Name:  "gotelnet",
				Usage: "complete a task on the list",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "timeout", Aliases: []string{"t"}},
				},
				Action: func(ctx *cli.Context) error {
					if len(ctx.Args().Slice()) < 1 {
						os.Exit(1)
					}
					obj := mytelnet.CreateTelnet(ctx.Args().Slice()[0], ctx.Args().Slice()[1], ctx.Int("t"))
					err := obj.Run()
					if err != nil {
						fmt.Println(err)
					}
					return err
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
