## Concept

**What is headless service ?** 

A headless service is a service with a service IP but instead of load-balancing it will return the IPs of our associated Pods. This allows us to interact directly with the Pods instead of a proxy. It's as simple as specifying `None` for `.spec.clusterIP` and can be utilized with or without selectors - you'll see an example with selectors in a moment.


When deploying services, you have the possibility to set three different service types to specify which kind of service you want. These are: </br>
**ClusterIP:** the default type to expose the server only on an cluster-internal ip address
**NodePort:** allows to expose the service through a static port on the node.
**LoadBalancer:** allows to expose the service using a cloud provider’s external load balancer.</br>
To avoid requests being load-balanced behind one single ip address, we can explicitly specifying “None” for the cluster IP when a single ip address is not desired. Kubernetes won’t allocate any IP address to the service. This type of service is termed as headless service.

For instance:
```
apiVersion: v1
kind: Service
metadata:
  name: my-headless-service
spec:
  clusterIP: None # <-- Headless Service
  selector:
    app: test-app
  ports:
    - protocol: TCP
      port: 80
      targetPort: 3000 
```

## Run

1. Build Dockerimage
2. Deploy pod.
3. Deploy normal service and headless service.
4. After deploying headless service, enter into pod `kubectl exec -it pod_name -n work -- sh` then do `nslookup servicename`. It will return the IP address.
5. Fire `curl http://ip_address_from_above_step:80`

<details>
<summary>References</summary>
A. https://medium.com/swlh/discovering-running-pods-by-using-dns-and-headless-services-in-kubernetes-7002a50747f4 </br>
B. https://stackoverflow.com/questions/52707840/what-is-a-headless-service-what-does-it-do-accomplish-and-what-are-some-legiti</br>
C. Kubernetes in Action By Marko Luksa
</details>
