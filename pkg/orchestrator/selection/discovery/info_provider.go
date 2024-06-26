package discovery

import (
	"context"

	"github.com/bacalhau-project/bacalhau/pkg/model"
	"github.com/bacalhau-project/bacalhau/pkg/orchestrator"
)

type discoveredNodesProvider struct {
	discoverer orchestrator.NodeDiscoverer
}

func NewDebugInfoProvider(discoverer orchestrator.NodeDiscoverer) model.DebugInfoProvider {
	return &discoveredNodesProvider{discoverer: discoverer}
}

// GetDebugInfo implements models.DebugInfoProvider
func (p *discoveredNodesProvider) GetDebugInfo(ctx context.Context) (info model.DebugInfo, err error) {
	nodes, err := p.discoverer.List(ctx)
	info.Component = "DiscoveredNodes"
	info.Info = nodes
	return info, err
}

var _ model.DebugInfoProvider = (*discoveredNodesProvider)(nil)
