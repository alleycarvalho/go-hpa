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

### Print da execução (CI) no Cloud Build:
![CI](/go-hpa.png)

## 2) Implementando o "hpa"
Aplicar hpa para o Kubernetes:
```
kubectl apply -f k8s/hpa.yaml
```

## 3) POD com looping infinito
Criar uma máquina busybox acessando o bash:
```
kubectl run -it loader --image=busybox /bin/sh
```

No terminal da máquina busybox, testar a aplicação:
```
wget -q -O- http://go-hpa-service.default.svc.cluster.local;
```

No terminal da máquina busybox, rodar o loop para escalar:
```
while true; do wget -q -O- http://go-hpa-service.default.svc.cluster.local; done;
```

Em outro terminal, verificar hpa a cada 15 segundos:
```
watch kubectl get hpa
```
