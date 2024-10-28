include ../.env

build-api:
	 docker build . -t ${API_IMAGE}

start-api:
	 docker run --name api --network=library-link_backend-net -dp 8080:8080 --env-file ../.env ${API_IMAGE}

stop-api:
	docker rm -f api

logs-api:
	docker logs --follow api

