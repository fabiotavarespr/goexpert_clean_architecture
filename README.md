# Clean Architecture - Go Expert
[![Go](https://img.shields.io/badge/go-1.21.6-informational?logo=go)](https://go.dev/)
[![MySQL](https://img.shields.io/badge/mysql-8.0.23-informational?logo=mysql)](https://www.mysql.com/)
[![RabbitMQ](https://img.shields.io/badge/rabbitmq-3.12.12-informational?logo=rabbitmq)](https://rabbitmq.com/)

This project implements the third challenge - Clean Architecture - for the Postgraduate in Go Expert.

# Index
- [Clean Architecture - Go Expert](#clean-architecture---go-expert)
- [Index](#index)
- [Stack](#stack)
  - [Use Case](#use-case)
- [Running the project](#running-the-project)
  - [Usage Examples](#usage-examples)
    - [Interacting with the GraphQL service](#interacting-with-the-graphql-service)
    - [Interacting with the gRPC service using Evans](#interacting-with-the-grpc-service-using-evans)
    - [Interacting with HTTP/Web API](#interacting-with-httpweb-api)
- [Stopping the project](#stopping-the-project)

# Stack
- [Golang](https://go.dev/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [MySQL](https://www.mysql.com/)
- [RabbitMQ](https://rabbitmq.com/)
- 

## Use Case

The application has the following use cases:

**Create Order (CreateOrderUseCase)**: This use case is responsible for creating a new order in the system. Receives order information, such as price and rate, validates and saves it in the database. It also triggers an "Order Created" event to notify other systems of the order creation.

**List Orders (ListOrderUseCase)**: This use case is responsible for listing all orders in the system. Retrieves orders from the database and returns a list of orders. It also triggers an event with the entire paylod returned.

# Running the project

To run the application, you will need to have Make, Docker and Docker Compose installed on your machine. If you don't already have them installed, you can download them from the following links:

- Go: [https://go.dev/](https://go.dev/)
- Make: [https://www.gnu.org/software/make/](https://www.gnu.org/software/make/)
- Docker: [https://docs.docker.com/get-docker/](https://docs.docker.com/get-docker/)
- Docker Compose: [https://docs.docker.com/compose/install/](https://docs.docker.com/compose/install/)

After installing all dependencies, follow the steps below to run project infrastructure:

1. Open a terminal in the project root.

2. Run the following command to build and start the services defined in the docker-compose file `make docker-compose-up`:

```bash
make docker-compose-up
```

When the entire infrastructure is running, run the following command to start the project.

1. Open a terminal in the project root.

2. Run the following command to start the project `make run`:

```bash
make run
```

To stop the application, press *Ctrl+C* in the terminal.

## Usage Examples

### Interacting with the GraphQL service

The application provides a GraphQL API for interacting with requests.
Below are some examples of queries and mutations that can be performed:
The GraphQL Playground can be accessed at http://localhost:8080/.

**List all orders**:

```graphql
query listOrder {
  listOrder{
    id
    Price
    Tax
    FinalPrice
  }
}
```

**Create a new order**:

```graphql
mutation createOrder {
  createOrder(input: {Price: 100.5, Tax: 0.12}){
    id
    Price
    Tax
    FinalPrice
  }
}
```

### Interacting with the gRPC service using Evans

To interact with your application's gRPC service using Evans, follow the instructions below:

This will start Evans in REPL mode, allowing you to interact with the application's gRPC service.

```bash
make grpc-run
```

*Chamando métodos do serviço OrderService*

**List all orders**:

```evans
pb.OrderService@127.0.0.1:50051> call ListOrders 
```

**Create a new order**:

```evans
pb.OrderService@127.0.0.1:50051> call CreateOrder
price (TYPE_FLOAT) => 100
tax (TYPE_FLOAT) => 1
```

### Interacting with HTTP/Web API

*HTTP/Web* interface for interacting with orders. Below are examples of how to use HTTP endpoints::

**Create a new order**:

*POST /order*
```bash
curl --location 'http://localhost:8000/order' \
--header 'Content-Type: application/json' \
--data '{
    "price": 102.6,
    "tax": 0.5
}'
```

**List All Orders**:

*GET /orders*
```bash
curl --location 'http://localhost:8000/order'
```

# Stopping the project

To stop the application, press *Ctrl+C* in the terminal.

To stop all dependencies, follow the steps below to stop project infrastructure:

1. Run the following command to stop the services defined in the docker-compose file `make docker-compose-down`:

```bash
make docker-compose-down
```