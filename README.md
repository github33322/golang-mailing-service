# Golang Mailing Service

The Golang Mailing Service is a proof of concept microservice developed in Go that allows storage of customers and sending them emails based on a mailing ID. The service uses Postgres as its database, and both the database and the application are containerized and prepared for deployment using Docker.

## Overview

The project uses various libraries such as `gorilla/mux` for routing, `gorm` for ORM, and `govalidator` for data validation. The project is structured into multiple packages for separation of concerns, including database, models, validators, router, and services.

## Features

The service allows storage of customer data, including their email, the title and content of the message, and a mailing ID. You can add records to the database using the `POST /api/messages` endpoint. You can also send a mocked message to customers with a specific mailing ID and subsequently delete those customers from the database using the `POST /api/messages/send` endpoint. Lastly, you can delete a specific mailing entry with the `DELETE /api/messages/{id}` endpoint.

## Getting started

### Requirements

- Docker
- Go
- Postgres

### Quickstart

1. Clone the repository.
2. Navigate to the project directory.
3. Run the containers with `docker compose up --build`

### TODO

To make this app functional you need to implement your email sending service.

