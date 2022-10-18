dockerbuild:
	sudo docker build -t vibhordubey/html-webapp .
dockerpublish:
	sudo docker run -d --name html-webapp -p 80:80 vibhordubey/html-webapp
logs:
	sudo docker container logs html-webapp -f
rmvdockercontr:
	sudo docker container rm -f html-webapp
deployDepSvc:
	sudo kubectl create ns work
	sudo kubectl apply -f html-webapp-deployment.yaml -n work -v 7
	sudo kubectl apply -f html-webapp-service.yaml -n work -v 7
svcDepStatus:
	sudo kubectl get svc -n work
	sudo kubectl get deployments -n work
cleanSvcDep:
	sudo kubectl delete deployments demo-web-page -n work
	sudo kubectl delete service html-webapp -n work
	sudo kubectl delete ns work
podStatus:
	sudo kubectl get po -n work -o wide


