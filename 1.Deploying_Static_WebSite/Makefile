dockerbuild:
	sudo docker build -t vibhordubey/html-webapp .
dockerpublish:
	sudo docker run -d --name html-webapp -p 7777:7777 vibhordubey/html-webapp
logs:
	sudo docker container logs html-webapp -f
