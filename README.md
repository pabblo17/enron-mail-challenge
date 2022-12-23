
#  Enron Mail Challenge

Application to search for a term in the content of an email in a database


![client](./cli-enron-mail/public/client.png?raw=true)

## Zinc Search (Installation)

Ref. [link](https://docs.zincsearch.com/quickstart/#installation)

### Installation - Local

**S.O.:** Ubuntu-22.04 (WSL Windows)

**Download version:** [releases-zinc-search](https://github.com/zinclabs/zinc/releases).

**Use in this case:** [zinc_0.3.5_Linux_arm64.tar.gz](https://github.com/zinclabs/zinc/releases/download/v0.3.5/zinc_0.3.5_Linux_arm64.tar.gz)

Unzip the contents, preferably in a new folder
```bash
  tar -xvf zinc_0.3.5_Linux_arm64.tar.gz
```
Create a data folder that will store the data
```bash
mkdir data
```
Use to start 
```bash
  ZINC_FIRST_ADMIN_USER=admin ZINC_FIRST_ADMIN_PASSWORD=Complexpass#123 ./zinc
```

Now point your browser to http://localhost:4080 and login

### Installation - Kubernetes (minikube)
Create a namespace:
```bash
kubectl create ns zinc
```
Create the deployment and port forward:
```bash
kubectl apply -f zinc-search/k8s-manifest/kube-deployment.yaml
```
The service is configured as NodePort

Now point your browser to http://localhost:30000 and login or can be made available with an ingress or port-forward:
```bash
kubectl -n zinc port-forward svc/zinc 4080:4080
```

### Installation - Helm Kubernetes (minikube)

Update Helm values located in [values.yaml](https://github.com/zinclabs/zinc/blob/main/helm/zinc/values.yaml)

Create the namespace:
```bash
kubectl create ns zinc
```
Install the chart:

```bash
helm install zinc -f zinc-search/helm/values.yaml --namespace zinc
```
The service is configured as NodePort

Now point your browser to http://localhost:30000 and login or can be made available with an ingress or port-forward:

```bash
kubectl -n zinc port-forward svc/zinc 4080:4080
```

Other option is create chart and install 

```bash
helm package /zinc-search/helm/zinc --version 1.0.0

helm install -f zinc-search/helm/values.yaml zinc zinc-1.0.0.tgz --namespace zinc
```
## Indexer Mails (zinc-search-indexer)

download data in [enron_mail_2011040](http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz)

* Configure and add enviroments use sample file `.env.sample`
    ```bash
    ZINC_SEARCH_HOST="http://localhost:[PORT]/api/"
    ZINC_SEARCH_INDEX="mail"
    ZINC_SEARCH_USER="admin"
    ZINC_SEARCH_PASSWORD="Complexpass#123"
    BATCH_SIZE=500
    ```

* Download dependencies
    ```bash
    go mod download
    ```

* Build app

    ```bash
    go build -o indexer
    ```
* Run indexer

    ```bash
    ./indexer [path of mails to indexer]
    ```

## Server API (api-search-enton-mail)

* Configure and add enviroments use sample file `.env.sample`
    ```bash
    ZINC_SEARCH_HOST="http://localhost:[PORT]/api/"
    ZINC_SEARCH_HOST_ES="http://localhost:[PORT]/es/"
    ZINC_SEARCH_USER="admin"
    ZINC_SEARCH_PASSWORD="Complexpass#123"
    PORT=8080
    ```

* Download dependencies
    ```bash
    go mod download
    ```

* Build app

    ```bash
    go build -o api-search
    ```
* Run indexer

    ```bash
    ./api-search
    ```

## Cliente APP (cli-enron-mail)

* Configure and add enviroments use sample file `.env.sample`
    ```bash
    VUE_APP_SEARCH_API_SERVICE_HOST=http://localhost:[PORT]/api/
    VUE_APP_SEARCH_API_SERVICE_QUERY=http://localhost:[PORT]/api/search
    VUE_APP_SEARCH_API_SERVICE_FIND=http://localhost:[PORT]/api/search/id
    VUE_APP_SEARCH_API_INDEX=mail
    ```

* Install project dependencies
    ```bash
    npm install
    ```

* Run app

    ```bash
    npm run serve -- --port 3000
    ```
* Go [localhost:3000](http:localhost:3000)
## Author
[@PabloChandi](https://www.github.com/pabblo1717)

