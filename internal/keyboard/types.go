package keyboard

import (
	"github.com/arvingarciabtw/ditto/internal/keyboard/base"
	"github.com/arvingarciabtw/ditto/internal/keyboard/standards"
)

const (
	Pinky  = base.Pinky
	Ring   = base.Ring
	Middle = base.Middle
	Index  = base.Index
	Thumb  = base.Thumb
	Any    = base.Any
)

type Finger = base.Finger

type Key = base.Key

type Data = standards.Data

var StandardListItems = standards.ListItems
