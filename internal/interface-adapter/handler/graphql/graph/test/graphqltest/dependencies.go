package graphqltest

import (
	"hackbar-copilot/internal/infrastructure/datasource/filesystem"
	"hackbar-copilot/internal/interface-adapter/handler/graphql"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph"
	"hackbar-copilot/internal/usecase/barcounter"
	"hackbar-copilot/internal/usecase/cashier"
	"hackbar-copilot/internal/usecase/copilot"
	"hackbar-copilot/internal/usecase/order"
)

const jwtSecret = "testsecret"

func Dependencies(dataDirPath string) (graph.Dependencies, graphql.Option) {
	fs, err := filesystem.NewRepository(dataDirPath)
	if err != nil {
		panic(err)
	}

	return graph.Dependencies{
		Copilot: copilot.New(copilot.Dependencies{
			Gateway: fs.CopilotGateway(),
		}),
		OrderService: order.New(order.Dependencies{
			Gateway: fs.OrderGateway(),
		}),
		BarCounter: barcounter.New(barcounter.Dependencies{
			Gateway: fs.BarCounterGateway(),
		}),
		Cashier: cashier.New(cashier.Dependencies{
			Gateway: fs.CashierGateway(),
		}),
	}, graphql.Option{JWTSecret: jwtSecret}
}
