package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

var (
	delayTrains []DelayTrain
)

type DelayTrain struct {
	Name          string `json:"name"`
	Company       string `json:"company"`
	LastUpdateGmt int    `json:"lastupdate_gmt"`
	Source        string `json:"source"`
}

type DelayTrains []DelayTrain

func main() {
	app := cli.NewApp()

	app.Name = "train"
	app.Usage = "This app will help you to know if the train which you use is delay."
	app.Version = "1.0.0"

	app.Action = func(context *cli.Context) error {

		targetTrain := context.Args().Get(0)
		jsonBytes, err := getJSON()
		if err != nil {
			log.Fatal("Getting response failed: %v", err)
		}

		if err := parseJSONtoDelayTrain(jsonBytes); err != nil {
			log.Fatal(err)
		}

		for _, train := range delayTrains {
			if targetTrain == train.Name {
				fmt.Println("遅延しています。最終更新: %v", train.LastUpdateGmt)
				continue
			}
		}

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func getJSON() ([]byte, error) {
	//FIXME 毎回APIを叩くと遅いので、最終レスポンスから１時間経過していない場合は処理をskipさせる
	response, err := http.Get("https://tetsudo.rti-giken.jp/free/delay.json")

	defer response.Body.Close()

	jsonBytes, err := ioutil.ReadAll(response.Body)

	return jsonBytes, err
}

func parseJSONtoDelayTrain(jsonBytes []byte) error {
	if err := json.Unmarshal(jsonBytes, &delayTrains); err != nil {
		return err
	}

	return nil
}
