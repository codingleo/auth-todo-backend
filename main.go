package main

import (
	"github.com/codingleo/auth-todo-backend/api"
	"github.com/codingleo/auth-todo-backend/database"
)

func main() {
	database.NewConnection("test.db").Connect()
	s := api.NewAPIServer(":3000")
	s.Start()
}
