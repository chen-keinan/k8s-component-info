# k8s-component-info

## Description
k8s-component-info is an open source project which collect component and version info from runnning k8s cluster and produce k8s bill of materials.

```sh
go build main.go

./main
```

```json
{
  "$schema": "http://cyclonedx.org/schema/bom-1.4.schema.json",
  "bomFormat": "CycloneDX",
  "specVersion": "1.4",
  "version": 1,
  "metadata": {
    "timestamp": "2023-04-04T16:46:25+03:00",
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
      "bom-ref": "pkg:oci/kube-proxy@4bbef4ca108cdc3b99fe23d487fa4fca933a62c4fc720626a3706df9cef63b21?repository_url=k8s.gcr.io/library/kube-proxy",
      "type": "container",
      "name": "kube-proxy",
      "version": "v1.21.1",
      "purl": "pkg:oci/kube-proxy@4bbef4ca108cdc3b99fe23d487fa4fca933a62c4fc720626a3706df9cef63b21?repository_url=k8s.gcr.io/library/kube-proxy"
    },
    {
      "type": "container",
      "name": "kind-control-plane",
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
          "value": "6.1.18-200.fc37.aarch64"
        },
        {
          "name": "operating_system",
          "value": "linux"
        }
      ]
    },
    {
      "bom-ref": "pkg:kube-proxy@v1.21.1",
      "type": "library",
      "name": "kubelet",
      "version": "6.1.18-200.fc37.aarch64",
      "purl": "pkg:kube-proxy@v1.21.1"
    },
    {
      "bom-ref": "pkg:kube-proxy@v6.1.18-200.fc37.aarch64",
      "type": "library",
      "name": "kube-proxy",
      "version": "6.1.18-200.fc37.aarch64",
      "purl": "pkg:kube-proxy@v6.1.18-200.fc37.aarch64"
    },
    {
      "bom-ref": "pkg:Ubuntu@v21.04",
      "type": "library",
      "name": "containerd",
      "version": "1.5.2",
      "purl": "pkg:Ubuntu@v21.04"
    }
  ],
  "dependencies": [
    {
      "ref": "pkg:kind-kind:1.21.1",
      "dependsOn": [
        "pkg:oci/etcd@05b738aa1bc6355db8a2ee8639f3631b908286e43f584a3d2ee0c472de033c28?repository_url=k8s.gcr.io/library/etcd",
        "pkg:oci/kube-apiserver@18e61c783b41758dd391ab901366ec3546b26fae00eef7e223d1f94da808e02f?repository_url=k8s.gcr.io/library/kube-apiserver",
        "pkg:oci/kube-controller-manager@0c6dccae49de8003ee4fa06db04a9f13bb46cbaad03977e6baa21174f2dba2fc?repository_url=k8s.gcr.io/library/kube-controller-manager",
        "pkg:oci/kube-scheduler@8c783dd2520887cc8e7908489ffc9f356c82436ba0411d554237a0b9632c9b87?repository_url=k8s.gcr.io/library/kube-scheduler",
        "pkg:oci/coredns/coredns@1a1f05a2cd7c2fbfa7b45b21128c8a4880c003ca482460081dc12d76bfa863e8?repository_url=k8s.gcr.io/library/coredns/coredns",
        "pkg:oci/coredns/coredns@1a1f05a2cd7c2fbfa7b45b21128c8a4880c003ca482460081dc12d76bfa863e8?repository_url=k8s.gcr.io/library/coredns/coredns",
        "pkg:oci/kindest/kindnetd@f37b7c809e5dcc2090371f933f7acb726bb1bffd5652980d2e1d7e2eff5cd301?repository_url=index.docker.io/library/kindest/kindnetd",
        "pkg:oci/kube-proxy@4bbef4ca108cdc3b99fe23d487fa4fca933a62c4fc720626a3706df9cef63b21?repository_url=k8s.gcr.io/library/kube-proxy"
      ]
    },
    {
      "ref": "",
      "dependsOn": [
        "pkg:Ubuntu@v21.04"
      ]
    },
    {
      "ref": "pkg:Ubuntu@v21.04",
      "dependsOn": [
        "pkg:kube-proxy@v1.21.1"
      ]
    },
    {
      "ref": "pkg:Ubuntu@v21.04",
      "dependsOn": [
        "pkg:kube-proxy@v6.1.18-200.fc37.aarch64"
      ]
    },
    {
      "ref": "pkg:Ubuntu@v21.04",
      "dependsOn": [
        "pkg:Ubuntu@v21.04"
      ]
    }
  ]
}
```
