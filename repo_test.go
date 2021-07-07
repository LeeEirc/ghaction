package ghaction

import (
	"context"
	"os"
	"testing"

	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"
)

func TestGetRepoTags(t *testing.T) {
	token, ok := os.LookupEnv("GITHUB_TOKEN")
	if !ok {
		t.Fatal("NO GITHUB_TOKEN")
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	repo := Repository{
		Owner:  "LeeEirc",
		Name:   "jmstool",
		Client: client,
	}
	master := "master"
	branch, err := repo.GetBranch(ctx, master)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("GET branch %s commit %+v\n", master, branch.GetCommit().GetSHA())
}

