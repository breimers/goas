package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
)


//OpenAPI 3.0.0 Structs


type Info struct {
  Title           string              `json:"title"`
  Version         string              `json:"version"`
  Description     string              `json:"description"`
  TermsOfService  string              `json:"termsOfService"`
  Contact         map[string]string   `json:"contact"`
  License         map[string]string   `json:"license"`
}


type Server struct {
  Url             string    `json:"url"`
  Description     string    `json:"description"`
}


type Tag struct {
  Name            string              `json:"name"`
  Description     string              `json:"description"`
  ExternalDocs    map[string]string   `json:"externalDocs"`
}


type Path struct {
  Route            string                   `json:"route"`
  Method           string                   `json:"method"`
  HandlerFunc      http.HandlerFunc         `json:"handler"`
  Tags             []*Tag                   `json:"tags"`
  Summary          string                   `json:"summary"`
  Responses        map[int]string           `json:"responses"`
  Security         []map[string][]string    `json:"security"`
  RequestBody      *Body                    `json:"responses"`
}


type Schema struct {

}


type Body struct {

}


type Security struct {
  Name            string              `jso`
}


//simply panics on error
func check(e error) {
  if e != nil {
    panic(e)
  }
}


//this should be pretty obvious
func readSpec(path string) {
  dat, err := ioutil.ReadFile(path)
  check(err)

  in := []byte(dat)
  spec := map[string][]string{}

  check(json.Unmarshal(in, &spec))
  fmt.Println(spec)
}

//returns the request headers of the client
func testConnection(w http.ResponseWriter, req *http.Request) {
  for name, headers := range req.Header {
      for _, h := range headers {
          fmt.Fprintf(w, "%v: %v\n", name, h)
      }
  }
}

//http server as func so it can be invoked as goroutine
func httpWorker() {
  http.HandleFunc("/test", testConnection)
  http.ListenAndServe(":8000", nil)
}

//calls main worker
func main() {
  readSpec("openapi.json")
//  httpWorker()

}
