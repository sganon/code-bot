package main

import (
	"github.com/urfave/cli"
)

var gitlabAddrFlag = cli.StringFlag{
	EnvVar: "GITLAB_ADDR",
	Name:   "gladdr",
}

var gitlabTokenFlag = cli.StringFlag{
	EnvVar: "GITLAB_API_TOKEN",
	Name:   "gltoken",
}

var releaseSlackHookFlag = cli.StringFlag{
	EnvVar: "RELEASE_SLACK_HOOK",
	Name:   "releasehook",
}
