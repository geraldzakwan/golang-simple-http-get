package main

import (
  "os"
  "fmt"
  "io/ioutil"
  "encoding/json"
)

type Data struct {
  Data []Datum `json:"data"`
}

type Datum struct {
  ID string `json:"id"`
  Name string `json:"name"`
}

func main() {
    jsonFile, err := os.Open("data.json");

    if err != nil {
        fmt.Println(err)
    }

    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)
    var result map[string]interface{}

    json.Unmarshal([]byte(byteValue), &result)
    fmt.Println(result["data"])
}
