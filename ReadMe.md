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


#### **Namespaces**<br/>
![image](https://github.com/vibhordubey333/Kubernetes/assets/22407855/1d3ce8b5-54f0-4d39-a59e-ca198973a6fd)

![image](https://github.com/vibhordubey333/Kubernetes/assets/22407855/42099241-e2be-4784-b675-ea75209f7d60)


	- **Use cases:**<br/>		
 		1. Resource grouped in a namespace. Like Monitoring,Database,Elasic Stack.<br/>
		2. Conflicts: Many teams and same application. If two teams are using same cluster then it creates the issue as they both may overwrite the cluster.<br/>
		3. Resource sharing. Staging and development. Say, we deployed monitoring tools in a different namespace then they can be shared by staging and development environment.<br/>
		4. Blue/Green Deployment : Using different-different versions.<br/>
		5. Access and resource limits on namespace. Gives team access to their own namespace. Limit their resources as well.<br/>
	
	- **Limitations:**
		1. Say we've namespaces A1 and B1. And A1 has config map that references DB service now B1 service cannot use A1's config map it needs to create it's own. So, each NS must define there own ConfigMap. Same rule applies for Secrets as well.<br/>
		2. Resources which are not bound to namespaces are: volume, node. 
			- To view resources which are not bound to namespaces can be checked using:  `kubectl api-resources --namespaced=false`
			- `kubectl api-resources --namespaced=true` [Which are bound to a namespace.]
	
	 - K8's has 4 default namespaces. 
	```
	NAME              STATUS   AGE
	default           Active   146m -> If no namespace is defined then by default resources are created in this namespace.
	kube-node-lease   Active   146m -> Recent addition, it checks for the heartbeat of node. Each node has associated lease object in namespace which determines the node availbility.
	kube-public       Active   146m -> Contains publicly accessible data like config maps which contains cluster information. 
	kube-system       Active   146m -> Avoid tinkering with it. Contains, system processes[Master and kubectl processes]
	```
	-> Various methods to create a new namespace
		A. kubectl create ns your_namespace_name_here 
		B. Or kubectl apply -f deployment-file-name --namespace=your_namespace_name
		C. Best Practise. Define it in your deployment file. 



#### **Selectors**<br/>
Selectors, on the other hand, are used to select a group of objects based on their labels. They are used in various Kubernetes objects like Services, ReplicaSets, and Deployments to specify the Pods they are associated with.<br>

![image](https://github.com/vibhordubey333/Kubernetes/assets/22407855/dbf761c1-6ec9-496d-abd0-ef60aaf030da)


#### Exposing deployment port to service.
Ports defined in deployment is connecting to container port and port defined in service file is connecting to deployment port.

![image](https://github.com/vibhordubey333/Kubernetes/assets/22407855/38b4b060-8871-4d30-860b-947c151e0e66)

To verify whether port information[above] is correct or not:

```
kubectl get pod -o wide
kubectl get pod -o wide your_pod_name
```

#### Verify if service is mapped to a correct pod.
Execute `kubectl get svc` then `kubectl describe svc pod_name`
```
Name:              mongodb-service
	Namespace:         default
	Labels:            <none>
	Annotations:       <none>
	Selector:          app=mongodb
	Type:              ClusterIP
	IP Families:       <none>
	IP:                10.105.123.117
	IPs:               10.105.123.117
	Port:              <unset>  27017/TCP
	TargetPort:        27017/TCP
	Endpoints:         172.17.0.3:27017
	Session Affinity:  None
	Events:            <none>
```

**Emphasise on Endpoints field**

Execute `kubectl get po -o wide`<br/>

```
	NAME                                  READY   STATUS    RESTARTS   AGE   IP           NODE                NOMINATED NODE   READINESS GATES
	mongodb-deployment-77f7698445-xd49j   1/1     Running   0          79m   172.17.0.3   vibhor-virtualbox   <none>           <none>
```

**_Details from above commands are same i.e Pod IP address is same. Hence it is referring correct pod._**

## Ingress

  - **Without Ingress**<br/>
    We’ve to expose the application using which can be accessed using ip address and port no. <br/>
		
    **Usability**
      - Good for testcases
      - To tryout things.
        <br/>
        ![image](https://github.com/vibhordubey333/Kubernetes/assets/22407855/5a5f32b9-ea11-4499-a29d-53e444a55015)
        <br/>
	But production ready application should not accessed using ip address and port.<br/>
 
 	  ![image](https://github.com/vibhordubey333/Kubernetes/assets/22407855/30eaa433-98df-4cc2-9070-8a9fa8d6d007)

        YAML file example<br/>
    
    ![image](https://github.com/vibhordubey333/Kubernetes/assets/22407855/37a5cc45-2d66-48dd-91c3-50ea1020eb4e)

  - **With Ingress**
     - Request coming from outside will land on Ingress and ingress will forward it to pod.<br/>
       ![image](https://github.com/vibhordubey333/Kubernetes/assets/22407855/22953166-0464-43dd-8efb-dbbb7d110321)<br/>
       ![image](https://github.com/vibhordubey333/Kubernetes/assets/22407855/f1312253-89a6-4579-88a9-c32ff3ebbfe6)<br/>
     - YAML file example<br/>
       ![image](https://github.com/vibhordubey333/Kubernetes/assets/22407855/c3bb0e96-b7c3-4980-b624-c4095c42c5b2)<br/>
     - Ingress configuration<br/>
       ![image](https://github.com/vibhordubey333/Kubernetes/assets/22407855/1e0e98a8-7a41-4b0b-80b5-8782e2a8e568)<br/>
       ![image](https://github.com/vibhordubey333/Kubernetes/assets/22407855/7e3230a3-2419-4834-a9a8-edaf1d6292c5)<br/>
       ![image](https://github.com/vibhordubey333/Kubernetes/assets/22407855/e04f3e17-0220-4457-a68d-dd11c846fb9d)<br/>
       ![image](https://github.com/vibhordubey333/Kubernetes/assets/22407855/00ad772b-0217-4a4c-b408-955d963263a8)<br/>
     - For GCE/AWS we can use cloud load balancer. Which doesn't require much of a configuration.<br/>
       ![image](https://github.com/vibhordubey333/Kubernetes/assets/22407855/963896c3-6c05-438f-9aec-9d8b6c9649ab)<br/>






<br/>


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
 Pods are allocated a private IP address by default and cannot be reached outside of the cluster. You can use the kubectl port-forward command to map a local    port to a port inside the pod like this:<br/>
 `kubectl port-forward pod-name 8080:8080`

- Admin Commands
   - `minikube ip`
   -  Enable dashboard `minikube dashboard`
   - `kubectl get nodes`
   - `kubectl cluster-info`
   - `kubectl cluster-info dump`
   
- Secrets
   - `kubectl get secret`
   - `kubectl describe secret secret_name`
- Namespace
   - kubectl create ns
- Pods
   - To get pods/services/deployments in a namespace `kubectl get all -n namespace_name`
   - To get pods `kubectl get po`
</details>

 
<details>
<summary>Minikube Installation</summary>

- Install Minikube.
- Then start `minikube start driver=docker` [Docker should be installed. Install from APT instead of snapd] 
```
minikube version
minikube version: v1.27.1
commit: fe869b5d4da11ba318eb84a3ac00f336411de7ba
```
- Avoid using docker from snapd store. <br/>
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

- If everything is good. Then pods should be up in `ingress-nginx` namespace.
  
```
v@vibhor-:~ $ kubectl get ns
NAME              STATUS   AGE
default           Active   3m47s
ingress-nginx     Active   2m12s
kube-node-lease   Active   3m47s
kube-public       Active   3m47s
kube-system       Active   3m47s


osboxes@osboxes:~$ kubectl get po -n ingress-nginx
NAME                                        READY   STATUS      RESTARTS   AGE
ingress-nginx-admission-create-cxh7x        0/1     Completed   0          3m42s
ingress-nginx-admission-patch-6pb78         0/1     Completed   0          3m42s
ingress-nginx-controller-67fd4fc6fd-hfbcj   1/1     Running     0          3m42s
```

</details>

