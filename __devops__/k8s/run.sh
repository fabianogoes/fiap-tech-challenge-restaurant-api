#!/bin/bash

clear

# Delete all
echo "---------------"
echo "Deleting all..."
echo "---------------"	
kubectl delete --all svc
kubectl delete --all deploy
kubectl delete --all cm
kubectl delete --all pvc
kubectl delete --all pv

# Running Postgres
echo "-------------------"
echo "Running Postgres..."
echo "-------------------"
kubectl apply -f postgres-persistentvolumeclaim.yaml
kubectl apply -f postgres-secret.yaml
kubectl apply -f postgres-deployment.yaml
kubectl apply -f postgres-service.yaml

# Running API
echo "--------------"
echo "Running API..."
echo "--------------"
kubectl apply -f app-secret.yaml
kubectl apply -f app-deployment.yaml
kubectl apply -f app-service.yaml

kubectl get deploy --watch