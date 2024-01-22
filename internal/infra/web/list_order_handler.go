package web

import (
	"encoding/json"
	"github.com/fabiotavarespr/goexpert_clean_architecture/internal/entity"
	"github.com/fabiotavarespr/goexpert_clean_architecture/internal/usecase"
	"github.com/fabiotavarespr/goexpert_clean_architecture/pkg/events"
	"net/http"
)

type WebListOrderHandler struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderListed     events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewWebListOrderHandler(
	OrderRepository entity.OrderRepositoryInterface,
	OrderListed events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *WebListOrderHandler {
	return &WebListOrderHandler{
		OrderRepository: OrderRepository,
		OrderListed:     OrderListed,
		EventDispatcher: EventDispatcher,
	}
}

func (h *WebListOrderHandler) List(res http.ResponseWriter, req *http.Request) {
	orderUseCase := usecase.NewListOrderUseCase(h.OrderRepository, h.OrderListed, h.EventDispatcher)
	output, err := orderUseCase.Execute()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(res).Encode(output)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}
