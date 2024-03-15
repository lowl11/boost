package object

type Object struct {
	Name         string
	ObjectCount  int
	Path         string
	RelativePath string
	IsFolder     bool

	Children []Object
}
