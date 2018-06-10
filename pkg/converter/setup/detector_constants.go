package setup

// PredefinedParticleTypesSet ...
var PredefinedParticleTypesSet = map[string]bool{
	"all":              true,
	"proton":           true,
	"he_4":             true,
	"heavy_ion":        true,
	"neutron":          true,
	"pion_pi_minus":    true,
	"pion_pi_plus":     true,
	"pion_pi_zero":     true,
	"anti_neutron":     true,
	"anti_proton":      true,
	"kaon_minus":       true,
	"kaon_plus":        true,
	"kaon_zero":        true,
	"kaon_anti":        true,
	"gamma":            true,
	"electron":         true,
	"positron":         true,
	"muon_minus":       true,
	"muon_plus":        true,
	"e_neutrino":       true,
	"e_anti_neutrino":  true,
	"mi_neutrino":      true,
	"mi_anti_neutrino": true,
	"deuteron":         true,
	"triton":           true,
	"he_3":             true,
}

// ScoringTypesSet ...
var ScoringTypesSet = map[string]bool{
	"dose":       true,
	"energy":     true,
	"fluence":    true,
	"crossflu":   true,
	"letflu":     true,
	"dlet":       true,
	"tlet":       true,
	"avg_energy": true,
	"avg_beta":   true,
	"ddd":        true,
	"spc":        true,
	"alanine":    true,
	"counter":    true,
}
