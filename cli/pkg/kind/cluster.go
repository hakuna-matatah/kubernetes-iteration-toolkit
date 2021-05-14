package kind

import (
	"fmt"
	"os"

	"github.com/prateekgogia/kit/cli/pkg/helm"
	"github.com/prateekgogia/kit/cli/pkg/util/functional"
	"github.com/prateekgogia/kit/cli/pkg/util/logging"
	"k8s.io/api/node/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	controllerruntime "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	kind "sigs.k8s.io/kind/pkg/cluster"
)

var log = logging.NewNamedLogger("Kind")

var (
	scheme = runtime.NewScheme()
)

func init() {
	for _, builder := range []func(*runtime.Scheme) error{
		clientgoscheme.AddToScheme,
		v1alpha1.AddToScheme,
	} {
		if err := builder(scheme); err != nil {
			log.Fatalf("Failed to build client go scheme, %w", err)
		}
	}
}

// Cluster wraps a Kind cluster
type Cluster struct {
	Name       string
	provider   *kind.Provider
	kubeClient client.Client
	helmClient *helm.Client
}

// NewCluster idempotently gets an existing cluster or creates a new one.
func NewCluster(name string) (*Cluster, error) {
	provider := kind.NewProvider()

	// 1. List clusters
	clusters, err := provider.List()
	if err != nil {
		return nil, fmt.Errorf("Failed to list kind clusters, %s", err.Error())
	}

	// 2. Create cluster if not exists
	if !functional.ContainsString(clusters, name) {
		log.Infof("Creating kind cluster %s", name)
		if err := provider.Create(name); err != nil {
			return nil, fmt.Errorf("Failed to create kind cluster %s, %s", name, err.Error())
		}
	}

	// 3. Connect to to cluster
	if err := provider.ExportKubeConfig(name, os.Getenv("KUBECONFIG")); err != nil {
		return nil, fmt.Errorf("Failed to export kubeconfig for cluster %s, %s", name, err.Error())
	}

	// 4. Construct kube client
	kubeClient, err := client.New(controllerruntime.GetConfigOrDie(), client.Options{})
	if err != nil {
		return nil, fmt.Errorf("Failed to build kube client for cluster %s, %s", name, err.Error())
	}

	// 5. Construct Helm Client
	helmClient, err := helm.NewClient()
	if err != nil {
		return nil, fmt.Errorf("Failed to build helm client for cluster %s, %s", name, err.Error())
	}

	return &Cluster{
		Name:       name,
		provider:   provider,
		kubeClient: kubeClient,
		helmClient: helmClient,
	}, nil
}

// ApplyYAML applies a file or directory of files to the cluster.
func (c *Cluster) ApplyYAML(path string) error {
	// TODO
	return nil
}

// ApplyChart applies a helm chart to the cluster.
func (c *Cluster) ApplyChart(options helm.ApplyOptions) error {
	return c.helmClient.Apply(options)
}

// Stop stops the cluster and cleans up all resources.
func (c *Cluster) Stop() error {
	if err := c.provider.Delete(c.Name, os.Getenv("KUBECONFIG")); err != nil {
		return fmt.Errorf("failed to stop kind cluster %s, %w", c.Name, err)
	}
	return nil
}
