# Introduction To Cloud Concepts & Getting Started with AWS

---

- ### What is Cloud Computing?
  - The practice of using a network of remote servers hosted on the internet to store,manage, and process data, rather than a local server or a personal 
    computer.
  - **On-Premise**:-
     - You own the servers
     - You hire the IT people
     - You pay or rent the real-estate
     - You take the all risk
  - **Cloud Providers**:-
     - Someome else owns the servers
     - Someone else hires the IT people
     - Someone else pays or rents the real-estate
     - You are responsible for your configuration cloud services and code, someone else takes care of the rest
 
 ---
 
 - ### The Evolution Of Cloud Hosting
   - **Dedicated Server:-**
   
      ![](https://imgur.com/hUej5tc.png)
      - One physical machine dedicated to single a business runs a single web-app/site.
      - Its very Expensive,High Maintenance, *High security

   - **Virtual Private Server (VPS):-**
     
      ![](https://imgur.com/vSm8eLj.png)
      - One physical machine dedicated to a single business.
      - The physical machine is virtualized into sub-machines runs multiple web-apps/sites
      - Better utilization and Isolation of Resources

   - **Shared Hosting:-**
   
      ![](https://imgur.com/XYYiSdR.png)
      - One physical machine, shared by hundred of businesses relies on most tenants undeer-utilizing their resources.
      - Very cheap, Limited functionality, Poor Isolation

   - **Cloud Hosting:-**

      ![](https://imgur.com/Sx3ySEi.png)
      - Multiple physical machines that act as one system. The system is abstracted into multiple cloud services.
      - Its very flexible, Scalable, Cost-Effective, High Configurability
     
---

- ### What is Amazon?
   - An American multinational computer technology corporation headquartered in Seattle, Washnigton.
- ### What is Amazon Web Services (AWS)?
   - Amazon class their cloud provider service *Amazon Web Services* commonly referred to just AWS
   - Its a collection of cloud services that can be used together under a single unified API to build a lot of kind of workloads.
   - **`AWS services`**:-
      - SQS (Simple Queue Service) was the first AWS service launched foer public use in 2004.
      - S3 (Simple Storage Service) was launched in March of 2006
      - EC2 (Elastic Compute Cloud) was launched in August of 2006
      - In November 2010, it was reported that all of Amazon.com's retail sites had migrated to AWS

---

- ### What is Cloud Service Provider?
   - A Cloud Service Provider (CSP) is a company which
     - provides multiple Cloud Services e.g. tens to hundreds of services
     - those Cloud services can be chained together to create cloud architectures
     - those Cloud services are accessible via Single Unified API e.g. AWS API
     - those Cloud services utilized metered billing based on usage e.g. per second, per hour
     - those Cloud services have rich monitoring built in e.g. AWS CloudTrail
     - those Cloud services have an Infrastructure as a Service (IaaS) offering
     - those Cloud services offers automation via Infrastructure as Code (IaC).
    
     ![](https://imgur.com/fvNONS8.png)
    > **Note**:- **If a company offers multiple cloud services under a single UI but do not meet most of or all of these requirements, 
                it would be referred to as a Cloud Platform e.g. Twilio, Hashicorp, Databricks**
 - ### Landscape of CSPs ⤵️
    - **`Tier-1 (Top Tier)`** ➡ Early to market, wide offering, strong synergies between services, well recognised in the industry
    
    ![](https://imgur.com/ZByomrV.png)
    
    - **`Tier-2 (Mid Tier)`** ➡ Backend by well-known tech companies, slow to innovate and turned to specialization.
    
    ![](https://imgur.com/0hg2rYq.png)
    
    - **`Tier-3 (Light Tier)`** ➡ Virtual Private Servers (VPS) turned to offer core IaaS offering. Simple, cost-effective
    
    ![](https://imgur.com/OXORIvR.png)
 
 ---
    
- ### Some Common Cloud Services
   - A cloud service provider can have hundreds of cloud services that are grouped into various types of services. The four most common types of        cloud services (the 4 core) for Infrastructure as a Service (IaaS) would be :-
     - **`Compute`**
        - Imagine having a virtual computer that can run application, programs and code.
     - **`Networking`**
        - Imagine having virtual network defining internet connections or network isolations between services or outbound to the internet.
     - **`Storage`**
        - Imagine having a virtual hard-drive that can store files
     - **`Databases`**
        - Imagine a virtual database for storing reporting data or a database for general purpose web-application
     
     ![](https://imgur.com/sfpsWQE.png)
     
     ![](https://imgur.com/31DyeB4.png)
     
- ### The Evolution Of Computing
   1. **Dedicated**
    
      ![]()
      - A physical server wholly utilized by a single customer.
      - You have to guess your capacity
      - you'all overpay for an underutilized server
      - You can't vertical scale, you need a manual migration
      - Replacing a server is very difficult
      - You are limited by your Host Operating System
      - Multiple apps can result in conflicts in resource sharing
      - You have a *gurantee of security, privacy, and full utility of underlying resources
     
   1. **VMs**
   
      ![]()
      - You can run multiple Virtual Machines on one machine.
      - Hypervisor is the software layer that lets you run the VMs
      - A physical server shared by multiple customers
      - You are to pay for a fraction of the server
      - You'all overpay for an underutilized Virtual Machine
      - You are limited by your Guest Opearating System
      - Multiple apps on a single Virtual Machine can result in conflicts in resource sharing
      - Easy to export or import images for migration
      - Easy to Vertical or Horizontally scale
      
    1. **Containers**

       ![]()
       - Virtual Machine running multiple containers
       - Docker Daemon is the name of the software layer that lets you run multiple containers
       - You can maximize the utilize of the available capacity which is more cost-effective
       - Your containers share the same underlying OS so containers are more efficient than multiple VMs
       - Multiple apps can run side by side without being limited to the same OS requirements and will not cause conflicts during resources sharing

    1. **Functions**

       ![]()
       - Are managed VMs running managed containers known as Serverless Compute
       - You upload a piece of code, choose the amount of memory and duration
       - Only responsible for code and data, nothing else
       - Very cost-effective, only pay for the timecode is running, VMs only run when there is code to be executed
       - Cold starts is a side-effect of this setup

---

- ### Types Of Cloud Computing
   - There are mainly three types of Cloud computing i.e. SaaS, PaaS, and IaaS

     ![](https://imgur.com/zUsJQc7.png)
     
- ### Cloud Computing Deployment Models
   - **`Public Cloud`**
     - Everything (the workload or project) is built on the CSP. Also known as: Cloud-Native or Cloud First
     
     ![]()
   
   - **`Private Cloud`**
      - Everything built on company's datacenters. Also known as On-Premise and the cloud could be Openstack
      
      ![]()
      
   - **`Hybrid Cloud`**
      - Using both On-Premise and a Cloud Service Provider

      ![]()
      
   - **`Cross-Cloud`**
      - Using multiple cloud providers aka multi-cloud

- ### Cloud Computing Deployment Models

  ![](https://imgur.com/WsyWJMo.png)
  
---

- ### Creating an AWS Account
   - At First go to [`AWS`](https://aws.amazon.com/)
   - Then click on `create an AWS Account` or `Sign in to the console` (if you already created account) button
   
   ![](https://imgur.com/bHH4W4C.png)
   
   - You need a credit card or debit card here to create your account.
   - After creating account successfully 🎉 go to AWS Management Console
   - When you sign in your account you can see there are two types
     - Root user - It needs email address to login
     - IAM user  - It needs 12 digit Account ID or Account Alias
   - Lets sign in with the root account and search `IAM (Identity Access Management)` on the search bar
     
     ![](https://imgur.com/R7a44Qr.png)
   - Then at the left panel click on the `Users` to create a new user and then click on `Add users`
   
     ![](https://imgur.com/wyF2SK6.png)
     
     ![](https://imgur.com/0P3SNng.png)
    - Then click `create group` button and check only `AdministratorAccess`

     ![](https://imgur.com/j2iXF9J.png)
    
    - After that click next - next - and then click on `create user`
    - after that copy the password to clipboard and remember your username/Account alias/Account ID and sign out from root account

      ![](https://imgur.com/t7zZrJG.png)
    - Again sign in with your username and password as IAM user and change your auto generated password to your own preferred password

      ![](https://imgur.com/zmNFYLf.png)
      
---
      
- ### AWS Region selector
  - You can select your region for using various region based AWS services (most of the time use **US East (N. Virginia)
    us-east-1**)
    
    ![](https://imgur.com/3BukYbP.png)
  - Some services are global and not dependable on region. Ex:- CloudFront

    ![](https://imgur.com/8gHSc95.png)
    
- ### AWS Billing
   - Billing is something like you have to pay according to whatever service you are using in the basis of time

- ### AWS Free Tier
   - You can use AWS services for free for 12 months
   - just you need to go to [`this site`](https://aws.amazon.com/free/?all-free-tier.) and take free tier
   - After that you need to checkout these as below image and then save preference

     ![](https://imgur.com/q2ojej1.png)
     
- ### AWS Billing
   - We can monitor our spend on a particular service through building alerts or alarms
   - Search `cloudwatch` on the search bar
   - `cloudwatch` is a collection of services such as **Cloudwatch alarms**, **Cloudwatch logs**, **Cloudwatch metrics**
   
   ![](https://imgur.com/hpsteFb.png)
   
- ### Turning on MFA (Multi Factor Authentication)
  - one of the strongest recommendationsthat aws gives you to set MFA on your aws root user account.
  - Make sure you logged in with your root user account and then go to `IAM` dashboard
  - Here you can see `Add MFA` button, click on that then `Activate MFA` and `Virtual MFA device` you can continue here.

   ![](https://imgur.com/nX0Ns5Y.png)
   
   ![](https://imgur.com/zYl92uX.png)
   
  - you can use here any of the app you like for authentication

   ![](https://imgur.com/0kVYuji.png)
   
  - Now scan the QR with your app and add two MFA key and there you go :)

- ### Evolution of Computing Power

  ![](https://imgur.com/Umy7IDB.png)
  
  - you can do quantam computing in AWS using [`Amazon Braket`](https://us-east-1.console.aws.amazon.com/braket/home?cp=bn&pg=ln&region=us-east-1#/devices) service

  ![](https://imgur.com/dFiHxmX.png)
  
- ### The Benefits of Cloud

   ![](https://imgur.com/dCIx7AJ.png)
 
---

- ### AWS Global Infrastructure
  - The **AWS global infrastructure** is globally distributed hardware and datacenters that are physically networked together to act as one large resources for the end customer.
    - The AWS global infrastructure is made up of the following resources :-
      - 25 launched Regions
      - 81 availbility Zones
      - 108 Direct connections locations
      - 275+ points of presence
      - 11 local zones   

- ### Regions
  
  ![](https://imgur.com/QSxbwOD.png)
  
- ### Regional vs Global Services

  ![](https://imgur.com/MU9xPoe.png)
  
- ### Global Infrastructure - AZs

  ![](https://imgur.com/D75B2s4.png)
  
  ![](https://imgur.com/Xud5jSe.png)
  
- ### What is Fault Tolerance?
  
  ![](https://imgur.com/LT3OdkC.png)
    
  ![](https://imgur.com/e0SJVmQ.png)
  
- ### AWS Global Network

  ![](https://imgur.com/w8KqTKQ.png)
  
- ### Point of Presence (PoP)

   ![](https://imgur.com/ifXfsKQ.png)
   
   ![](https://imgur.com/WKLLQ9f.png)
 
---

- ### AWS Direct Connect service

   ![](https://imgur.com/D3V1BlB.png)
   
- ### AWS sustainbility

  ![](https://imgur.com/PYMa22C.png)
  
- ### AWS Ground station

   ![](https://imgur.com/VwUqYQb.png)
   
- ### AWS outposts service

  ![](https://imgur.com/dWck5eS.png)
  
---

- ### Cloud Architecture Terminologies

   ![](https://imgur.com/z9vCuL8.png)
   
   1. ### High Availibility
      
      ![](https://imgur.com/6VfvoI3.png)
      
   2. ### High Scalability

      ![](https://imgur.com/FCGLFbJ.png)
      
   3. ### High Elasticity

      ![](https://imgur.com/XcEYPjy.png)
      
   4. ### Highly Fault Tolerance

       ![](https://imgur.com/coliuWu.png)
       
   5. ### High Durability

       ![](https://imgur.com/bUkh4mC.png)
  
