# Movies Demo Application

## Architecture

Movies is composed of 5 microservices in different programming languages. They talk to each other with rest callings but on the next step they are going to talk each other over gRPC. The application uses MongoDB.

|  Service       | Language                      
|----------------|------------------------------
|Frontend        | `VueJS`                   
|Operations      | `Golang`               
|Favorites       | `Pyhton`
|Authentication  | `Golang`
|Review          | `NodeJS`


## Features

- Repository has deployment files for running on **Kubernetes**. Deployment files prepared for local Kubernetes deployments.
- [Istio](https://istio.io/latest/)  is an open source service mesh and I added yamls to repo for istio configuration. 
-  I prepared a [Helm](https://helm.sh/)  chart. It helps us to manage our kubernetes applications. Also with [Helmfile](https://github.com/roboll/helmfile) we can manage helm charts.

## Todos
- Add helm chart and helmfile
- gRPC Implementation
- Handle logs
