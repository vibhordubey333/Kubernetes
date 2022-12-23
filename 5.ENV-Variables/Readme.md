### Concepts Used: 
1. Exposing deployent using `LoadBalancer` Service.
2. Env use.
### To Run

1. Build docker image `docker build -t kub-data-demo .`
2. Deploy pv,pvc,service and deployment file.
    - `kubectl apply -f environment.yaml`
    - `kubectl apply -f=deployment.yaml -f=service.yaml` 
3. Pod will go into crash then manually edit the deployment and Make RestartPolicy as Never. Because we've not pushed our image to private/public repo.
    - `kubectl edit deployment deployment-name`
3. To check port
    - `kubectl get svc`
    - `kubectl get pv`
    - `kubectl get pvc`


### APIs

1. `curl -XGET http://localhost:31721/story`
Output: {"story":""}

2. `curl -X POST http://localhost:31721/story -H 'Content-Type: application/json' -d '{"text":"My Text"}'`
Output: {"message":"Text was stored!"}

3. Crashing pod intentionally. 
    `curl -XGET http://localhost:31721/error`
