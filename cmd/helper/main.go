package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	FlagVolumePath = "p"
	FlagVolumeSize = "s"
	FlagVolumeMode = "m"
	cmdSetup       = "setup"
	cmdTeardown    = "teardown"
	envVolDir      = "VOL_DIR"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	a := cli.NewApp()
	a.Usage = "helper"

	a.Commands = []cli.Command{
		{
			Name: cmdSetup,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  FlagVolumePath,
					Usage: "Required. Volume absolute path.",
					Value: "",
				},
				cli.IntFlag{
					Name:  FlagVolumeSize,
					Usage: "Required. Volume size in bytes.",
					Value: 0,
				},
				cli.StringFlag{
					Name:  FlagVolumeMode,
					Usage: "Required. Volume mode.",
					Value: "",
				},
			},
			Action: func(c *cli.Context) {
				if err := setupVolume(c); err != nil {
					logrus.Fatalf("Error setup volume: %v", err)
				}
			},
		},
		{
			Name: cmdTeardown,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  FlagVolumePath,
					Usage: "Required. Volume absolute path.",
					Value: "",
				},
				cli.IntFlag{
					Name:  FlagVolumeSize,
					Usage: "Required. Volume size in bytes.",
					Value: 0,
				},
				cli.StringFlag{
					Name:  FlagVolumeMode,
					Usage: "Required. Volume mode.",
					Value: "",
				},
			},
			Action: func(c *cli.Context) {
				if err := teardownVolume(c); err != nil {
					logrus.Fatalf("Error setup volume: %v", err)
				}
			},
		},
	}

	if err := a.Run(os.Args); err != nil {
		logrus.Fatalf("Critical error: %v", err)
	}
}

func setupVolume(c *cli.Context) error {
	volPath := c.String(FlagVolumePath)
	volSize := c.Int(FlagVolumeSize)
	volMode := c.String(FlagVolumeMode)

	logrus.Infof("volume path: %v, size: %v, mode: %v", volPath, volSize, volMode)

	return os.MkdirAll(os.Getenv(envVolDir), 0777)
}

func teardownVolume(c *cli.Context) error {
	volPath := c.String(FlagVolumePath)
	volSize := c.Int(FlagVolumeSize)
	volMode := c.String(FlagVolumeMode)

	logrus.Infof("volume path: %v, size: %v, mode: %v", volPath, volSize, volMode)

	return os.RemoveAll(os.Getenv(envVolDir))
}
