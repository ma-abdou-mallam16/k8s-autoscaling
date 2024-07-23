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

## Installation d'Argo CD

```bash
kubectl create namespace argocd
```

## Application de la configuration

```bash
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

## Vérification de l'installation

```bash
kubectl get all -n argocd 
```

## Connexion à Argo CD

```bash
kubectl port-forward svc/argocd-server -n argocd 8080:443
```

## Créer le mot de passe admin avec le CLI 

```bash
argocd admin initial-password -n argocd
```

## Récupération du mot de passe lors d'une future connexion

```bash
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo
```

## Création de la clé SSH de déploiement avec un accès en écriture 

```bash
ssh-keygen -t ed25519 -C "your_email@example.com"
```

## Ajout de la clé publique sur Github

  1. Copiez le contenu de la clé publique (fichier .pub).

  2. Allez sur GitHub, puis sur la page du dépôt infra.

  3. Cliquez sur "Settings" dans le menu en haut à droite.

  4. Cliquez sur "Deploy keys" dans le menu latéral gauche.

  5. Cliquez sur "Add deploy key".

  6. Collez la clé publique, donnez un titre à la clé (Clé pour les Actions du répertoire code), cochez la case "Allow write access", puis cliquez sur "Add key".

## Utiliser la clé dans GitHub Actions

  1. Copiez le contenu de la clé privée

  2. Allez sur la page du dépôt du code applicatif.

  3. Cliquez sur "Settings", puis sur "Secrets and variables" dans le menu latéral gauche.

  4. Cliquez sur "Actions" puis sur "New repository secret".

  5. Nommez le secret GH_SSH_KEY, et collez le contenu de la clé privée. Cliquez ensuite sur "Add secret".

## Création du job .github/workflows/workflow.yml

