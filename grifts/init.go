package grifts

import (
	"github.com/bestcodyever/githubauth/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
