package main

import (
	"fmt"
	"net/http"

	router "github.com/kazuki0924/go-chi/infrastructure/router"
)

var (
	httpRouter router.Router = router.NewChiRouter()
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
