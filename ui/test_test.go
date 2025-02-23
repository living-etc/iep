package ui_test

import (
	"testing"

	"github.com/living-etc/iep/ui"
)

func TestRun(t *testing.T) {
	testcases := []struct {
		name string
		test ui.Test
	}{
		{
			name: "run a test",
			test: ui.Test{
				Id:                     1,
				Name:                   "nginx is installed",
				ExerciseId:             "0001-deploy-a-webapp",
				ResourceType:           "Package",
				ResourceName:           "nginx",
				ResourceAttribute:      "Status",
				ResourceAttributeValue: "install ok installed",
				Negation:               true,
				Result:                 false,
			},
		},
	}

	_ = testcases
}
