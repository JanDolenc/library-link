include ../.env

# Build docker image
build-db:
	 docker build . -t ${DATABASE_IMAGE}

# Login to psql client as custom user
query-db:
	docker exec -it database psql -U ${POSTGRES_USER} ${POSTGRES_DB}

# Launch shell in postgres container
sh-db:
	docker exec -it database sh

