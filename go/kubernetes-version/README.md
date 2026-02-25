This example demonstrates fetching the Kubernetes version using the client-go
library. It is designed to be run either within or external to a kubernetes
cluster.  The example exposes HTTP on port 8080 and reports the Kubernetes
version there.  Each query to the HTTP endpoint will trigger a new fetch of the
Kubernetes version.
