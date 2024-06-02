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
# write demo secrets
vault kv put -mount=secret username user_1=jimmy user_2=takashi
vault kv put -mount=secret password user_1=hello user_2=world

# read demo secrets
vault kv get -mount=secret username
vault kv get -mount=secret password

# you can create secrets also in the vault UI
kubectl port-forward vault-0 8200:8200
open localhost:8200
```