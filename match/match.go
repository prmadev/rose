// Match package ﳑ find the color code required.
package match

// Ingredient struct ﳑ holds the different values of a color.
// it is not meant to be used directly, but rather as part of a swatch.
type Ingredient struct {
	Name string
	Hex  string
}

// Swatch map ﳑ holds different ingredients(colors) of a swatch.
// each swatch must be filled with [<name of the ingredient>]<ingredient>.
// it is not meant to be used directly, but rather as part of a pallet.
type Swatch []Ingredient

// Palette map ﳑ holds different swatches of a color palette.
// each palette must be filled with [<name of the swatch>]<swatch>.
type Palette map[string]Swatch

// Get function ﳑ returns the specific type of value in a given swatch and palette.
func Get(name, swatchName string) string {
	s := rosePine[swatchName]
	for _, ing := range s {
		if ing.Name == name {
			return ing.Hex
		}
	}
	return ""
}
func List(swatchName string) Swatch {
	s := rosePine[swatchName]
	return s
}
