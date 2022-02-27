package ping

import (
	"context"
	"testing"
)

func TestPing(t *testing.T) {
	tests := []struct {
		name   string
		config Config
		err    error
	}{
		{
			name: "test-success",
			config: Config{
				Version:     "some-version",
				Environment: "some-environment",
				AWSRegion:   "some-region",
			},
			err: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resp, err := test.config.Ping(context.Background())
			if resp.Version != test.config.Version {
				t.Errorf("version mismatch. expected: %q; got: %q", test.config.Version, resp.Version)
			}

			if resp.Version != test.config.Version {
				t.Errorf("environment mismatch. expected: %q; got: %q", test.config.Environment, resp.Environment)
			}

			if resp.Version != test.config.Version {
				t.Errorf("awsregion mismatch. expected: %q; got: %q", test.config.AWSRegion, resp.AWSRegion)
			}

			if err != test.err {
				t.Errorf("err mismatch. expected: %v; got: %v", test.err, err)
			}
		})
	}
}
