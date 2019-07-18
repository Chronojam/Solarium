package hardware

type TemperatureEffect interface {
	HWAlterTemp(v int) int
}
type RadioactivityEffect interface {
	HWAlterRads(v int) int
}
