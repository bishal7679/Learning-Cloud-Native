# Kubernetes Objects

---

- ### Pods:-
  - It is the smallest unit in Kubernetes. 
  - It is a group of one or more containers and a specification telling how the containers will be running inside the Kubernetes cluster.
  - A pod gets its own IP address.

  ![]()
  
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
     - When we give a command, kubectl apply -f or create -f, and provide a YAML file, it's first converted to JSON and sent to API Server. Here, the request is authenticated using our kube config credentials, and then it's authorized whether the user is actually authorized to perform this particular command to create a pod. Then if any admission controllers are there, they are checked before it gets persisted to etcd datastore.
       Now, after it is persisted in the etcd datastore, it goes to pending state. Then the scheduler pitches in and tries to find the best match for the node where it has to be spawned. It will go through the cluster and find the best fit node based on the resources, etc., and if the image is already present on the node because that has some reference. Once it gets the node, it will fill the label spec node name and send it to the API server, and now, that particular request is also stored in etcd.
     
   - ### API server in a pod's lifecycle :- 
     - Now, the API server instructs the kubelet, "Hey, there is one pod that has to be spawned on this particular node." The kubelet is responsible for fetching the image from the image registry. It can be any image registry, and then the CRI does that, and the CNI will get the IP attached to the pod, and then that particular IP will be sent again back to the API Server, then again, it is stored in etcd. From there, it goes to ContainerCreating when the scheduler selects the node, and after that, when everything is there, the image is pulled, and then it goes to the running state. 
       We can see the state when we use the `kubectl get pods` command.
       
   - ### Probes and checks :- 
     - Now there can be other advanced things that happen, or whenever the process dies too many times within a pod, it can also go to crashloopbackoff, and whenever it is succeeded, it will be in the succeeded state. Another advanced kind of lifecycle includes liveliness, so we can have a /health and a /ready. These are known as liveness and readiness probes. 
       The pod will keep on checking that, and if it fails, it can lead to the crashloopbackoff.
       
   - ### Hooks and its usages :- 
     - There are some of the hooks that can be implemented. For example, if there are some actions that you want to perform after the main container just starts, you can have a post-start hook, and if you want to perform before the main container gets terminated, you will have to have a pre-stop hook. There is also something called the init container that will be running before actually starting the main container.
   
 - ### Init Containers
