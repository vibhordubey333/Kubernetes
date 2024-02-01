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

### Theory of Jobs
kubernetes job object basically deploys a pod but it runs for completion as opposed to objects like deployment, replicasets, replication controllers, and DaemonSets, which runs continuously.

For example, if you pass 100 as an argument, the shell script will echo the message 100 times and the container will exit.

#### TODO:
Read more about using completions and parallelism feature.


### Multiple Job Pods and Parallelism

When a job is deployed you can make it run on multiple pods with parallelism.

For example, in a job if you want to run 6 pods and run 2 pods in parallel, you need to add the following two parameters to your job manifest.

```
completions: 6
parallelism: 2
```

Sample yaml: https://devopscube.com/create-kubernetes-jobs-cron-jobs/

```
apiVersion: batch/v1
kind: Job
metadata:
  name: kubernetes-parallel-job
  labels:
    jobgroup: jobexample
spec:
  completions: 6
  parallelism: 2
  template:
    metadata:
      name: kubernetes-parallel-job
      labels:
        jobgroup: jobexample
    spec:
      containers:
      - name: c
        image: devopscube/kubernetes-job-demo:latest
        args: ["100"]
      restartPolicy: OnFailure
```


