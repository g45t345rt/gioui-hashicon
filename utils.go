package gioui_hashicon

import "math"

func convertHSLtoRGB(h, s, l float64) (uint8, uint8, uint8) {
	var r, g, b float64

	if s == 0 {
		r = l
		g = l
		b = l
	} else {
		var q, p float64

		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - (l * s)
		}

		p = 2*l - q

		r = hueToRGB(p, q, h+1.0/3.0)
		g = hueToRGB(p, q, h)
		b = hueToRGB(p, q, h-1.0/3.0)
	}

	return uint8(r * 255), uint8(g * 255), uint8(b * 255)
}

func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	if t < 1.0/6.0 {
		return p + (q-p)*6*t
	}
	if t < 1.0/2.0 {
		return q
	}
	if t < 2.0/3.0 {
		return p + (q-p)*(2.0/3.0-t)*6
	}
	return p
}

func clampNormalize(v float64, min float64, max float64) float64 {
	return min + math.Mod(v, max-min)
}
