# Simple Go application

This is a simple Go web application.

## Techstack

* psql database
* Echo for server
* GORM
* HTMX and Templ for frontend

## JSON vs HTML REST

routes under the api/v1 namespace will correspond to JSON REST API and api/v2 for HTML REST API.

## Get started

copy the docker-compose.yaml file and run
```sh
docker compose -f docker-compose.yaml up
```