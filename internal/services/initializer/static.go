package initializer

import "github.com/lowl11/boost/pkg/system/config"

func Run() {
	if config.IsProduction() {
		return
	}

	// create .gitignore file if not exists
	initGitignore()

	// create profiles folder if not exists
	initProfiles()

	// create controllers folder if not exists
	initControllers()

	// create services folder if not exists
	initServices()

	// create cmd folder if not exists
	initCMD()
}
