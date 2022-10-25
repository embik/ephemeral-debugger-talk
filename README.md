# Ephemeral Containers in Action - Running a Go Debugger in Kubernetes

This repository serves as companion repository to the talk [Ephemeral Containers in Action - Running a Go Debugger in Kubernetes](https://cfp.cloud-native.rejekts.io/cloud-native-rejekts-eu-amsterdam-2023/talk/XSGW8F/), given at [Cloud Native Rejekts](https://cloud-native.rejekts.io) EU 2023.

## Talk Description

**Ephemeral containers are an amazing recent feature in Kubernetes with great potential. We will explore that potential by running a live debugger session alongside an application pod and debug it remotely.**

The modern observability stack has transformed the way you troubleshoot issues in a microservice environment. Some situations however ask for investigation of a single application pod that seems to be misbehaving. Ephemeral containers provide a way to attach to a seemingly problematic pod without restarting it. This allows developers to observe an issue in a live or staging environment running on top of Kubernetes.

We will discuss the practicality of launching a Go debugger (Delve) within an ephemeral container to remotely debug an application both on the CLI and in VS Code, highlighting the requirements and possible limitations one might encounter when trying to set up a similar troubleshooting routine. As part of that, we will explore the API for ephemeral containers and the current implementation in kubectl.

While the talk will use Go and Delve as an example, the considerations and steps presented are of universal importance to running a debugger for your language stack of choice.

## Content

- [slides](./slides/): the talk slides in PDF format.
- [sample-app](./sample-app/): very, very simple Go webserver that is debugged during the talk.
- [deploy](./deploy/): deployment manifests for the sample application.
- [images](./images/): Dockerfile for a [Delve](https://github.com/go-delve/delve) container image, used as part of the talk.
- [kubectl-debug](./kubectl-debug/): a simple patch to `kubectl debug` that was hacked together in the early days of this talk.

## Other Resources

- [embik/kubectl-ephemeral](https://github.com/embik/kubectl-ephemeral): the custom `kubectl` plugin used for launching Delve as ephemeral container.
