package covidApi

import (
	"bansosman/bussiness/covid"
	"encoding/json"
	"net/http"
)

type CovidApi struct {
	httpClient http.Client
}

func NewCovidApi() covid.Repository {
	return &CovidApi{
		httpClient: http.Client{},
	}
}

func (cov *CovidApi) GetCaseCovid() (covid.Domain, error) {
	req, _ := http.NewRequest("GET", "https://api.kawalcorona.com/indonesia/provinsi/", nil)
	req.Header.Set("User-Agent", "api.kawalcorona.com/#go-v1.17")
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
	return data.toDomain(), nil
}
