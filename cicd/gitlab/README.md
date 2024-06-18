# GitLab agent for Kubernetes

![Gitlab Architecture](/images/gitlab-architecture.png)

This document contains instructions and configurations for integrating a Kubernetes cluster with GitLab using the GitLab Agent to enable GitLab CI/CD pipelines.

## Step 1 (k8s-connection): Setting Up GitLab and Kubernetes Integration

1. Create a GitLab Repository
    - Create a new repository named k8s-connection on GitLab.
2. Set Up Kubernetes Cluster (Minikube)
    - Ensure you have some cluster running locally (e.g. Minikube)
3. Create Agent Configuration
    - Create an empty configuration file for the GitLab Agent:
    ```bash
    mkdir -p .gitlab/agents/k8s-connection
    touch .gitlab/agents/k8s-connection/config.yaml
    ```
    - Copy-Paste -> microservices/cicd/gitlab/agents/k8s-connection/config.yaml

4. Register Agent with GitLab
    - Navigate to GitLab UI:
        - Go to k8s-connection > Operate > Kubernetes clusters > Connect a cluster > k8s-connection.
    - Run Helm Chart to deploy GitLab Agent:
    ```bash
    helm repo add gitlab https://charts.gitlab.io
    helm repo update

    export AGENT_ACCESS_TOKEN=glagent-xxx-xxxx
    export IMAGE_TAG=v16.4.0
    export NAMESPACE=gitlab-agent

    helm upgrade --install k8s-connection gitlab/gitlab-agent \
        --namespace $NAMESPACE \
        --create-namespace \
        --set image.tag=$IMAGE_TAG \
        --set config.token=$AGENT_ACCESS_TOKEN \
        --set config.kasAddress=wss://kas.gitlab.com
    ```

5. Verify Agent Deployment in Terminal
```bash
kubectl get pods -n $NAMESPACE
```

6. Verify Agent Connection in Gitlab

Menu > Operate > Kubernetes Clusterss
![Gitlab k8s Agent](/images/gitlab-k8s-agent.png)

## Step 2 (k8s-microservices):

1. Create a GitLab Repository
    - Create a new repository named k8s-microservices on GitLab.

2. Configure CI/CD:
    - Create a file named `.gitlab-ci.yml` in the root directory of your project.

3. Define your CI/CD pipeline in `.gitlab-ci.yml`
    - Docker Build / Push Stage
    - K8S Deployment Stage

Gitlab Container Repository is used for the Docker repo.
![Gitlab Container Registry](/images/gitlab-container-registry.png)

4. Create k8s manifest 

    -  Deployment manifests are located in `cicd/gitlab/depl`

5. Add a Git remote named gitlab pointing to the URL of your GitLab repository

    ```bash
    git remote add gitlab git@gitlab.com:<GITLAB_GROUP_ID>/<GITLAB_PROJECT_ID>.git
    ```

6. Push Changes

    ```bash
    git push gitlab main
    ```

7. Verify the deployment stages

![Gitlab CICD](/images/gitlab-cicd.png)
