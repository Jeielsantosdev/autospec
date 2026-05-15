package config

type Config struct {
	Name        string
	Version     string
	Title       string
	Description string
	ServerURL   string
	DocsPath    string
}

func Default() Config {
	return Config{
		Name:        "autospec",
		Version:     "dev",
		Title:       "autospec",
		Description: "Automatic OpenAPI generation for Go",
		ServerURL:   "http://localhost:8080",
		DocsPath:    "/docs",
	}
}
