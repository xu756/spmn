package main

import (
	"github.com/xu756/spmn/models"
	"github.com/xu756/spmn/router"
)

func main() {
	models.InitMysqlDB()
	router.InitRouter()
}
