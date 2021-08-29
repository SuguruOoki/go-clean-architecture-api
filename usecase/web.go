package main

import (
	"fmt"
	"todo/infra/handler"
	"todo/injector"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("sever start")
	todoHandler := injector.InjectTodoHandler()
	e := echo.New()
	handler.InitRouting(e, todoHandler)
	e.Logger.Fatal(e.Start(":58000"))
}
