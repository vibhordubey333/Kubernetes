### Problem-Statement
Our goal here will be to create a service that load balances connections to two different deployments. You might use this as a simplistic way to run two versions of your apps in parallel for a period of time.

In the real world, you'll likely use a 3rd party load balancer to provide advanced blue/green or canary-style deployments, but this assignment will help further your understanding of how service selectors are used to find pods and use them as service endpoints.

For simplicity, version 1 of our application will be using the NGINX image, and version 2 of our application will be using the Apache image. They both listen on port 80 by default and have different HTML by default so that it's easy to distinguish which is being accessed.

Once properly set up, when we connect to the service we expect to see some requests being served by NGINX and some requests being served by Apache.

Objectives:
We need to create two deployments: one for v1 (NGINX), another for v2 (Apache).

They will be exposed through a single service.

The selector of that service will need to match the pods created by *both* deployments.

For that, we will need to change the deployment specification to add an extra label, to be used solely by the service.

That label should be different from the pre-existing labels of our deployments, otherwise, our deployments will step on each other's toes.

We're not at the point of writing our own YAML from scratch, so you'll need to use the kubectl edit command to modify existing resources.

Questions for this assignment
What commands did you use to perform the following?

1. Create a deployment running one pod using the official NGINX image.

2. Expose that deployment with a ClusterIP service.

3. Check that you can successfully connect to the exposed service.

What commands did you use to perform the following?

1. Change (edit) the service definition to use a label/value of myapp: web

2. Check that you cannot connect to the exposed service anymore.

3. Change (edit) the deployment definition to add that label/value to the pod template.

4. Check that you *can* connect to the exposed service again.

What commands did you use to perform the following?

1. Create a deployment running one pod using the official Apache image (httpd).

2. Change (edit) the deployment definition to add the label/value picked previously.

3. Connect to the exposed service again.

(It should now yield responses from both Apache and NGINX.)

TL;DR <br/>
We're having single service to expose our two deployments.

### Solution

Using declarative approach.

1. Create a deployment running one pod using the official NGINX image.
    `kubectl create deployment v1-nginx --image=nginx`

2. Expose that deployment with a ClusterIP service.
    `kubectl expose deployment v1-nginx --port=80 or kubectl create service clusterip v1-nginx --tcp=80`

3. Check that you can successfully connect to the exposed service.
    Access the pod. List the pod `kubectl get po`
   ```
   Avoid this as it will be also reachable inside.
   vibhor@vibhor-virtualbox:$ kubectl exec -it v1-nginx-7fbd5f74f-qn4wj -- sh
    # curl localhost

   ```

or
    
```   
   Using sh pod
   kubectl get svc v1-nginx
   curl http://A.B.C.D
```

or
  
  ```
    vibhor@vibhor-virtualbox:$ kubectl run --restart=Never --image=alpine -it --rm testcontainer1
    If you don't see a command prompt, try pressing enter.

     # apk add curl
     # curl v1-nginx[This is the pod name]
   ```

4. Change (edit) the service definition to use a label/value of myapp: web <br/>
   Edit the YAML manifest of the service with `kubectl edit service v1-nginx` . Look for the 
   selector: section, and change app: v1-nginx to myapp: web. Make sure to change the 
  selector: section, not the labels: section! After making the change, save and quit.

5. Check that you cannot connect to the exposed service anymore.
    `Repeat steps mentioned on step 3`
6. Change (edit) the deployment definition to add that label/value to the pod template.
    `Done on step 4`
7. Check that you *can* connect to the exposed service again.
   ```
   vibhor@vibhor-virtualbox:$ kubectl run --restart=Never --image=alpine -it --rm testcontainer1
   If you don't see a command prompt, try pressing enter.

   # apk add curl
   # curl v1-nginx[This is the pod name]
   ```
8.  Create a deployment running one pod using the official Apache image (httpd).
    `kubectl create deployment v2-apache --image=httpd`
9.  Change (edit) the deployment definition to add the label/value picked previously.
    
    ```
    Same as previously: kubectl edit deployment v2-apache, then add the label myapp: web below 
    app: v2-apache. Again, make sure to change the labels in the pod template, not of the deployment 
    itself.
    ```
10. Connect to the exposed service again.

## TODO:

Need to debug, why httpd is not accessible after adding selector.
