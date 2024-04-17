package Router

import (
	"TodoApp/DB"
	"fmt"

	"TodoApp/InputValidation"
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
func RouterHander(e *echo.Echo){
	e.GET("/",greating)
	e.POST("/todo",createTodo,InputValidation.CreateTodoMiddleware)
	e.PUT("/completed/:id",completedTodo)
	e.GET("/todo",getTodo)
}
func greating (c echo.Context) error {
	return c.JSON(http.StatusOK,"server started ")
}
func createTodo(c echo.Context) error {
    title:=c.Get("title").(string)
	description:=c.Get("description").(string)

    // Create a new Todo with the retrieved values
    todo := &DB.Todo{
        Title:       title,
        Description: description,
        Completed:   false,
    }

    collection := DB.Client.Database("TodoApp").Collection("Todos")
    result, err := collection.InsertOne(context.TODO(), todo)
    if err != nil {
        log.Println(err)
        return c.JSON(http.StatusInternalServerError, "internal server error in putting document in mongo")
    }
fmt.Println(result)
    return c.JSON(http.StatusOK, "todo has been created")
}


func getTodo(c echo.Context)error{
	collection:=DB.Client.Database("TodoApp").Collection("Todos")
	cursor,err:= collection.Find(context.TODO(),bson.D{})
	if err!=nil{
		log.Println(err)
		return c.JSON(http.StatusInternalServerError,"error while creating cursor")

	}
	defer cursor.Close(context.TODO())
	var todos []DB.Todo
	for cursor.Next(context.TODO()){
		var todo DB.Todo
		if err:=cursor.Decode(&todo);err!=nil{
			log.Println(err)
			return c.JSON(http.StatusInternalServerError,"not able to decode")

		}
		todos=append(todos, todo)
	}
	return c.JSON(http.StatusOK,todos)
}
func completedTodo(c echo.Context)error{
	id:=c.Param("id")
	objId,err:= primitive.ObjectIDFromHex(id)
	if err!=nil{
		log.Println(err)
		return c.JSON(http.StatusBadRequest,"invalid id")
	}
	collection:= DB.Client.Database("TodoApp").Collection("Todos")
	fliter:= bson.M{"_id":objId}
	update:=bson.M{"$set":bson.M{"completed":true}}
	_,err=collection.UpdateOne(context.TODO(),fliter,update)
	if err!=nil{
		log.Println(err)
		return c.JSON(http.StatusInternalServerError,"error updating")
	}
	return c.JSON(http.StatusOK,"todo completed")
}