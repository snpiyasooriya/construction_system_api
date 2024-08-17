package main

import (
	"github.com/snpiyasooriya/construction_design_api/config"
	"github.com/snpiyasooriya/construction_design_api/internal/infrastructre/database"
)

func main() {
	conf := config.GetConfig()
	_ = database.NewPostgres(conf)
}
