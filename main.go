package main

import (
	"github.com/joshiaj7/simple-user-app/internal/config"
	"github.com/joshiaj7/simple-user-app/internal/controller"

	"fmt"
	"net/http"
)

func main() {
	err := config.SetupDB()
	if err != nil {
		fmt.Println("Failed to set up DB, error : ", err)
		return
	}

	controller.SetupRoute()

	fmt.Println("listening to port 8080")
	http.ListenAndServe(":8080", nil)
}
