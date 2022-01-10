package main

import (
  "fmt"

  . "github.com/ihch/opendata-corona-jp-fetch"
)

func main() {
  fmt.Println("Hello, World!")

  covidItemList, err := FetchCovid19JapanAll()
  if err != nil {
    panic(err)
  }

  fmt.Println(covidItemList)
}
