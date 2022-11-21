package test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/filecoin-project/bacalhau/ops/aws/canary/pkg/models"
	"github.com/filecoin-project/bacalhau/ops/aws/canary/pkg/router"
	"github.com/filecoin-project/bacalhau/pkg/computenode"
	"github.com/filecoin-project/bacalhau/pkg/requesternode"
	"github.com/filecoin-project/bacalhau/pkg/test/devstack"
	"github.com/stretchr/testify/require"
)

func TestScenarios(t *testing.T) {
	stack, _ := devstack.SetupTest(context.Background(), t, 1, 0, false, computenode.NewDefaultComputeNodeConfig(), requesternode.NewDefaultRequesterNodeConfig())

	os.Setenv("BACALHAU_HOST", stack.Nodes[0].APIServer.Host)
	os.Setenv("BACALHAU_PORT", fmt.Sprint(stack.Nodes[0].APIServer.Port))

	for name := range router.TestcasesMap {
		t.Run(name, func(t *testing.T) {
			event := models.Event{Action: name}
			err := router.Route(context.Background(), event)
			require.NoError(t, err)
		})
	}
}