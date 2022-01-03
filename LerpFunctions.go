package GoLerpLib

import "math"

func Clamp(val float32, low float32, high float32) float32 {
	val = float32(math.Max(float64(val), float64(low)))
	val = float32(math.Min(float64(val), float64(high)))

	return val
}

func Lerp(a float32, b float32, t float32, clamped bool) float32 {
	if clamped {
		t = Clamp(t, 0, 1)
	}
	return a + ((b - a) * t)
}

func InverseLerp(a float32, b float32, val float32, clamped bool) float32 {
	t := ((val - a) / (b - a))
	if clamped {
		t = Clamp(t, 0, 1)
	}
	return a + ((b - a) * t)
}

func Rescale(a float32, b float32, val float32, c float32, d float32, clamped bool) float32 {
	t := InverseLerp(a, b, val, clamped)
	return Lerp(c, d, t, clamped)
}
