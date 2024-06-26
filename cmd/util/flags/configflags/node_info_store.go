package configflags

import (
	"github.com/bacalhau-project/bacalhau/pkg/config/types"
)

var NodeInfoStoreFlags = []Definition{
	{
		FlagName:     "node-info-store-ttl",
		ConfigPath:   types.NodeRequesterNodeInfoStoreTTL,
		DefaultValue: Default.Node.Requester.NodeInfoStoreTTL,
		Description: `Sets the duration for which node information is retained in the node info store after which it
is automatically removed from the store.`,
	},
}
