# Boilerplate REST service using Gin web framework (golang)

[![Go](https://github.com/lakhinsu/gin-boilerplate/actions/workflows/go.yml/badge.svg)](https://github.com/lakhinsu/gin-boilerplate/actions/workflows/go.yml) 

[![Go Report Card](https://goreportcard.com/badge/github.com/lakhinsu/gin-boilerplate)](https://goreportcard.com/report/github.com/lakhinsu/gin-boilerplate)

## Introduction

This repository contains a boilerplate REST API service using [Gin](https://gin-gonic.com/) web framework.

This boilerplate can be used to quickly bootstrap a backend project with minimal changes. It uses most common and highly adpated libraries like [zerolog](https://github.com/rs/zerolog) for logging and [viper](https://github.com/spf13/viper) for env configuration.

I have also written a [blog](https://medium.com/@hinsulak/rest-api-service-boilerplate-using-gin-web-framework-golang-c4eeb48b14f) where i explain the process behind creating this boilerplate, feel free to check it out.

## Features
- HTTP and HTTPS support
- .env file and OS environment variables support
- Models
- Controllers
- Routers
- Request ID middleware
- Request logging middleware
- CORS middleware


## Run locally
This repo does not contain a Dockerfile at the moment so i'm listing down steps for running this service locally.
- Get dependencies
`go get .`
- Set environment variables
  Fill out the sample .env file provided in the repository

Execute `go run main.go` command at the repository root to start the service.
## Running in Docker
The repo includes sample Dockerfile and docker-compose.yaml files.
