## Kubernetes Dashboard Ingress on Minikube

This example demonstrates how to set up an Ingress resource to access the Kubernetes Dashboard on minikube.

### Prerequisites
- minikube installed and running
- kubectl configured
- ingress-nginx addon enabled: `minikube addons enable ingress`

### Quick Start (Using Makefile)

```bash
# Show all available commands
make help

# Complete installation
make install

# Verify everything is working
make test

# Show access instructions
make access

# Add to /etc/hosts automatically
make add-hosts

# Cleanup when done
make cleanup
```

### Manual Setup Instructions

#### 0. Enable the dashboard addon:
```bash
minikube addons enable dashboard
minikube addons enable metrics-server
```

#### 1. Change ingress-nginx service to LoadBalancer (if not already):
```bash
kubectl patch svc ingress-nginx-controller -n ingress-nginx -p '{"spec": {"type": "LoadBalancer"}}'
```

#### 2. Apply the ingress:
```bash
kubectl apply -f dashboard-ingress.yaml
```

#### 3. Verify the ingress has an ADDRESS:
```bash
kubectl get ingress -n kubernetes-dashboard
```

Expected output:
```
NAME                CLASS                  HOSTS           ADDRESS          PORTS   AGE
dashboard-ingress   cluster-wide-ingress   dashboard.com   10.104.169.155   80      2m
```

#### 4. Access the dashboard:

**Option A**: Via NodePort (Recommended - no additional setup needed)
```bash
# Add to /etc/hosts (example with actual IP)
echo "192.168.64.3 dashboard.com" | sudo tee -a /etc/hosts

# Or use dynamic command
echo "$(minikube ip) dashboard.com" | sudo tee -a /etc/hosts

# Then visit: http://dashboard.com:31640
```

**Option B**: Via minikube tunnel (requires sudo password)
```bash
# In a separate terminal
sudo minikube tunnel

# Then access at: http://dashboard.com (port 80)
```

### Testing the Ingress

Test with curl before accessing via browser:
```bash
# Get minikube IP and NodePort
MINIKUBE_IP=$(minikube ip)
NODEPORT=$(kubectl get svc -n ingress-nginx ingress-nginx-controller -o jsonpath='{.spec.ports[0].nodePort}')

# Test the ingress
curl -H "Host: dashboard.com" http://$MINIKUBE_IP:$NODEPORT -I
```

Expected response: `HTTP/1.1 200 OK`

### Troubleshooting

**If ADDRESS field is empty:**
- Verify the backend service exists: `kubectl get svc -n kubernetes-dashboard`
- Check if the service is type LoadBalancer: `kubectl get svc -n ingress-nginx ingress-nginx-controller`
- Restart the ingress controller: `kubectl rollout restart deployment ingress-nginx-controller -n ingress-nginx`
- Check controller logs: `kubectl logs -n ingress-nginx -l app.kubernetes.io/component=controller --tail=50`

**To get the NodePort:**
```bash
kubectl get svc -n ingress-nginx ingress-nginx-controller
```
Look for the port mapping like `80:31640/TCP` - the 31640 is your NodePort.

**If getting 503 errors:**
- Verify dashboard pods are running: `kubectl get pods -n kubernetes-dashboard`
- Check endpoints exist: `kubectl get endpoints -n kubernetes-dashboard kubernetes-dashboard`
- Verify endpointslices: `kubectl get endpointslices -n kubernetes-dashboard`
- Restart ingress controller to force resync: `kubectl delete pod -n ingress-nginx -l app.kubernetes.io/component=controller`

**If ingress-nginx is not installed:**
```bash
# Enable the ingress addon
minikube addons enable ingress

# Wait for it to be ready
kubectl wait --for=condition=ready pod -l app.kubernetes.io/component=controller -n ingress-nginx --timeout=90s
```

### Cleanup
```bash
kubectl delete -f dashboard-ingress.yaml
minikube addons disable dashboard
```

### Makefile Commands Reference

The Makefile provides convenient commands for common operations:

**Installation:**
- `make install` - Complete installation (addons + ingress + verification)
- `make setup-addons` - Enable dashboard and metrics-server addons only
- `make patch-service` - Patch ingress-nginx service to LoadBalancer only
- `make apply-ingress` - Apply the ingress resource only

**Verification & Testing:**
- `make verify` - Check status of all resources
- `make status` - Detailed status including endpoints
- `make test` - Test ingress with curl
- `make access` - Show access instructions with current IP/port

**Utilities:**
- `make add-hosts` - Automatically add dashboard.com to /etc/hosts
- `make show-ip` - Display minikube IP
- `make show-nodeport` - Display the ingress NodePort
- `make logs` - Show ingress controller logs
- `make restart-controller` - Restart ingress controller (troubleshooting)

**Cleanup:**
- `make cleanup` - Remove ingress and disable addons
- `make clean-ingress` - Remove only the ingress resource
- `make clean-all` - Full cleanup including service patch revert

### Notes
- The ADDRESS field shows the LoadBalancer IP (typically the cluster IP like 10.104.169.155)
- With NodePort access, use the minikube IP and NodePort (typically 31640 for HTTP)
- With tunnel access, you can use port 80 directly
- The ingress routes traffic to the kubernetes-dashboard service in the kubernetes-dashboard namespace
- The ingress uses the `cluster-wide-ingress` IngressClass
- The Makefile automatically detects your minikube IP and NodePort for convenience