package water

type Idle struct{}

func New(useFuel bool) *Idle {
	return &Idle{}
}

func (w *Idle) Simulate() {
}

func (w *Idle) Done() bool {
	return true
}
func (w *Idle) Amount() int {
	return 0
}
