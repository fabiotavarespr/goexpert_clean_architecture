//go:build wireinject
// +build wireinject

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

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	event.NewOrderList,
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventInterface), new(*event.OrderListed)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

var setOrderListedEvent = wire.NewSet(
	event.NewOrderList,
	wire.Bind(new(events.EventInterface), new(*event.OrderListed)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewListOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.ListOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderListedEvent,
		usecase.NewListOrderUseCase,
	)
	return &usecase.ListOrderUseCase{}
}

func NewWebCreateOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebCreateOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		web.NewWebCreateOrderHandler,
	)
	return &web.WebCreateOrderHandler{}
}

func NewWebListOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebListOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderListedEvent,
		web.NewWebListOrderHandler,
	)
	return &web.WebListOrderHandler{}
}
