// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"database/sql"
	"github.com/fabiotavarespr/goexpert_clean_architecture/internal/entity"
	"github.com/fabiotavarespr/goexpert_clean_architecture/internal/event"
	"github.com/fabiotavarespr/goexpert_clean_architecture/internal/infra/database"
	"github.com/fabiotavarespr/goexpert_clean_architecture/internal/infra/web"
	"github.com/fabiotavarespr/goexpert_clean_architecture/internal/usecase"
	"github.com/fabiotavarespr/goexpert_clean_architecture/pkg/events"
	"github.com/google/wire"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	orderRepository := database.NewOrderRepository(db)
	orderCreated := event.NewOrderCreated()
	createOrderUseCase := usecase.NewCreateOrderUseCase(orderRepository, orderCreated, eventDispatcher)
	return createOrderUseCase
}

func NewListOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.ListOrderUseCase {
	orderRepository := database.NewOrderRepository(db)
	orderListed := event.NewOrderList()
	listOrderUseCase := usecase.NewListOrderUseCase(orderRepository, orderListed, eventDispatcher)
	return listOrderUseCase
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	orderRepository := database.NewOrderRepository(db)
	orderCreated := event.NewOrderCreated()
	webOrderHandler := web.NewWebOrderHandler(eventDispatcher, orderRepository, orderCreated)
	return webOrderHandler
}

// wire.go:

var setOrderRepositoryDependency = wire.NewSet(database.NewOrderRepository, wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)))

var setEventDispatcherDependency = wire.NewSet(events.NewEventDispatcher, event.NewOrderCreated, event.NewOrderList, wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)), wire.Bind(new(events.EventInterface), new(*event.OrderCreated)), wire.Bind(new(events.EventInterface), new(*event.OrderListed)))

var setOrderCreatedEvent = wire.NewSet(event.NewOrderCreated, wire.Bind(new(events.EventInterface), new(*event.OrderCreated)))

var setOrderListedEvent = wire.NewSet(event.NewOrderList, wire.Bind(new(events.EventInterface), new(*event.OrderListed)))
