package codeforces

import (
	"encoding/json"
	"io"
	"net/http"
)

const BaseURI string = "https://codeforces.com/api/"

const ContestStandingURI string = "contest.standings"

type ContestStandingResponse struct {
	Status string  `json:"status"`
	Result *Result `json:"result"`
}

type Result struct {
	Contest  *Contest   `json:"contest"`
	Problems []*Problem `json:"problems"`
}

type Contest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Problem struct {
	ContestID int    `json:"contestId"`
	Index     string `json:"index"`
	Name      string `json:"name"`
}

// get required data from codeforces api
//
// @returns Name of Contest, Number of Problems, Error
func CodeforcesStandings(contestId string) (string, int, error) {

	resp, err := http.Get(BaseURI + ContestStandingURI + "?contestId=" + contestId)
	if err != nil {
		return "", 0, err
	}

	respData, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", 0, err
	}

	response := &ContestStandingResponse{}
	err = json.Unmarshal(respData, response)
	if err != nil {
		return "", 0, err
	}

	return response.Result.Contest.Name, len(response.Result.Problems), nil
}
