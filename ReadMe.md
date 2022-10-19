<details>
<summary>Kubernetes Concepts </summary>

<details>
<summary>Interview Questions</summary>
1. How many ways we can access our microservice deployed using k8's ?<br/>

   A. Inside the pod using `exec` command
   B. Using pod IP `kubectl get po -n work -o wide`
   C. Using services like NodePort,ClusterIP,LoadBalancer and ExternalName

2. 

</details>

</details>


<details> 
<summary> Cheat Sheet</summary>

- Create namespace `kubectl create ns namespace`
- Run deployment and service `kubectl apply -f deployment/service_file_name -n namespace_name`
- Check deployments `kubcetl get deployments -n namespace_name`
- Check services `kubectl get svc -n namespace_name`
- Check IP of pods deployed `kubectl get po -n namespace -o wide`
- Edit the deployment of running pod `kubectl edit deploy pod_name_avoid_hexa_decimal_value -n namespace`
- Delete deployment `kubectl delete deployment deployment_name -n namespace`
- Delete service `kubectl delete svc service_name -n namespace`
- Access pod `kubectl exec -it pod_name -n namespace -- sh`
</details>

 
<details>
<summary>Minikube Installation</summary>

- Install Minikube.
```
minikube version
minikube version: v1.27.1
commit: fe869b5d4da11ba318eb84a3ac00f336411de7ba
```
- Run `minikube --vm-driver=none start --kubernetes-version=v1.17.0`

```
osboxes@osboxes:~$ minikube status
[sudo] password for osboxes:            
minikube
type: Control Plane
host: Running
kubelet: Running
apiserver: Running
kubeconfig: Configured

```
</details>

<details> 
<summary> Enabling Ingress In Minikube </summary>

- Refer documentation `https://kubernetes.io/docs/tasks/access-application-cluster/ingress-minikube/`
```
osboxes@osboxes:~$ minikube addons enable ingress
💡  ingress is an addon maintained by Kubernetes. For any concerns contact minikube on GitHub.
You can view the list of minikube maintainers at: https://github.com/kubernetes/minikube/blob/master/OWNERS
[sudo] password for osboxes:
    ▪ Using image k8s.gcr.io/ingress-nginx/controller:v0.49.3
    ▪ Using image docker.io/jettech/kube-webhook-certgen:v1.5.1
    ▪ Using image docker.io/jettech/kube-webhook-certgen:v1.5.1
🔎  Verifying ingress addon...
🌟  The 'ingress' addon is enabled
```
```
osboxes@osboxes:~$ kubectl get po -n ingress-nginx
NAME                                        READY   STATUS      RESTARTS   AGE
ingress-nginx-admission-create-cxh7x        0/1     Completed   0          3m42s
ingress-nginx-admission-patch-6pb78         0/1     Completed   0          3m42s
ingress-nginx-controller-67fd4fc6fd-hfbcj   1/1     Running     0          3m42s
```
</details>

