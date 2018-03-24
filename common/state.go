package common

// State holds the state of the editor to display the user interface.
type State struct {
	Mode              Mode
	WindowStates      map[int]*WindowState
	Layout            Layout
	Cmdline           []rune
	CmdlineCursor     int
	CompletionResults []string
	CompletionIndex   int
	SearchMode        rune
	Error             error
	ErrorType         int
}

// WindowState holds the state of one window.
type WindowState struct {
	Name          string
	Width         int
	Offset        int64
	Cursor        int64
	Bytes         []byte
	Size          int
	Length        int64
	Mode          Mode
	Pending       bool
	PendingByte   byte
	EditedIndices []int64
	FocusText     bool
}

// Message types
const (
	MessageInfo = iota
	MessageError
)
