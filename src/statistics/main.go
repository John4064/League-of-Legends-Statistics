package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func processUserName(name string) summonerDTO {

	///lol/summoner/v4/summoners/by-name/
	url := fmt.Sprintf("https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s/?api_key=%s", name, apiKey)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("KeyValue", apiKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
	var test summonerDTO
	fmt.Println(3)
	fmt.Println(test.name)
	err := json.Unmarshal(body, &test)
	if err != nil {
		fmt.Println("test")
	} else {
		fmt.Println(test)
	}
	//Todo is map to a summonerDTO
	return test
}

func main() {
	temp := processUserName("CarnageOutlaw")
	fmt.Println(temp.name)
	fmt.Println("Done")
	//user := summonerDS.summonerDTO();
	//Order of api call, first get puid frm name, then match history and then display said matches

	//{
	//	"id": "lAU1O3_lvZCj34BMVvaEgpVPwVl_T-LWy2lJCdGpCL6jNO0",
	//	"accountId": "lqI1skui5x1IRr4KpOM-BeIAuPifhBCVWJJI6HWZ0V6MEtA",
	//	"puuid": "PRUJY8y4lZla4Bi92B3xgvPtMSOd_rAbJ0REljcvatwga8WkWNxTvAh4JqzvmYLC2bU7hsl-PJoIoQ",
	//	"name": "CarnageOutlaw",
	//	"profileIconId": 1374,
	//	"revisionDate": 1658863073228,
	//	"summonerLevel": 159
	//}

	//Match data
	//	"NA1_4386498033",
	//"NA1_4386273759",
	//	"NA1_4386086372",
	//	"NA1_4385501336",
	//	"NA1_4385355885",
	//	"NA1_4385013287",
	//	"NA1_4384962970",
	//	"NA1_4383938973",
	//	"NA1_4383105676",
	//	"NA1_4381944482",
	//	"NA1_4380950497",
	//	"NA1_4380870072",
	//	"NA1_4378635337",
	//	"NA1_4378550424",
	//	"NA1_4378017427",
	//	"NA1_4377736600",
	//	"NA1_4377741954",
	//	"NA1_4377659194",
	//	"NA1_4377613399",
	//	"NA1_4377602088"]

}
