# Cmd de test

## Créer un Pod de test nommé 'testcom'

```bash
kubectl run -it --rm testcom --image alpine -- sh
```

## Installer curl dans le conteneur

```bash
apk add curl
```

## Envoyer requête sur un Pod en utilisant le service headless

```bash
curl http://web-0.webserv
```
