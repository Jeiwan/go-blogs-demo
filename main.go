package main

import (
	"os"

	"github.com/Jeiwan/goblogs/cmd"
	"github.com/sirupsen/logrus"
)

func main() {
	if os.Getenv("DEBUG") != "" {
		logrus.SetLevel(logrus.DebugLevel)
	}

	if err := cmd.Run(); err != nil {
		logrus.Fatalln(err)
	}
}
