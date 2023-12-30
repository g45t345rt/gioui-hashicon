package gioui_hashicon

type MinMaxValue struct {
	Min float64
	Max float64
}

type VariationValue struct {
	MinMax  MinMaxValue
	Enabled bool
}

type LightValue struct {
	Map     map[Light]float64
	Enabled bool
}

type Config struct {
	HashKey     string
	Hue         MinMaxValue
	Saturation  MinMaxValue
	Lightness   MinMaxValue
	Variation   VariationValue
	Shift       MinMaxValue
	FigureAlpha MinMaxValue
	Light       LightValue
}

var DefaultConfig = Config{
	HashKey:    "emerald/hashicon",
	Hue:        MinMaxValue{Min: 0, Max: 360},  // max 360
	Saturation: MinMaxValue{Min: 70, Max: 100}, // max 100
	Lightness:  MinMaxValue{Min: 45, Max: 65},  // max 100
	Variation: VariationValue{
		MinMax: MinMaxValue{
			Min: 5, Max: 20,
		},
		Enabled: true,
	},
	Shift:       MinMaxValue{Min: 60, Max: 300},
	FigureAlpha: MinMaxValue{Min: .7, Max: 1.2},
	Light: LightValue{
		Map: map[Light]float64{
			Top:   10,
			Right: -8,
			Left:  -4,
		},
		Enabled: true,
	},
}
