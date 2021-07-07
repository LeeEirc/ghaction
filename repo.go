package ghaction

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v35/github"
)

type Repository struct {
	Owner  string
	Name   string
	Client *github.Client
}

func (r *Repository) GetBranch(ctx context.Context, branch string, ) (*github.Branch, error) {
	gitBranch, _, err := r.Client.Repositories.GetBranch(ctx, r.Owner, r.Name, branch)
	if err != nil {
		return nil, err
	}
	return gitBranch, nil
}
func (r *Repository) CheckoutBranch(ctx context.Context, baseBranch, branch string) error {
	return r.createRefFromBranch(ctx, baseBranch, fmt.Sprintf("refs/heads/%s", branch))
}

func (r *Repository) CreateTagByBranch(ctx context.Context, branch string, tag string) error {
	return r.createRefFromBranch(ctx, branch, fmt.Sprintf("refs/tags/%s", tag))
}

func (r *Repository) createRefFromBranch(ctx context.Context, branch string, ref string) error {
	gitBranch, err := r.GetBranch(ctx, branch)
	if err != nil {
		return err
	}
	gitCommit := gitBranch.GetCommit()
	newRef := github.Reference{
		Ref: github.String(ref),
		URL: github.String(gitCommit.GetURL()),
		Object: &github.GitObject{
			Type: github.String("commit"),
			SHA:  gitCommit.SHA,
			URL:  github.String(gitCommit.GetURL()),
		},
		NodeID: gitCommit.NodeID,
	}
	_, _, err = r.Client.Git.CreateRef(ctx, r.Owner, r.Name, &newRef)
	return err
}

func (r *Repository) CreateRelease(ctx context.Context, tagName string, branch string, ) (*github.RepositoryRelease, error) {
	gitRelease := github.RepositoryRelease{
		TagName:         github.String(tagName),
		TargetCommitish: github.String(branch),
		Name:            github.String(tagName),
		Draft:           github.Bool(true),
	}
	resGitRelease, _, err := r.Client.Repositories.CreateRelease(ctx, r.Owner, r.Name, &gitRelease)
	if err != nil {
		return nil, err
	}
	return resGitRelease, nil
}

func (r *Repository) UploadAssetToRelease(ctx context.Context, gitRelease *github.RepositoryRelease, file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	opt := github.UploadOptions{
		Name:      f.Name(),
	}
	_, _, err = r.Client.Repositories.UploadReleaseAsset(ctx, r.Owner, r.Name, gitRelease.GetID(), &opt, f)
	return err
}
