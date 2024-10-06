# Variable
DOCKER_IMAGE = crud-service:local
DOCKER_NAME = crud-service
DOCKER_NETWORK = app-network
TEST_CONTAINER = crud-service-test

# Rules!!!
.PHONY: build test deploy clean

all : build test deploy clean

build: build-stage
build-stage: 
	@echo "Membuat Docker Image!!!"
	docker build -t $(DOCKER_IMAGE) -f ./Dockerfile ./
	@echo "Selesai di build!, diberikan jeda menunggu selama 5 detik!"
	sleep 5

test: test-stage
test-stage: 
# Jalankan container untuk testing
	@echo "Proses UNIT TEST di mulai!!!"
	docker run --name $(TEST_CONTAINER) --network $(DOCKER_NETWORK) $(DOCKER_IMAGE) go test -v ./...

# Hentikan dan hapus container setelah test selesai
	docker container stop $(TEST_CONTAINER) && docker container rm $(TEST_CONTAINER)
	@echo "Tests Sudah Selesai Menunggu Lagi 5 Detik"


deploy: deploy-stage
deploy-stage:
	@echo "Sekarang DEPLOY!!!... "
	docker run -d --name $(DOCKER_NAME) --network $(DOCKER_NETWORK) -p 50052:50052 $(DOCKER_NAME)
	@echo "DEPLOYMENT Selesai. Menunggu 5 detik untuk proses selanjutnya "CLEAN" "

clean: stop-stage-container clean-stage-container clean-stage-image
stop-stage-container:
	@echo "Stopping container..."
	docker container stop $(DOCKER_NAME)
	@echo "Container stopped. Waiting for 5 seconds..."
	sleep 5

clean-stage-container:
	@echo "Hapus container..."
	docker container rm $(DOCKER_NAME)
	@echo "Container removed. Waiting for 5 seconds..."
	sleep 5

clean-stage-image: # Script Atau Perintah!
	@echo "Hapus Docker image..."
	docker rmi $(DOCKER_IMAGE)
	@echo "Image removed."
	sleep 5