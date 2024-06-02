## Cmd pour voir les classes disponibles

```bash 
kubectl get sc
```

## Ou la version plus longue

```bash
kubectl get storageclasses
```

## Vérification après création de PVC

```bash
kubectl get pvc
```

## Vérifier que le PV a été créée dynamiquement

```bash
kubectl get pv
```

## Afficher les logs

```bash
kubectl exec nginx-deployment-57445776fd-7qshw -- cat /var/log/nginx/error.log
```

## Suppression du déploiement

```bash 
kubectl delete -f mon-deploiement.yaml
```

## Suppression du PVC

```bash
kubectl delete pvc nginx-logs
```

