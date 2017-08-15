REGISTRY_SERVER=192.168.98.100:5000
NAME=my-prometheus
DOCKER_IMAGE=$(REGISTRY_SERVER)/$(NAME):latest

all: deploy_local

deploy_clean:
		kubectl delete deployment $(NAME)
		kubectl delete service $(NAME)
deploy_local:
	docker build -t $(DOCKER_IMAGE) .
	docker push $(DOCKER_IMAGE)
	kubectl run $(NAME) --image=$(DOCKER_IMAGE) --port=9090
	kubectl expose deployment $(NAME) --type=LoadBalancer

.PHONY:
	all \
