package planet

import (
	"log"
	"testing"
)

type testExtremeEffect struct{}

func (t *testExtremeEffect) ID() string {
	return "ExtremeTest"
}
func (t *testExtremeEffect) Name() string {
	return "ExtremeTest"
}
func (t *testExtremeEffect) AlterTemp(v int) int {
	// Add 100oC to temp
	return 100
}

type testHeatShields struct{}

func (t *testHeatShields) ID() string {
	return "foobar"
}
func (t *testHeatShields) Name() string {
	return "foobar"
}
func (t *testHeatShields) HWAlterTemp(v int) int {
	// half any effect
	if v > 0 {
		return v / 2
	}
	return v
}

func TestUpdateTemperature(t *testing.T) {
	t.Run("NoEffectsNoHardware", func(t *testing.T) {
		p := New("foobar")
		// Absolute Zero in C
		p.DistanceFromSun = 275
		p.AtmosphericDensity = 0

		p.UpdateTemperature()
		if p.Temperature != -273 {
			t.Fail()
		}
	})

	t.Run("ExtremeEffectNoHardware", func(t *testing.T) {
		p := New("foobar")
		p.DistanceFromSun = 275
		p.AtmosphericDensity = 0
		p.AddEffect(&testExtremeEffect{})

		p.UpdateTemperature()
		if p.Temperature != -173 {
			t.Fail()
		}
	})

	t.Run("ExtremeEffectHeatShield", func(t *testing.T) {
		p := New("foobar")
		p.DistanceFromSun = 275
		p.AtmosphericDensity = 0
		p.AddEffect(&testExtremeEffect{})
		p.AddEffect(&testHeatShields{})
		p.UpdateTemperature()
		if p.Temperature != -223 {
			t.Fail()
		}
	})
}

type testExtremeEffectRads struct{}

func (t *testExtremeEffectRads) ID() string {
	return "ExtremeTest"
}
func (t *testExtremeEffectRads) Name() string {
	return "ExtremeTest"
}
func (t *testExtremeEffectRads) AlterRads(v int) int {
	// Add 5 rads
	return 5
}

type testRadAmplifier struct{}

func (t *testRadAmplifier) ID() string {
	return "foobar"
}
func (t *testRadAmplifier) Name() string {
	return "foobar"
}
func (t *testRadAmplifier) HWAlterRads(v int) int {
	return v * 2
}

func TestUpdateRadioactivity(t *testing.T) {
	t.Run("NoEffectsNoHardware", func(t *testing.T) {
		p := New("foobar")
		p.DistanceFromSun = 10

		p.UpdateRadioactivity()
		if p.Radioactivity != 5 {
			t.Fail()
		}
	})

	t.Run("ExtremeEffectNoHardware", func(t *testing.T) {
		p := New("foobar")
		p.DistanceFromSun = 10
		p.AddEffect(&testExtremeEffectRads{})

		p.UpdateRadioactivity()
		if p.Radioactivity != 10 {
			t.Fail()
		}
	})

	t.Run("ExtremeEffectRadAmplifier", func(t *testing.T) {
		p := New("foobar")
		p.DistanceFromSun = 10
		p.AddEffect(&testExtremeEffectRads{})
		p.AddEffect(&testRadAmplifier{})
		p.UpdateRadioactivity()
		if p.Radioactivity != 15 {
			log.Printf("%v", p.Radioactivity)
			t.Fail()
		}
	})
}
