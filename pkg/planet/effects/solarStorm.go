package effects

type SolarStorm struct {
	Identifier  string
	Intensity   int
	Description string
}

func (s *SolarStorm) ID() string {
	return s.Identifier
}

func (s *SolarStorm) Name() string {
	return "SolarStorm"
}

func (s *SolarStorm) Desc() string {
	return s.Description
}

func (s *SolarStorm) AlterRads(v int) int {
	return 50 * s.Intensity
}
func (s *SolarStorm) AlterTemp(v int) int {
	return 100 * s.Intensity
}
