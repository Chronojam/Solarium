package hardware

type HeatShields struct {
	Quality float64
}

func (h *HeatShields) AlterTemp(v int) int {
	// Dont do anything if we're cooling.
	if v <= 0 {
		return v
	}
	vf := float64(v)
	// q = 1 shields block all positive temp changes
	return int(vf - (vf * h.Quality))
}
