1. `kubectl apply -f ic-nginx-hn.yaml`
2. `curl http://192.168.49.2:80` or on browser `http://192.168.49.2:80`
3. Fire deployment
```
kubectl create deployment cheddar --image=errm/cheese:cheddar
kubectl create deployment stilton --image=errm/cheese:stilton
kubectl create deployment wensleydale --image=errm/cheese:wensleydale
```
4. Create service
```
kubectl expose deployment cheddar --port=80
kubectl expose deployment stilton --port=80
kubectl expose deployment wensleydale --port=80
```

5. Deploy Ingress `kubectl apply -f ingress.yaml`
6. Verify:  You should see image instead of 404
```
http://cheddar.192.168.49.2.nip.io/
http://stilton.192.168.49.2.nip.io/
http://wensleydale.192.168.49.2.nip.io/
```

