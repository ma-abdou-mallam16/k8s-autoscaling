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