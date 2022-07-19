package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := fmt.Sprintf("https://na1.api.riotgames.com/lol/platform/v3/champion-rotations?api_key=%s", apiKey)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("KeyValue", apiKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
