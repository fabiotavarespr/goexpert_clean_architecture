package service

import (
	"context"

	"github.com/fabiotavarespr/goexpert_clean_architecture/internal/infra/grpc/pb"
	"github.com/fabiotavarespr/goexpert_clean_architecture/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrderUseCase   usecase.ListOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, listOrderUseCase usecase.ListOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrderUseCase:   listOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, in *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	output, err := s.ListOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}
	var orders []*pb.OrderResponse
	for _, order := range output {
		orders = append(orders, &pb.OrderResponse{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		})
	}
	return &pb.ListOrdersResponse{
		Orders: orders,
	}, nil
}
