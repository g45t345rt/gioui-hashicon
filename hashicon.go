package gioui_hashicon

import (
	"image"
	"image/color"
	"math"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"golang.org/x/crypto/blake2s"
)

type Hashicon struct {
	Config Config
}

func createHash(hash string, config Config) []byte {
	hasher, _ := blake2s.New128([]byte(config.HashKey)) // 128 bits for 16 bytes long hash
	hasher.Write([]byte(hash))
	return hasher.Sum(nil)
}

func (r Hashicon) Layout(gtx layout.Context, size float32, hash string) layout.Dimensions {
	config := r.Config
	value := createHash(hash, config)
	hue := clampNormalize(float64(value[0]), config.Hue.Min, config.Hue.Max)
	saturation := clampNormalize(float64(value[1]), config.Saturation.Min, config.Saturation.Max)
	saturation /= 100
	lightness := clampNormalize(float64(value[2]), config.Lightness.Min, config.Lightness.Max)
	lightness /= 100
	shift := clampNormalize(float64(value[3]), config.Shift.Min, config.Shift.Max)
	figureAlpha := clampNormalize(float64(value[4]), config.FigureAlpha.Min, config.FigureAlpha.Max)
	figureIndex := int(math.Mod(float64(value[5]), float64(len(Figures))))

	for i := 0; i < len(Sprites); i++ {
		line := Sprites[i]
		shape := Shapes[line.shape]
		var path clip.Path

		path.Begin(gtx.Ops)
		if !line.hidden {
			path.MoveTo(f32.Pt(size*(shape.x1+line.x), size*(shape.y1+line.y)))
			path.LineTo(f32.Pt(size*(shape.x2+line.x), size*(shape.y2+line.y)))
			path.LineTo(f32.Pt(size*(shape.x3+line.x), size*(shape.y3+line.y)))
		}
		path.Close()
		x := math.Round(float64(value[6]) / (float64(i) + 1))
		variation := float64(0)
		if config.Variation.Enabled {
			minMax := config.Variation.MinMax
			variation = clampNormalize(x, minMax.Min, minMax.Max)
		}

		light := float64(0)
		if config.Light.Enabled {
			light = config.Light.Map[line.light] / 100
		}

		hue2 := (hue + variation) / 360
		r, g, b := convertHSLtoRGB(hue2, saturation, lightness+light)
		shapePath := path.End()
		paint.FillShape(gtx.Ops, color.NRGBA{R: r, G: g, B: b, A: 255}, clip.Outline{
			Path: shapePath,
		}.Op())

		figure := Figures[figureIndex]
		if figure[i] > 0 {
			alpha := float64(figure[i]) * figureAlpha / 10
			hue3 := (hue + shift + variation) / 360
			r, g, b := convertHSLtoRGB(hue3, saturation, lightness+light)
			paint.FillShape(gtx.Ops, color.NRGBA{R: r, G: g, B: b, A: uint8(alpha * 255)}, clip.Outline{
				Path: shapePath,
			}.Op())
		}
	}

	return layout.Dimensions{Size: image.Pt(int(size), int(size))}
}
