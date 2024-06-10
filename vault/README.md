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
vault kv put secret/yuya_password/config username='mango' password='wave'

# read the demo secret
vault kv get -mount=secret yuya_password/config
vault kv get -format=json secret/yuya_password/config | jq ".data.data"

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
# deploy with tilt
cd vault
tilt up

# deploy with k8s
kubectl apply -f k8s/service-mesh/vault/vault.yaml
```

## Verify 
```bash
# see the website
open localhost:8899
# change the secret
vault kv put secret/yuya_password/config username='mango' password='aceno'
```
