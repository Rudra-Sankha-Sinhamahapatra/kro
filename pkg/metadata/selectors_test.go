package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// testObject is a mock implementation specific to selector tests
type testObject struct {
	metav1.ObjectMeta
}

func newTestObject(name string, uid types.UID) *testObject {
	return &testObject{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			UID:  uid,
		},
	}
}

func TestNewInstanceSelector(t *testing.T) {
	instance := newTestObject("test-instance", "instance-123")
	selector := NewInstanceSelector(instance)

	assert.Equal(t, map[string]string{
		InstanceIDLabel: "instance-123",
	}, selector.MatchLabels)
}

func TestNewResourceGraphDefinitionSelector(t *testing.T) {
	rgd := newTestObject("test-rgd", "rgd-123")
	selector := NewResourceGraphDefinitionSelector(rgd)

	assert.Equal(t, map[string]string{
		ResourceGraphDefinitionIDLabel: "rgd-123",
	}, selector.MatchLabels)
}

func TestNewInstanceAndResourceGraphDefinitionSelector(t *testing.T) {
	instance := newTestObject("test-instance", "instance-123")
	rgd := newTestObject("test-rgd", "rgd-123")

	selector := NewInstanceAndResourceGraphDefinitionSelector(instance, rgd)

	assert.Equal(t, map[string]string{
		InstanceIDLabel:                "instance-123",
		ResourceGraphDefinitionIDLabel: "rgd-123",
	}, selector.MatchLabels)
}

func TestNewNodeAndInstanceAndResourceGraphDefinitionSelector(t *testing.T) {
	node := newTestObject("test-node", "node-123")
	instance := newTestObject("test-instance", "instance-123")
	rgd := newTestObject("test-rgd", "rgd-123")

	selector := NewNodeAndInstanceAndResourceGraphDefinitionSelector(node, instance, rgd)

	assert.Equal(t, map[string]string{
		NodeIDLabel:                    "test-node",
		InstanceIDLabel:                "instance-123",
		ResourceGraphDefinitionIDLabel: "rgd-123",
	}, selector.MatchLabels)
}
