package main

import (
	"github.com/chaitanyabsprip/bonzai"
	"github.com/chaitanyabsprip/bonzai/comp/completers/git"
)

var Cmd = &bonzai.Cmd{
	Name: `test`,
	Comp: git.CompTags,
}
