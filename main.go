package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "os"
  "encoding/json"
  "github.com/gorilla/mux"
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


type ExternalDocs struct {
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
  Name            string                          `json:"name"`
  Properties      []map[string]map[string]string  `json:"properties"`
}


type Body struct {
  Name            string                  `json:"name"`
  Content         []map[string]*Schema    `json:"content"`
}

//need to implement oauth2 and bearer token instead of qs key
type Security struct {
  Name            string    `json:"name"`
  Type            string    `json:"type"`
  Scopes          []string  `json:"scopes"`
  Key             string    `json:"key"`
}

type Components struct {
  Schemas         []*Schema   `json:"schemas"`
  RequestBodies   []*Body     `json:"requestBodies"`
  SecuritySchemes []*Security `json:"securitySchemes"`
}

type Swagger struct {
  OpenAPI         string        `json:"openapi"`
  Info            *Info         `json:"info"`
  ExternalDocs    *ExternalDocs `json:"externaldocs"`
  Servers         []*Server     `json:"servers"`
  Tags            []*Tag        `json:"tags"`
  Paths           []*Path       `json:"paths"`
  Components      *Components   `json:"components"`

}

//simply panics on error
func check(e error) {
  if e != nil {
    panic(e)
  }
}


//this should be pretty obvious
func readSpec(path string) (Swagger) {

  // Open our jsonFile
  jsonFile, err := os.Open("openapi.json")
  // if we os.Open returns an error then handle it
  check(err)
  defer jsonFile.Close()
  byteValue, _ := ioutil.ReadAll(jsonFile)

  var swagger Swagger

  json.Unmarshal([]byte(byteValue), &swagger)

  //
  // for item := range spec["servers"] {
  //   var server Server
  //   check(json.Unmarshal([]byte(item), &server))
  //   servers = append(servers, server)
  // }
  //
  // for item := range spec["tags"] {
  //   var tag Tag
  //   check(json.Unmarshal([]byte(item), &tag))
  //   tags = append(tags, tag)
  // }
  //
  // for item := range spec["paths"] {
  //   var path Path
  //   check(json.Unmarshal([]byte(item), &path))
  //   paths = append(paths, path)
  // }
  //
  // // var ret_spec map[string][]interface{}
  //
  return swagger
}

//returns the request headers of the client
func testConnection(w http.ResponseWriter, req *http.Request) {
  for name, headers := range req.Header {
      for _, h := range headers {
          fmt.Fprintf(w, "%v: %v\n", name, h)
      }
  }
}

func serverInfo(i Info, w http.ResponseWriter, req *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(i)
}

//calls main worker
func main() {
  swagger := readSpec("openapi.json")
  fmt.Println(swagger)
  r := mux.NewRouter()
  r.HandleFunc("/api/{label}/{id}", testConnection)

}
