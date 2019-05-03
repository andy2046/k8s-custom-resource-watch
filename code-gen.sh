# ROOT_PACKAGE the target package (relative to $GOPATH/src) for code generation
ROOT_PACKAGE="github.com/andy2046/k8s-custom-resource-watch"
# CUSTOM_RESOURCE_NAME the name of the custom resource which the client code is generated for
CUSTOM_RESOURCE_NAME="customresource"
# CUSTOM_RESOURCE_VERSION the version of the resource
CUSTOM_RESOURCE_VERSION="v1"

# retrieve the code-generator scripts and bins
go get -u k8s.io/code-generator/...
cd $GOPATH/src/k8s.io/code-generator

# run the code-generator entrypoint script
./generate-groups.sh all "$ROOT_PACKAGE/pkg/client" "$ROOT_PACKAGE/pkg/apis" \
    "$CUSTOM_RESOURCE_NAME:$CUSTOM_RESOURCE_VERSION"

# view the newly generated files
tree $GOPATH/src/$ROOT_PACKAGE/pkg/client
