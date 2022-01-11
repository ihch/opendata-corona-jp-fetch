package root

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

/*
 * 全国の感染者数 オープンデータからデータを取得する関数
 */
func FetchCovid19JapanAll() (covidItemList APICovidItemList, err error) {
  response, err := http.Get(Covid19JapanAllDataUrl)
  if err != nil {
    log.Fatalf("Error: failed fetch to %s. %v\n", Covid19JapanAllDataUrl, err)
    return covidItemList, err
  }

  json.NewDecoder(response.Body).Decode(&covidItemList)

  return covidItemList, nil
}

func TransformToCoronaItem(data *APICovidItem) CovidItem {
  totalPatients, err := strconv.Atoi(data.Npatients)
  if err != nil {
    totalPatients = 0
  }

  return CovidItem{
    Date: data.Date,
    Prefecture: data.NameJp,
    Patients: 0,
    TotalPatients: totalPatients,
  }
}

func TransformToCoronaItemWithBeforePatients(data *APICovidItem, beforePatients int) CovidItem {
  totalPatients, err := strconv.Atoi(data.Npatients)
  if err != nil {
    totalPatients = 0
  }

  return CovidItem{
    Date: data.Date,
    Prefecture: data.NameJp,
    Patients: totalPatients - beforePatients,
    TotalPatients: totalPatients,
  }
}
