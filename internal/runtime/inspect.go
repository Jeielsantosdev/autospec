package runtime

type Route struct {
	Method      string
	Path        string
	Handler     string
	Middlewares []string
}

type Snapshot struct {
	Routes []Route
	Auth   bool
}

func Inspect() Snapshot {
	return Snapshot{
		Routes: make([]Route, 0),
	}
}

func DetectAuth(Middlewares []string) bool {
	for _, middleware := range Middlewares {
		if middleware == "AuthMiddleware" || middleware ==
			"JWTMiddleware" {
			return true
		}
	}

	return false
}
