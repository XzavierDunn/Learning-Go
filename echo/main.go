package main

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    )

func main() {
    // Echo Instance
    e := echo.New()

    //Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Route => handler
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello World!\n")
    })

    e.GET("/json", func(c echo.Context) error {
        return c.JSONBlob(
            http.StatusOK,
            []byte(`{"id": "1", "msg": "Hello"}`),
        )
    })

    e.GET("/html", func(c echo.Context) error {
        return c.HTML(
            http.StatusOK,
            "<h1>Hello</h1>",
        )
        })

    //Start Server
    e.Logger.Fatal(e.Start(":1323"))
}
