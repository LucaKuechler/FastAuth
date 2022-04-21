# FAST AUTH

## BACKEND CONFIGURATION
* Fill in your Parameters for all uppercase words in the deploy folder.

* Build image.
```bash
$ docker build -t <your-registry-url>/fastauth:latest src/
```

* Login into your private registry.
```bash
$ docker login <your-registry-url>
``` 

* Push to your private registry.
```bash
$ docker push <your-registry-url>/fastauth:latest 
``` 

* Apply configuration.
```bash
$ kubectl apply -f deploy/
```

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

PRIVATE_REGISTRY_USERNAME = username
PRIVATE_REGISTRY_PASSWORD = password

KUBE_CONFIG = {}
```
