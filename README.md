# CNPG-I Hello World Plugin

A [CNPG-I](https://github.com/cloudnative-pg/cnpg-i) plugin to easily
patch the Pods created by CNPG using a PodTemplate object.

This plugin uses
the [pluginhelper](https://github.com/cloudnative-pg/cnpg-i-machinery/tree/main/pkg/pluginhelper)
from [`cnpg-i-machinery`](https://github.com/cloudnative-pg/cnpg-i-machinery) to
simplify the plugin's implementation.

## Show me how it works

``` yaml
apiVersion: postgresql.cnpg.io/v1
kind: Cluster
metadata:
  name: cluster-example
  namespace: default
spec:
  instances: 3

  plugins:
  - name: cnpg-i-podspec.cloudnative-pg.io
    parameters:
      instanceTemplateName: instance-pod
      instanceTemplateNamespace: cnpg-system

  storage:
    size: 1Gi
---
apiVersion: v1
kind: PodTemplate
metadata:
  namespace: cnpg-system
  name: instance-pod
template:
  metadata:
    annotations:
      test: tests
      tooth: teeth
  spec:
    containers:
      - name: postgres
        image: ghcr.io/cloudnative-pg/postgresql:18.0-system-trixie
        env:
          - name: CONTAINER_NAME
            value: postgres

```

## Running the plugin

To install it, start with
[CloudNativePG](https://github.com/cloudnative-pg/cloudnative-pg)
checkout and create a development environment with:

``` shell
cloudnative-pg $ ./hack/setup-cluster -n 1 -r create load deploy
...
```

After that you need to install Cert-Manager:

```shell
# Choose the latest version of cert-manager
kubectl apply -f \
  https://github.com/cert-manager/cert-manager/releases/download/vX.Y.Z/cert-manager.yaml
```

Then install the plugin by compiling it:

```
cnpg-i-podspec $ task local-kind-deploy
```

Finally, create a cluster resource to see the plugin in action.

```
cnpg-i-podspec $ kubectl apply -f doc/examples/cluster-example.yaml
```

## Plugin development

For additional details on the plugin implementation refer to
the [development documentation](doc/development.md).
