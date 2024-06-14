## Installation du Cli Azure

```bash
curl -sL https://aka.ms/InstallAzureCLIDeb | sudo bash
```

## Connexion à Azure

```bash
az login
```

## Définition de l'abonnement au Cluster

```bash
az account set --subscription 904fe061-1956-488a-8e31-2de8692a187a
```

## Obtention des identifiants pour le Cluster

```bash
az aks get-credentials --resource-group aks_group --name projet-aks
```

## Vérification du contexte

```bash
kubectl config get-contexts
```

## Visualiser les Nodes

```bash
kubectl get nodes
```


# ---------------------------------------------------------------------


# Installation de Grafana, Loki et Promtail 

## - Ajouter le répertoire Helm de Grafana

```bash
helm repo add grafana https://grafana.github.io/helm-charts
```

## - Mise à jour de nos dépôts disponibles

```bash
helm repo update
```

## - Lister les Charts disponibls sur le répertoire 

```bash
helm search repo grafana
```

## - Installer Loki

## - Créer le namespace

```bash
kubectl create namespace loki
```

## - Se placer dans le même dossier que le fichier loki et installer le Chart 

```bash 
helm upgrade --install --namespace loki logging grafana/loki -f values.yaml --set loki.auth_enabled=false 
```

## - Lister les éléments que Helm a installé 

```bash 
kubectl -n loki get all
```

## - Lister les PVC créées par MinIO 

```bash
kubectl -n loki get pvc
```

## - Installer Grafana 

```bash
helm install --namespace=loki grafana grafana/grafana
```

## - Passer le fichier promtail-values.yaml 

```bash
helm install --namespace loki --set loki.serviceName=loki-gateway promtail grafana/promtail -f promtail-values.yml
```

## Connexion à Grafana et récupération du mot de passe de Grafana

```bash
kubectl get secret --namespace loki grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo
```

## Créer un tunnel pour nous connecter en local sur le Service de Grafana 

```bash
kubectl port-forward --namespace loki service/grafana 3000:80 
```

## Sur Loki, entrer l'url suivante :

```bash
http://logging-loki-gateway.loki.svc.cluster.local/
```

## Installation de Prometheus

 ### - Ajout du répertoire

```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
```

### - Installation

```bash
helm install prometheus prometheus-community/prometheus -n loki
```

## Dans Grafana allez dans Data Source et sélectionnez Prometheus

## Entrez http://prometheus-server.loki.svc.cluster.local dans Prometheus server URL

## Nettoyage des installations Helm

```bash
helm list -n loki

helm delete grafana -n loki

kubectl get all -n loki

kubectl delete all --all -n loki --force
```