package main

import (
	bkey "charm.land/bubbles/v2/key"
)

type bindings struct {
	Layout  bkey.Binding
	Size    bkey.Binding
	HideKey bkey.Binding
}

var commands = bindings{
	Layout: bkey.NewBinding(
		bkey.WithKeys("ctrl+shift+l"),
		bkey.WithHelp("^l", "layout"),
	),
	Size: bkey.NewBinding(
		bkey.WithKeys("ctrl+shift+s"),
		bkey.WithHelp("^s", "size"),
	),
	HideKey: bkey.NewBinding(
		bkey.WithKeys("ctrl+shift+h"),
		bkey.WithHelp("^h", "hide"),
	),
}
