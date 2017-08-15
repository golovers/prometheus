# Monitor REST API with Prometheus

## Run without Docker and Kubernetes

```sh
cd my-restapi
go run main.go
```

## Deploy with private Docker and Kubernetes

```sh
cd my-prometheus
make

cd ../my-restapi
make

minikube service my-restapi
minikube service my-prometheus
```
## Redeploy
```
cd my-prometheus
make deploy_cleann

cd ../my-restapi
make deploy_clean
```
Then run the deployment above again
