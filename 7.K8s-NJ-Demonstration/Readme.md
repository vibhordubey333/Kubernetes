#### Big Picture To Achieve. 
![image](https://github.com/vibhordubey333/Kubernetes/assets/22407855/57b62b12-d90d-4cd6-b982-624687fb9880)

##### Request Flow From Browser
![image](https://github.com/vibhordubey333/Kubernetes/assets/22407855/f5611e42-91ae-4a06-bf16-86f29a9ba2cd)

#### Secrets
![image](https://github.com/vibhordubey333/Kubernetes/assets/22407855/fb23591a-ea0e-49bb-af68-e0d3ceb0cbc1)

To generate password and username, which is value itself in below command. They are stored in base64 encoding.<br/>

`echo -n 'username' | base64`<br/>
`echo -n 'password' | base64`<br/>

**Create secret before deployment otherwise there will be error.**

#### Run

1. Create secret `kubectl apply -f mongodb-secret.yaml `<br/>
secret/mongodb-secret created<br/>

2. Get created secret.<br/>
   `kubectl get secret`<br/>
NAME                  TYPE                                  DATA   AGE
default-token-ntphp   kubernetes.io/service-account-token   3      25h
mongodb-secret        Opaque                                2      8s

2. `kubectl apply -f mongodb-configmap.yaml`
3. `kubectl get cm`
4. `kubectl apply -f mongo-express.yaml`
5. `kubectl get svc`
6. To launch mongo-express `minikube service mongo-express-service`.


