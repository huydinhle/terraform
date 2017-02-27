package terraform

import (
	"fmt"

	"github.com/hashicorp/terraform/config"
)

// NodeTaintableResource represents a resource that is "taintable":
// it is ready to be applied and is represented by a diff.
type NodeTaintableResource struct {
	*NodeAbstractResource
}

// GraphNodeDestroyer
func (n *NodeTaintableResource) DestroyAddr() *ResourceAddress {
	return n.Addr
}

// GraphNodeEvalable
func (n *NodeTaintableResource) EvalTree() EvalNode {
	// Eval info is different depending on what kind of resource this is
	switch mode := n.Addr.Mode; mode {
	case config.ManagedResourceMode:
		return n.evalTreeManagedResource()

	default:
		panic(fmt.Errorf("unsupported resource mode %s", mode))
	}
}

func (n *NodeTaintableResource) evalTreeManagedResource() EvalNode {
	addr := n.NodeAbstractResource.Addr

	// stateId is the ID to put into the state
	stateId := addr.stateId()

	// Build the instance info. More of this will be populated during eval
	info := &InstanceInfo{
		Id:   stateId,
		Type: addr.Type,
	}

	// Declare a bunch of variables that are used for state during
	// evaluation. Most of this are written to by-address below.
	var provider ResourceProvider
	var state *InstanceState

	// This happened during initial development. All known cases were
	// fixed and tested but as a sanity check let's assert here.
	if n.ResourceState == nil {
		err := fmt.Errorf(
			"No resource state attached for addr: %s\n\n"+
				"This is a bug. Please report this to Terraform with your configuration\n"+
				"and state attached. Please be careful to scrub any sensitive information.",
			addr)
		return &EvalReturnError{Error: &err}
	}

	return &EvalSequence{
		Nodes: []EvalNode{
			&EvalGetProvider{
				Name:   n.ProvidedBy()[0],
				Output: &provider,
			},
			&EvalReadState{
				Name:   stateId,
				Output: &state,
			},
			&EvalTaint{
				Info:     info,
				Provider: &provider,
				State:    &state,
				Output:   &state,
			},
			&EvalWriteState{
				Name:         stateId,
				ResourceType: n.ResourceState.Type,
				Provider:     n.ResourceState.Provider,
				Dependencies: n.ResourceState.Dependencies,
				State:        &state,
			},
		},
	}
}
