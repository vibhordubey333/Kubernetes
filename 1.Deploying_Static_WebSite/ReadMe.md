### Steps To Run
0. Refer Makefile to deploy.
1. Using NodePort service
    `minikube ip` then curl -v http://minikubeip:port pick port from `kubectl get svc -n work`

   -    Command:
         ```
            vibhor@vibhor-virtualbox:~/code-repositories/Kubernetes/1.Deploying_Static_WebSite (master)$ kubectl get svc -n work
            NAME          TYPE       CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE
            html-webapp   NodePort   10.107.183.127   <none>        80:32765/TCP   101m

         ```
    -    Curl requests:
          `curl http://192.168.58.2:32765`
         
   3. Using ClusterIP service.
       ClusterIP service is not exposed outside the cluster. In order to access. <br/>
           -   `kubectl exec --namespace=work -it demo-web-page-7d44d8fd-jztqq -- sh` <br/>
           -   Use CluserIP from command `kubectl get svc -n work` <br/>
           -   Then fire `curl 10.107.183.127` or `curl localhost`

```
osboxes@osboxes:~/Documents/Kubernetes/1.Deploying_Static_WeblhhhhhhostSite$ kubectl get svc -n work
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

If above is not working then alternative solution is: <br/>
    -   `kubectl proxy` <br/>
    -   In another terminal   `curl -v http://localhost:80` <br/>

Minikube version

```
minikube version: v1.32.0
commit: 8220a6eb95f0a4d75f7f2d7b14cef975f050512d
```

Verify internal Minikube pods are up.

```
vibhor@vibhor-virtualbox:~/code-repositories/Kubernetes/1.Deploying_Static_WebSite (master *)$ kubectl get po --all-namespaces -o wide
NAMESPACE     NAME                               READY   STATUS    RESTARTS      AGE     IP             NODE       NOMINATED NODE   READINESS GATES
kube-system   coredns-5dd5756b68-bkdgn           1/1     Running   0             3h21m   10.244.0.2     minikube   <none>           <none>
kube-system   etcd-minikube                      1/1     Running   0             3h21m   192.168.58.2   minikube   <none>           <none>
kube-system   kube-apiserver-minikube            1/1     Running   0             3h21m   192.168.58.2   minikube   <none>           <none>
kube-system   kube-controller-manager-minikube   1/1     Running   0             3h21m   192.168.58.2   minikube   <none>           <none>
kube-system   kube-proxy-txpv7                   1/1     Running   0             3h21m   192.168.58.2   minikube   <none>           <none>
kube-system   kube-scheduler-minikube            1/1     Running   0             3h21m   192.168.58.2   minikube   <none>           <none>
kube-system   storage-provisioner                1/1     Running   2 (35m ago)   3h21m   192.168.58.2   minikube   <none>           <none>
work          demo-web-page-7d44d8fd-jztqq       1/1     Running   0             3h18m   10.244.0.3     minikube   <none>           <none>
work          demo-web-page-7d44d8fd-rx4l5       1/1     Running   0             3h18m   10.244.0.4     minikube   <none>           <none>

```