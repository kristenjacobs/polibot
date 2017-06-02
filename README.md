# polibot
Example of using Wercker with a really simple go application

To run locally

```
make wercker-dev
```

To run the build pipeline locally

```
make wercker-build
```

To run the deploy pipeline locally

```
make wercker-deploy
```

## Building a Kubernetes Cluster on Oracle BMCS

In order to deploy this application, we need a running Kubernetes cluster.

See: https://kubernetes.io/docs/getting-started-guides/kubeadm/

Create 3 machines, one master and 2 slaves. Note: The network that the BMC machines use must be configured to allow UDP traffic, i.e. add the following stateless ingress and egress rules:

```
Source: 10.0.0.0/8
IP Protocol: UDP 
Source Port Range: All	 
Destination Port Range: 8472  
Allows: UDP traffic for ports: 8472
```

Note: This opens UDP traffic on port 8472, thus allowing the encapsulated
fannel traffic between nodes in the cluster. 

### To Install Kubernetes (On masters and slaves)

```
sudo su -
ARCH=x86_64
cat <<EOF > /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-${ARCH}
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg
        https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
EOF
setenforce 0
yum install -y docker kubelet kubeadm kubectl kubernetes-cni
systemctl enable docker && systemctl start docker
systemctl enable kubelet && systemctl start kubelet
systemctl stop firewalld
exit
```

### To Start Kubernetes (On the master)

Start kube (Note: Need to plugin the public IP of the master)

```
sudo kubeadm init --pod-network-cidr=10.244.0.0/16 --apiserver-cert-extra-sans=<MASTER-IP> | tee $HOME/notes.txt
```

Point kubectl at the kubernetes cluster config  

```
sudo cp /etc/kubernetes/admin.conf $HOME/
sudo chown $(id -u):$(id -g) $HOME/admin.conf
export KUBECONFIG=$HOME/admin.conf
```

Creates the networking pods Using fannel

```
kubectl create -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel-rbac.yml
kubectl create -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
```

Start the kube dashboard

```
kubectl create -f https://git.io/kube-dashboard
```

Now run from the master, we can see the nodes/pods running as follows:

```
kubectl get nodes
kubectl get pods --all-namespaces
```

### To Join the Cluster (On the slave)

Run this from each slave node to join the cluster (Get token from 'kubeadm init' output)
Note: You have to do the (install) steps above to install kubeadm on each slave.

```
sudo kubeadm join --token <TOKEN> <MASTER-IP>:6443 
```

Running ```kubectl get nodes``` from the master should show the connected slave node.


### To Connect Form Another Machine

Copy the config from /etc/kubernetes/admin.conf on the master to somewhere local.
Update the private IP addres in the config with the public IP of the master machine.
Set the location of the config file in the environment:

```
export KUBECONFIG=<PATH_TO_ADMIN.CONF>
```

kubectl can then be used as normal, e.g:

```
kubectl cluster-info
kubectl get pods --all-namespaces
etc...
```

#### To accces API va a browser, run in a terminal (on your local machine):

```
export KUBECONFIG=<PATH_TO_ADMIN.CONF>
kubectl proxy
```

The kube api is now accessable locally via: http://localhost:8001/api/v1

#### To accces UI va a browser:
Ssh to the master machine while setting up a tunnel from the local machine to the master.

```
ssh -i ~/keys/bmc -L 8001:localhost:8001 opc@<MASTER-IP>
export KUBECONFIG=<PATH_TO_ADMIN.CONF>
kubectl proxy
```

The kube ui is now accessable locally via: http://localhost:8001/ui

### To Reset and Start Again

On the master and slave nodes:

```
sudo kubeadm reset
```