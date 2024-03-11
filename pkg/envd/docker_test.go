package envd

import (
	"github.com/docker/docker/api/types/container"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestGpusOptAll(t *testing.T) {
	for _, testcase := range []string{
		"all",
		"-1",
		"count=all",
		"count=-1",
	} {
		requests, err := deviceRequests(testcase)
		if err != nil {
			t.Error(err)
		}

		assert.Assert(t, is.Len(requests, 1))
		assert.Check(t, is.DeepEqual(requests[0], container.DeviceRequest{
			Count:  -1,
			Driver: "nvidia",
			Capabilities: [][]string{
				{"gpu"},
				{"nvidia"},
				{"compute"},
				{"compat32"},
				{"graphics"},
				{"utility"},
				{"video"},
				{"display"},
			},
			Options: map[string]string{},
		}))
	}
}

func TestGpusOpts(t *testing.T) {
	for _, testcase := range []string{
		"driver=nvidia,\"capabilities=compute,utility\",\"options=foo=bar,baz=qux\",\"device=0,2\",count=1",
		"1,driver=nvidia,\"capabilities=compute,utility\",\"options=foo=bar,baz=qux\",\"device=0,2\"",
		"count=1,driver=nvidia,\"capabilities=compute,utility\",\"options=foo=bar,baz=qux\",\"device=0,2\"",
		"driver=nvidia,\"capabilities=compute,utility\",\"options=foo=bar,baz=qux\",count=1,\"device=0,2\"",
	} {

		requests, err := deviceRequests(testcase)
		if err != nil {
			t.Error(err)
		}
		assert.Assert(t, is.Len(requests, 1))
		assert.Check(t, is.DeepEqual(requests[0], container.DeviceRequest{
			Driver:       "nvidia",
			Count:        1,
			DeviceIDs:    []string{"0", "2"},
			Capabilities: [][]string{{"compute", "utility", "gpu"}},
			Options:      map[string]string{"foo": "bar", "baz": "qux"},
		}))
	}
}
