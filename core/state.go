package core

// State holds the state of the editor to display the user interface.
type State struct {
	Name   string
	Width  int
	Offset int64
	Cursor int64
	Bytes  []byte
	Size   int
	Length int64
}