package main

import (
  "os"
  "log"
  "io/ioutil"
  "encoding/json"
)

type Data struct {
  Data []Datum `json:"data"`
}

type Datum struct {
  ID int `json:"id"`
  Name string `json:"name"`
}

func main() {
  jsonFile, err := os.Open("data.json")
  if err != nil {
    log.Println(err)
  }

  byteValue, _ := ioutil.ReadAll(jsonFile)
  jsonFile.Close()

  log.Println("Data is succesfully loaded")

  var result Data
  json.Unmarshal([]byte(byteValue), &result)

  for i := range result.Data {
    log.Println(result.Data[i].ID)
    log.Println(result.Data[i].Name)
  }
}
