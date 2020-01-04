package cmd

import (
	"github.com/prongbang/gomtl/pkg/api"
	"github.com/prongbang/gomtl/pkg/database"
)

func Run() {
	conn := database.NewDbConnection()
	apis := api.CreateAPI(conn)
	apis.Register()
}
