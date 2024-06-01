# Demo Microservices

## About
This repository provides a simple demo of microservices to be used for learning cloud-native technologies and frameworks such as Istio, ArgoCD, Tilt, Flux, Prometheus, Helm, and Vitess.

## Architecture
![Microservice Architecture](/images/microservice-architecture.png)

The architecture consists of a Python frontend application that interacts with three services:
- **Details Service:** Written in Java.
- **Reviews Service:** Written in Golang.
- **Payment Service:** Written in Golang.

## Getting Started

### Prerequisites
Ensure you have the following installed:
- Docker
- Kubernetes

### Installation

1. **Clone the repository:**
```bash
git clone https://github.com/yourusername/demo-microservices.git
cd demo-microservices
```

2. **Deploy Apps with Tilt File:**
```bash
tilt up
```
