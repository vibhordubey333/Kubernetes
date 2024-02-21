#### Traefik

1. Replacing nginx with traefik.
2. Traefik is a proxy with built-in kubernetes ingress support.
3. Traefik  releases are named after cheese.
4. Pick between deployment and daemonset.
5. Implementing daemonset so that each node can accept connections.
6. We provide a YAML file which is essentially the sum of 
	-	Traefik Daemonset resources[patched with hostNetwork and tolerations]
	-	Traefik RBAC rules allowing it to watch necessary API objects.


#### Demo

1. Create resource
For Docker Desktop with LoadBalancer `kubectl apply -f ic-traefik-lb.yaml` 
For Minikube/Microk8s with hostNetwork `kubectl apply -f ic-traefik-hn.yaml`
2. Check pod status
`kubectl describe -n kube-system `
3. 

```
vibhor@vibhor-virtualbox:~/code-repositories/Kubernetes/17.Traefik (master %)$ kubectl describe -n kube-system ds/traefik-ingress-controller
Name:           traefik-ingress-controller
Selector:       k8s-app=traefik-ingress-lb
Node-Selector:  <none>
Labels:         k8s-app=traefik-ingress-lb
Annotations:    deprecated.daemonset.template.generation: 2
Desired Number of Nodes Scheduled: 1
Current Number of Nodes Scheduled: 1
Number of Nodes Scheduled with Up-to-date Pods: 1
Number of Nodes Scheduled with Available Pods: 0
Number of Nodes Misscheduled: 0
Pods Status:  0 Running / 1 Waiting / 0 Succeeded / 0 Failed
Pod Template:
  Labels:           k8s-app=traefik-ingress-lb
                    name=traefik-ingress-lb
  Service Account:  traefik-ingress-controller
  Containers:
   traefik-ingress-lb:
    Image:       traefik:v2.7
    Ports:       80/TCP, 8080/TCP
    Host Ports:  80/TCP, 8080/TCP
    Args:
      --api.dashboard
      --accesslog
      --api
      --api.insecure
      --log.level=INFO
      --metrics.prometheus
      --providers.kubernetesingress
      --entrypoints.http.Address=:80
      --providers.kubernetesingress.ingressendpoint.hostname=localhost
    Environment:  <none>
    Mounts:       <none>
  Volumes:        <none>
Events:
  Type    Reason            Age   From                  Message
  ----    ------            ----  ----                  -------
  Normal  SuccessfulCreate  29m   daemonset-controller  Created pod: traefik-ingress-controller-fvxxx
  Normal  SuccessfulDelete  29m   daemonset-controller  Deleted pod: traefik-ingress-controller-fvxxx
  Normal  SuccessfulCreate  29m   daemonset-controller  Created pod: traefik-ingress-controller-jzj7p

```
4. To test 
http://cheddar.192.168.49.2.nip.io/

5. Traefik Web UI
It is difficult for microk8s and minikube.
TODO: for minikube need to try couple of workarounds by editing deployment of ic-traefik-hn.yaml `kubectl edit -n kube-system ds/traefik-ingress-controller`
