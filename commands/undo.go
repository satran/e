package commands

import (
	"github.com/satran/e/editor"
)

type Undo struct{}

func (u Undo) Apply(e *editor.Editor) {
	e.ActiveView().Buffer().Undo()
}

type Redo struct{}

func (r Redo) Apply(e *editor.Editor) {
	e.ActiveView().Buffer().Redo()
}
