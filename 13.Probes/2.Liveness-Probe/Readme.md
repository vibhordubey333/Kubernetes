#### Liveness Probe ?

Allow Kubernetes to check if your app is alive. The kubelet agent that runs on each node uses the liveness probes to ensure that the containers are running as expected. If a container app is no longer serving requests, kubelet will intervene and restart the container.

For example, if an application is not responding and cannot make progress because of a deadlock, the liveness probe detects that it is faulty. Kubelet then terminates and restarts the container. Even if the application carries defects that cause further deadlocks, the restart will increase the container's availability. It also gives your developers time to identify the defects and resolve them later.

Startup probe forces liveness and readiness checks to wait.


### References

	-	https://newrelic.com/blog/how-to-relic/kubernetes-health-checks