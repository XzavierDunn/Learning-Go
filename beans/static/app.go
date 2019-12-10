package main

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
    "net/http"
    "html/template"
)

func getData() (string, string) {
    req, err := http.NewRequest("GET", "https://swapi.co/api/people/1/", nil)
    if err != nil {
        fmt.Println(err)
    }
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }

    var response map[string]string
    json.Unmarshal(body, &response)

    return response["name"], response["birth_year"]

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello</h1>")
}

func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/yes", yesHandler)
    http.ListenAndServe(":5000", nil)
}

type yesPage struct {
    Name string
    Birth string
}

func yesHandler(w http.ResponseWriter, r *http.Request) {
    name, birth := getData()
    p := yesPage{Name: name, Birth: birth}
    t, err := template.ParseFiles("yes.html")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(t.Execute(w, p))
}
