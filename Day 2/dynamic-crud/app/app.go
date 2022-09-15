package app

import "alterra-agmc-dynamic-crud/config"

type Application struct {
	Config *config.Config
}

func Init() *Application {
	application := &Application{
		Config: config.Init(),
	}

	return application
}
