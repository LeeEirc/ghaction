package ghaction

import (
	"testing"
)

func TestParseRepoURL(t *testing.T) {
	testDatas := []struct {
		target string
		result []string
	}{
		{
			target: "https://github.com/LeeEirc/jms-guacamole",
			result: []string{"LeeEirc", "jms-guacamole"},
		},
		{
			target: "LeeEirc/jms-guacamole",
			result: []string{"LeeEirc", "jms-guacamole"},
		},
	}
	for i := range testDatas {
		owner, repo, err := ParseRepoURL(testDatas[i].target)
		if err != nil {
			t.Fatalf("parser repo URL %s failed: %s", testDatas[i].target, err)
		}
		if owner != testDatas[i].result[0] || repo != testDatas[i].result[1] {
			t.Fatalf("parser repo URL %s result failed: %s - %s", testDatas[i].target,
				owner, repo)
		}
	}
}
