package autospec

import (
	"errors"
	"fmt"

	"github.com/Jeielsantosdev/autospec/internal/commands"
	"github.com/Jeielsantosdev/autospec/internal/config"
	"github.com/Jeielsantosdev/autospec/internal/openapi"
	"github.com/Jeielsantosdev/autospec/internal/version"
)

type Config struct {
	Name string
}

type App struct {
	config config.Config
}

func New() *App {
	return &App{
		config: config.Default(),
	}
}

func (a *App) Spec() *openapi.Spec {
	spec := openapi.NewSpec(a.config.Title, a.config.Version, a.config.Description)
	spec.AddServer(a.config.ServerURL, "Local development server")
	spec.AddSecurityScheme("bearerAuth", openapi.SecurityScheme{
		Type:         "http",
		Scheme:       "bearer",
		BearerFormat: "JWT",
		Description:  "Bearer token authentication",
	})
	return spec
}

func (a *App) Run(args []string) error {
	if len(args) == 0 {
		fmt.Printf("%s %s\n", a.config.Name, version.String())
		fmt.Println("Use: autospec [version|dev|watch]")
		return nil
	}

	switch args[0] {
	case "dev":
		return commands.Dev()
	case "watch":
		return commands.Watch()
	case "version", "--version", "-v":
		fmt.Println(version.String())
		return nil
	default:
		return errors.New("comando desconhecido: " + args[0])
	}
}
