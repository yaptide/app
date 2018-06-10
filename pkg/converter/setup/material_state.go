package setup

// StateOfMatter represent state of material.
type StateOfMatter int

const (

	// UndefinedStateOfMatter ...
	UndefinedStateOfMatter StateOfMatter = iota

	// Solid state of matter.
	Solid
	// Gas state of matter.
	Gas
	// Liquid state of matter.
	Liquid
)
