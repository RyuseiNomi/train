package json

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type DelayTrain struct {
	Name          string `json:"name"`
	Company       string `json:"company"`
	LastUpdateGmt int    `json:"lastupdate_gmt"`
	Source        string `json:"source"`
}

type DelayTrains []DelayTrain

// GetJSON APIを叩いてJSONを取得する
func GetJSON() ([]byte, error) {
	//FIXME 毎回APIを叩くと遅いので、最終レスポンスから１時間経過していない場合は処理をskipさせる
	response, err := http.Get("https://tetsudo.rti-giken.jp/free/delay.json")

	defer response.Body.Close()

	jsonBytes, err := ioutil.ReadAll(response.Body)

	return jsonBytes, err
}

// ParseJSONtoDelayTrain 取得したJSONを構造体にパースする
func ParseJSONtoDelayTrain(jsonBytes []byte) (DelayTrains, error) {
	var delayTrains DelayTrains
	if err := json.Unmarshal(jsonBytes, &delayTrains); err != nil {
		return delayTrains, err
	}

	return delayTrains, nil
}
