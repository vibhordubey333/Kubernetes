
#downloaad kubebuilder and install locally.
0. curl -L -o kubebuilder "https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)"
1. chmod +x kubebuilder && sudo mv kubebuilder /usr/local/bin/
2. kubebuilder init --domain my.domain --repo my.domain/guestbook
3. make manifests
4. make install
5. make run
6. kubectl apply -k config/samples/
---

Reference:
https://book.kubebuilder.io/quick-start

