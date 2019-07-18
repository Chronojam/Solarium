package effects

type TemperatureEffect interface {
	AlterTemp(v int) int
}
type RadioactivityEffect interface {
	AlterRads(v int) int
}
