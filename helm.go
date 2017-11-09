package helm

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"k8s.io/helm/pkg/helm"
	"k8s.io/helm/pkg/helm/portforwarder"

	"fmt"
)

func getHelmClient() (*helm.Client, error) {
	var config *rest.Config
	config, _ = rest.InClusterConfig()
	client, _ := kubernetes.NewForConfig(config)
	tillerTunnel, _ := portforwarder.New("kube-system", client, config)
	tillerTunnelAddress := fmt.Sprintf("localhost:%d", tillerTunnel.Local)
	hclient := helm.NewClient(helm.Host(tillerTunnelAddress))
	return hclient, nil
}
