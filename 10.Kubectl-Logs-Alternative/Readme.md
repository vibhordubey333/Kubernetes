### Limitation With kubectl-logs
You can see that kubectl can read logs from multiple pods using label selector but there is a limitation of this solution. Firstly, logs from different pods are mixed, which prohibits you from knowing which log line came from which pod. Secondly, it doesn’t work in tail mode (using –follow (-f)). This is because –follow streams the logs from the API server. You open a connection to the API server per pod, which will open a connection to the corresponding kubelet itself to stream the logs continuously. This does not scale well and translates to a lot of inbound and outbound connections to the API server, therefore, it became a design decision to limit multiple connections. So you can either stream the logs of one pod, or select a bunch of pods at the same time without streaming.
source: https://blog.knoldus.com/read-k8s-logs-using-stern/

### Official Repo
https://github.com/stern/stern

### Binary
-  tar -C /usr/local /usr/bin/ -xzf tar_ball_name
-  https://github.com/stern/stern/releases/download/v1.28.0/stern_1.28.0_linux_amd64.tar.gz
