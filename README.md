# kubernetes-configmap-exporter
Exports Kubernetes configmap to a directory from within a Kubernetes container to a directory as files.
The config map keys are the filenames with corresponding values as the file content.

# Usage

```sh
kubernetes-configmap-exporter <namespace> <configmap-name> <output-directory>
```
