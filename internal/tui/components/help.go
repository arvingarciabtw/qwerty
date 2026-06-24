package components

import (
	bkey "charm.land/bubbles/v2/key"
)

type Bindings struct {
	Layout   bkey.Binding
	Size     bkey.Binding
	Standard bkey.Binding
	HideKey  bkey.Binding
}

var Commands = Bindings{
	Layout: bkey.NewBinding(
		bkey.WithKeys("l"),
		bkey.WithHelp("l", "layout"),
	),
	Size: bkey.NewBinding(
		bkey.WithKeys("s"),
		bkey.WithHelp("s", "size"),
	),
	Standard: bkey.NewBinding(
		bkey.WithKeys("d"),
		bkey.WithHelp("d", "standard"),
	),
	HideKey: bkey.NewBinding(
		bkey.WithKeys("h"),
		bkey.WithHelp("h", "hide"),
	),
}
