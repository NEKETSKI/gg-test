# gg-test

## Building
Application can be build using `make build` command. This command creates a Docker image with an application.

## Testing
To run a unit tests execute `make test` command. This command also starts the Postgres container.

## Running
Service can be run locally using a `docker compose`. File located in [deployments/compose](https://github.com/NEKETSKY/gg-test/blob/main/deployments/compose/docker-compose.yml).