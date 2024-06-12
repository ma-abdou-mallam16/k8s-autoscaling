## Utiliser les teintes 

```bash
kubectl taint nodes node1 clé1=valeur1:NoSchedule

kubectl taint nodes node1 app=blue:NoSchedule
```

## Supprimer une teinte 

```bash
kubectl taint nodes <Nom_Node> <Clé_Teinte>-<Effet>- 

kubectl taint nodes node1 app=blue:NoSchedule-
```

## Utiliser tolérances avec l'opérateur 'Equal'

```bash
tolerations:
  - key: "app"
    operator: "Equal"
    value: "blue"
    effect: "NoSchedule"
```

## Utiliser tolérances avec l'opérateur 'Exists'

```bash
tolerations:
  - key: "app"
    operator: "Exists"
    effect: "NoSchedule"
```

## Utiliser tolérances en spécifiant le champ 'tolerationSeconds'

```bash
tolerations:
  - key: "app"
    operator: "Exists"
    effect: "NoExecute"
    tolerationSeconds: 3600
```

