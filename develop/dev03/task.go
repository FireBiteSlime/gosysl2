package main

import (
	"develop/dev03/task3/mysort"
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
				Name:  "gosort",
				Usage: "complete a task on the list",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "key", Aliases: []string{"k"}},
					&cli.BoolFlag{Name: "number", Aliases: []string{"n"}},
					&cli.BoolFlag{Name: "reverse", Aliases: []string{"r"}},
					&cli.BoolFlag{Name: "unique", Aliases: []string{"u"}},
					&cli.BoolFlag{Name: "ignore-leading-blanks", Aliases: []string{"b"}},
					&cli.BoolFlag{Name: "check", Aliases: []string{"c"}},
				},
				Action: func(ctx *cli.Context) error {
					if len(ctx.Args().Slice()) < 1 {
						os.Exit(1)
					}
					obj := mysort.CreateSort(ctx.String("key"), ctx.Bool("number"), ctx.Bool("reverse"), ctx.Bool("unique"), ctx.Bool("ignore-leading-blanks"),
						ctx.Bool("check"), ctx.Args().Slice())

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
