package main

import (
	"fmt"
	"net/http"

	infrastructure "github.com/kazuki0924/go-chi/infrastructure/router"
)

var (
	httpRouter infrastructure.Router = infrastructure.NewChiRouter()
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	httpRouter.GET("/carDetails", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "test")
	})

	httpRouter.SERVE(port)
}
