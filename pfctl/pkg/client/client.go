package client

import (

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	restclient "k8s.io/client-go/rest"
	"os"
	"path/filepath"

	"k8s.io/client-go/tools/clientcmd"
)

// Client object for all operations
type Client struct {
	clientsetGetter
	contextsGetter
}

// GetPlaceholderClient returns an empty client
func GetPlaceholderClient() *Client {
	return &Client{}
}

func clientsetHelper(getConfig func() (*restclient.Config, error)) (kubernetes.Interface, error) {
	config, err := getConfig()

	if err != nil {
		return nil, err
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	return clientset, err
}

// GetConfigClient returns a client with a specific kubeconfig path
func GetConfigClient(path string) *Client {
	contexts := GetContexts(path)
	return &Client{
		clientsetGetter: &configClientsetGetter{
			clientsets: make(map[string]clusterFunctionality),
			config:     path,
		},
		contextsGetter: StaticContextsGetter{
			contexts: contexts,
		},

	}
}

// GetKubeConfigPath returns the default location of a kubeconfig file
func GetKubeConfigPath() string {
	if fl := os.Getenv("KUBECONFIG"); fl != "" {
		return fl
	}
	home, err := os.UserHomeDir()
	if err != nil { // Can't find home dir
		panic(err.Error())
	}

	return filepath.Join(home, ".kube", "config")
}

// GetContexts returns a list of clusters from a config file
func GetContexts(configpath string) []string {
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: configpath},
		&clientcmd.ConfigOverrides{}).RawConfig()

	if err != nil {
		panic(err.Error())
	}

	ctxs := make([]string, 0, len(config.Contexts))
	for k := range config.Contexts {
		ctxs = append(ctxs, k)
	}

	return ctxs
}