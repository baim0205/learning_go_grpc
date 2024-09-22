# Variabel
AUTH_SERVICE_IMAGE = auth-service:local
CRUD_SERVICE_IMAGE = crud-service:local
DOCKER_COMPOSE = docker-compose

# Target untuk membangun image
.PHONY: build
build: build-auth build-crud

build-auth:
	docker build -t $(AUTH_SERVICE_IMAGE) -f ./auth-service/Dockerfile ./auth-service

build-crud:
	docker build -t $(CRUD_SERVICE_IMAGE) -f ./crud-service/Dockerfile ./crud-service

# Target untuk menjalankan container
.PHONY: up
up:
	$(DOCKER_COMPOSE) up -d

# Target untuk menghentikan dan menghapus container
.PHONY: down
down:
	$(DOCKER_COMPOSE) down

# Target untuk membersihkan image
.PHONY: clean
clean:
	docker rmi $(AUTH_SERVICE_IMAGE) $(CRUD_SERVICE_IMAGE) || true

# Target untuk menjalankan build dan up
.PHONY: run
run: build up
