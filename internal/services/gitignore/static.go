package gitignore

import (
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boost/pkg/io/file"
)

func Init() {
	if file.Exist(".gitignore") {
		return
	}

	if err := file.New(".gitignore", []byte(`.idea
.vs_code
.vscode
.DS_Store
*.exe
*.log
logs
.env
cdn
`)); err != nil {
		log.Error(err, "Create .gitignore file error")
	}
}
