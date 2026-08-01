package main

import (
	"unsia.ac.id/akademic_be/internal/component"
	"unsia.ac.id/akademic_be/internal/config"
)

func main() {
	viper := config.NewViperConfig()
	cnf := config.NewModel(viper)
	logger := config.NewLogger(cnf)
	validate := config.NewValidateConfig()
	db := config.NewDatabase(cnf, logger)
	config.RunAutoSeeder(db, logger)
	app := config.NewFiber(cnf)
	minio := config.NewMinio(cnf)
	redisClient := config.NewCached(cnf)
	config.InitSentry(cnf)

	component.Bootstrap(&component.Component{
		App:         app,
		DB:          db,
		Log:         logger,
		Validate:    validate,
		Config:      cnf,
		Minio:       minio,
		RedisClient: redisClient,
	})

	if err := app.Listen(cnf.Server.Host + ":" + cnf.Server.Port); err != nil {
		logger.Fatal(err.Error())
	}
}
