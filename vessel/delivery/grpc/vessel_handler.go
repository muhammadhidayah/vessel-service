package grpc

import (
	"context"

	pb "github.com/muhammadhidayah/vessel-service/proto/vessel"
	"github.com/muhammadhidayah/vessel-service/vessel"
)

type VesselHandler struct {
	VUsecase vessel.Usecase
}

func (handler *VesselHandler) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	vessel, err := handler.VUsecase.FindAvailable(req)

	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}
