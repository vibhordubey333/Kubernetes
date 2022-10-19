## Run
0. Refer Makefile to deploy.
1. Using NodePort service
    `minikube ip` then curl -v http://minikubeip:port pick port from `kubectl get svc -n work`
2. Using ClusterIP service.

```
osboxes@osboxes:~/Documents/Kubernetes/1.Deploying_Static_WebSite$ kubectl get svc -n work
NAME          TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)   AGE
html-webapp   ClusterIP   10.96.61.75   <none>        80/TCP    9s
osboxes@osboxes:~/Documents/Kubernetes/1.Deploying_Static_WebSite$ curl http://10.96.61.75:80
<!DOCTYPE html>

<html>

<head>

<title>HTML</title>

</head>

<body>

<h1>Kubernetes</h1>

<p> Hello World </p>

</body>

</html>
```