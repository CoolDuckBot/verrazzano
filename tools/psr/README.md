# PSR Testing Tool

The PSR testing tool is used to test Verrazzano performance, scalability, and reliability.  The tool works by doing some
work on the Verrazzano cluster, then collecting and analyzing results.  The type of work that is done depends on the goal of the test,
and it can be done at any layer.  For example, you could test OpenSearch REST API while doing a Verrazzano upgrade, or even
a Kubernetes upgrade.

## Concepts
### Backend
The tool consists of backend pods that are deployed using Helm charts.  The backend consists of a single image
that has all the worker code.  When a pod is started, the worker config, is passed in as a set of Env vars.
The main.go code in the pod gets the worker type, creates an instance of the worker, then invokes it
to run.  The pod only runs a single worker, which executes until the pod terminates, by default.

### Use Cases
The term `use case` just describes what the worker does.  Each backend pod runs exactly one `use case` at a time, 
in the context of a worker. Each use case is a focused task, like log generation, making HTTP GET requests against 
some endpoint, or repeatedly scaling a component.

To run a use case, just do a Helm install or upgrade.  Each use case has a set of Env vars stored in override files
that define the configuration. For example, to deploy a use case to 
generate logs using 10 replicas, you would run the following command:
```
helm install psr-log-gen manifests/charts/k8s -f manifests/usecases/loggen.yaml --set replicas=10
```
**NOTE** The worker code for a use case expects all dependencies to exist at the time of execution.  This is the job of the scenario, 
described next.  Workers should not know about other workers, or explicitly depend on them (in the worker code itself).  With that said,
a worker could be written to gracefully wait until some required dependency was ready before doing the work, but it would never create or
provision a dependency.

### Scenarios
Scenarios are a combination of one or more use cases running concurrently.
Consider the scenario where one use case generates logs and another that scales OpenSearch out and in, repeatedly.  

To run a scenario, you install two Helm releases as follows (note some use cases might not exist yet): 
```
helm install psr-log-gen manifests/charts/k8s -f manifests/usecases/opensearch/loggen.yaml --set replicas=5
helm install psr-log-gen manifests/charts/k8s -f manifests/usecases/opensearch/scale-out-in.yaml --set replicas=1
```
The override file for the scale use case might look like this:
```
envVars:
  PSR_WORKER_TYPE: WT_OPENSEARCH_SCALE_OUT_IN
  PSR_OS_TIER: master
  PSR_MAX_PODS: 5
  PSR_MIN_PODS: 3  
  PSR_INTERATION_DELAY: 5s
  PSR_DURATION : 1h
```
The OpenSearch worker code would read those Env vars at run time and behave accordingly.

### Metrics
The workers themselves will generator metrics so that we know what kind of load they are putting on the system,
and to also have insight to their progress to be sure that they working as expected.

### Remote Execution
There is nothing preventing a worker accessing a cluster different than the one it is running in.  You
would just need to put the KUBECONFIG in a secret and pass the secret info to as an Env var.  However,
the default mode and intention is that these workers run in the Verrazzano cluster that is being measure.

## Usage
Use the Makefile to build the backend image or execute other targets.  If you just want to try 
the example use case on a local Kind cluster, then run the following which builds the code, builds the docker image, 
loads the docker image into the Kind cluster, and deploys the example use case.
```
make run-example-k8s
```
After you run the make command, run `helm list` to see the `psr` release.  The get the logs for backend pods in the default namespace 
to see that they are just logging then sleeping in a loop.

The other Make targets are:
* go-build - build the code
* docker-build - go-build, then build the image
* docker-push - docker-build, then push the image to ghcr.io. 
* kind-load-image - docker-build, then load to local Kind cluster
* run-example-k8s - kind-load-image, then deploy the example use case using Helm k8s chart

### Example Helm Installs

Install the example worker as a Kubernetes deployment with 10 replicas:
```
helm install  psr manifests/charts/k8s --set imageName=ghcr.io/verrazzano/psr-backend:local-582bfcfcf --replicas=10
```

Install the logging generator worker as an OAM application using the default image in the helm chart:
```
helm install  psr2 manifests/charts/oam -f manifests/helm/workers/opensearch.yaml
```

Install the example worker as a Kubernetes deployment using the default image in the helm chart, providing an imagePullSecret
```
helm install  psr-3  manifests/charts/k8s/ --set imagePullSecrets[0].name=verrazzano-container-registry
```

### The Backend Image
The backend image is a private image that should never be made public in ghcr.io.  The image is
built using scratch.

## Source Code
The `backend` directory has the go code which consists of a few packages:
* config - configuration code
* metrics - metrics server for metrics generated by the workers
* spi - the worker interface
* workers - the various workers

The `manifests/charts` directory has the Helm charts.  There is a chart for using plain Kubernetes
resources to deploy the backend, and there is a chart to deploy the backend as an oam app. Any use case 
can be deployed with either chart.

The `manifests/usecases` directory has the Helm override files for every use case. These files must
contain the configuration, as key:value pairs, required by the worker.
