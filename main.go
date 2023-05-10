package main

import (
	"fmt"
	"vkbot/client/handler"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	h := handler.InitHandler()
	err := h.Start()
	fmt.Println(err)
}
