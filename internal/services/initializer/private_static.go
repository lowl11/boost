package initializer

import (
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boost/pkg/io/file"
	"github.com/lowl11/boost/pkg/io/folder"
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

func initControllers() {
	if folder.Exist("controllers") {
		return
	}

	createFolder("controllers")
	createFolder("controllers/hello_controller")
	createFile("controllers/hello_controller/controller.go", []byte(`package hello_controller

import (
	"github.com/lowl11/boost"
	"github.com/lowl11/boost/pkg/web/base/controller"
)

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
