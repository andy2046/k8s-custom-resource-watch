# k8s-custom-resource-watch

> a custom K8s controller example to watch for the creation, update, deletion of user defined custom resources.

## Workflow

### 1 Define custom resource

#### 1.1 define API group name / Version / Resource name
* The API group name `nokube.xyz`
* The Version `v1`
* Resource name `customresource`

#### 1.2 create the directory structure
```sh
mkdir -p pkg/apis/customresource/v1
```

#### 1.3 add API group name const in register file
```sh
touch pkg/apis/customresource/register.go
```
> package has the same name as the custom resource
```go
package customresource

// GroupName for customresource
const GroupName = "nokube.xyz"
```

#### 1.4 create the resource structs
```sh
touch pkg/apis/customresource/v1/types.go
```
> `// +<tag_name>[=value]` are indicators for code generator

> `+genclient` — generate a client for the package

> `+genclient:noStatus` — when generating the client, there is no status stored for the package

> `+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object` — generate deepcopy logic (required) implementing the runtime.Object interface (for both `CustomResource` and `CustomResourceList`)

#### 1.5 create a doc source file for the package
```sh
touch pkg/apis/customresource/v1/doc.go
```
```go
// +k8s:deepcopy-gen=package
// +groupName=nokube.xyz

package v1
```
> `// +groupName=nokube.xyz` inform the generator what the API group name is

> `// +k8s:deepcopy-gen=package` deepcopy should be generated for all types in the package

#### 1.6 add functions to handle adding types to the schemes
```sh
touch pkg/apis/customresource/v1/register.go
```

### 2 Run the code generator
run `code-gen.sh` shell script to do all the heavy lifting via `k8s.io/code-generator` package

### 3 Wire up the generated code

#### 3.1 to return custom resource clientset instance to interact with `CustomResource` in `main.go`
```go
// retrieve the Kubernetes cluster client from outside of the cluster.
func getKubeClient() (kubernetes.Interface, resourceclientset.Interface) {
	var kubeConfigPath string
	if !inCluster {
		// resolve path to `$HOME/.kube/config`
		kubeConfigPath = path.Join(userHomeDir(), "/.kube/config")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		logger.Fatalf("BuildConfigFromFlags: %v", err)
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Fatalf("client NewForConfig: %v", err)
	}

	resourceClient, err := resourceclientset.NewForConfig(config)
	if err != nil {
		logger.Fatalf("resourceClient NewForConfig: %v", err)
	}

	logger.Println("Successfully get k8s client")
	return client, resourceClient
}
```

#### 3.2 to return custom resource informer instance in `controller.go`
```go
// New creates a Controller instance.
func New(client kubernetes.Interface, resourceClient resourceclientset.Interface) *Controller {
	informer := resourceinformerV1.NewCustomResourceInformer(
		resourceClient,
		nameSpace,
		0,
		cache.Indexers{},
    )
    
    // ...
}
```

### 4 Add Custom Resource Definition
```sh
mkdir crd
# touch crd/customresource.yml
kubectl apply -f crd/customresource.yml
kubectl get customresourcedefinition
```

### 5 Run the controller
```sh
go run main.go
# touch customresource-deploy.yml
kubectl apply -f customresource-deploy.yml
kubectl get CustomResource
```

### 6 Clean up
```sh
kubectl delete CustomResource example-customresource
kubectl delete -f crd/customresource.yml
```
