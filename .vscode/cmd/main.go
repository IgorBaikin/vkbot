package main

import (
	"log"
	"vkbot/client/handler"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	h := handler.InitHandler()
	err := h.Start()
	log.Println(err)
}
