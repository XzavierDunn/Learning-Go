package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "fmt"
 )

    type User struct {
        gorm.Model
        Name string
        Age int64
    }

func main() {
    db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user= dbname= password=")
    if err != nil {
        fmt.Println(err)
    }

//    db.CreateTable(&User{})
//    db.DropTable(&User{})
    

    defer db.Close()

//    user := User{Name: "Tall Bitch ass", Age: 900324}
//    db.Create(&user)

    db.Where("Age = 900324").Delete(&User{})
}
