REGISTRY_SERVER=192.168.98.100:5000
NAME=my-restapi
DOCKER_IMAGE=$(REGISTRY_SERVER)/$(NAME):latest

all: vet format build deploy_local

vet:
	go vet .

build:
	CGO_ENABLED=0 go build -tags netgo -o $(NAME).bin

format:
	gofmt -l -w .

deploy_clean:
	kubectl delete deployment $(NAME)
	kubectl delete service $(NAME)

deploy_docker_kubernetes:
	docker build -t $(DOCKER_IMAGE) .
	docker push $(DOCKER_IMAGE)
	kubectl run $(NAME) --image=$(DOCKER_IMAGE) --port=8080
	kubectl expose deployment $(NAME) --type=LoadBalancer
deploy_local: build deploy_docker_kubernetes


.PHONY:
	all \
	vet \
	build
