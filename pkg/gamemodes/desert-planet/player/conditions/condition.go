package conditions

type State int

const (
	Fatal State = iota
	NeedsAttention
	Normal
	Good
)
