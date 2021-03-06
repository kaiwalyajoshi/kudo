package check

import (
	"testing"

	"github.com/kudobuilder/kudo/pkg/kudoctl/util/vars"
)

func TestKubeConfigPath(t *testing.T) {
	vars.KubeConfigPath = "/tmp/;"

	testNonExisting := []struct {
		expected string
	}{
		{"failed to find kubeconfig file: stat /tmp/;: no such file or directory"}, // 1
	}

	for _, tt := range testNonExisting {
		actual := KubeConfigPath()
		if actual != nil {
			if actual.Error() != tt.expected {
				t.Errorf("non existing test:\nexpected: %v\n     got: %v", tt.expected, actual)
			}
		}
	}

	vars.KubeConfigPath = ""

	testZero := []struct {
		expected *string
	}{
		{nil}, // 1
	}

	for _, tt := range testZero {
		actual := KubeConfigPath()
		if actual != nil {
			t.Errorf("empty path test:\nexpected: %v\n     got: %v", tt.expected, actual)
		}
	}
}

func TestGithubCredentials(t *testing.T) {
	vars.GithubCredentialPath = "/tmp/;"

	testNonExisting := []struct {
		expected string
	}{
		{"failed to find github credential file: stat /tmp/;: no such file or directory"}, // 1
	}

	for _, tt := range testNonExisting {
		actual := GithubCredentials()
		if actual != nil {
			if actual.Error() != tt.expected {
				t.Errorf("non existing test:\nexpected: %v\n     got: %v", tt.expected, actual)
			}
		}
	}

	vars.GithubCredentialPath = ""

	testZero := []struct {
		expected *string
	}{
		{nil}, // 1
	}

	for _, tt := range testZero {
		actual := GithubCredentials()
		if actual != nil {
			t.Errorf("empty path test:\nexpected: %v\n     got: %v", tt.expected, actual)
		}
	}
}
