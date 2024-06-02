## Création du statefulset

```bash
kubectl apply -f my-sql-statefulset.yaml
```

## Liste tous les objets k8s 

```bash
kubectl get all
```

## Afficher PVC

```bash
kubectl get pvc
```

## Afficher PV

```bash
kubectl get pv
```

## Suppression du statefulset

```bash
kubectl delete -f my-sql-statefulset.yaml
```

## Suppression de PVC et PV

```bash
kubectl delete pvc mysql-volume-mysql-statefulset-0
```

## Visualiser les logs

```bash
kubectl logs pods/mysql-statefulset-0
```

## Création du service clusterIP 

```bash
kubectl apply -f my-sql-service.yaml
```

## Déployer le fichier php-nginx-deployment.yaml

```bash 
kubectl apply -f php-nginx-deploy.yaml
```

## Vérifions les conteneurs nginx et php-fpm

```bash
kubectl get pods
```

## Vérification du conteneur nginx déployé

```bash
kubectl logs php-nginx-deployment-67cf94bd98-wr9cb -c init-mysql
```

## Vérification du conteneur php-fpm déployé

```bash
kubectl logs php-nginx-deployment-5db8fb8766-bxjdh -c php-fpm
```

## Se connecter au conteneur nginx

```bash
kubectl exec -it php-nginx-deployment-5db8fb8766-bxjdh -c nginx -- sh
```

## Vérifier la configuration de nginx

```bash
nginx -T
```

## Se connecter au conteneur php-fpm

```bash
kubectl exec -it php-nginx-deployment-5db8fb8766-bxjdh -c php-fpm -- bash
```

## Appliquer le service loadbalancer.yaml

```bash
kubectl apply -f loadbalancer.yaml
```

## Afficher le service loadbalancer

```bash
kubectl get svc
```

## Supprimer les éléments du cluster

```bash
kubectl delete -f k8s/
```

## Récréation des éléments du cluster

```bash
kubectl apply -f k8s/
```