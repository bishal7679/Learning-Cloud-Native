## All argocd command

### 1. Install Argo CD

```js
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

### 2. Download Argo CD CLI

- Download the latest Argo CD version from <https://github.com/argoproj/argo-cd/releases/latest>.


### 3. Access The Argo CD API Server

- By default, the Argo CD API server is not exposed with an external IP. To access the API server, choose one of the following techniques to expose the Argo CD API server:

- port forwarding
```js
kubectl port-forward svc/argocd-server -n argocd 8080:443
```

### 4. Login Using The CLI in Terminal

The initial password for the admin account is auto-generated and stored as clear text in the field password in a secret named argocd-initial-admin-secret in your Argo CD installation namespace. You can simply retrieve this password using kubectl:

```js
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo
```

```js
argocd login localhost:8081
```
```js
argocd account update-password
```

### 5. Register A Cluster To Deploy Apps To (Optional)

```js
kubectl config get-contexts -o name
```
```js
argocd cluster add <put here context-name>
```

### 6. Create An Application From A Git Repository

```js
kubens argocd
```
```js
argocd app create demo2 \  
--project default \
--repo https://github.com/codefresh-contrib/gitops-certification-examples \
--path "./simple-app" \
--dest-namespace default \
--dest-server https://kubernetes.default.svc
```
```js
kubectl get applications
```
#### Syncing the created application named demo2
```js
argocd app sync demo2
```