# Kubernetes Volume

---

Kubernetes provides various volume plug-ins that can be mounted inside the container to store some data temporarily or in a persistent manner, or locally as well. 
The Kubernetes volumes can be divided into four sections on a very high level.

![](https://imgur.com/SpzjLak.png)

- ### `remote storage` :- 
  - Remote storage can be the Cloud volume plug-ins, GlusterFs, and NFS. All these can be created as remote storage. 
  - We have external storage outside of our cluster, and data can be saved over there. So even if the cluster goes down, your data is still safe.

- ### `Ephemeral Storage` :-
  - Ephemeral storage is something that the container needs when it's running. 
  - For the caching or the Secrets mounting, the ConfigMap usage, all these volumes are there as ephemeral **storage**, **emptyDir**, **Secrets**, **ConfigMap**, **CSI ephemeral volumes**.

- ### `Hostpath` :- 
  - hostPath mounts the node's disk inside the container. And so, it is very much restricted inside the cluster. 
  - If your pod is pinned up on another node where the hostPath is not there, there can be inconsistencies, but there are certain use cases for the usage of hostPath.

- ### `Persistent Volume Claim` :- 
  - persistent volume claim persist the data in Kubernetes.(these are separate Kubernetes objects) 
  - The persistent volume claim is the volume plug-in, and what it does is it talks to persistent volume and claims that persistent volume. 
  - Persistent volume can be **statically provisioned** by the admin, or **dynamic configuration and provisioning** can be based on the storage class.

---

- ## EmptyDir 
  
