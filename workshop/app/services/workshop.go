package services

import (
	"context"
	"fmt"
	"github.com/go-bricks/bricks/interfaces/monitor"

	workshop "github.com/go-bricks/bricks-demo/workshop/api"
	"github.com/go-bricks/bricks-demo/workshop/app/controllers"
	"github.com/go-bricks/bricks-demo/workshop/app/validations"
	"github.com/go-bricks/bricks/interfaces/log"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/fx"
)

type workshopServiceDeps struct {
	fx.In

	Logger      log.Logger
	Controller  controllers.WorkshopController
	Validations validations.WorkshopValidations
	Metrics     monitor.Metrics `optional:"true"`
}

type workshopImpl struct {
	deps                                 workshopServiceDeps
	workshop.UnimplementedWorkshopServer // if keep this one added even when you change your interface this code will compile
}

func CreateWorkshopService(deps workshopServiceDeps) workshop.WorkshopServer {
	return &workshopImpl{
		deps: deps,
	}
}

func (w *workshopImpl) AcceptCar(ctx context.Context, car *workshop.Car) (*empty.Empty, error) {
	if err := w.deps.Validations.AcceptCar(ctx, car); err != nil {
		return nil, err
	}
	w.deps.Logger.WithField("car", car).Debug(ctx, "accepting car")
	return w.deps.Controller.AcceptCar(ctx, car)
}

func (w *workshopImpl) PaintCar(ctx context.Context, request *workshop.PaintCarRequest) (result *empty.Empty, err error) {
	defer func() {
		counter := w.deps.Metrics.WithTags(monitor.Tags{
			"color":   request.GetDesiredColor(),
			"success": fmt.Sprintf("%t", err == nil),
		}).Counter("paint_desired_color", "New paint color for car")
		counter.Inc()
	}()
	err = w.deps.Validations.PaintCar(ctx, request)
	if err == nil {
		result, err = w.deps.Controller.PaintCar(ctx, request)
	}
	w.deps.Logger.WithError(err).Debug(ctx, "sending car to be painted")
	return
}

func (w *workshopImpl) RetrieveCar(ctx context.Context, request *workshop.RetrieveCarRequest) (*workshop.Car, error) {
	if err := w.deps.Validations.RetrieveCar(ctx, request); err != nil {
		return nil, err
	}
	w.deps.Logger.Debug(ctx, "retrieving car")
	return w.deps.Controller.RetrieveCar(ctx, request)
}

func (w *workshopImpl) CarPainted(ctx context.Context, request *workshop.PaintFinishedRequest) (*empty.Empty, error) {
	if err := w.deps.Validations.CarPainted(ctx, request); err != nil {
		return nil, err
	}
	w.deps.Logger.Debug(ctx, "car painted by sub workshop")
	return w.deps.Controller.CarPainted(ctx, request)
}
