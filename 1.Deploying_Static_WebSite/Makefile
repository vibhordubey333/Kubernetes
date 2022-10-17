dockerbuild:
	sudo docker build -t vibhordubey/html-webapp .
dockerpublish:
	sudo docker run -d --name html-webapp -p 80:80 vibhordubey/html-webapp
logs:
	sudo docker container logs html-webapp -f
rmvcontr:
	sudo docker container rm -f html-webapp
