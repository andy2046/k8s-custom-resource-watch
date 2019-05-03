package handler

import (
	"log"
	"os"

	coreV1 "k8s.io/api/core/v1"
)

var (
	logger = log.New(os.Stdout, "handler:", log.LstdFlags)
)

// Handler interface.
type Handler interface {
	ObjectCreated(obj interface{})
	ObjectDeleted(obj interface{})
	ObjectUpdated(objOld, objNew interface{})
}

// KubeHandler implements Handler interface.
type KubeHandler struct{}

// ObjectCreated for object creation.
func (t *KubeHandler) ObjectCreated(obj interface{}) {
	logger.Println("KubeHandler.ObjectCreated")
	// assert the type to a Pod object
	pod, ok := obj.(*coreV1.Pod)
	if ok {
		logger.Printf("    Name: %s", pod.Name)
		logger.Printf("    NodeName: %s", pod.Spec.NodeName)
		logger.Printf("    Phase: %s", pod.Status.Phase)
	}
}

// ObjectDeleted for object deletion.
func (t *KubeHandler) ObjectDeleted(obj interface{}) {
	logger.Println("KubeHandler.ObjectDeleted")
}

// ObjectUpdated for object update.
func (t *KubeHandler) ObjectUpdated(objOld, objNew interface{}) {
	logger.Println("KubeHandler.ObjectUpdated")
}
