<details>
<summary>Kubernetes Concepts </summary> <br/>

### **1. Kubernetes Terminologies**

**Control Plane:** The collection of processes that control kubernetes nodes. This is where all task assignment originates.

**Node:** A Kubernetes node is a small collection of resources that support one or more containers. Each node contains Docker, kube-proxy and kubelet -- services that help create the runtime environment and support Kubernetes pods.
A Node Controller manages all aspects of the node throughout its lifecycle. It maintains a list of nodes and available machines and resources, and it can delete unhealthy nodes or remove pods from unavailable nodes. You can use the command-line kubectl to run commands against the node.

**Cluster:** A cluster is a group of servers or computing resources that behave as a single system. For the purposes of Kubernetes, a cluster usually means the set of nodes you use to manage and run your containerized applications.

A Kubernetes cluster is made up of one primary node and a number of secondary nodes. The primary node controls the state of the entire cluster. It also issues all task assignments for the cluster, including scheduling, maintenance and updates.

**Kubelet:** *_**This service runs on nodes, reads the container manifests, and ensures the defined containers are started and running.**_*Kubelet is the agent that handles Kubernetes pods for each Kubernetes node. It registers nodes with the API server, and it ensures all containers on a pod are running and healthy. It reports to the primary node regarding the health of its host, and it conveys information to and from the API server. When the control plane requires something from a node, kubelet executes the action.

**Kubectl** The command line configuration tool for kubernetes.

**Kube-proxy.** Kube-proxy facilitates networking services for a Kubernetes environment. It handles networking communications both inside and outside of a Kubernetes cluster, and maintains network rules on nodes. It uses your OS's packet filtering layer when available, and when it can't use the packet filtering layer, it forwards network traffic itself.

**Kubernetes scheduler** The Kubernetes scheduler controls performance, capacity and availability of resources and containers throughout a given Kubernetes environment. It matches each pod you create to a suitable set of resources on a node and distributes copies of pods across different nodes to increase availability. It upholds affinity and anti-affinity rules and quality of service settings.

You can configure Kubernetes scheduler in one of two ways. The PriorityFunction policy directs the scheduler to rank machines based on best fit for a specific node, whereas the FitPredicate policy follows required rules.

**Namespaces:** In Kubernetes, the namespaces is effectively your working area. It's like a project in GCP or a similar thing in AWS.

**Etcd:** Etcd is the primary data store that Kubernetes uses. It contains all configuration data and information about the state of a given cluster, and it stores and replicates all cluster states. You can deploy etcd as either pods on the primary node or as an external cluster.

Etcd is defined as distributed, reliable key-value store for the most critical data of a distributed system.

**Pods:** A pod is effectively a unit of work. It is a way to describe a series of containers, the volumes they might share, and interconnections that those containers within the pod may need. You can have a pod that has a single container in it (or more than one container). Pods are flexible, too: Update one and it becomes version two, and version one is taken out, giving you a rolling update. As Jason spells out, "It gives us a way to say, 'I always want to have three and still be able to migrate an application live from one version to another version without having downtime.'

**Service:** It can be thought of as like a load balancer for pods. It knows which pods are alive, healthy, and ready to respond so that when we try to access whatever pod we want to get to instead of to connect to the deployment and getting the one we get, and then always asking that pod for work."

**Ingress:** This works with the service to make sure everything ends up in the right place. Ingress can also provide load balancing. 

Ingress is not a load balancer, but performs load balancing functions for a Kubernetes environment. It controls traffic to and from services, as well as external access to services. It performs load balancing tasks by setting up an external load balancer and directs traffic to that service based on a set of rules. This enables you to use multiple back-end services via the same IP address.

**Volume:** A Kubernetes volume is a directory containing all data accessible for containers in a given pod. Volumes provide a method for connecting containers and pods -- which only exist as long as you use them -- to a more permanent set of data stored elsewhere. When you delete a pod, the volume associated with it is destroyed as well. However, the data within that volume outlasts the containers or pods that use it.
Kubernetes supports about 20 different varieties of volumes, including emptyDir volumes, local volumes and specialty platform-specific volumes.

**ConfigMaps:** This is an API object for storing information in key-value pairs. "A ConfigMap is very useful for doing things like pre-stashing environment variables or files that can actually be mounted directly into pods without actually having to have an actual file system somewhere," Jason says, adding that they're not meant for confidential data.

**Secrets:** Secrets are an object and a place to store confidential information as the name implies.

## **2. Kubernetes Architecture ?**

## 3. **What about Docker?** <br/>
Docker can be used as a container runtime that Kubernetes orchestrates. When Kubernetes schedules a pod to a node, the kubelet on that node will instruct Docker to launch the specified containers.
The kubelet then continuously collects the status of those containers from Docker and aggregates that information in the control plane. Docker pulls containers onto that node and starts and stops those containers.
The difference when using Kubernetes with Docker is that an automated system asks Docker to do those things instead of the admin doing so manually on all nodes for all containers.
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
- To preview object without actually sending it to API server `kubectl apply -f file_name --dry-run=client`
- To find IP of `headless service`. 
   - Enter into pod `kubectl exec -it pod_name -n namespace -- sh`
   - Fire `nslookup serviceName.namespace.svc.cluster.local` 
- Accessing the pod with private IP using port forwarding.
 Pods are allocated a private IP address by default and cannot be reached outside of the cluster. You can use the??kubectl port-forward??command to map a local    port to a port inside the pod like this:<br/>
 `kubectl port-forward pod-name 8080:8080`
 
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
????  ingress is an addon maintained by Kubernetes. For any concerns contact minikube on GitHub.
You can view the list of minikube maintainers at: https://github.com/kubernetes/minikube/blob/master/OWNERS
[sudo] password for osboxes:
    ??? Using image k8s.gcr.io/ingress-nginx/controller:v0.49.3
    ??? Using image docker.io/jettech/kube-webhook-certgen:v1.5.1
    ??? Using image docker.io/jettech/kube-webhook-certgen:v1.5.1
????  Verifying ingress addon...
????  The 'ingress' addon is enabled
```
```
osboxes@osboxes:~$ kubectl get po -n ingress-nginx
NAME                                        READY   STATUS      RESTARTS   AGE
ingress-nginx-admission-create-cxh7x        0/1     Completed   0          3m42s
ingress-nginx-admission-patch-6pb78         0/1     Completed   0          3m42s
ingress-nginx-controller-67fd4fc6fd-hfbcj   1/1     Running     0          3m42s
```
</details>

