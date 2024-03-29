package constants

const (
	LabelValueManagedBy   = "mcs.nauti.io"
	OriginName            = "origin-name"
	OriginNamespace       = "origin-namespace"
	LabelSourceNamespace  = "submariner-io/originatingNamespace"
	LabelSourceCluster    = "syncer.nri.io/sourceCluster"
	LabelSourceName       = "syncer.io/sourceName"
	LabelOriginNameSpace  = "syncer.io/sourceNamespace"
	MCSLabelServiceName   = "multicluster.kubernetes.io/service-name"
	MCSLabelSourceCluster = "multicluster.kubernetes.io/source-cluster"
	KubernetesServiceName = "kubernetes.io/service-name"
)

const (
	MaxNamespaceLength = 10
	MaxNameLength      = 10
	MaxClusternName    = 10
)

type RouteOperation int

const (
	Add RouteOperation = iota
	Delete
)
