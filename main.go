package main

import (
	"api-rest-gin/models"
	"api-rest-gin/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{"Andresa", "12345678912", "12345678"},
		{"Denis", "12345678912", "12345678"},
	}
	routes.HandleRequests()
}
