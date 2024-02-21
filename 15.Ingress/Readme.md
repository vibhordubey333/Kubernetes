1. `kubectl apply -f ic-nginx-hn.yaml`
2. `curl http://192.168.49.2:80` or on browser `http://192.168.49.2:80`
3. Fire deployment
```
kubectl create deployment cheddar --image=errm/cheese:cheddar
kubectl create deployment stilton --image=errm/cheese:stilton
kubectl create deployment wensleydale --image=errm/cheese:wensleydale
```
4. Create service
```
kubectl expose deployment cheddar --port=80
kubectl expose deployment stilton --port=80
kubectl expose deployment wensleydale --port=80
```

5. Deploy Ingress `kubectl apply -f ingress.yaml`
6. Verify:  You should see image instead of 404
```
http://cheddar.192.168.49.2.nip.io/
http://stilton.192.168.49.2.nip.io/
http://wensleydale.192.168.49.2.nip.io/
```

7. Annotations are not the standard, say if we're using any ingress annotation and deploying our file then there are chances it will not redirect as annotation may be not the part of cloud platform.

kubectl apply -f https://k8smastery.com/redirect.yaml
Verify using: http://minikube_ip.nip.io or localhost or minikube_ip
It should redirect at google.com

8. kubectl get ingress

```
vibhor@vibhor-virtualbox:$ kubectl get ingress
NAME          CLASS                  HOSTS                             ADDRESS   PORTS   AGE
cheddar       cluster-wide-ingress   cheddar.192.168.49.2.nip.io                 80      2d15h
my-google     cluster-wide-ingress   *                                           80      14h
stilton       cluster-wide-ingress   stilton.192.168.49.2.nip.io                 80      2d15h
wensleydale   cluster-wide-ingress   wensleydale.192.168.49.2.nip.io             80      2d15h
```

* means it's a default one.
Let's say cheddar or stilton are not able to server then my-google will come into play.

9. Testing
`http://192.168.49.2.nip.io`
and it should redirect to google.com