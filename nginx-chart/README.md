## Installer le chart pour dev et prod

```bash
helm install nginx-dev ./nginx-chart -f values-dev.yaml

helm install nginx-prod ./nginx-chart -f values-prod.yaml
```

## Installation, MàJ et désinstallation du Chart 

```bash
helm install mon-installation ./mon-chart-simple

helm upgrade mon-installation ./mon-chart-simple
helm rollback mon-installation 1

helm uninstall mon-installation
``` 

