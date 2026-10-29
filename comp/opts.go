package comp

import (
	"github.com/chaitanyabsprip/bonzai"
	"github.com/chaitanyabsprip/bonzai/fn/filt"
)

type opts struct{}

// Opts is a [bonzai.Completer] for all available [bonzai.Cmd.Opts].
var Opts = new(opts)

func (opts) Complete(args ...string) []string {
	if len(args) == 0 {
		return []string{}
	}
	return filt.HasPrefix(current.OptsSlice(), args[0])
}

func (opts) SetCmd(a *bonzai.Cmd) { current = a }
func (opts) Cmd() *bonzai.Cmd     { return current }
