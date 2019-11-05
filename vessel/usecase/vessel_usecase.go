package usecase

import (
	pb "github.com/muhammadhidayah/vessel-service/proto/vessel"
	"github.com/muhammadhidayah/vessel-service/vessel"
)

type vesselUsecase struct {
	repo vessel.Repository
}

func NewVesselUsecase(repo vessel.Repository) vessel.Usecase {
	return &vesselUsecase{repo}
}

func (usecase *vesselUsecase) FindAvailable(specification *pb.Specification) (*pb.Vessel, error) {
	vessel, err := usecase.repo.FindAvailable(specification)

	if err != nil {
		return nil, err
	}

	return vessel, nil
}
