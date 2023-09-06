deploy:
	kubectl create ns k8s && kubectl apply -f mongodb-configmap.yaml -f mongodb-secret.yaml -f mongo.yaml -f mongo-express.yaml && sleep 5 && watch kubectl get po -n k8s
