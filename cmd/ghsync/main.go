package main

import (
	"context"
	"log"

	"github.com/google/go-github/v35/github"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"

	"ghaction"
)

type Config struct {
	RepoURL        string `mapstructure:"REPO_URL" yaml:"REPO_URL"`
	FilePattern    string `mapstructure:"FILE_PATTERN" yaml:"FILE_PATTERN"`
	GitAccessToken string `mapstructure:"ACCESS_TOKEN" yaml:"ACCESS_TOKEN"`
	TagName        string `mapstructure:"TAG_NAME" yaml:"tag_name"`
	Branch         string `mapstructure:"BRANCH" yaml:"branch"`
}

var globalConfig Config

func init() {
	pflag.String("file_pattern", "", "Assets to upload")
	pflag.String("repo_url", "", "Repo to sync")
	pflag.String("access_token", "", "Github Access Token")
	pflag.String("tag_name", "", "Tag name")
	pflag.String("branch", "master", "Branch name")
}

func main() {
	globalConfig = LoadConfig()

	owner, repo, err := ghaction.ParseRepoURL(globalConfig.RepoURL)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Create REPO %s  Tag %s from Branch %s\n", globalConfig.RepoURL,
		globalConfig.TagName, globalConfig.Branch)
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: globalConfig.GitAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	gitClient := github.NewClient(tc)
	repoService := ghaction.Repository{
		Owner:  owner,
		Name:   repo,
		Client: gitClient,
	}
	gitRelease, err := repoService.CreateRelease(ctx, globalConfig.TagName, globalConfig.Branch)
	if err != nil {
		log.Fatal(err)
	}
	if globalConfig.FilePattern == "" {
		return
	}
	files, err := ghaction.GetMatchedFiles(globalConfig.FilePattern)
	if err != nil {
		log.Fatal(err)
	}
	for i := range files {
		err = repoService.UploadAssetToRelease(ctx, gitRelease, files[i])
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("upload asset file %s\n", files[i])
	}

}

func LoadEnvConfig() Config {
	return Config{
		RepoURL:        viper.GetString("REPO_URL"),
		FilePattern:    viper.GetString("FILE_PATTERN"),
		GitAccessToken: viper.GetString("ACCESS_TOKEN"),
		TagName:        viper.GetString("TAG_NAME"),
		Branch:         viper.GetString("BRANCH"),
	}
}

func LoadConfig() Config {
	pflag.Parse()
	viper.AutomaticEnv()
	conf := LoadEnvConfig()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Fatal(err)
	}
	return conf
}
