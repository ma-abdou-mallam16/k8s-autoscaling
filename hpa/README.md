## Création des objets 

```bash
kubectl apply -f config.yaml
```

## Mise à l'échelle

```bash
kubectl autoscale deployment php-apache --cpu-percent=50 --min=1 --max=10

kubectl get hpa
```

## Lancer un Pod utilisant une image alpine

```bash
kubectl run -it --rm load-generator --image=alpine /bin/sh
```

## Installer wrk

```bash
apk update && apk add --no-cache wrk

kubectl get hpa --watch

wrk -t4 -c1000 -d30s http://php-apache
```