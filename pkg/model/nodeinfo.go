package model

import (
	"context"
)

//go:generate stringer -type=NodeType -trimprefix=NodeType -output=nodeinfo_string.go
type NodeType int

const (
	NodeTypeRequester NodeType = iota
	NodeTypeCompute
)

type NodeInfoProvider interface {
	GetNodeInfo(ctx context.Context) NodeInfo
}

type ComputeNodeInfoProvider interface {
	GetComputeInfo(ctx context.Context) ComputeNodeInfo
}

type NodeInfo struct {
	BacalhauVersion BuildVersionInfo  `json:"BacalhauVersion"`
	NodeType        NodeType          `json:"NodeType"`
	Labels          map[string]string `json:"Labels"`
	ComputeNodeInfo *ComputeNodeInfo  `json:"ComputeNodeInfo"`
}

// IsComputeNode returns true if the node is a compute node
func (n NodeInfo) IsComputeNode() bool {
	return n.NodeType == NodeTypeCompute
}

type ComputeNodeInfo struct {
	ExecutionEngines   []string          `json:"ExecutionEngines"`
	Publishers         []string          `json:"Publishers"`
	StorageSources     []string          `json:"StorageSources"`
	MaxCapacity        ResourceUsageData `json:"MaxCapacity"`
	AvailableCapacity  ResourceUsageData `json:"AvailableCapacity"`
	MaxJobRequirements ResourceUsageData `json:"MaxJobRequirements"`
	RunningExecutions  int               `json:"RunningExecutions"`
	EnqueuedExecutions int               `json:"EnqueuedExecutions"`
}
