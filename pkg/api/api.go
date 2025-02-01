package api

type UserRoute struct {
	Method string
	Path   string
}

func (r UserRoute) GetMethod() string {
	return r.Method
}

func (r UserRoute) FullPath() string {
	return r.Path
}
