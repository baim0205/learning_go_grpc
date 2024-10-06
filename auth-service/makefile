# ==============================
# Buat Variable
# ==============================
DOCKER_IMAGE = auth-service:local
DOCKER_NAME = auth-service
DOCKER_NETWORK = app-network
TEST_CONTAINER = auth-service-test

# ==============================
# Ini sama seperti stage 
# ==============================
.PHONY: build deploy clean test

# =================================================
# Ini untuk menjalankan perintahnya sekaligus
# -------------------------------------------------
all: build test deploy clean

#=== Stage Build ====#
build: build-stage
build-stage: # Script Atau Perintah!
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) -f ./Dockerfile ./
	@echo "Build completed. Waiting for 5 seconds..."
	sleep 5

#=== Stage Test ====#
test: test-stage
test-stage: 
	@echo "Running unit tests..."
# Jalankan container untuk testing
	docker run --name $(TEST_CONTAINER) --network $(DOCKER_NETWORK) $(DOCKER_IMAGE) go test -v ./...
# Hentikan dan hapus container setelah test selesai
	docker container stop $(TEST_CONTAINER) && docker container rm $(TEST_CONTAINER)
	@echo "Tests completed. Waiting for 5 seconds..."
	sleep 5

#=== Stage Deploy ====#
deploy: deploy-stage
deploy-stage:
	@echo "Deploying Docker container..."
	docker run -d --name $(DOCKER_NAME) --network $(DOCKER_NETWORK) -p 50051:50051 $(DOCKER_IMAGE)
	@echo "Deployment completed. Waiting for 5 seconds..."
	sleep 5

#=== Stage Clean ====#
clean: stop-stage-container clean-stage-container clean-stage-image
stop-stage-container:
	@echo "Stopping container..."
	docker container stop $(DOCKER_NAME)
	@echo "Container stopped. Waiting for 5 seconds..."
	sleep 5

clean-stage-container:
	@echo "Removing container..."
	docker container rm $(DOCKER_NAME)
	@echo "Container removed. Waiting for 5 seconds..."
	sleep 5

clean-stage-image: # Script Atau Perintah!
	@echo "Removing Docker image..."
	docker rmi $(DOCKER_IMAGE)
	@echo "Image removed."
	sleep 5
