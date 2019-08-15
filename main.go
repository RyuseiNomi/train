package main

import (
	"fmt"
	"log"
	"os"

	job "github.com/RyuseiNomi/train/job"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "train"
	app.Usage = "This app will help you to know if the train which you use is delay."
	app.Version = "1.0.0"

	app.Action = func(context *cli.Context) error {

		targetTrain := context.Args().Get(0)
		jsonBytes, err := job.GetJSON()
		if err != nil {
			log.Fatal("Getting response failed: %v", err)
		}

		delayTrains, err := job.ParseJSONtoDelayTrain(jsonBytes)
		if err != nil {
			log.Fatal(err)
		}

		operationStatus := job.GetOperationStatus(targetTrain, delayTrains)

		fmt.Println(operationStatus)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
