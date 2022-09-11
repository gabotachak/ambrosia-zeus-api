package main

import (
	"ambrosia-zeus-api/cmd/api/router"
	"ambrosia-zeus-api/cmd/api/util"
)

func main() {
	db := util.InitializeDatabase()
	if db == nil {
		panic("DB connection failed")
	}

	router.SetupEndpoints(db)
}
