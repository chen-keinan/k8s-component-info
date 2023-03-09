# k8s-component-info

```json
{
  "cluster_name": "kind-kind",
  "version": {
    "major": "1",
    "minor": "21",
    "gitVersion": "v1.21.1",
    "gitCommit": "5e58841cce77d4bc13713ad2b91fa0d961e69192",
    "gitTreeState": "clean",
    "buildDate": "2021-05-21T23:06:30Z",
    "goVersion": "go1.16.4",
    "compiler": "gc",
    "platform": "linux/arm64"
  },
  "control_plane": {
    "components": [
      {
        "name": "etcd",
        "container": "k8s.gcr.io/etcd:3.4.13-0"
      },
      {
        "name": "kube-controller-manager",
        "container": "k8s.gcr.io/kube-controller-manager:v1.21.1"
      },
      {
        "name": "kube-apiserver",
        "container": "k8s.gcr.io/kube-apiserver:v1.21.1"
      },
      {
        "name": "kube-scheduler",
        "container": "k8s.gcr.io/kube-scheduler:v1.21.1"
      }
    ]
  },
  "node_info": [
    {
      "node_role": "master",
      "node_name": "kind-control-plane",
      "kubelet_version": "v1.21.1",
      "container_runtime_version": "containerd://1.5.2",
      "os_image": "Ubuntu 21.04",
      "host_name": "kind-control-plane",
      "kernel_version": "6.1.14-200.fc37.aarch64",
      "kube_proxy_version": "6.1.14-200.fc37.aarch64",
      "operating_system": "linux",
      "architecture": "arm64"
    },
    {
      "node_role": "worker",
      "node_name": "kind-worker",
      "kubelet_version": "v1.21.1",
      "container_runtime_version": "containerd://1.5.2",
      "os_image": "Ubuntu 21.04",
      "host_name": "kind-worker",
      "kernel_version": "6.1.14-200.fc37.aarch64",
      "kube_proxy_version": "6.1.14-200.fc37.aarch64",
      "operating_system": "linux",
      "architecture": "arm64"
    },
    {
      "node_role": "worker",
      "node_name": "kind-worker2",
      "kubelet_version": "v1.21.1",
      "container_runtime_version": "containerd://1.5.2",
      "os_image": "Ubuntu 21.04",
      "host_name": "kind-worker2",
      "kernel_version": "6.1.14-200.fc37.aarch64",
      "kube_proxy_version": "6.1.14-200.fc37.aarch64",
      "operating_system": "linux",
      "architecture": "arm64"
    }
  ],
  "addons": [
    {
      "name": "coredns-558bd4d5db-d4tx8",
      "container": "k8s.gcr.io/coredns/coredns:v1.8.0"
    },
    {
      "name": "coredns-558bd4d5db-vsqqp",
      "container": "k8s.gcr.io/coredns/coredns:v1.8.0"
    },
    {
      "name": "kindnet-2s5ds",
      "container": "docker.io/kindest/kindnetd:v20210326-1e038dc5"
    },
    {
      "name": "kindnet-d8lt5",
      "container": "docker.io/kindest/kindnetd:v20210326-1e038dc5"
    },
    {
      "name": "kindnet-wh5h6",
      "container": "docker.io/kindest/kindnetd:v20210326-1e038dc5"
    },
    {
      "name": "kube-proxy-lcrc6",
      "container": "k8s.gcr.io/kube-proxy:v1.21.1"
    },
    {
      "name": "kube-proxy-n8rt5",
      "container": "k8s.gcr.io/kube-proxy:v1.21.1"
    },
    {
      "name": "kube-proxy-tdpgn",
      "container": "k8s.gcr.io/kube-proxy:v1.21.1"
    }
  ]
}
```
