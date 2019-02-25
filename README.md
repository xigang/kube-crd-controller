An example of a custom Kubernetes controller that's only purpose is to watch for the creation, updating, or deletion of all custom resource of type MyCustomResource (in the all namespaces). This was created as an exercise to understand how Kubernetes controllers work and interact with the cluster and resources.

## Installation

```
# go get github.com/xigang/kube-crd-controller
# go build -o kube-crd-controller
```

## Getting Started

#### register the custom resource definition

```
# kubectl create -f example/mainifest/crd-example.yaml
```

#### Create a Mycustomresource type of resource

```
# kubectl create -f example/mainifest/k8s-custom-resource-examle.yaml
```

#### build and run example

```
# ./kube-crd-controller -kubeconfig=/etc/kubernetes/kubeconfig -master=https://kubernetes-master:6443
```
