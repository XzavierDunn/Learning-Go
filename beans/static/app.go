package main

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
    "net/http"
    "html/template"
    "strings"
    "bufio"
    "os"
)

func getData() (string, string) {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter a num to find a person: ")
    text, _ := reader.ReadString('\n')
    text = strings.Replace(text, "\n", "", -1)

    req, err := http.NewRequest("GET", "https://swapi.co/api/people/" + text + "/", nil)
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
    fmt.Println("Listening 127.0.0.1:5000/yes")
    http.ListenAndServe(":5000", nil)
}

type yesPage struct {
    Name string
    Birth string
    Other []string
}

func yesHandler(w http.ResponseWriter, r *http.Request) {
    name, birth := getData()
    p := yesPage{Name: name, Birth: birth, Other: []string{"Beans", "yeet"}}
    t, err := template.ParseFiles("yes.html")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(t.Execute(w, p))
}

