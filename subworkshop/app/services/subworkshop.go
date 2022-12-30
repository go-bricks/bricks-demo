package services

import (
	"context"

	subworkshop "github.com/go-bricks/bricks-demo/subworkshop/api"
	"github.com/go-bricks/bricks-demo/subworkshop/app/controllers"
	"github.com/go-bricks/bricks-demo/subworkshop/app/validations"
	"github.com/go-bricks/bricks/interfaces/log"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/fx"
)

type subWorkshopServiceDeps struct {
	fx.In

	Logger      log.Logger
	Controller  controllers.SubWorkshopController
	Validations validations.SubWorkshopValidations
}

type subWorkshopImpl struct {
	deps subWorkshopServiceDeps
	subworkshop.UnimplementedSubWorkshopServer
}

func CreateSubWorkshopService(deps subWorkshopServiceDeps) subworkshop.SubWorkshopServer {
	return &subWorkshopImpl{
		deps: deps,
	}
}

func (s *subWorkshopImpl) PaintCar(ctx context.Context, request *subworkshop.SubPaintCarRequest) (*empty.Empty, error) {
	if err := s.deps.Validations.PaintCar(ctx, request); err != nil {
		return nil, err
	}
	s.deps.Logger.Debug(ctx, "sub workshop - actually painting the car")
	return s.deps.Controller.PaintCar(ctx, request)
}
