# Docker and Postgres
```shell
# Create Network
docker network create --driver bridge postgres-network

# Postgres
docker run \
    --name postgres \
    --network=postgres-network \
    -e POSTGRES_USER=chipman \
    -e POSTGRES_PASSWORD=123mudar \
    -e POSTGRES_DB=northwind \
    -p 5432:5432 \
    -d \
    postgres

# Pgadmin
docker run \
    --name pgadmin \
    --network=postgres-network \
    -p 15432:80 \
    -e PGADMIN_DEFAULT_EMAIL=email@email.com \
    -e PGADMIN_DEFAULT_PASSWORD=123mudar \
    -d dpage/pgadmin4

```