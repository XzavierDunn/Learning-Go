package main

import (
    "fmt"
    "./Config"
    "./Models"
    "./Routes"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

var err error

func main() {
    // Create db connect
    Config.DB, err = gorm.Open("postgres", Config.DbURL(
        Config.BuildDBConfig()))

    if err != nil {
        fmt.Println("status: ", err)
    }

    defer Config.DB.Close()

    // Run migrations
    Config.DB.AutoMigrate(&Models.Todo{})

    //setup Routes
    r := Routes.SetupRouter()
    r.LoadHTMLGlob("templates/*")

    // run
    r.Run()
}
