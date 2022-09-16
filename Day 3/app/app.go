package app

import "alterra-agmc-day3/config"

type Application struct {
	Config *config.Config
}

func Init() *Application {
	application := &Application{
		Config: config.Init(),
	}

	return application
}
