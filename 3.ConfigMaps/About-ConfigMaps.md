### To Run

1. Build docker image `docker build -t kub-data-demo .`
2. Deploy pv,pvc,service and deployment file.
    - `kubectl apply -f environment.yaml`
    - `kubectl apply -f=deployment.yaml -f=service.yaml` 
3. Pod will go into crash then manually edit the deployment and Make RestartPolicy as Never. Because we've not pushed our image to private/public repo.
    - `kubectl edit deployment deployment-name`

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

#### References

1. https://matthewpalmer.net/kubernetes-app-developer/articles/configmap-example-yaml.html
2. https://cloud.google.com/kubernetes-engine/docs/concepts/configmap#:~:text=ConfigMaps%20bind%20non%2Dsensitive%20configuration,helps%20keep%20your%20workloads%20portable.
3. https://www.weave.works/blog/kubernetes-configmap
4. https://www.bmc.com/blogs/kubernetes-configmap
