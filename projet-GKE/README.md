## Création du secret pour l'accès aux répertoires privés

```bash
kubectl create secret docker-registry docker-hub-secret --docker-server=https://index.docker.io/v1/ --docker-username=aminou --docker-password=xxxxxxxxxxxxxx --docker-email=contact@k8s.fr
```

## Création de la configuration NGINX (configMap) : nginx-configmap.yaml

```bash
kubectl apply -f nginx-configmap.yaml

kubectl get cm
```

## Déploiement frontend : frontend-deployment.yaml

```bash
kubectl apply -f frontend-deployment.yaml

kubectl get pod
```

## Création du service

```bash
kubectl apply -f frontend-service.yaml

kubectl get sc
```

## Installation de MongoDB sur notre cluster GKE

```bash
helm repo add bitnami https://charts.bitnami.com/bitnami
```

## Créer un fichier mongodb.yaml et placer le dans le dossier helm 

```bash
helm show values bitnami/mongodb > mongodb.yaml
```

## Installer le Chart

```bash
helm install -f mongodb.yaml mongodb bitnami/mongodb
```

## Suivre la création du ReplicaSet

```bash
kubectl logs pod/mongodb-0 -f
```

## Obtenir le mot de passe afin de se connecter ensuite à MongoDB Compass

```bash
kubectl get secret --namespace default mongodb -o jsonpath="{.data.mongodb-root-password}" | base64 -d
```

## Créer un tunnel pour accéder au cluster

```bash
kubectl port-forward svc/mongodb-headless 27017
```

## Lancer le script generate_secrets.sh créée

```bash
bash generate_secrets.sh

kubectl get secrets
```

## Configuration TLS sur GKE

## - Créer une adresse IP publique réservée

### Définisser le projet

```bash
gcloud config set project node-337618
```

### Pour rechercher le nom du projet ou taper la commande 

```bash
gcloud projects list
```

## Créer l'adresse IP

```bash
gcloud compute addresses create nginx-aminou --global
```

## Trouver l'adresse statique créée

```bash
gcloud compute addresses describe nginx-aminou --global

  # ip : 34.54.4.148
```

## Création d'un gestionnaire de certificat Google avec le fichier : gke-managed-cert.yaml

```bash
kubectl apply -f gke-managed-cert.yaml
```

## Redirection HTTPS avec la création du fichier : redirect-https.yaml

```bash
kubectl apply -f redirect-https.yaml
```

## Modification de l'Ingress

```bash
kubectl apply -f ingress.yaml
```

## Lancer cette commande après quelques minutes

```bash
kubectl describe managedcertificate managed-cert
```

## Créer une ressource PodMonitoring

## - Lister les projets Gcloud

```bash
gcloud config list
```

## Créer fichier mongodb-podmonitoring.yaml

```bash
kubectl apply -f mongodb-podmonitoring.yaml
```

## Créer l'opérateur OperatorConfig

```bash
kubectl apply -f operator-monitoring.yaml

kubectl describe podmonitorings
```

## Sauvegarde et restauration de MongoDB

```bash
kubectl apply -f mongodb-backup.yaml
```

## Création d'un accès sur la console

## - Lister les services accounts

```bash
gcloud iam service-accounts list
```

## Création d'une clé privée pour se connecter

```bash
gcloud iam service-accounts keys create --iam-account cronjob-mongodb@premium-archery-396214.iam.gserviceaccount.com service-account.json
```

## Création du secret à partir de la précédente clé créée

```bash
kubectl create secret generic gcs-credentials --from-file service-account.json
```

## Création d'une image pour l'upload des dumps sur GCS

```bash
docker build -t aminou12345/mongobackupgcs .
```

## Puis 

```bash
docker push aminou12345/mongobackupgcs
```

