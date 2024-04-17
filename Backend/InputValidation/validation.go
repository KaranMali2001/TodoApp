package InputValidation

import (
	"TodoApp/DB"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func CreateTodoMiddleware(next echo.HandlerFunc)echo.HandlerFunc{
	return func(c echo.Context) error {
		todo:=new(DB.Todo)
		if err:=c.Bind(todo);err!=nil{
			return err;
		}
		validate:=validator.New()
		if err:= validate.StructExcept(todo,"Id");err!=nil{
			return c.JSON(http.StatusBadRequest,err.Error())
		}
		fmt.Println("validation sucessful")
		c.Set("title",todo.Title)
		c.Set("description",todo.Description)
		return next(c)
	}
}