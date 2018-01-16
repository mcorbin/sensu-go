package agent

import (
	"encoding/json"
	"testing"

	"github.com/sensu/sensu-go/testing/testutil"
	"github.com/sensu/sensu-go/types"
	"github.com/stretchr/testify/assert"
)

func TestTokenSubstitution(t *testing.T) {
	testCases := []struct {
		name            string
		data            interface{}
		input           interface{}
		expectedCommand string
		expectedError   bool
	}{
		{
			name:            "empty data",
			data:            &types.Entity{},
			input:           *types.FixtureCheckConfig("check"),
			expectedCommand: "command",
			expectedError:   false,
		},
		{
			name:            "empty input",
			data:            types.FixtureEntity("entity"),
			input:           types.CheckConfig{},
			expectedCommand: "",
			expectedError:   false,
		},
		{
			name:            "invalid input",
			data:            types.FixtureEntity("entity"),
			input:           make(chan int),
			expectedCommand: "",
			expectedError:   true,
		},
		{
			name:            "invalid template",
			data:            types.FixtureEntity("entity"),
			input:           types.CheckConfig{Command: "{{nil}}"},
			expectedCommand: "",
			expectedError:   true,
		},
		{
			name:            "simple template",
			data:            types.FixtureEntity("entity"),
			input:           types.CheckConfig{Command: "{{ .ID }}"},
			expectedCommand: "entity",
			expectedError:   false,
		},
		{
			name:            "default value for existing field",
			data:            map[string]interface{}{"ID": "foo", "Check": map[string]interface{}{"Name": "check_foo"}},
			input:           types.CheckConfig{Command: `{{ .ID | default "bar" }}`},
			expectedCommand: "foo",
			expectedError:   false,
		},
		{
			name:            "default value for missing field",
			data:            map[string]interface{}{"ID": "foo", "Check": map[string]interface{}{"Name": "check_foo"}},
			input:           types.CheckConfig{Command: `{{ .Check.Foo | default "bar" }}`},
			expectedCommand: "bar",
			expectedError:   false,
		},
		{
			name:            "default int value for missing field",
			data:            map[string]interface{}{"ID": "foo", "Check": map[string]interface{}{"Name": "check_foo"}},
			input:           types.CheckConfig{Command: `{{ .Check.Foo | default 1 }}`},
			expectedCommand: "1",
			expectedError:   false,
		},
		{
			name:          "unmatched token",
			data:          map[string]interface{}{"ID": "foo"},
			input:         types.CheckConfig{Command: `{{ .System.Hostname }}`},
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := tokenSubstitution(tc.data, tc.input)
			testutil.CompareError(err, tc.expectedError, t)

			if !tc.expectedError {
				checkResult := types.CheckConfig{}
				err = json.Unmarshal(result, &checkResult)
				assert.NoError(t, err)

				assert.Equal(t, tc.expectedCommand, checkResult.Command)
			}
		})
	}
}
