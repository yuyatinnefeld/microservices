# Integrate a Kubernetes cluster with an external Vault

## Run Vault Server
```bash
vault server -dev -dev-root-token-id root -dev-listen-address 0.0.0.0:8200
export VAULT_ADDR='http://127.0.0.1:8200'
export VAULT_TOKEN=root
vault status
```

## Create Secrets
```bash
# write  ademo secret
vault kv put secret/devwebapp/config username='yuya' password='salsa'

# read the demo secret
vault kv get -mount=secret devwebapp/config
vault kv get -format=json secret/devwebapp/config | jq ".data.data"

# check secret in ui
open http://127.0.0.1:8200

# start k8s dashbaord 
minikube dashboard --url
```

## Determine the Vault address
```bash
minikube ssh

# retrieve the value of the minikube host 
cat /etc/resolv.conf
ip addr show dev eth0

# retrieve the status of the Vault server
echo 192.168.64.1 | xargs -I{} curl -s http://{}:8200/v1/sys/seal-status
exit

EXTERNAL_VAULT_ADDR=192.168.64.1
```

## Deploy app with hard-coded Vault address

```bash
# create a service account
kubectl create sa internal-app

# create a demo pod
cat > devwebapp.yaml <<EOF
apiVersion: v1
kind: Pod
metadata:
  name: devwebapp
  labels:
    app: devwebapp
spec:
  serviceAccountName: internal-app
  containers:
    - name: app
      image: burtlo/devwebapp-ruby:k8s
      env:
      - name: VAULT_ADDR
        value: "http://$EXTERNAL_VAULT_ADDR:8200"
      - name: VAULT_TOKEN
        value: root
EOF

kubectl apply -f devwebapp.yaml

# Request content served at localhost:8080
kubectl exec devwebapp -- curl -s localhost:8080 ; echo
```