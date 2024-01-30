1. To apply `kubectl apply -f cronjob.yaml`
2. To check cronjob
   
   ```
   vibhor@vibhor-virtualbox:~/code-repositories/Kubernetes/9.CronJob (master *%)$ kubectl get cronjob -n work
    NAME    SCHEDULE    SUSPEND   ACTIVE   LAST SCHEDULE   AGE
    hello   * * * * *   False     0        29s             45s
    ```
3.  To check pod which cronjob has created

```
vibhor@vibhor-virtualbox:~/code-repositories/Kubernetes/9.CronJob (master *%)$ kubectl get po -n work
NAME                           READY   STATUS      RESTARTS      AGE
demo-web-page-7d44d8fd-jztqq   1/1     Running     1 (18h ago)   43h
demo-web-page-7d44d8fd-rx4l5   1/1     Running     1 (18h ago)   43h
hello-28443395-42mcr           0/1     Completed   0             74s
hello-28443396-5khbq           0/1     Completed   0             14s
```

4. To check jobs in realtime:

```
vibhor@vibhor-virtualbox:~/code-repositories/Kubernetes (master)$ kubectl get jobs -n work --watch
NAME             COMPLETIONS   DURATION   AGE
hello-28443402   1/1           3m48s      4m48s
hello-28443405   1/1           15s        67s
hello-28443406   1/1           4s         48s
```

5. To check logs of cron jobs

```
vibhor@vibhor-virtualbox:~/code-repositories/Kubernetes/9.CronJob (master)$ kubectl get po -n work
NAME                           READY   STATUS      RESTARTS      AGE
demo-web-page-7d44d8fd-jztqq   1/1     Running     1 (19h ago)   43h
demo-web-page-7d44d8fd-rx4l5   1/1     Running     1 (19h ago)   43h
hello-28443428-kq8n9           0/1     Completed   0             114s
hello-28443429-zt42r           0/1     Completed   0             69s
hello-28443430-kxtc2           0/1     Completed   0             5s
```

```
vibhor@vibhor-virtualbox:~/code-repositories/Kubernetes/9.CronJob (master)$ kubectl logs hello-28443430-kxtc2 -n work
Tue Jan 30 09:10:06 UTC 2024
Hello from the Kubernetes cluster

```
   
### References:
  -  https://medium.com/@pranay.shah/how-to-get-logs-from-cron-job-in-kubernetes-last-completed-job-7957327c7e76

### Cron Issues
  -  https://stackoverflow.com/questions/75974501/kubernetes-finding-logs-of-a-failed-cronjob

