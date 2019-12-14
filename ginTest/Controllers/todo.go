package Controllers

import (
    "net/http"
    "../Models"
    "github.com/gin-gonic/gin"
    "fmt"
)

// Index
func Index(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{
        "title": "YEET",
    })
    
}

// list all todos
func GetTodos(c *gin.Context) {
    var todo []Models.Todo
    err := Models.GetAllTodos(&todo)
    fmt.Println(err)
    if err != nil {c.AbortWithStatus(http.StatusNotFound)} else {
        c.JSON(http.StatusOK, todo)
    }
}

// Create Todo
func CreateTodo(c *gin.Context) {
    var todo Models.Todo
    c.BindJSON(&todo)
    err := Models.NewTodo(&todo)
    if err != nil {
        c.AbortWithStatus(http.StatusNotFound)
    } else {
        c.JSON(http.StatusOK, todo)
    }
}

// Get specific todo
func GetTodo(c *gin.Context) {
    id := c.Params.ByName("id")
    var todo Models.Todo
    err := Models.GetTodo(&todo, id)
    if err != nil {
        c.AbortWithStatus(http.StatusNotFound)
    } else {
        c.JSON(http.StatusOK, todo)
    }
}

// Update todo
func UpdateTodo(c *gin.Context) {
    var todo Models.Todo
    id := c.Params.ByName("id")
    err := Models.GetTodo(&todo, id)
    if err != nil {
        c.JSON(http.StatusNotFound, todo)
    }
    c.BindJSON(&todo)
    err = Models.UpdateTodo(&todo, id)
    if err != nil {
        c.AbortWithStatus(http.StatusNotFound)
    } else {
        c.JSON(http.StatusOK, todo)
    }
}


// Del todo
func DelTodo(c *gin.Context) {
    var todo Models.Todo
    id := c.Params.ByName("id")
    err := Models.DelTodo(&todo, id)
    if err != nil {
        c.AbortWithStatus(http.StatusNotFound)
    } else {
        c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
    }
}

// Del everything
func DelAll(c *gin.Context) {
    var todo []Models.Todo
    err := Models.DelAll(&todo)
    if err != nil {
        c.AbortWithStatus(http.StatusNotFound)
    } else {
        c.JSON(http.StatusOK, err)
    }
}

