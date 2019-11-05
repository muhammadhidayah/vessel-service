package repository

import (
	"errors"

	pb "github.com/muhammadhidayah/vessel-service/proto/vessel"
	"github.com/muhammadhidayah/vessel-service/vessel"
)

type repoVessel struct {
	vessels []*pb.Vessel
}

func NewRepoVessel(vessels []*pb.Vessel) vessel.Repository {
	return &repoVessel{vessels}
}

func (repo *repoVessel) FindAvailable(specification *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if specification.Capacity <= vessel.Capacity && specification.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}

	return nil, errors.New("No vessel found")
}
