package main

import (
	"fmt"

	"github.com/micro/go-micro"
	pb "github.com/muhammadhidayah/vessel-service/proto/vessel"
	deliveryGRPC "github.com/muhammadhidayah/vessel-service/vessel/delivery/grpc"
	"github.com/muhammadhidayah/vessel-service/vessel/repository"
	"github.com/muhammadhidayah/vessel-service/vessel/usecase"
)

func main() {
	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}

	repo := repository.NewRepoVessel(vessels)
	usecase := usecase.NewVesselUsecase(repo)
	handler := &deliveryGRPC.VesselHandler{usecase}

	srv := micro.NewService(
		micro.Name("vessel"),
	)

	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(), handler)

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
