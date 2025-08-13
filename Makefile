DOCKERCOMPOSECMD=docker-compose

.PHONY: docker-compose-up docker-compose-down docker-compose-restart

docker-compose-up:
	$(DOCKERCOMPOSECMD) up -d --force-recreate

docker-compose-down:
	$(DOCKERCOMPOSECMD) down --remove-orphans

docker-compose-restart: docker-compose-down docker-compose-up