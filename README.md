# ORDER-SERVICE

This repository implements a GoLang microservice for managing orders within a larger application. It utilizes the [GoFr](https://gofr.dev/) framework for building a RESTful API.

### Features
- Create, retrieve, update, and delete orders.
- Manage order lifecycle states (e.g., placed, confirmed, shipped, delivered, cancelled).
- Integrate with other microservices for functionalities like inventory management and payment processing.

### Architecture
This service adheres to a microservice architecture, promoting:
- Scalability: Independent scaling based on order volume.
- Resilience: Failure in one service minimizes impact on others.
- Maintainability: Smaller codebases for focused development.

### Technologies
- GoLang: The primary programming language for the service.
- GoFr: An Opinionated Go Framework, for accelerated microservice development.

### Getting Started
- Ensure you have GoLang and Git installed on your system.
- Clone the repository:
```bash
git clone git@github.com:srijan-27/order-service.git
```

- Run the service:
```bash
go run main.go
```

By default, the service listens on port 8080. You can adjust this by setting the HTTP_PORT environment variable.

### API Documentation
Detailed API documentation with endpoints and request/response formats will be added in a separate file (e.g., docs.md).

### Contribution
We welcome contributions to this project. Please refer to the CONTRIBUTING.md file for guidelines.
