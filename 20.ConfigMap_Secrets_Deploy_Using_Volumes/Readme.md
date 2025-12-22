# Apply the files:
kubectl apply -f app-config.yaml
kubectl apply -f app-secret.yaml
kubectl apply -f pod-volumes.yaml

# List the files in the config directory
kubectl exec hvac-worker -- ls /etc/config
# Output: settings.conf

# Read the content of the config file
kubectl exec hvac-worker -- cat /etc/config/settings.conf

# Read the secret key
kubectl exec hvac-worker -- cat /etc/keys/api-key

# To decode the encrypted secret
echo your_encrypted_secret_output | base64 --decode