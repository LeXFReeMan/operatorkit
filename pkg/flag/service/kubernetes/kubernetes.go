package kubernetes

import (
	"github.com/LeXFReeMan/operatorkit/v7/pkg/flag/service/kubernetes/tls"
	"github.com/LeXFReeMan/operatorkit/v7/pkg/flag/service/kubernetes/watch"
)

// Kubernetes is a data structure to hold Kubernetes specific command line
// configuration flags.
type Kubernetes struct {
	Address        string
	InCluster      string
	KubeConfig     string
	KubeConfigPath string
	TLS            tls.TLS
	Watch          watch.Watch
}
