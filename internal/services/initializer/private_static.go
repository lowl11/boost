package initializer

import (
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boost/pkg/io/file"
	"github.com/lowl11/boost/pkg/io/folder"
	"github.com/lowl11/boost/pkg/system/types"
	"regexp"
	"strings"
)

func initGitignore() {
	createFile(".gitignore", []byte(`.idea
.vs_code
.vscode
.DS_Store
*.exe
*.log
logs
.env
cdn
`))
}

func initProfiles() {
	createFolder("profiles")
	createFile("profiles/config.yml", []byte(`port: ":8080"`))
	createFile("profiles/dev.yml", nil)
	createFile("profiles/production.yml", nil)
}

func initCMD() {
	createFolder("cmd")
	createFolder("cmd/api")

	content, err := file.Read("go.mod")
	if err != nil {
		log.Error(err, "Read go.mod file error")
		return
	}

	appName := regexp.MustCompile("module\\s(.*)").FindAllStringSubmatch(types.ToString(content), -1)[0][1]

	createFile("cmd/api/main.go", []byte(strings.ReplaceAll(`package main

import (
	"github.com/lowl11/boost"
	"github.com/lowl11/boost/pkg/system/config"
	"%APP_NAME%/controllers/hello_controller"
)

func main() {
	app := boost.New()

	di.MapControllers(
		hello_controller.New,
	)

	app.Run(config.Get("port"))
}
`, "%APP_NAME%", appName)))
}

func initControllers() {
	createFolder("controllers")
	createFolder("controllers/hello_controller")
	createFile("controllers/hello_controller/controller.go", []byte(`package hello_controller

import "github.com/lowl11/boost/pkg/web/base/controller"

type Controller struct {
	controller.Base
}

func New() *Controller {
	return &Controller{}
}

func (controller Controller) RegisterEndpoints(router boost.Router) {
	group := router.Group("/base/group/endpoint")

	group.GET("/hello", func(ctx boost.Context) error {
		return controller.Ok(ctx, "Hello world")
	})
}
`))
}

func initServices() {
	createFolder("services")
}

func createFile(path string, body []byte) {
	if !file.Exist(path) {
		if err := file.New(path, body); err != nil {
			log.Error(err, "Create "+path+" file error")
		}
	}
}

func createFolder(path string) {
	if !folder.Exist(path) {
		if err := folder.Create(".", path); err != nil {
			log.Error(err, "Create "+path+" folder error")
		}
	}
}
