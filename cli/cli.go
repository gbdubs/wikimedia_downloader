package main

import (
	"errors"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/gbdubs/wikimedia_downloader"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "Wikimedia Downloader",
		Usage:   "A CLI for downloading files from the wikimedia archive.",
		Version: "1.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "wikimedia_file_name",
				Aliases: []string{"w"},
				Usage:   "the file to download from wikimedia.",
			},
			&cli.StringFlag{
				Name:    "output_file_path",
				Aliases: []string{"o"},
				Usage:   "where to place output, defaults to /tmp/wikimedia_downloader.",
			},
			&cli.BoolFlag{
				Name:  "verbose",
				Usage: "Whether to print the output or silently succeed, if the command succeeds.",
			},
		},
		Action: func(c *cli.Context) error {
			if c.String("wikimedia_file_name") == "" {
				return errors.New("query must be provided")
			}
			v := c.Bool("verbose")
			input := &wikimedia_downloader.Input{
				WikimediaFileName: c.String("wikimedia_file_name"),
				OutputFilePath:    c.String("output_file_path"),
			}
			output, err := input.Execute()
			if err != nil {
				return err
			}
			if v {
				spew.Dump(*output)
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
