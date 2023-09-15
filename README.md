# latency-test

Microsserviço para teste de latência de services meshes Kubernetes

# Docker

Antes de realizar do deploy do serviço no Kubernetes deve-se gerar a imagem docker e exportá-la para o Kubernetes:
- Gerar imagem docker
	> docker build -t latency-test .
- Exportar imagem para o minikube
	> minikube image load latency-test

# Kubernetes

Para realizar o deploy da aplicação e expor o serviço por meio do loadBalancer execute os procedimentos abaixo:

## Deploy
- Inicializar o tunel do minikube
	> minikube tunnel --bind-address 127.0.0.1
	
- Deploy da aplicação no kubernetes:
	> kubectl apply -f k8s/deployments/latency1.yaml
	
- Expor o serviço na porta especificada
	> kubectl expose deployments latency-test-1 --type="LoadBalancer" --port 8080

## Dashboard

O minikube inicializa o dashboard do kubernetes por meio do comando abaixo
> minikube dashboard