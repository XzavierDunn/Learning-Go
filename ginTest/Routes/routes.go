package Routes

import (
    "github.com/gin-gonic/gin"
    "../Controllers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.GET("/", Controllers.Index)
    v1 := r.Group("/v1")
    {
        v1.GET("/todo", Controllers.GetTodos)
        v1.POST("/todo", Controllers.CreateTodo)
        v1.DELETE("/todo", Controllers.DelAll)
        v1.GET("/todo/:id", Controllers.GetTodo)
        v1.PUT("/todo/:id", Controllers.UpdateTodo)
        v1.DELETE("/todo/:id", Controllers.DelTodo)
    }
    return r
}
