package params

type ClusterConfig struct {
	Name        string `json:"name"`
	ClusterType string `json:"type"`
	KubeConfig  string `json:"kubeConfig"`
}
