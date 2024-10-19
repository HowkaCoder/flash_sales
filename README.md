# Flash Sales Microservices Project

## Overview

Flash Sales is a microservices-based project designed for managing flash sales campaigns, where users can participate in limited-time promotions with high demand. This project aims to demonstrate an event-driven architecture for handling high-concurrency environments, providing real-time updates and seamless user experiences. It was built with scalability, resilience, and simplicity in mind.

## Microservices Overview

1. **User Service**  
   The User Service handles user registration, authentication, and management. It ensures that only verified users can participate in flash sales events.

2. **Product Service**  
   The Product Service manages the product catalog, inventory, and pricing. It provides real-time updates on product availability and handles dynamic pricing during flash sales events.

3. **Sales Service**  
   The Sales Service is responsible for processing transactions during flash sales. It ensures fairness by managing the queue of purchase requests and securing orders for users as stock becomes available.

4. **Notification Service**  
   The Notification Service sends real-time updates to users about their orders, stock availability, and promotional offers. It supports multiple notification channels like email and SMS.

5. **Order Service**  
   The Order Service handles order processing and management. It coordinates with the Sales Service to ensure successful transaction completion and updates the order history for each user.

## Technology Stack

This project employs a modern stack for building distributed microservices:

- **Golang**: The core programming language used for all microservices, known for its simplicity, efficiency, and concurrency support.
- **gRPC**: A high-performance RPC framework that facilitates communication between microservices with contract-based service definitions.
- **Kafka**: An event-streaming platform that ensures reliable communication between services in an event-driven architecture.
- **SQLite**: A lightweight and fast embedded database used for local persistence in microservices, which can be swapped with a more robust database like PostgreSQL for production environments.
- **Docker & Docker Compose**: Containerization technologies used to simplify the development and deployment process. Docker Compose manages the orchestration of all services and dependencies.
- **Event-Driven Architecture**: The project follows an event-driven design where services communicate asynchronously using Kafka. This architecture allows for scalable and decoupled microservices.

### System Architecture

The project uses an event-driven architecture where the microservices are loosely coupled and communicate through Kafka. Each service operates independently but subscribes to relevant events from other services. This architecture ensures that the system can handle high levels of concurrent transactions during flash sales while maintaining consistency and integrity.

## Project Development Roadmap

### Future Improvements

1. **Kubernetes (K8s) Integration**  
   In the next phase, we plan to introduce Kubernetes (K8s) for managing and orchestrating our microservices. A Kafka cluster will be deployed on K8s to handle distributed event streaming. This will improve the scalability and reliability of our messaging system.
   
2. **API Gateway and Reverse Proxy**  
   An API Gateway will be introduced to provide a single entry point for all user-facing services. Additionally, a reverse proxy will be added to handle SSL termination, ensuring all traffic is encrypted.

3. **CI/CD Pipeline**  
   Continuous Integration and Continuous Deployment (CI/CD) pipelines will be integrated into the project. This will enable automated testing, building, and deployment of microservices, ensuring fast and reliable rollouts.

### Potential Enhancements:
- **Advanced Monitoring**: Incorporating Prometheus and Grafana for real-time monitoring of microservice performance and health.
- **Horizontal Scaling**: Adding support for autoscaling microservices based on traffic and resource usage.

## Developer's Comments

This project was developed in a short amount of time, and while it is currently minimalistic, it provides a solid foundation for a larger, more robust system. The current implementation demonstrates key microservice patterns and uses essential technologies like gRPC, Kafka, and Docker to handle event-driven architecture.

In the future, I aim to expand the system by implementing a fully-fledged CI/CD pipeline, adding K8s support for service orchestration, and improving overall system security and scalability. Although it’s a basic prototype, it can serve as a starting point for handling large-scale flash sales applications.

