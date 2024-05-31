package main

import (
	"github.com/gabotachak/ambrosia-zeus-api/cmd/api/router"
	"github.com/gabotachak/ambrosia-zeus-api/cmd/api/util"
)

func main() {
	db := util.InitializeDatabase()
	if db == nil {
		panic("DB connection failed")
	}

	router.SetupEndpoints(db)
}
