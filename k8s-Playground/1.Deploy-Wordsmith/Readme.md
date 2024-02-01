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

**Q. Assignment is to create the kubectl create commands to make this all work:**
This is what we should see when we bring up the web frontend on our browser:
![image](https://github.com/vibhordubey333/Kubernetes/assets/22407855/fcf33231-78c4-4fdf-93bb-2dcce3ff3945)

(You will probably see a different sentence, though.)

- Yes, there is some repetition in that sentence; that's OK for now

- If you see empty LEGO bricks, something's wrong ...

**Questions for this assignment**

Q. **What deployment commands did you use to create the pods?** <br/>
```
kubectl create deployment db --image=bretfisher/wordsmith-db
kubectl create deployment web --image=bretfisher/wordsmith-web
kubectl create deployment api --image=bretfisher/wordsmith-api
```
Q. **What service commands did you use to make the pods available on a friendly DNS name?** <br/>
```
kubectl expose deployment db --port=5432
kubectl expose deployment web --port=80 --type=NodePort
kubectl expose deployment api --port=8080
or
kubectl create service clusterip db --tcp=5432
kubectl create service nodeport web --tcp=80
kubectl create service clusterip api --tcp=8080
```
Q. **What if we needed to add more wordsmith-api pods. What is the command to scale that deployment up to 5 pods?** <br/>
`kubectl scale deployment api --replicas=5`

<br/> To check the gui, `minikube ip` then use web service exposed port.

