package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

func main() {
    //Defining request and error checking
    req, err := http.NewRequest("GET", "https://swapi.co/api/people/1/", nil)
    if err != nil {
        fmt.Println(err)
    }

    //Creating client to send http request
    client := &http.Client{}

    //Using the client to send the request
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
    }

    // Close http request after function has finished
    defer resp.Body.Close()

    //Using ioutil to get out the body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }

    //Creating a map to save the json
    var response map[string]string
    json.Unmarshal(body, &response)

    fmt.Println("Response Status:", resp.Status, "\n")
    fmt.Println("Name:", response["name"])
    fmt.Println("Birth Year:", response["birth_year"])
}

