# Kubernetes Objects

---

- ### `Pods` :-
  - It is the smallest unit in Kubernetes. 
  - It is a group of one or more containers and a specification telling how the containers will be running inside the Kubernetes cluster.
  - A pod gets its own IP address.

  ![](https://imgur.com/dNqrrkc.png)
  
  ```bash
    # to get the running pods inside the cluster
    $ kubectl get pods
    
    # to create a pod with declarative way (yaml manifest/json file will be applied)
    $ kubectl create -f sample-pod.yaml
    pod/sample-pod created
    
    # to get the pod list with more details (`-o wide` means output wide)
    $ kubectl get pods -o wide
    
    # the imperative way to create a pod (if you dont specify the namespace in the yaml file then it will gets created in the default namespace)
    $ kubectl run demo --image=nginx --port=80
    pod/demo created
    
    # to get the logs/history of of a pod
    $ kubectl logs demo
    /docker-entrypoint.sh: /docker-entrypoint.d/ is not empty, will attempt to perform configuration
    /docker-entrypoint.sh: Looking for shell scripts in /docker-entrypoint.d/
    /docker-entrypoint.sh: Launching /docker-entrypoint.d/10-listen-on-ipv6-by-default.sh
    10-listen-on-ipv6-by-default.sh: info: Getting the checksum of /etc/nginx/conf.d/default.conf
    10-listen-on-ipv6-by-default.sh: info: Enabled listen on IPv6 in /etc/nginx/conf.d/default.conf
    /docker-entrypoint.sh: Launching /docker-entrypoint.d/20-envsubst-on-templates.sh
    /docker-entrypoint.sh: Launching /docker-entrypoint.d/30-tune-worker-processes.sh
    /docker-entrypoint.sh: Configuration complete; ready for start up
    2022/10/20 15:06:18 [notice] 1#1: using the "epoll" event method
    2022/10/20 15:06:18 [notice] 1#1: nginx/1.23.2
    2022/10/20 15:06:18 [notice] 1#1: built by gcc 10.2.1 20210110 (Debian 10.2.1-6)
    2022/10/20 15:06:18 [notice] 1#1: OS: Linux 4.19.202
    2022/10/20 15:06:18 [notice] 1#1: getrlimit(RLIMIT_NOFILE): 1048576:1048576
    2022/10/20 15:06:18 [notice] 1#1: start worker processes
    2022/10/20 15:06:18 [notice] 1#1: start worker process 31
    2022/10/20 15:06:18 [notice] 1#1: start worker process 32
    
    # to describe a pod and get to see all the kubelet events
    $ kubectl describe pod demo
    Events:
    Type    Reason     Age    From               Message
    ----    ------     ----   ----               -------
    Normal  Scheduled  4m33s  default-scheduler  Successfully assigned default/demo to minikube
    Normal  Pulling    4m33s  kubelet            Pulling image "nginx"
    Normal  Pulled     4m29s  kubelet            Successfully pulled image "nginx" in 3.851198621s
    Normal  Created    4m29s  kubelet            Created container demo
    Normal  Started    4m29s  kubelet            Started container demo
    
    # to delete a pod
    $ kubectl delete pod demo
    pod "demo" deleted
    
    # to get a shell inside a pod like docker -exec
    $ kubectl exec -it sample-pod bash
  ```
  
  ---
  
   - ### Lifecycle of a Pod :- 
     - When we give a command, `kubectl apply -f` or `create -f`, and provide a YAML file, it's first converted to JSON and sent to API Server. Here, the request is authenticated using our kube config credentials, and then it's authorized whether the user is actually authorized to perform this particular command to create a pod. Then if any admission controllers are there, they are checked before it gets persisted to etcd datastore.
       Now, after it is persisted in the etcd datastore, it goes to pending state. Then the scheduler pitches in and tries to find the best match for the node where it has to be spawned. It will go through the cluster and find the best fit node based on the resources, etc., and if the image is already present on the node because that has some reference. Once it gets the node, it will fill the label spec node name and send it to the API server, and now, that particular request is also stored in etcd.
     
   - ### API server in a pod's lifecycle :- 
     - Now, the API server instructs the kubelet, "Hey, there is one pod that has to be spawned on this particular node." The kubelet is responsible for fetching the image from the image registry. It can be any image registry, and then the CRI does that, and the CNI will get the IP attached to the pod, and then that particular IP will be sent again back to the API Server, then again, it is stored in etcd. From there, it goes to ContainerCreating when the scheduler selects the node, and after that, when everything is there, the image is pulled, and then it goes to the running state. 
       We can see the state when we use the `kubectl get pods` command.
       
   - ### Probes and checks :- 
     - Now there can be other advanced things that happen, or whenever the process dies too many times within a pod, it can also go to crashloopbackoff, and whenever it is succeeded, it will be in the succeeded state. Another advanced kind of lifecycle includes liveliness, so we can have a /health and a /ready. These are known as liveness and readiness probes. 
       The pod will keep on checking that, and if it fails, it can lead to the crashloopbackoff.
       
   - ### Hooks and its usages :- 
     - There are some of the hooks that can be implemented. For example, if there are some actions that you want to perform after the main container just starts, you can have a post-start hook, and if you want to perform before the main container gets terminated, you will have to have a pre-stop hook. There is also something called the init container that will be running before actually starting the main container.

   ![](https://imgur.com/njNTvAb.png)

---

- ### `Init Containers` :- 
  - Init containers are just like regular containers, but they run to completion and run before the main container starts. There are various use cases for init         containers, such as init containers can contain certain utilities or custom code for the setup that is not present for the app container. 
  - It can also change the file system before starting the app container based on certain logic. It can also be used to limit the attack surface by keeping certain     tools as part of init containers. Init containers can also delay the start of the main containers by having some precondition check; unless they are met, it         will keep trying the init containers.
  
- ### `Container Probes` :-
  - The kubelet uses `liveness probes` to know when to restart a container. For example, liveness probes could catch a deadlock, where an application is running,       but unable to make progress. Restarting a container in such a state can help to make the application more available despite bugs.
  - The kubelet uses `readiness probes` to know when a container is ready to start accepting traffic. A Pod is considered ready when all of its containers are           ready.One use of this signal is to control which Pods are used as backends for Services. When a Pod is not ready, it is removed from Service load balancers.
  - The kubelet uses `startup probes` to know when a container application has started. If such a probe is configured, it disables liveness and readiness checks         until it succeeds, making sure those probes don't interfere with the application startup. This can be used to adopt liveness checks on slow starting containers,     avoiding them getting killed by the kubelet before they are up and running.

  ![](https://imgur.com/hIMoWKZ.png)

---

- ### `Deployment` :-
   - Creating Deployment by command method
     
     ```bash
     kubectl create deployment demo --image=nginx --replicas=3 --port=80
     ```
     it will create a deployment named "demo" with 3 replicas and image is nginx:latest
   - Get the deployment list

     ```bash
     kubectl get deploy
     ```
   - We can see rolling out of a deployment, and we can check the status using this command
   
     ```bash
     kubectl rollout status deployment demo
     ```
   - Setting an image in a deployment
   
     ```bash
     kubectl set image deployment/demo --nginx=nginx:1.15.0 --record
     ```
   - What happens when a wrong image is set in a deployment?
   
     ```bash
     kubectl set image deployment/demo --nginx=nginx:1.15.abc --record
     ```
     we can see that the pods are running, but it created an extra pod. So the max surge came into the picture, creating an extra pod, and the image ErrPull came.        Hence, this pod never got ready, and that's why the previous one never got terminated.
   - Rollbacking from a wrong image deployment
     - We can first see the status through the 
     
     ```bash
     kubectl rollout history deployment demo
     ```
     
     ```bash
     kubectl rollout undo deployment demo --to-revision=2
     ```
     
   - Scaling a deployment
    
     ```bash
     kubectl scale deployment demo --replicas=5
     ```
   - Delete and Edit a deployment
   
     ```bash
     kubectl delete deployment demo
     ```
     
     ```bash
     kubectl edit deployment demo
     ```
---
 - ### `StatefulSets` :- 
    - StatefulSet is a Kubernetes object which is used for Stateful applications such as databases. However, there are cases just like databases where we need           persistence, or we need to store the state of the application, which Deployments cannot serve.

    ![](https://imgur.com/W0U69tJ.png)
    - Creating and scaling up a StatefulSet 
       - We will use the `kubectl get storageclass` command, and we can see that we are using a rancher local path provisioner, which is very simple to install with          the YAML file. [`checkout`](https://github.com/rancher/local-path-provisioner)
       - now create the statefulset [statefulset.yaml](https://github.com/bishal7679/Learning-Cloud-Native/blob/main/Kubernetes/StatefulSet/StatefulSet.yml)

         ```bash
         kubectl create -f statefulset.yaml
         ```
         
         ```bash
         kubectl get statefulset
         ```
         
         We can see that the name is the web. Hence, we define the name as web, and the pod names are very simple, web-0 and web-1.
         
       - scaling up to 3 relicas

         ```bash
         kubectl scale --replicas=3 statefulset/web
         ```
    - > **Note**:- whenever we scale up the replicas of statefulsets, it gets replicated with sticky order like web-0, web-1, web-2 etc. and next replica will be not created until and unless previous replica got created succesfully. Likewise pods of statefulsets gets deleted in reverse order like web-2, web-1 and then web-0 and one pod will be not deleted until and unless its next replica got deleted successfully!
    - Now, we have over here that whenever we create a headless service, it will also have created a network entity, so it has its own DNS name that can be referred       to. How we can check that is we can use the curl command. Now go inside a pod and run this command 

      ```bash
      curl web-0.nginx.default.svc.cluster.local
      ```
      you will see hello world!
    - Scaling down a StatefulSet

      ```bash
      kubectl scale --replicas=1 statefulset/web
      ```
      
      we can see that the two pods web-2 & web-1 is getting deleted, and we only have one pod remaining, which is web-0.
---
 - ### `Daemonset` :- 
    - DaemonSets in Kubernetes ensure that a copy of the pod runs on all or some nodes. Whenever a new node joins, the same copy of the pod is pinned over there,         and whenever a node is removed from the cluster, the pod also gets removed.
    - The DaemonSet controller controls DaemonSet, which is scheduled on all the pods, except for the ones where you cannot schedule the pods, or it's not                 schedulable, like on a master node.
    - first see what all DaemonSets we already have running on our cluster by this below command :- 

      ```bash
      kubectl get ds -A
      ```
      
      create a daemonset by the yaml manifest
      
      ```bash
      kubectl create -f ds.yaml
      ```
      
      ```bash
      kubectl get ds
      ```
      
      ```bash
      kubectl logs -f #daemonset-pod-name here
      ```
      
      ```bash
      kubectl get ds -owide
      ```
---
