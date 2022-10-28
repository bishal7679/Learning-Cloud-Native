# Kubernetes Secret

---

- ### `ConfigMap` :- 
   - ConfigMaps decouple containerized image and provide different configurations to run in different environments.
   - #### Creating a ConfigMap :- 
     - ```bash
       kubectl create configmap
       ```
       We can create configmap from a file with a key, from a literal, or the env file.
       
       ```bash
       kubectl create configmap test --from-file=file.prop
       ```
       ```bash
       kubectl describe configmap test
       ```
