package main

import (
  "fmt"

  . "opendata-corona-jp-batch"
)

func main() {
  fmt.Println("Hello, World!")

  covidItemList, err := FetchCovid19JapanAll()
  if err != nil {
    panic(err)
  }

  fmt.Println(covidItemList)
}
