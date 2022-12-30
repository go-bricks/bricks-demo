package bricks

import (
	"context"

	workshop "github.com/go-bricks/bricks-demo/workshop/api"
	"github.com/go-bricks/bricks-demo/workshop/app/controllers"
	"github.com/go-bricks/bricks-demo/workshop/app/data"
	"github.com/go-bricks/bricks-demo/workshop/app/services"
	"github.com/go-bricks/bricks-demo/workshop/app/validations"
	serverInt "github.com/go-bricks/bricks/interfaces/http/server"
	"github.com/go-bricks/bricks/providers/groups"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type workshopServiceDeps struct {
	fx.In

	// API Implementations
	Workshop workshop.WorkshopServer
}

func WorkshopAPIsAndOtherDependenciesFxOption() fx.Option {
	return fx.Options(
		// GRPC Service APIs registration
		fx.Provide(fx.Annotated{
			Group:  groups.GRPCServerAPIs,
			Target: workshopGRPCServiceAPIs,
		}),
		// GRPC Gateway Generated Handlers registration
		fx.Provide(fx.Annotated{
			Group:  groups.GRPCGatewayGeneratedHandlers + ",flatten", // "flatten" does this [][]serverInt.GRPCGatewayGeneratedHandlers -> []serverInt.GRPCGatewayGeneratedHandlers
			Target: workshopGRPCGatewayHandlers,
		}),
		// All other tutorial dependencies
		workshopDependencies(),
	)
}

func workshopGRPCServiceAPIs(deps workshopServiceDeps) serverInt.GRPCServerAPI {
	return func(srv *grpc.Server) {
		workshop.RegisterWorkshopServer(srv, deps.Workshop)
		// Any additional gRPC Implementations should be called here
	}
}

func workshopGRPCGatewayHandlers() []serverInt.GRPCGatewayGeneratedHandlers {
	return []serverInt.GRPCGatewayGeneratedHandlers{
		// Register workshop REST API
		func(mux *runtime.ServeMux, endpoint string) error {
			return workshop.RegisterWorkshopHandlerFromEndpoint(context.Background(), mux, endpoint, []grpc.DialOption{grpc.WithInsecure()})
		},
		// Any additional gRPC gateway registrations should be called here
	}
}

func workshopDependencies() fx.Option {
	return fx.Provide(
		services.CreateWorkshopService,
		controllers.CreateWorkshopController,
		data.CreateCarDB,
		validations.CreateWorkshopValidations,
	)
}
