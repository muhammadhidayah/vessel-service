package vessel

import (
	pb "github.com/muhammadhidayah/vessel-service/proto/vessel"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}
