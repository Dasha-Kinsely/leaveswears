package main

import (
	"log"
	//"github.com/Dasha-Kinsely/leaveswears/loggers"
	"github.com/Dasha-Kinsely/leaveswears/databases"
	"github.com/Dasha-Kinsely/leaveswears/routers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var r *gin.Engine

func init() {
	/* Create SimpleHTTP and Error Loggers
	eS := loggers.SetUp("s")
	if eS != nil {
		log.Printf(">>> Simple Loggers creation unsuccessful...\n")
	}
	eE := loggers.SetUp("e")
	if eE != nil {
		log.Printf(">>> Error Loggers creation unsuccessful...\n")
	}*/
	// Initialize Database
	databases.InitDB()
	db := databases.GetDB()
	databases.FirstMigration(db)
}

func main() {
	// gin Defaults
	r = gin.Default()
	routers.initializeRoutes(r)
	
	r.Run()
}
