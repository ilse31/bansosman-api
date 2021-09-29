package covidApi

import (
	"bansosman/bussiness/covid"
	"encoding/json"
	"net/http"
)

type COVIDAPI struct {
	httpClient http.Client
}

func CovidApi() covid.Repository {
	return &COVIDAPI{
		httpClient: http.Client{},
	}
}

func (cov *COVIDAPI) GetAllCaseProv() (covid.Domain, error) {
	req, _ := http.NewRequest("GET", "https://api.kawalcorona.com/indonesia/provinsi/", nil)
	req.Header.Set("User-Agent", "api.kawalcorona.com/#go-1.1.7")
	resp, err := cov.httpClient.Do(req)
	if err != nil {
		return covid.Domain{}, err
	}
	defer resp.Body.Close()

	data := Response{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return covid.Domain{}, err
	}
	return covid.Domain{}, nil
}
