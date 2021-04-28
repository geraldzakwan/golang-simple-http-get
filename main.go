package main

import (
  "os"
  "log"
  "strings"
  "io/ioutil"
  "encoding/json"
  "net/http"
)

type Data []Datum

type Datum struct {
  ID int `json:"id"`
  Name string `json:"name"`
}

func loadData() Data {
  jsonFile, err := os.Open("data.json")
  if err != nil {
    log.Println(err)
  }

  byteValue, _ := ioutil.ReadAll(jsonFile)
  jsonFile.Close()

  log.Println("Data is succesfully loaded")

  var data Data
  json.Unmarshal([]byte(byteValue), &data)

  // for i := range data.Data {
  //   log.Println(data.Data[i].ID)
  //   log.Println(data.Data[i].Name)
  // }

  return data
}

func dataHandler(data Data) func(w http.ResponseWriter, r *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
      http.Error(w, "404 not found.", http.StatusNotFound)
      return
    }

    if r.Method != "GET" {
      http.Error(w, "Method is not supported.", http.StatusNotFound)
      return
    }

    w.Header().Set("Content-Type", "application/json")

    ids := r.URL.Query().Get("id")

    // Case 1: Request without parameter "id"
    if ids == "" {
      for i := range data {
        log.Println(data[i].ID)
        log.Println(data[i].Name)
      }

      jsonData, err := json.Marshal(data)
      if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
      }

      w.Write(jsonData)
    } else {
      idList := strings.Split(ids, ",")
      log.Println(idList)
    }
  }
}

func main() {
  data := loadData()

  http.HandleFunc("/", dataHandler(data))

  log.Println("Starting server at port 8080")
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal(err)
  }
}
