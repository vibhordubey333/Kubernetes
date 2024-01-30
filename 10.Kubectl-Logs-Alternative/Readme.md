### Limitation With kubectl-logs
You can see that kubectl can read logs from multiple pods using label selector but there is a limitation of this solution. Firstly, logs from different pods are mixed, which prohibits you from knowing which log line came from which pod. Secondly, it doesn’t work in tail mode (using –follow (-f)). This is because –follow streams the logs from the API server. You open a connection to the API server per pod, which will open a connection to the corresponding kubelet itself to stream the logs continuously. This does not scale well and translates to a lot of inbound and outbound connections to the API server, therefore, it became a design decision to limit multiple connections. So you can either stream the logs of one pod, or select a bunch of pods at the same time without streaming.
source: https://blog.knoldus.com/read-k8s-logs-using-stern/

![image](https://github.com/vibhordubey333/Kubernetes/assets/22407855/7bc8254a-0c6a-46f6-bc01-ac1a2386a081)


Note: Potential downside of using tool like stern is it could overload the API server if many pods are running.

### Official Repo
https://github.com/stern/stern

### Binary
-  tar -C /usr/local /usr/bin/ -xzf tar_ball_name
-  https://github.com/stern/stern/releases/download/v1.28.0/stern_1.28.0_linux_amd64.tar.gz
-  More flags: `stern demo-web-page-7d44d8fd-jztqq -n work --tail 1 --timestamps`
-  TODO: Explore, selectors/labels
