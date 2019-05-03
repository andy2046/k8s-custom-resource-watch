package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"os/user"
	"path"
	"syscall"

	"github.com/andy2046/k8s-custom-resource-watch/internal/controller"
	resourceclientset "github.com/andy2046/k8s-custom-resource-watch/pkg/client/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	logger         = log.New(os.Stdout, "main:", log.LstdFlags)
	usageInCluster = "flag if it's in cluster or outside of the cluster"
	inCluster      bool
)

func init() {
	flag.BoolVar(&inCluster, "in-cluster",
		getEnv("kube-in-cluster"), usageInCluster)
}

func main() {
	client, resourceClient := getKubeClient()
	controller := controller.New(client, resourceClient)
	stopCh := make(chan struct{})
	defer close(stopCh)

	// run the controller to process items
	go controller.Run(stopCh)

	stopSig := make(chan os.Signal, 1)
	signal.Notify(stopSig, syscall.SIGTERM, syscall.SIGINT)
	<-stopSig
}

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

func userHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		logger.Printf("userHomeDir: %v", err)
		return os.Getenv("HOME")
	}
	return usr.HomeDir
}

func getEnv(key string) bool {
	if _, ok := os.LookupEnv(key); ok {
		return true
	}
	return false
}
