# Projeto: Kubernetes e hpa

## 1) Fase 1
Imagem publicada no Docker Hub:

**[https://hub.docker.com/r/alleycarvalho/go-hpa](https://hub.docker.com/r/alleycarvalho/go-hpa)**

Executando container localmente para testes:
```
docker run -d --name go-hpa -p 8000:80 --rm alleycarvalho/go-hpa
```

**[Verificar no browser](http://localhost:8000/)**

Aplicar deployment para o Kubernetes:
```
kubectl apply -f k8s/deployment.yaml
```

Aplicar service para o Kubernetes:
```
kubectl apply -f k8s/service.yaml
```

**[Acessar aplicação rodando com Kubernetes](http://35.226.155.206/)**
