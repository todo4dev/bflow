// presentation/http/server/group_definition.go
package server

type GroupType struct {
	Path    string
	Factory func(*Grouper)
}

func Group(path string, factory func(*Grouper)) GroupType {
	return GroupType{Path: path, Factory: factory}
}
