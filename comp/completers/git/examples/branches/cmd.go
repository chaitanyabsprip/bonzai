package main

import (
	"github.com/chaitanyabsprip/bonzai"
	"github.com/chaitanyabsprip/bonzai/comp/completers/git"
)

var Cmd = &bonzai.Cmd{Name: `branches`, Comp: git.CompBranches}
