# If facing ErrImage
dockerbuild:
	sudo docker build -t vibhordubey333/html-webapp:latest .
	sudo docker push vibhordubey333/html-webapp:latest
dockerpublish:
	sudo docker run -d --name html-webapp -p 80:80 vibhordubey/html-webapp
logs:
	sudo docker container logs html-webapp -f
rmvdockercontr:
	sudo docker container rm -f html-webapp
deployDepSvc:
	kubectl config current-context
	kubectl config use-context minikube
	kubectl create ns work
	kubectl apply -f html-webapp-deployment.yaml -n work -v 7
	kubectl apply -f html-webapp-NodePort-service.yaml -n work -v 7
svcDepStatus:
	sudo kubectl get svc -n work
	sudo kubectl get deployments -n work
cleanSvcDep:
	sudo kubectl delete deployments demo-web-page -n work
	sudo kubectl delete service html-webapp -n work
	sudo kubectl delete ns work
podStatus:
	sudo kubectl get po -n work -o wide

clusterIPDep:
	sudo kubectl apply -f html-webapp-ClusterIP-service.yaml -n work

# Access NodePort service (macOS compatible - uses minikube service)
accessNodePort:
	@echo "Accessing NodePort service via minikube service..."
	@echo "This will open the service URL in your browser or show the URL to curl"
	@minikube service html-webapp -n work

# Get service URL for NodePort (macOS compatible)
getServiceURL:
	@echo "Service URL:"
	@minikube service html-webapp -n work --url

# Port forward alternative (works on all platforms)
portForward:
	@echo "Setting up port-forward on localhost:8080"
	@echo "Access via: curl http://localhost:8080"
	@kubectl port-forward -n work svc/html-webapp 8080:80
