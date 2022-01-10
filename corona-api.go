package root

import (
  "encoding/json"
  "net/http"
  "log"
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
