package main

import (
	"context"
	"errors"
	"fmt"
	"hackbar-copilot/internal/infrastructure/api/http"
	"hackbar-copilot/internal/infrastructure/datasource/filesystem"
	"hackbar-copilot/internal/interface-adapter/handler/graphql"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph"
	"hackbar-copilot/internal/usecase/copilot"
	"hackbar-copilot/internal/usecase/order"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	httpgo "net/http"

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

	option := getCLIOption(os.Args)
	deps, close, err := loadDependencies(option.DataDirPath)
	if err != nil {
		return err
	}
	defer close()

	server := http.NewServer(
		fmt.Sprintf("%s:%s", option.Host, option.Port),
		deps.Usecase.GraphQL,
		graphql.Option{
			JWTSecret: option.JWTSecret,
		},
	)
	slog.Info(fmt.Sprintf("Starting HTTP Server. Listening at %s.", server.Addr))
	err = serveGraceful(ctx, server, deps.Datasources)
	slog.Info("Server closed.")
	return err
}

type server interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

func serveGraceful(ctx context.Context, server server, datasources depsDatasources) (err error) {
	errServeChan := make(chan error, 1)

	go func() {
		err := server.ListenAndServe()
		if !errors.Is(err, httpgo.ErrServerClosed) {
			errServeChan <- err
		} else {
			close(errServeChan)
		}
	}()

	<-ctx.Done()
	errShutdown := server.Shutdown(ctx)
	errServe := <-errServeChan
	errSave := datasources.fs.SavePersistently()
	return errors.Join(errServe, errShutdown, errSave)
}

type option struct {
	Host        string
	Port        string
	DataDirPath string
	JWTSecret   string
}

func getCLIOption(osArgs []string) option {
	flag := pflag.NewFlagSet(osArgs[0], pflag.ExitOnError)
	host := flag.IP("host", net.IPv4(127, 0, 0, 1), "")
	port := flag.StringP("port", "p", "8080", "")
	dataDirPath := flag.StringP("data", "d", "/var/lib/hackbar-copilot", "")
	jwtSecret := flag.String("jwt.secret", "", "JWT secret key")
	// *dataDirPath = "/Users/taku_ting/workspaces/hackbar/hackbar-copilot/data"
	flag.Parse(osArgs[1:])

	return option{
		host.String(),
		*port,
		*dataDirPath,
		*jwtSecret,
	}
}

type dependencies struct {
	Datasources depsDatasources
	Usecase     depsUsecase
}
type depsDatasources struct {
	fs filesystem.Filesystem
}
type depsUsecase struct {
	GraphQL graph.Dependencies
}

func loadDependencies(dataDirPath string) (dependencies, func() /* close func */, error) {
	fs, err := filesystem.NewRepository(dataDirPath)
	if err != nil {
		return dependencies{}, nil, err
	}

	orderRepo, close := fs.Order()

	return dependencies{
		Datasources: depsDatasources{fs},
		Usecase: depsUsecase{
			GraphQL: graph.Dependencies{
				Copilot: copilot.New(copilot.Dependencies{
					Recipe: fs.Recipe(),
					Menu:   fs.Menu(),
					Stock:  fs.Stock(),
				}),
				OrderService: order.New(order.Dependencies{
					Menu:  fs.Menu(),
					Order: orderRepo,
				}),
			},
		},
	}, close, nil
}
