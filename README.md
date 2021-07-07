# ghaction

This repo package will help to create some command line tools, which interactive with GITHUB API.

## ghrcreate

this is a command line tool to create release and upload assets.

usage
```shell
$ ./ghrcreate --help

  --access_token string   Github Access Token
  --branch string         Branch name (default "master")
  --file_pattern string   Assets to upload
  --repo_url string       Repo to sync
  --tag_name string       Tag name

```

Example:
`$ ./ghrcreate --file_pattern '*.tar.gz' --repo_url LeeEirc/jmstool --tag_name v0.0.6 --access_token $GITHUB_TOKEN`
it will create v0.0.6 draft release on LeeEirc/jmstool, and upload all files which match '*.tar.gz' pattern.

