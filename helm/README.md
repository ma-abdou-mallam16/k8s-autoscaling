## Rechercher l'artifact hub

```bash
helm search hub <nom>

helm search hub mongodb
```

## Rechercher dans les dépôts locaux

```bash
helm repo add dyma https://aminou.github.io/charts
helm search repo aminou
```

## Installer un nouveau package

```bash
helm install happy-panda wordpress

helm install wordpress wordpress
```


## Ajouter un répo bitnami 

```bash
helm repo add bitnami https://charts.bitnami.com/bitnami
```

## Puis 

```bash 
helm show values bitnami/mongodb
```

## Utilisation d'un fichier mongo-values.yaml

```bash
helm install -f mongo-values.yaml bitnami/mongodb --generate-name
```

## Mise à jour d'une release 

```bash
helm upgrade -f mongo-values.yaml mongodb bitnami/mongodb
``` 

## Vérifier que la nouvelle config a été appliquée

```bash
helm get values
```

## Retour en arrière (rollback)

```bash
helm rollback mongodb 1
```

## Voir les numéros de révision d'une version spécifique 

```bash
helm history [RELEASE]
```

## Lister les releases installées 

```bash
helm list --all
```

## Lister les releases désinstallées 

```bash
helm list --uninstalled
```

## Désinstaller une release 

```bash
helm uninstall mongodb

# pour conserver un enregistrement de la suppression de la release

helm uninstall --keep-history
```

## Créer un chart (package kubernetes) avec Helm 

```bash 
helm create nom_du_chart
```

