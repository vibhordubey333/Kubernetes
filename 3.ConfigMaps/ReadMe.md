### ConfigMaps
A ConfigMap stores configuration settings that your Kubernetes pods consume.

#### Working of ConfigMap ?

A ConfigMap is a dictionary of key-value pairs that store configuration settings for your application.

ConfigMaps bind non-sensitive configuration artifacts such as configuration files, command-line arguments and environment variables to your pod containers and system components at runtime.
A ConfigMap separates your configuration from your pod and components which helps keep your workloads portable.
This maks their configuration easier to change and manage, and prevents hardcoding configuration data to pod specifications.

Environment variables can be saved in 3 manners.
    <br/> - Define the ConfigMap in YAML file.
    <br/> - Mount the ConfigMap through a volume.
    <br/> - Create the ConfigMap through a volume.

#### Run

##### Method 1  - Defining ConfigMap In Deployment File.

1. Build docker image `make build`
2. Create namespace `kubctl create ns work`
3. Deploy pods `kubectl apply -f file_name.yaml -n work`
4. Check status `kubectl get po -n work -o wide`
5. Use IP address from step 4 and `curl -XGET http://ip:4444/hello`

```
~/Documents$ curl http://172.17.0.6:4444/hello
Hello, World! Name:Abracadabara %!(EXTRA string= Country:%s, string=)
```
Environment variable `Name:Abracadabara`  is coming deployment file `go-deployment.yaml` 

6. If we describe our pod we can see the environment variable.

```
osboxes@osboxes:~/Documents$ kubectl describe po go-deployment-cfd5fffd8-b4wj7 -n work1
Name:             go-deployment-cfd5fffd8-b4wj7
Namespace:        work1
Priority:         0
Service Account:  default
Node:             osboxes/192.168.1.6
Start Time:       Sat, 29 Oct 2022 13:20:41 -0400
Labels:           app=api
                  pod-template-hash=cfd5fffd8
Annotations:      <none>
Status:           Running
IP:               172.17.0.6
IPs:
  IP:           172.17.0.6
Controlled By:  ReplicaSet/go-deployment-cfd5fffd8
Containers:
  api:
    Container ID:   docker://8dc95cda2d69cbe625b22bdc5356dde40f48dfc1dd2c3c63dd4f9e695357d6af
    Image:          dockerenv:latest
    Image ID:       docker://sha256:9bc25262aca096c871d51145783a9e25437d4169e62543eb4052fd5336b99ccb
    Port:           4444/TCP
    Host Port:      0/TCP
    State:          Running
      Started:      Sat, 29 Oct 2022 13:20:45 -0400
    Ready:          True
    Restart Count:  0
    Environment:
      name:  Abracadabara
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-frzpz (ro)
Conditions:
  Type              Status
  Initialized       True
  Ready             True
  ContainersReady   True
  PodScheduled      True
Volumes:
  default-token-frzpz:
    Type:        Secret (a volume populated by a Secret)
    SecretName:  default-token-frzpz
    Optional:    false
QoS Class:       BestEffort
Node-Selectors:  <none>
Tolerations:     node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                 node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:          <none>
```
##### Method 2  - Defining ConfigMap In Seperate Deployment File Of Type ConfigMap.

7. Navigate to dir `3.ConfigMaps/ConfigMap_From_Manifests`
8. Create namespace if it doesn't exists `kubectl create ns work1`
9. Create ConfigMap using `kubectl create -f manifests/cm-country.yaml -n work1`
10. Verify created ConfigMap
`kubectl get cm -n work1`
```
Kubernetes/3.ConfigMaps/ConfigMap_From_Manifests$ kubectl get cm -n work1
NAME               DATA   AGE
manifest-country   1      2m2s
manifest-name      1      114s
```
11. To see data defined in ConfigMap `kubectl get configmap manifest-example -n work1 -o yaml`
```
Kubernetes/3.ConfigMaps/ConfigMap_From_Manifests$ kubectl get configmap manifest-country -n work1 -o yaml
apiVersion: v1
data:
  country: Country_Config-Map_From_Manifest
kind: ConfigMap
metadata:
  creationTimestamp: "2022-11-01T04:18:49Z"
  name: manifest-country
  namespace: work1
  resourceVersion: "43535"
  selfLink: /api/v1/namespaces/work1/configmaps/manifest-country
  uid: 4285d475-21a8-4497-88c9-2140c9de4d46

```
```
kubectl get configmap manifest-name -n work1 -o yaml
apiVersion: v1
data:
  name: Name_Config-Map_From_Manifest
kind: ConfigMap
metadata:
  creationTimestamp: "2022-11-01T04:18:57Z"
  name: manifest-name
  namespace: work1
  resourceVersion: "43542"
  selfLink: /api/v1/namespaces/work1/configmaps/manifest-name
  uid: 67debc0b-4941-46ef-8b86-941752cde863

```

#### References

1. https://matthewpalmer.net/kubernetes-app-developer/articles/configmap-example-yaml.html
2. https://cloud.google.com/kubernetes-engine/docs/concepts/configmap#:~:text=ConfigMaps%20bind%20non%2Dsensitive%20configuration,helps%20keep%20your%20workloads%20portable.
3. https://www.weave.works/blog/kubernetes-configmap
4. https://www.bmc.com/blogs/kubernetes-configmap/
5. https://github.com/jetstack/kubernetes.github.io/blob/master/docs/tasks/configure-pod-container/configure-pod-configmap.md