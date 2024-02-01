1. Create yaml. <br/>
   It will echo "hello world"
```
apiVersion: batch/v1
kind: Job
metadata:
  name: hello-world-job
spec:
  template:
    metadata:
      name: hello-world-job
    spec:
      containers:
      - name: hello-world
        image: centos:7
        command:
         - "bin/bash"
         - "-c"
         - "echo hello world"
      restartPolicy: Never
```

2. `kubectl apply -f hello_world_job.yaml`
3. `kubectl get jobs`
NAME              COMPLETIONS   DURATION   AGE
hello-world-job   0/1           6s         6s

4. `kubectl get po`
```
vibhor@vibhor-virtualbox:~/code-repositories/Kubernetes (master)$ kubectl get po
NAME                      READY   STATUS      RESTARTS       AGE
hello-world-job-r57d2     0/1     Completed   0              18m

```
5. Checking logs
```
kubectl logs hello-world-job-r57d2
hello world

```
