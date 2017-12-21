package commands

import (
	"github.com/satran/e/editor"
)

type Repeat struct {
	Command editor.Command
	Count   int
}

func (r Repeat) Apply(e *editor.Editor) {
	for i := 0; i < r.Count; i++ {
		r.Command.Apply(e)
	}
}
