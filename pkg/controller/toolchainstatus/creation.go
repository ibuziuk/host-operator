package toolchainstatus

import (
	"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1"
	commonclient "github.com/codeready-toolchain/toolchain-common/pkg/client"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// CreateOrUpdateResources creates a toolchainstatus resource with the given name in the given namespace
func CreateOrUpdateResources(client client.Client, s *runtime.Scheme, namespace, toolchainStatusName string) error {
	toolchainStatus := &v1alpha1.ToolchainStatus{
		ObjectMeta: v1.ObjectMeta{
			Namespace: namespace,
			Name:      toolchainStatusName,
		},
		Spec: v1alpha1.ToolchainStatusSpec{},
	}
	cl := commonclient.NewApplyClient(client, s)
	_, err := cl.CreateOrUpdateObject(toolchainStatus, false, nil)
	return err
}
