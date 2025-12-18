.PHONY: help install setup-addons patch-service apply-ingress verify access test status cleanup clean-ingress clean-all

# Default target
help:
	@echo "Kubernetes Dashboard Ingress - Makefile Commands"
	@echo "=================================================="
	@echo ""
	@echo "Installation:"
	@echo "  make install          - Complete installation (addons + ingress)"
	@echo "  make setup-addons     - Enable dashboard and metrics-server addons"
	@echo "  make patch-service    - Patch ingress-nginx service to LoadBalancer"
	@echo "  make apply-ingress    - Apply the ingress resource"
	@echo ""
	@echo "Verification:"
	@echo "  make verify           - Verify installation status"
	@echo "  make status           - Show current status of all resources"
	@echo "  make test             - Test the ingress with curl"
	@echo "  make access           - Show how to access the dashboard"
	@echo ""
	@echo "Cleanup:"
	@echo "  make cleanup          - Complete cleanup (remove ingress + disable addons)"
	@echo "  make clean-ingress    - Remove only the ingress resource"
	@echo "  make clean-all        - Full cleanup including service patch revert"
	@echo ""
	@echo "Utilities:"
	@echo "  make add-hosts        - Add dashboard.com to /etc/hosts"
	@echo "  make show-nodeport    - Show the NodePort for ingress-nginx"
	@echo "  make show-ip          - Show minikube IP address"
	@echo "  make logs             - Show ingress controller logs"

# Complete installation
install: setup-addons patch-service apply-ingress verify
	@echo ""
	@echo "Installation complete."
	@echo ""
	@$(MAKE) access

# Enable dashboard and metrics-server addons
setup-addons:
	@echo "Enabling Kubernetes Dashboard addon..."
	minikube addons enable dashboard
	@echo ""
	@echo "Enabling metrics-server addon..."
	minikube addons enable metrics-server
	@echo ""
	@echo "Waiting for dashboard pods to be ready..."
	@kubectl wait --for=condition=ready pod -l k8s-app=kubernetes-dashboard -n kubernetes-dashboard --timeout=90s 2>/dev/null || true
	@echo "Addons enabled successfully."

# Patch ingress-nginx service to LoadBalancer type
patch-service:
	@echo "Patching ingress-nginx service to LoadBalancer type..."
	@kubectl patch svc ingress-nginx-controller -n ingress-nginx -p '{"spec": {"type": "LoadBalancer"}}' 2>/dev/null || \
		echo "Warning: Could not patch service (may not exist or already patched)"
	@echo "Service patched."

# Apply the ingress resource
apply-ingress:
	@echo "Applying dashboard ingress..."
	kubectl apply -f dashboard-ingress.yaml
	@echo "Ingress applied."

# Verify installation
verify:
	@echo ""
	@echo "Verifying installation..."
	@echo ""
	@echo "Dashboard Pods:"
	@kubectl get pods -n kubernetes-dashboard
	@echo ""
	@echo "Dashboard Services:"
	@kubectl get svc -n kubernetes-dashboard
	@echo ""
	@echo "Ingress Resources:"
	@kubectl get ingress -n kubernetes-dashboard
	@echo ""
	@echo "Ingress Controller Service:"
	@kubectl get svc -n ingress-nginx ingress-nginx-controller
	@echo ""

# Show complete status
status: verify
	@echo "Endpoints:"
	@kubectl get endpoints -n kubernetes-dashboard kubernetes-dashboard
	@echo ""
	@echo "Ingress Details:"
	@kubectl describe ingress dashboard-ingress -n kubernetes-dashboard | head -20

# Test the ingress with curl
test:
	@echo "Testing ingress with curl..."
	@MINIKUBE_IP=$$(minikube ip); \
	NODEPORT=$$(kubectl get svc -n ingress-nginx ingress-nginx-controller -o jsonpath='{.spec.ports[0].nodePort}'); \
	echo "Testing: http://$$MINIKUBE_IP:$$NODEPORT with Host: dashboard.com"; \
	echo ""; \
	curl -H "Host: dashboard.com" http://$$MINIKUBE_IP:$$NODEPORT -I

# Show access instructions
access:
	@echo "Access Instructions:"
	@echo "===================="
	@echo ""
	@MINIKUBE_IP=$$(minikube ip); \
	NODEPORT=$$(kubectl get svc -n ingress-nginx ingress-nginx-controller -o jsonpath='{.spec.ports[0].nodePort}' 2>/dev/null || echo "31640"); \
	echo "Option 1: Via NodePort (Recommended)"; \
	echo "-------------------------------------"; \
	echo "1. Add to /etc/hosts:"; \
	echo "   echo \"$$MINIKUBE_IP dashboard.com\" | sudo tee -a /etc/hosts"; \
	echo ""; \
	echo "2. Access in browser:"; \
	echo "   http://dashboard.com:$$NODEPORT"; \
	echo ""; \
	echo "Option 2: Via minikube tunnel"; \
	echo "------------------------------"; \
	echo "1. Run in separate terminal:"; \
	echo "   sudo minikube tunnel"; \
	echo ""; \
	echo "2. Access in browser:"; \
	echo "   http://dashboard.com"; \
	echo ""; \
	echo "Quick test with curl:"; \
	echo "   curl -H \"Host: dashboard.com\" http://$$MINIKUBE_IP:$$NODEPORT"

# Add dashboard.com to /etc/hosts
add-hosts:
	@echo "Adding dashboard.com to /etc/hosts..."
	@MINIKUBE_IP=$$(minikube ip); \
	if grep -q "dashboard.com" /etc/hosts; then \
		echo "Note: dashboard.com already exists in /etc/hosts"; \
	else \
		echo "$$MINIKUBE_IP dashboard.com" | sudo tee -a /etc/hosts; \
		echo "Added $$MINIKUBE_IP dashboard.com to /etc/hosts"; \
	fi

# Show NodePort
show-nodeport:
	@echo "NodePort for ingress-nginx:"
	@kubectl get svc -n ingress-nginx ingress-nginx-controller -o jsonpath='{.spec.ports[0].nodePort}' && echo ""

# Show minikube IP
show-ip:
	@echo "Minikube IP:"
	@minikube ip

# Show ingress controller logs
logs:
	@echo "Ingress Controller Logs (last 50 lines):"
	@kubectl logs -n ingress-nginx -l app.kubernetes.io/component=controller --tail=50

# Remove only the ingress resource
clean-ingress:
	@echo "Removing dashboard ingress..."
	@kubectl delete -f dashboard-ingress.yaml 2>/dev/null || echo "Warning: Ingress not found"
	@echo "Ingress removed."

# Complete cleanup
cleanup: clean-ingress
	@echo ""
	@echo "Disabling dashboard addon..."
	@minikube addons disable dashboard 2>/dev/null || echo "Warning: Dashboard addon not enabled"
	@echo ""
	@echo "Disabling metrics-server addon..."
	@minikube addons disable metrics-server 2>/dev/null || echo "Warning: Metrics-server addon not enabled"
	@echo ""
	@echo "Cleanup complete."
	@echo ""
	@echo "Note: /etc/hosts entry for dashboard.com was NOT removed."
	@echo "To remove it manually, edit /etc/hosts and delete the dashboard.com line."

# Full cleanup including service patch revert
clean-all: cleanup
	@echo ""
	@echo "Reverting ingress-nginx service to NodePort..."
	@kubectl patch svc ingress-nginx-controller -n ingress-nginx -p '{"spec": {"type": "NodePort"}}' 2>/dev/null || \
		echo "Warning: Could not revert service (may not exist)"
	@echo "Full cleanup complete."

# Restart ingress controller (useful for troubleshooting)
restart-controller:
	@echo "Restarting ingress controller..."
	kubectl delete pod -n ingress-nginx -l app.kubernetes.io/component=controller
	@echo "Waiting for new pod to be ready..."
	@kubectl wait --for=condition=ready pod -l app.kubernetes.io/component=controller -n ingress-nginx --timeout=90s 2>/dev/null || true
	@echo "Controller restarted."

# Enable ingress addon if not already enabled
enable-ingress:
	@echo "Enabling ingress-nginx addon..."
	minikube addons enable ingress
	@echo "Waiting for ingress controller to be ready..."
	@kubectl wait --for=condition=ready pod -l app.kubernetes.io/component=controller -n ingress-nginx --timeout=90s 2>/dev/null || true
	@echo "Ingress addon enabled."

