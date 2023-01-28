package types

type ClusterType string

const (
	ClusterKubernetes ClusterType = "kubernetes"
	ClusterOpenshift  ClusterType = "openshift"
)
