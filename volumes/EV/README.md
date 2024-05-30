## Créer POD

```bash 
kubectl apply -f myPod.yaml
```

## Commnde pour accéder au fichier depuis le premier conteneur

```bash
kubectl exec myPod -c conteneur-1 -- cat /usr/share/nginx/html/index.html
```

## Cmd pour se connecter au premier conteneur

```bash 
kubectl exec -it myPod -c conteneur-1 -- sh
```

## Cmd pour se connecter au second conteneur

```bash
kubectl exec -it myPod -c conteneur-2 -- sh
```