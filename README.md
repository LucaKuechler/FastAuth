# Strength Tracker

## CICD CONFIGURATION
* Set following github secrets:
```
GH_JWT_SECRET_KEY = key
GH_DB_USER = user
GH_DB_PASSWORD = password
GH_DB_HOST = hostname:port
GH_DB_NAME = database

GH_REGISTRY_URL = url.com
GH_PULL_SECRET_NAME = name

PRIVATE_REGISTRY_USERNAME
PRIVATE_REGISTRY_PASSWORD

KUBE_CONFIG = {}
```

## BACKEND CONFIGURATION

* Create IngressRoute and Secret
```
kubectl apply -f kubernetes/ingressRoute.yaml -f kubernetes/secret.yaml
```
