package main

import(
    "fmt"
)

func main() {
    // Define map
    emails := make(map[string]string)
    fart := "Fart"

    // Assign kv
    emails["Xzavier"] = "xd@gmail.com"
    emails["Beans"] = "Beans@gmail.com"
    emails[fart] = "yes@gmail.com"

    fmt.Println(emails)
    fmt.Println(len(emails))
    fmt.Println(emails["Xzavier"])

    // Delete from map
    delete(emails, "Beans")

    fmt.Println("\n", emails)
    fmt.Println(len(emails))
}
