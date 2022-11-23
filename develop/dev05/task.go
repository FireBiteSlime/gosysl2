package main

import (
	"develop/dev05/task5/mygrep"
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
				Name:  "gogrep",
				Usage: "complete a task on the list",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "after", Aliases: []string{"A"}},
					&cli.StringFlag{Name: "before", Aliases: []string{"B"}},
					&cli.StringFlag{Name: "context", Aliases: []string{"C"}},
					&cli.BoolFlag{Name: "count", Aliases: []string{"c"}},
					&cli.BoolFlag{Name: "ignoreCase", Aliases: []string{"i"}},
					&cli.BoolFlag{Name: "invert", Aliases: []string{"v"}},
					&cli.BoolFlag{Name: "fixed", Aliases: []string{"F"}},
					&cli.BoolFlag{Name: "lineNum", Aliases: []string{"n"}},
				},
				Action: func(ctx *cli.Context) error {
					if len(ctx.Args().Slice()) < 1 {
						os.Exit(1)
					}

					obj := mygrep.CreateGrep(ctx.String("after"), ctx.String("before"), ctx.String("context"),
						ctx.Bool("count"), ctx.Bool("ignoreCase"), ctx.Bool("invert"), ctx.Bool("fixed"),
						ctx.Bool("lineNum"), ctx.Args().Slice()[1:], ctx.Args().Slice()[0])

					err := obj.Run()

					if err != nil {
						fmt.Errorf("Error:", err)
					}
					return obj.Output()
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
