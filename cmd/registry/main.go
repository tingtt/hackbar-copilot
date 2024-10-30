package main

import (
	"context"
	"fmt"
	"hackbar-copilot/internal/infrastructure/api/http"
	"hackbar-copilot/internal/infrastructure/datasource/filesystem"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph"
	"hackbar-copilot/internal/usecase/recipes"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/pflag"
)

func main() {
	err := run()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func run() error {
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt, os.Kill)

	option := getCLIOption()
	deps, err := loadDependencies(option.DataDirPath)
	if err != nil {
		return err
	}

	s := http.NewServer(
		fmt.Sprintf("%s:%s", option.Host, option.Port),
		deps.Usecase.GraphQL,
	)
	go s.ListenAndServe()

	<-ctx.Done()
	err = s.Shutdown(ctx)
	if err != nil {
		return err
	}
	err = deps.Datasources.Recipes.SavePersistently()
	if err != nil {
		return err
	}
	return nil
}

type option struct {
	Host        string
	Port        string
	DataDirPath string
}

func getCLIOption() option {
	host := pflag.String("host", "127.0.0.1", "")
	port := pflag.StringP("port", "p", "8080", "")
	dataDirPath := pflag.StringP("data", "d", "/var/lib/hackbar-copilot", "")
	pflag.Parse()

	return option{
		*host,
		*port,
		*dataDirPath,
	}
}

type dependencies struct {
	Datasources depsDatasources
	Usecase     depsUsecase
}
type depsDatasources struct {
	Recipes *filesystem.Filesystem
}
type depsUsecase struct {
	GraphQL graph.Dependencies
}

func loadDependencies(dataDirPath string) (dependencies, error) {
	recipeRepository, err := filesystem.NewRepository(dataDirPath)
	if err != nil {
		return dependencies{}, err
	}
	recipes, err := recipes.NewService(recipeRepository)
	if err != nil {
		return dependencies{}, err
	}

	return dependencies{
		Datasources: depsDatasources{
			Recipes: recipeRepository,
		},
		Usecase: depsUsecase{
			GraphQL: graph.Dependencies{
				Orders:  nil,
				Recipes: recipes,
			},
		},
	}, nil
}
