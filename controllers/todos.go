package controllers

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "gopkg.in/mgo.v2" 
    "gopkg.in/mgo.v2/bson"    
    . "github.com/blitzkriegcoding/todo-gin-mgo/models"
    . "github.com/blitzkriegcoding/todo-gin-mgo/config"
)

var todosCollection *mgo.Collection
var session *mgo.Session

func init() {
    session, todosCollection = InitDB("test-todo", "todo")
}

func CreateTodo(context *gin.Context) {
    title := context.PostForm("Title")
    completed, _ := strconv.ParseBool(context.PostForm("Completed"))
    var todo = Todo{bson.NewObjectId(), title, completed}
    fmt.Println("" + todo.Title + " completed: " + strconv.FormatBool(todo.Completed))
    err := todosCollection.Insert(&todo)
    if err != nil {
        log.Fatal(err)
    }

    context.JSON(http.StatusCreated, gin.H{
        "status":  http.StatusCreated,
        "message": "todo item created successfully",
    })
}

func FetchAllTodo(context *gin.Context) {
    var todos []Todo
    err := todosCollection.Find(nil).All(&todos)
    if err != nil {
        log.Fatal(err)
    }

    if len(todos) <= 0 {
        context.JSON(http.StatusNotFound, gin.H{
            "status":  http.StatusNotFound,
            "message": "no todo found",
        })
        return
    }

    context.JSON(http.StatusOK, gin.H{
        "status": http.StatusOK,
        "data":   todos,
    })
}

func FetchSingleTodo(context *gin.Context) {
    todo := Todo{}
    id := bson.ObjectIdHex(context.Param("id"))
    err := todosCollection.FindId(id).One(&todo)

    if err != nil || todo == (Todo{}) {
        fmt.Println("Error: " + err.Error())
        context.JSON(http.StatusNotFound, gin.H{
            "status":  http.StatusNotFound,
            "message": "todo not found",
        })
        return
    }

    context.JSON(http.StatusOK, gin.H{
        "status": http.StatusOK,
        "data":   todo,
    })
}

func UpdateTodo(context *gin.Context) {
    id := bson.ObjectIdHex(context.Param("id"))
    title := context.PostForm("title")
    completed, _ := strconv.ParseBool(context.PostForm("completed"))

    err := todosCollection.UpdateId(id, bson.M{"title": title, "completed": completed})

    fmt.Printf("completed: %t\n\n", completed)

    if err != nil {
        fmt.Println("Error: " + err.Error())
        context.JSON(http.StatusNotFound, gin.H{
            "status":  http.StatusNotFound,
            "message": "todo not found",
        })
        return
    }

    context.JSON(http.StatusOK, gin.H{
        "status":  http.StatusOK,
        "message": "Todo updated successfully!",
    })
}

func DeleteTodo(context *gin.Context) {
    id := bson.ObjectIdHex(context.Param("id"))

    fmt.Printf("id: %v", id)

    err := todosCollection.RemoveId(id)

    if err != nil {
        fmt.Println("Error: " + err.Error())
        context.JSON(http.StatusNotFound, gin.H{
            "status":  http.StatusNotFound,
            "message": "todo not found",
        })
        return
    }

    context.JSON(http.StatusOK, gin.H{
        "status":  http.StatusOK,
        "message": "Todo deleted successfully!",
    })
}