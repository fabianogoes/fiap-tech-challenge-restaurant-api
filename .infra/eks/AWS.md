# Deploy AWS
How to deploy and test the application into AWS using EKS

## Creating the Cluster

- [ ] AWS Academy `Start Lab`
- [ ] Update `~./aws/credentials`
- [ ] EKS - create cluster `tech-challenge-f2` - wait for status `active`
- [ ] EKS - create nodes - wait for status `active`
- [ ] AWS cli - `aws sts get-caller-identity`
- [ ] AWS cli - `aws eks update-kubeconfig --region us-east-1 --name tech_challenge_eks_cluster`

## Deploying the Application

- [ ] `kubectl cluster-info`
- [ ] `kubectl apply -f metrics.yaml`
- [ ] `kubectl apply -f postgres-persistentvolumeclaim.yaml`
- [ ] `kubectl apply -f postgres-secret.yaml`
- [ ] `kubectl apply -f postgres-deployment.yaml`
- [ ] `kubectl apply -f postgres-service.yaml`
- [ ] Encripting database host secrets
  - [ ] Host: `echo 'tech-challenge-db.czj1nxgd8v9c.us-east-1.rds.amazonaws.com' | base64`
  - [ ] Update `app-secret.yaml` with `DB_HOST` value
- [ ] `kubectl apply -f app-secret.yaml`
- [ ] `kubectl apply -f app-deployment.yaml`
- [ ] `kubectl apply -f app-hpa.yaml`
- [ ] `kubectl apply -f app-service.yaml`
- [ ] `kubectl get deploy --watch`

## Testing the Application Order Flow

- [ ] `kubectl get svc` and copy ExternalIP and paste into Postman environment
- [ ] Test `/health`
- [ ] Test Order Flow


## Testing HPA autoscaling

- [ ] `./stress.sh`
- [ ] `kubectl get pods --watch`

## Deleting all into AWS

- [ ] `./delete-all.sh`
- [ ] AWS EKS - Delete Nodes
- [ ] AWS EKS - Delete Cluster
- [ ] AWS Academy `End Lab`

## References

- [Encrypting Confidential Data at Rest](https://kubernetes.io/docs/tasks/administer-cluster/encrypt-data/)


