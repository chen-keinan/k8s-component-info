# k8s-component-info

## Description
k8s-component-info is an open source project which collect component and version info from runnning k8s cluster and produce k8s bill of materials.

```sh
go build main.go

./main
```

```json
{
  "bomFormat": "CycloneDX",
  "specVersion": "1.4",
  "serialNumber": "urn:uuid:3e671687-395b-41f5-a30f-a58921a69b79",
  "metadata": {
    "timestamp": "2023-03-27T16:29:30+03:00",
    "component": {
      "bom-ref": "kubernetes:1.21.1",
      "type": "Cluster",
      "name": "kind-kind",
      "version": "v1.21.1"
    },
    "tools": [
      {
        "vendor": "aquasecurity",
        "name": "trivy",
        "version": "0.38.1"
      }
    ]
  },
  "version": 1,
  "control_plane": {
    "components": [
      {
        "bom-ref": "pkg:oci/etcd@05b738aa1bc6355db8a2ee8639f3631b908286e43f584a3d2ee0c472de033c28?repository_url=k8s.gcr.io/library/etcd",
        "type": "container",
        "name": "etcd",
        "purl": "pkg:oci/etcd@05b738aa1bc6355db8a2ee8639f3631b908286e43f584a3d2ee0c472de033c28?repository_url=k8s.gcr.io/library/etcd"
      },
      {
        "bom-ref": "pkg:oci/kube-apiserver@18e61c783b41758dd391ab901366ec3546b26fae00eef7e223d1f94da808e02f?repository_url=k8s.gcr.io/library/kube-apiserver",
        "type": "container",
        "name": "kube-apiserver",
        "purl": "pkg:oci/kube-apiserver@18e61c783b41758dd391ab901366ec3546b26fae00eef7e223d1f94da808e02f?repository_url=k8s.gcr.io/library/kube-apiserver"
      },
      {
        "bom-ref": "pkg:oci/kube-controller-manager@0c6dccae49de8003ee4fa06db04a9f13bb46cbaad03977e6baa21174f2dba2fc?repository_url=k8s.gcr.io/library/kube-controller-manager",
        "type": "container",
        "name": "kube-controller-manager",
        "purl": "pkg:oci/kube-controller-manager@0c6dccae49de8003ee4fa06db04a9f13bb46cbaad03977e6baa21174f2dba2fc?repository_url=k8s.gcr.io/library/kube-controller-manager"
      },
      {
        "bom-ref": "pkg:oci/kube-scheduler@8c783dd2520887cc8e7908489ffc9f356c82436ba0411d554237a0b9632c9b87?repository_url=k8s.gcr.io/library/kube-scheduler",
        "type": "container",
        "name": "kube-scheduler",
        "purl": "pkg:oci/kube-scheduler@8c783dd2520887cc8e7908489ffc9f356c82436ba0411d554237a0b9632c9b87?repository_url=k8s.gcr.io/library/kube-scheduler"
      }
    ]
  },
  "nodes": [
    {
      "node_role": "master",
      "node_name": "kind-control-plane",
      "kubelet_version": "v1.21.1",
      "container_runtime_version": "containerd://1.5.2",
      "os_image": "Ubuntu 21.04",
      "host_name": "kind-control-plane",
      "kernel_version": "6.1.18-200.fc37.aarch64",
      "kube_proxy_version": "6.1.18-200.fc37.aarch64",
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
      "kernel_version": "6.1.18-200.fc37.aarch64",
      "kube_proxy_version": "6.1.18-200.fc37.aarch64",
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
      "kernel_version": "6.1.18-200.fc37.aarch64",
      "kube_proxy_version": "6.1.18-200.fc37.aarch64",
      "operating_system": "linux",
      "architecture": "arm64"
    }
  ],
  "addons": [
    {
      "bom-ref": "pkg:oci/coredns/coredns@1a1f05a2cd7c2fbfa7b45b21128c8a4880c003ca482460081dc12d76bfa863e8?repository_url=k8s.gcr.io/library/coredns/coredns",
      "type": "container",
      "name": "coredns",
      "purl": "pkg:oci/coredns/coredns@1a1f05a2cd7c2fbfa7b45b21128c8a4880c003ca482460081dc12d76bfa863e8?repository_url=k8s.gcr.io/library/coredns/coredns"
    },
    {
      "bom-ref": "pkg:oci/coredns/coredns@1a1f05a2cd7c2fbfa7b45b21128c8a4880c003ca482460081dc12d76bfa863e8?repository_url=k8s.gcr.io/library/coredns/coredns",
      "type": "container",
      "name": "coredns",
      "purl": "pkg:oci/coredns/coredns@1a1f05a2cd7c2fbfa7b45b21128c8a4880c003ca482460081dc12d76bfa863e8?repository_url=k8s.gcr.io/library/coredns/coredns"
    },
    {
      "bom-ref": "pkg:oci/kindest/kindnetd@f37b7c809e5dcc2090371f933f7acb726bb1bffd5652980d2e1d7e2eff5cd301?repository_url=index.docker.io/library/kindest/kindnetd",
      "type": "container",
      "name": "kindnet-cni",
      "purl": "pkg:oci/kindest/kindnetd@f37b7c809e5dcc2090371f933f7acb726bb1bffd5652980d2e1d7e2eff5cd301?repository_url=index.docker.io/library/kindest/kindnetd"
    },
    {
      "bom-ref": "pkg:oci/kindest/kindnetd@f37b7c809e5dcc2090371f933f7acb726bb1bffd5652980d2e1d7e2eff5cd301?repository_url=index.docker.io/library/kindest/kindnetd",
      "type": "container",
      "name": "kindnet-cni",
      "purl": "pkg:oci/kindest/kindnetd@f37b7c809e5dcc2090371f933f7acb726bb1bffd5652980d2e1d7e2eff5cd301?repository_url=index.docker.io/library/kindest/kindnetd"
    },
    {
      "bom-ref": "pkg:oci/kindest/kindnetd@f37b7c809e5dcc2090371f933f7acb726bb1bffd5652980d2e1d7e2eff5cd301?repository_url=index.docker.io/library/kindest/kindnetd",
      "type": "container",
      "name": "kindnet-cni",
      "purl": "pkg:oci/kindest/kindnetd@f37b7c809e5dcc2090371f933f7acb726bb1bffd5652980d2e1d7e2eff5cd301?repository_url=index.docker.io/library/kindest/kindnetd"
    },
    {
      "bom-ref": "pkg:oci/kube-proxy@4bbef4ca108cdc3b99fe23d487fa4fca933a62c4fc720626a3706df9cef63b21?repository_url=k8s.gcr.io/library/kube-proxy",
      "type": "container",
      "name": "kube-proxy",
      "purl": "pkg:oci/kube-proxy@4bbef4ca108cdc3b99fe23d487fa4fca933a62c4fc720626a3706df9cef63b21?repository_url=k8s.gcr.io/library/kube-proxy"
    },
    {
      "bom-ref": "pkg:oci/kube-proxy@4bbef4ca108cdc3b99fe23d487fa4fca933a62c4fc720626a3706df9cef63b21?repository_url=k8s.gcr.io/library/kube-proxy",
      "type": "container",
      "name": "kube-proxy",
      "purl": "pkg:oci/kube-proxy@4bbef4ca108cdc3b99fe23d487fa4fca933a62c4fc720626a3706df9cef63b21?repository_url=k8s.gcr.io/library/kube-proxy"
    },
    {
      "bom-ref": "pkg:oci/kube-proxy@4bbef4ca108cdc3b99fe23d487fa4fca933a62c4fc720626a3706df9cef63b21?repository_url=k8s.gcr.io/library/kube-proxy",
      "type": "container",
      "name": "kube-proxy",
      "purl": "pkg:oci/kube-proxy@4bbef4ca108cdc3b99fe23d487fa4fca933a62c4fc720626a3706df9cef63b21?repository_url=k8s.gcr.io/library/kube-proxy"
    }
  ],
  "dependencies": [
    {
      "ref": "kind-kind",
      "dependsOn": [
        "pkg:container/k8s.gcr.io/etcd@3.4.13-0",
        "pkg:container/k8s.gcr.io/kube-apiserver@v1.21.1",
        "pkg:container/k8s.gcr.io/kube-controller-manager@v1.21.1",
        "pkg:container/k8s.gcr.io/kube-scheduler@v1.21.1",
        "pkg:container/k8s.gcr.io/coredns/coredns@v1.8.0",
        "pkg:container/k8s.gcr.io/coredns/coredns@v1.8.0",
        "pkg:container/docker.io/kindest/kindnetd@v20210326-1e038dc5",
        "pkg:container/docker.io/kindest/kindnetd@v20210326-1e038dc5",
        "pkg:container/docker.io/kindest/kindnetd@v20210326-1e038dc5",
        "pkg:container/k8s.gcr.io/kube-proxy@v1.21.1",
        "pkg:container/k8s.gcr.io/kube-proxy@v1.21.1",
        "pkg:container/k8s.gcr.io/kube-proxy@v1.21.1"
      ]
    }
  ]
}
```
