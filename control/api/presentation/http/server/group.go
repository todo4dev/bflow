// presentation/http/server/group_definition.go
package server

type Group struct {
	Path    string
	Factory func(*Grouper)
}

func NewGroup(path string, factory func(*Grouper)) Group {
	return Group{Path: path, Factory: factory}
}
