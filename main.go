package main

import (
  "os"
  "log"
  "strings"
  "strconv"
  "io/ioutil"
  "encoding/json"
  "net/http"
)

type Data []Datum

type Datum struct {
  ID int `json:"id"`
  Name string `json:"name"`
}

type Response struct {
  Code int `json:"code"`
  Data Data `json:"data"`
}

type Error struct {
  Code int `json:"code"`
  Message string `json:"message"`
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

  for i := range data {
    log.Println("ID: ", data[i].ID, "Name: ", data[i].Name)
  }

  return data
}

func dataHandler(data Data) func(w http.ResponseWriter, r *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
      http.Error(w, "Invalid URL path, use root index", http.StatusNotFound)
      return
    }

    if r.Method != "GET" {
      http.Error(w, "Method is not supported, use GET", http.StatusNotFound)
      return
    }

    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.Header().Set("X-Content-Type-Options", "nosniff")

    ids, exist := r.URL.Query()["id"]

    // Case 1: Request without parameter id
    if !exist {
      jsonData, err := json.Marshal(Response{
        Code: http.StatusOK,
        Data: data,
      })

      if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
      }

      w.WriteHeader(http.StatusOK)
      w.Write(jsonData)
      return
    }

    idList := strings.Split(ids[0], ",")

    // Case 2 and 3: Request with single or multiple ids
    var returnData Data

    for i := range idList {
      idx, err := strconv.Atoi(idList[i])

      // Case 4: Request with invalid ID
      if err != nil {
        jsonError, err := json.Marshal(Error{
          Code: http.StatusBadRequest,
          Message: "Invalid or empty ID: \"" + idList[i] + "\"",
        })

        if err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
        }

        w.WriteHeader(http.StatusBadRequest)
        w.Write(jsonError)
        return
      }

      returnData = append(returnData, data[idx - 1])
    }

    jsonData, err := json.Marshal(Response{
      Code: http.StatusOK,
      Data: returnData,
    })

    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(jsonData)
    return
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
