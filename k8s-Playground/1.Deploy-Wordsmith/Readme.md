Wordsmith has 3 components: <br/>
   -  a web frontend bretfisher/wordsmith-web <br/>
   -  an API backend bretfisher/wordsmith-api NOTE: previously called "words" <br/>
   -   a postgres DB bretfisher/wordsmith-db <br/>

These images have been built and pushed to Docker Hub
We want to deploy all 3 components on Kubernetes

Here's how the parts of this app communicate with each other:

 -  The web frontend listens on port 80  
 -  The web frontend should be public (available on a high port from outside the cluster)
 -  The web frontend connects to the API at the address http://api:8080
 -  The API backend listens on port 8080
 -  The API connects to the database with the connection string pgsql://db:5432
 -  The database listens on port 5432
 -  Your Assignment is to create the kubectl create commands to make this all work
