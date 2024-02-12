
0. Startup probes

A startup probe is used to determine whether a container is ready to accept traffic. It differs from liveness and readiness probes in that it runs only when a container first starts, and it's disabled after the first successful check. This is useful for applications that take some time to initialize before they are ready to serve traffic.
Startup probes are used to determine when a container application has been initialized successfully. If a startup probe fails, the pod is restarted. When pod containers take too long to become ready, readiness probes may fail repeatedly. In this case, containers risk being terminated by kubelet before they are up and running. This is where the startup probe comes to the rescue.

The startup probe forces liveness and readiness checks to wait until it succeeds so that the application startup is not compromised. That is especially beneficial for slow-starting legacy applications.

1. Three types of Startup Probes are there
	-	HTTP Get Startup Probe
		Consider this example where the startup probe is defined in a Deployment. The probe uses an HTTP GET request to the /healthz endpoint on the liveness-port. If the endpoint returns a success code, the container is considered alive.

	-	Exec Startup Probe
		The startup probe is a shell script that uses the cat command to read the /tmp/healthy file. If the cat command returns a zero exit status, the container is considered alive.

	-	TCP Socket Startup Probe
		The startup probe checks if the TCP port 3000 is open. If the port is open, the container is considered alive.

2. To apply yaml file `kubectl apply -f yaml_file_name`

3. To check status 
	-	`kubectl get po`
	-	`kubectl describe po liveness-exec-failure`

4. To check pod events
	-	`kubectl events liveness-exec-failure`

---

-	References
	-	https://kubebyexample.com/concept/health-checks
	-	https://www.linkedin.com/pulse/understanding-kubernetes-startup-probes-examples-use-cases-feyz-sari-bc2de/
	-	https://newrelic.com/blog/how-to-relic/kubernetes-health-checks