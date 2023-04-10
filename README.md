# k8s-component-info

## Description
k8s-component-info is an open source project which collect component and version info from runnning k8s cluster and produce k8s bill of materials.

```sh
go build main.go

## Native
./main 

```json
{
 "metadata": {
  "timestamp": "2023-04-10T14:36:07+03:00",
  "component": {
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
 "control_plane": {
  "components": [
   {
    "name": "etcd:3.4.13-0",
    "version": "3.4.13-0",
    "repository": "etcd",
    "registry": "k8s.gcr.io",
    "digest": "05b738aa1bc6355db8a2ee8639f3631b908286e43f584a3d2ee0c472de033c28"
   },
   {
    "name": "kube-apiserver:v1.21.1",
    "version": "v1.21.1",
    "repository": "kube-apiserver",
    "registry": "k8s.gcr.io",
    "digest": "18e61c783b41758dd391ab901366ec3546b26fae00eef7e223d1f94da808e02f"
   },
   {
    "name": "kube-controller-manager:v1.21.1",
    "version": "v1.21.1",
    "repository": "kube-controller-manager",
    "registry": "k8s.gcr.io",
    "digest": "0c6dccae49de8003ee4fa06db04a9f13bb46cbaad03977e6baa21174f2dba2fc"
   },
   {
    "name": "kube-scheduler:v1.21.1",
    "version": "v1.21.1",
    "repository": "kube-scheduler",
    "registry": "k8s.gcr.io",
    "digest": "8c783dd2520887cc8e7908489ffc9f356c82436ba0411d554237a0b9632c9b87"
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
   "kernel_version": "6.2.8-200.fc37.aarch64",
   "kube_proxy_version": "6.2.8-200.fc37.aarch64",
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
   "kernel_version": "6.2.8-200.fc37.aarch64",
   "kube_proxy_version": "6.2.8-200.fc37.aarch64",
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
   "kernel_version": "6.2.8-200.fc37.aarch64",
   "kube_proxy_version": "6.2.8-200.fc37.aarch64",
   "operating_system": "linux",
   "architecture": "arm64"
  }
 ],
 "addons": [
  {
   "name": "coredns/coredns:v1.8.0",
   "version": "v1.8.0",
   "repository": "coredns/coredns",
   "registry": "k8s.gcr.io",
   "digest": "1a1f05a2cd7c2fbfa7b45b21128c8a4880c003ca482460081dc12d76bfa863e8"
  },
  {
   "name": "coredns/coredns:v1.8.0",
   "version": "v1.8.0",
   "repository": "coredns/coredns",
   "registry": "k8s.gcr.io",
   "digest": "1a1f05a2cd7c2fbfa7b45b21128c8a4880c003ca482460081dc12d76bfa863e8"
  },
  {
   "name": "kindest/kindnetd:v20210326-1e038dc5",
   "version": "v20210326-1e038dc5",
   "repository": "kindest/kindnetd",
   "registry": "index.docker.io",
   "digest": "f37b7c809e5dcc2090371f933f7acb726bb1bffd5652980d2e1d7e2eff5cd301"
  },
  {
   "name": "kindest/kindnetd:v20210326-1e038dc5",
   "version": "v20210326-1e038dc5",
   "repository": "kindest/kindnetd",
   "registry": "index.docker.io",
   "digest": "f37b7c809e5dcc2090371f933f7acb726bb1bffd5652980d2e1d7e2eff5cd301"
  },
  {
   "name": "kindest/kindnetd:v20210326-1e038dc5",
   "version": "v20210326-1e038dc5",
   "repository": "kindest/kindnetd",
   "registry": "index.docker.io",
   "digest": "f37b7c809e5dcc2090371f933f7acb726bb1bffd5652980d2e1d7e2eff5cd301"
  },
  {
   "name": "kube-proxy:v1.21.1",
   "version": "v1.21.1",
   "repository": "kube-proxy",
   "registry": "k8s.gcr.io",
   "digest": "4bbef4ca108cdc3b99fe23d487fa4fca933a62c4fc720626a3706df9cef63b21"
  },
  {
   "name": "kube-proxy:v1.21.1",
   "version": "v1.21.1",
   "repository": "kube-proxy",
   "registry": "k8s.gcr.io",
   "digest": "4bbef4ca108cdc3b99fe23d487fa4fca933a62c4fc720626a3706df9cef63b21"
  },
  {
   "name": "kube-proxy:v1.21.1",
   "version": "v1.21.1",
   "repository": "kube-proxy",
   "registry": "k8s.gcr.io",
   "digest": "4bbef4ca108cdc3b99fe23d487fa4fca933a62c4fc720626a3706df9cef63b21"
  }
 ]
}
```


./main cyclonedx

## Cyclonedx 

```json
{
  "$schema": "http://cyclonedx.org/schema/bom-1.4.schema.json",
  "bomFormat": "CycloneDX",
  "specVersion": "1.4",
  "version": 1,
  "metadata": {
    "timestamp": "2023-04-10T14:34:52+03:00",
    "tools": [
      {
        "vendor": "aquasecurity",
        "name": "trivy",
        "version": "0.38.1"
      }
    ],
    "component": {
      "bom-ref": "pkg:kind-kind:1.21.1",
      "type": "application",
      "name": "kind-kind",
      "version": "v1.21.1"
    }
  },
  "components": [
    {
      "bom-ref": "pkg:oci/etcd@05b738aa1bc6355db8a2ee8639f3631b908286e43f584a3d2ee0c472de033c28?repository_url=k8s.gcr.io/library/etcd",
      "type": "container",
      "name": "etcd",
      "version": "3.4.13-0",
      "purl": "pkg:oci/etcd@05b738aa1bc6355db8a2ee8639f3631b908286e43f584a3d2ee0c472de033c28?repository_url=k8s.gcr.io/library/etcd"
    },
    {
      "bom-ref": "pkg:oci/kube-apiserver@18e61c783b41758dd391ab901366ec3546b26fae00eef7e223d1f94da808e02f?repository_url=k8s.gcr.io/library/kube-apiserver",
      "type": "container",
      "name": "kube-apiserver",
      "version": "v1.21.1",
      "purl": "pkg:oci/kube-apiserver@18e61c783b41758dd391ab901366ec3546b26fae00eef7e223d1f94da808e02f?repository_url=k8s.gcr.io/library/kube-apiserver"
    },
    {
      "bom-ref": "pkg:oci/kube-controller-manager@0c6dccae49de8003ee4fa06db04a9f13bb46cbaad03977e6baa21174f2dba2fc?repository_url=k8s.gcr.io/library/kube-controller-manager",
      "type": "container",
      "name": "kube-controller-manager",
      "version": "v1.21.1",
      "purl": "pkg:oci/kube-controller-manager@0c6dccae49de8003ee4fa06db04a9f13bb46cbaad03977e6baa21174f2dba2fc?repository_url=k8s.gcr.io/library/kube-controller-manager"
    },
    {
      "bom-ref": "pkg:oci/kube-scheduler@8c783dd2520887cc8e7908489ffc9f356c82436ba0411d554237a0b9632c9b87?repository_url=k8s.gcr.io/library/kube-scheduler",
      "type": "container",
      "name": "kube-scheduler",
      "version": "v1.21.1",
      "purl": "pkg:oci/kube-scheduler@8c783dd2520887cc8e7908489ffc9f356c82436ba0411d554237a0b9632c9b87?repository_url=k8s.gcr.io/library/kube-scheduler"
    },
    {
      "bom-ref": "pkg:oci/coredns/coredns@1a1f05a2cd7c2fbfa7b45b21128c8a4880c003ca482460081dc12d76bfa863e8?repository_url=k8s.gcr.io/library/coredns/coredns",
      "type": "container",
      "name": "coredns/coredns",
      "version": "v1.8.0",
      "purl": "pkg:oci/coredns/coredns@1a1f05a2cd7c2fbfa7b45b21128c8a4880c003ca482460081dc12d76bfa863e8?repository_url=k8s.gcr.io/library/coredns/coredns"
    },
    {
      "bom-ref": "pkg:oci/coredns/coredns@1a1f05a2cd7c2fbfa7b45b21128c8a4880c003ca482460081dc12d76bfa863e8?repository_url=k8s.gcr.io/library/coredns/coredns",
      "type": "container",
      "name": "coredns/coredns",
      "version": "v1.8.0",
      "purl": "pkg:oci/coredns/coredns@1a1f05a2cd7c2fbfa7b45b21128c8a4880c003ca482460081dc12d76bfa863e8?repository_url=k8s.gcr.io/library/coredns/coredns"
    },
    {
      "bom-ref": "pkg:oci/kindest/kindnetd@f37b7c809e5dcc2090371f933f7acb726bb1bffd5652980d2e1d7e2eff5cd301?repository_url=index.docker.io/library/kindest/kindnetd",
      "type": "container",
      "name": "kindest/kindnetd",
      "version": "v20210326-1e038dc5",
      "purl": "pkg:oci/kindest/kindnetd@f37b7c809e5dcc2090371f933f7acb726bb1bffd5652980d2e1d7e2eff5cd301?repository_url=index.docker.io/library/kindest/kindnetd"
    },
    {
      "bom-ref": "pkg:oci/kindest/kindnetd@f37b7c809e5dcc2090371f933f7acb726bb1bffd5652980d2e1d7e2eff5cd301?repository_url=index.docker.io/library/kindest/kindnetd",
      "type": "container",
      "name": "kindest/kindnetd",
      "version": "v20210326-1e038dc5",
      "purl": "pkg:oci/kindest/kindnetd@f37b7c809e5dcc2090371f933f7acb726bb1bffd5652980d2e1d7e2eff5cd301?repository_url=index.docker.io/library/kindest/kindnetd"
    },
    {
      "bom-ref": "pkg:oci/kindest/kindnetd@f37b7c809e5dcc2090371f933f7acb726bb1bffd5652980d2e1d7e2eff5cd301?repository_url=index.docker.io/library/kindest/kindnetd",
      "type": "container",
      "name": "kindest/kindnetd",
      "version": "v20210326-1e038dc5",
      "purl": "pkg:oci/kindest/kindnetd@f37b7c809e5dcc2090371f933f7acb726bb1bffd5652980d2e1d7e2eff5cd301?repository_url=index.docker.io/library/kindest/kindnetd"
    },
    {
      "bom-ref": "pkg:oci/kube-proxy@4bbef4ca108cdc3b99fe23d487fa4fca933a62c4fc720626a3706df9cef63b21?repository_url=k8s.gcr.io/library/kube-proxy",
      "type": "container",
      "name": "kube-proxy",
      "version": "v1.21.1",
      "purl": "pkg:oci/kube-proxy@4bbef4ca108cdc3b99fe23d487fa4fca933a62c4fc720626a3706df9cef63b21?repository_url=k8s.gcr.io/library/kube-proxy"
    },
    {
      "bom-ref": "pkg:oci/kube-proxy@4bbef4ca108cdc3b99fe23d487fa4fca933a62c4fc720626a3706df9cef63b21?repository_url=k8s.gcr.io/library/kube-proxy",
      "type": "container",
      "name": "kube-proxy",
      "version": "v1.21.1",
      "purl": "pkg:oci/kube-proxy@4bbef4ca108cdc3b99fe23d487fa4fca933a62c4fc720626a3706df9cef63b21?repository_url=k8s.gcr.io/library/kube-proxy"
    },
    {
      "bom-ref": "pkg:oci/kube-proxy@4bbef4ca108cdc3b99fe23d487fa4fca933a62c4fc720626a3706df9cef63b21?repository_url=k8s.gcr.io/library/kube-proxy",
      "type": "container",
      "name": "kube-proxy",
      "version": "v1.21.1",
      "purl": "pkg:oci/kube-proxy@4bbef4ca108cdc3b99fe23d487fa4fca933a62c4fc720626a3706df9cef63b21?repository_url=k8s.gcr.io/library/kube-proxy"
    },
    {
      "bom-ref": "pkg:kind-control-plane",
      "type": "container",
      "name": "kind-control-plane",
      "purl": "pkg:kind-control-plane",
      "properties": [
        {
          "name": "node-role",
          "value": "master"
        },
        {
          "name": "host_name",
          "value": "kind-control-plane"
        }
      ]
    },
    {
      "bom-ref": "pkg:Ubuntu@v21.04",
      "type": "operating-system",
      "name": "Ubuntu",
      "version": "21.04",
      "purl": "pkg:Ubuntu@v21.04",
      "properties": [
        {
          "name": "architecture",
          "value": "arm64"
        },
        {
          "name": "kernel_version",
          "value": "6.2.8-200.fc37.aarch64"
        },
        {
          "name": "operating_system",
          "value": "linux"
        }
      ]
    },
    {
      "bom-ref": "pkg:kubelet@v1.21.1",
      "type": "library",
      "name": "kubelet",
      "version": "6.2.8-200.fc37.aarch64",
      "purl": "pkg:kubelet@v1.21.1"
    },
    {
      "bom-ref": "pkg:kube-proxy@v6.2.8-200.fc37.aarch64",
      "type": "library",
      "name": "kube-proxy",
      "version": "6.2.8-200.fc37.aarch64",
      "purl": "pkg:kube-proxy@v6.2.8-200.fc37.aarch64"
    },
    {
      "bom-ref": "pkg:containerd@v1.5.2",
      "type": "library",
      "name": "containerd",
      "version": "1.5.2",
      "purl": "pkg:containerd@v1.5.2"
    },
    {
      "bom-ref": "pkg:kind-worker",
      "type": "container",
      "name": "kind-worker",
      "purl": "pkg:kind-worker",
      "properties": [
        {
          "name": "node-role",
          "value": "worker"
        },
        {
          "name": "host_name",
          "value": "kind-worker"
        }
      ]
    },
    {
      "bom-ref": "pkg:Ubuntu@v21.04",
      "type": "operating-system",
      "name": "Ubuntu",
      "version": "21.04",
      "purl": "pkg:Ubuntu@v21.04",
      "properties": [
        {
          "name": "architecture",
          "value": "arm64"
        },
        {
          "name": "kernel_version",
          "value": "6.2.8-200.fc37.aarch64"
        },
        {
          "name": "operating_system",
          "value": "linux"
        }
      ]
    },
    {
      "bom-ref": "pkg:kubelet@v1.21.1",
      "type": "library",
      "name": "kubelet",
      "version": "6.2.8-200.fc37.aarch64",
      "purl": "pkg:kubelet@v1.21.1"
    },
    {
      "bom-ref": "pkg:kube-proxy@v6.2.8-200.fc37.aarch64",
      "type": "library",
      "name": "kube-proxy",
      "version": "6.2.8-200.fc37.aarch64",
      "purl": "pkg:kube-proxy@v6.2.8-200.fc37.aarch64"
    },
    {
      "bom-ref": "pkg:containerd@v1.5.2",
      "type": "library",
      "name": "containerd",
      "version": "1.5.2",
      "purl": "pkg:containerd@v1.5.2"
    },
    {
      "bom-ref": "pkg:kind-worker2",
      "type": "container",
      "name": "kind-worker2",
      "purl": "pkg:kind-worker2",
      "properties": [
        {
          "name": "node-role",
          "value": "worker"
        },
        {
          "name": "host_name",
          "value": "kind-worker2"
        }
      ]
    },
    {
      "bom-ref": "pkg:Ubuntu@v21.04",
      "type": "operating-system",
      "name": "Ubuntu",
      "version": "21.04",
      "purl": "pkg:Ubuntu@v21.04",
      "properties": [
        {
          "name": "architecture",
          "value": "arm64"
        },
        {
          "name": "kernel_version",
          "value": "6.2.8-200.fc37.aarch64"
        },
        {
          "name": "operating_system",
          "value": "linux"
        }
      ]
    },
    {
      "bom-ref": "pkg:kubelet@v1.21.1",
      "type": "library",
      "name": "kubelet",
      "version": "6.2.8-200.fc37.aarch64",
      "purl": "pkg:kubelet@v1.21.1"
    },
    {
      "bom-ref": "pkg:kube-proxy@v6.2.8-200.fc37.aarch64",
      "type": "library",
      "name": "kube-proxy",
      "version": "6.2.8-200.fc37.aarch64",
      "purl": "pkg:kube-proxy@v6.2.8-200.fc37.aarch64"
    },
    {
      "bom-ref": "pkg:containerd@v1.5.2",
      "type": "library",
      "name": "containerd",
      "version": "1.5.2",
      "purl": "pkg:containerd@v1.5.2"
    }
  ],
  "dependencies": [
    {
      "ref": "kind-kind",
      "dependsOn": [
        "pkg:oci/etcd@05b738aa1bc6355db8a2ee8639f3631b908286e43f584a3d2ee0c472de033c28?repository_url=k8s.gcr.io/library/etcd",
        "pkg:oci/kube-apiserver@18e61c783b41758dd391ab901366ec3546b26fae00eef7e223d1f94da808e02f?repository_url=k8s.gcr.io/library/kube-apiserver",
        "pkg:oci/kube-controller-manager@0c6dccae49de8003ee4fa06db04a9f13bb46cbaad03977e6baa21174f2dba2fc?repository_url=k8s.gcr.io/library/kube-controller-manager",
        "pkg:oci/kube-scheduler@8c783dd2520887cc8e7908489ffc9f356c82436ba0411d554237a0b9632c9b87?repository_url=k8s.gcr.io/library/kube-scheduler",
        "pkg:oci/coredns/coredns@1a1f05a2cd7c2fbfa7b45b21128c8a4880c003ca482460081dc12d76bfa863e8?repository_url=k8s.gcr.io/library/coredns/coredns",
        "pkg:oci/coredns/coredns@1a1f05a2cd7c2fbfa7b45b21128c8a4880c003ca482460081dc12d76bfa863e8?repository_url=k8s.gcr.io/library/coredns/coredns",
        "pkg:oci/kindest/kindnetd@f37b7c809e5dcc2090371f933f7acb726bb1bffd5652980d2e1d7e2eff5cd301?repository_url=index.docker.io/library/kindest/kindnetd",
        "pkg:oci/kindest/kindnetd@f37b7c809e5dcc2090371f933f7acb726bb1bffd5652980d2e1d7e2eff5cd301?repository_url=index.docker.io/library/kindest/kindnetd",
        "pkg:oci/kindest/kindnetd@f37b7c809e5dcc2090371f933f7acb726bb1bffd5652980d2e1d7e2eff5cd301?repository_url=index.docker.io/library/kindest/kindnetd",
        "pkg:oci/kube-proxy@4bbef4ca108cdc3b99fe23d487fa4fca933a62c4fc720626a3706df9cef63b21?repository_url=k8s.gcr.io/library/kube-proxy",
        "pkg:oci/kube-proxy@4bbef4ca108cdc3b99fe23d487fa4fca933a62c4fc720626a3706df9cef63b21?repository_url=k8s.gcr.io/library/kube-proxy",
        "pkg:oci/kube-proxy@4bbef4ca108cdc3b99fe23d487fa4fca933a62c4fc720626a3706df9cef63b21?repository_url=k8s.gcr.io/library/kube-proxy"
      ]
    },
    {
      "ref": "pkg:kind-control-plane",
      "dependsOn": [
        "pkg:Ubuntu@v21.04"
      ]
    },
    {
      "ref": "pkg:Ubuntu@v21.04",
      "dependsOn": [
        "pkg:kubelet@v1.21.1"
      ]
    },
    {
      "ref": "pkg:Ubuntu@v21.04",
      "dependsOn": [
        "pkg:kube-proxy@v6.2.8-200.fc37.aarch64"
      ]
    },
    {
      "ref": "pkg:Ubuntu@v21.04",
      "dependsOn": [
        "pkg:containerd@v1.5.2"
      ]
    },
    {
      "ref": "pkg:kind-worker",
      "dependsOn": [
        "pkg:Ubuntu@v21.04"
      ]
    },
    {
      "ref": "pkg:Ubuntu@v21.04",
      "dependsOn": [
        "pkg:kubelet@v1.21.1"
      ]
    },
    {
      "ref": "pkg:Ubuntu@v21.04",
      "dependsOn": [
        "pkg:kube-proxy@v6.2.8-200.fc37.aarch64"
      ]
    },
    {
      "ref": "pkg:Ubuntu@v21.04",
      "dependsOn": [
        "pkg:containerd@v1.5.2"
      ]
    },
    {
      "ref": "pkg:kind-worker2",
      "dependsOn": [
        "pkg:Ubuntu@v21.04"
      ]
    },
    {
      "ref": "pkg:Ubuntu@v21.04",
      "dependsOn": [
        "pkg:kubelet@v1.21.1"
      ]
    },
    {
      "ref": "pkg:Ubuntu@v21.04",
      "dependsOn": [
        "pkg:kube-proxy@v6.2.8-200.fc37.aarch64"
      ]
    },
    {
      "ref": "pkg:Ubuntu@v21.04",
      "dependsOn": [
        "pkg:containerd@v1.5.2"
      ]
    }
  ]
}
```
