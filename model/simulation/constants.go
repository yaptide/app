package simulation

// PredefinedParticleRecord ...
type PredefinedParticleRecord struct {
	Value   string `json:"value"`
	Name    string `json:"name"`
	Disable bool   `json:"-"`
}

// ScoringTypeRecord ...
type ScoringTypeRecord struct {
	Value   string `json:"value"`
	Name    string `json:"name"`
	Disable bool   `json:"-"`
}

// PredefinedMaterialRecord contains data needed on frontend related to PredefinedMaterials.
type PredefinedMaterialRecord struct {
	Value   string `json:"value"`
	Name    string `json:"name"`
	Color   Color  `json:"color"`
	Disable bool   `json:"-"`
}

// IsotopeRecord contains data needed on frontend related to Isotopes.
type IsotopeRecord struct {
	Value   string `json:"value"`
	Name    string `json:"name"`
	Disable bool   `json:"-"`
}

// PredefinedParticlesList ...
// TODO write test checking if all ids where used
var PredefinedParticlesList = []PredefinedParticleRecord{
	PredefinedParticleRecord{Value: "all", Name: "All particles"},
	PredefinedParticleRecord{Value: "proton", Name: "Proton"},
	PredefinedParticleRecord{Value: "he_4", Name: "He-4"},
	PredefinedParticleRecord{Value: "heavy_ion", Name: "Heavy ion"},
	PredefinedParticleRecord{Value: "neutron", Name: "Neutron"},
	PredefinedParticleRecord{Value: "pion_pi_minus", Name: "Pion π−"},
	PredefinedParticleRecord{Value: "pion_pi_plus", Name: "Pion π+"},
	PredefinedParticleRecord{Value: "pion_pi_zero", Name: "Pion π0"},
	PredefinedParticleRecord{Value: "anti_neutron", Name: "Anti-neutron"},
	PredefinedParticleRecord{Value: "anti_proton", Name: "Anti-proton"},
	PredefinedParticleRecord{Value: "kaon_minus", Name: "Kaon κ-"},
	PredefinedParticleRecord{Value: "kaon_plus", Name: "Kaon κ+"},
	PredefinedParticleRecord{Value: "kaon_zero", Name: "Kaon κ0"},
	PredefinedParticleRecord{Value: "kaon_anti", Name: "Kaon κ~"},
	PredefinedParticleRecord{Value: "gamma", Name: "Gamma ray"},
	PredefinedParticleRecord{Value: "electron", Name: "Electron"},
	PredefinedParticleRecord{Value: "positron", Name: "Positron"},
	PredefinedParticleRecord{Value: "muon_minus", Name: "Moun µ-"},
	PredefinedParticleRecord{Value: "muon_plus", Name: "Muon µ+"},
	PredefinedParticleRecord{Value: "e_neutrino", Name: "Neutrino e-"},
	PredefinedParticleRecord{Value: "e_anti_neutrino", Name: "Anti-Neutrino e-"},
	PredefinedParticleRecord{Value: "mi_neutrino", Name: "Neutrino µ−"},
	PredefinedParticleRecord{Value: "mi_anti_neutrino", Name: "Anti-Neutrino µ−"},
	PredefinedParticleRecord{Value: "deuteron", Name: "Deuteron"},
	PredefinedParticleRecord{Value: "triton", Name: "Triton"},
	PredefinedParticleRecord{Value: "he_3", Name: "He-3"},
}

// ScoringTypesList ...
var ScoringTypesList = []ScoringTypeRecord{
	ScoringTypeRecord{Value: "dose", Name: "Dose"},
	ScoringTypeRecord{Value: "energy", Name: "Energy"},
	ScoringTypeRecord{Value: "fluence", Name: "Fluence"},
	ScoringTypeRecord{Value: "crossflu", Name: "Cross", Disable: false},
	ScoringTypeRecord{Value: "letflu", Name: "Letflu", Disable: false},
	ScoringTypeRecord{Value: "dlet", Name: "Dlet"},
	ScoringTypeRecord{Value: "tlet", Name: "Tlet"},
	ScoringTypeRecord{Value: "avg_energy", Name: "Avg energy"},
	ScoringTypeRecord{Value: "avg_beta", Name: "Avg beta"},
	ScoringTypeRecord{Value: "ddd", Name: "DDD", Disable: false},
	ScoringTypeRecord{Value: "spc", Name: "SPC"},
	ScoringTypeRecord{Value: "alanine", Name: "Alanine"},
	ScoringTypeRecord{Value: "counter", Name: "Counter"},
}

var waterColor = NewColor(0x00, 0x93, 0xDD, 0xFF)

// PredefinedMaterialsList ...
var PredefinedMaterialsList = []PredefinedMaterialRecord{
	PredefinedMaterialRecord{Value: "hydrogen", Name: "Hydrogen (Z: 1)", Color: White},                                                          // nolint: lll
	PredefinedMaterialRecord{Value: "helium", Name: "Helium (Z: 2)", Color: Gray},                                                               // nolint: lll
	PredefinedMaterialRecord{Value: "lithium", Name: "Lithium (Z: 3)", Color: Gray},                                                             // nolint: lll
	PredefinedMaterialRecord{Value: "beryllium", Name: "Beryllium (Z: 4)", Color: Gray},                                                         // nolint: lll
	PredefinedMaterialRecord{Value: "boron", Name: "Boron (Z: 5)", Color: Gray},                                                                 // nolint: lll
	PredefinedMaterialRecord{Value: "carbon", Name: "Carbon (Z: 6), Amorphous (density 2.0 g/cm3)", Color: Gray},                                // nolint: lll
	PredefinedMaterialRecord{Value: "graphite", Name: "Graphite (Z: 6) (density 1.7 g/cm3)", Color: Gray},                                       // nolint: lll
	PredefinedMaterialRecord{Value: "nitrogen", Name: "Nitrogen (Z: 7)", Color: Gray},                                                           // nolint: lll
	PredefinedMaterialRecord{Value: "oxygen", Name: "Oxygen (Z: 8)", Color: Gray},                                                               // nolint: lll
	PredefinedMaterialRecord{Value: "fluorine", Name: "Fluorine (Z: 9)", Color: Gray},                                                           // nolint: lll
	PredefinedMaterialRecord{Value: "neon", Name: "Neon (Z: 10)", Color: Gray},                                                                  // nolint: lll
	PredefinedMaterialRecord{Value: "sodium", Name: "Sodium (Z: 11)", Color: Gray},                                                              // nolint: lll
	PredefinedMaterialRecord{Value: "magnesium", Name: "Magnesium (Z: 12)", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "aluminum", Name: "Aluminum (Z: 13)", Color: Gray},                                                          // nolint: lll
	PredefinedMaterialRecord{Value: "silicon", Name: "Silicon (Z: 14)", Color: Gray},                                                            // nolint: lll
	PredefinedMaterialRecord{Value: "phosphorus", Name: "Phosphorus (Z: 15)", Color: Gray},                                                      // nolint: lll
	PredefinedMaterialRecord{Value: "sulfur", Name: "Sulfur (Z: 16)", Color: Gray},                                                              // nolint: lll
	PredefinedMaterialRecord{Value: "chlorine", Name: "Chlorine (Z: 17)", Color: Gray},                                                          // nolint: lll
	PredefinedMaterialRecord{Value: "argon", Name: "Argon (Z: 18)", Color: Gray},                                                                // nolint: lll
	PredefinedMaterialRecord{Value: "potassium", Name: "Potassium (Z: 19)", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "calcium", Name: "Calcium (Z: 20)", Color: Gray},                                                            // nolint: lll
	PredefinedMaterialRecord{Value: "scandium", Name: "Scandium (Z: 21)", Color: Gray},                                                          // nolint: lll
	PredefinedMaterialRecord{Value: "titanium", Name: "Titanium (Z: 22)", Color: Gray},                                                          // nolint: lll
	PredefinedMaterialRecord{Value: "vanadium", Name: "Vanadium (Z: 23)", Color: Gray},                                                          // nolint: lll
	PredefinedMaterialRecord{Value: "chromium", Name: "Chromium (Z: 24)", Color: Gray},                                                          // nolint: lll
	PredefinedMaterialRecord{Value: "manganese", Name: "Manganese (Z: 25)", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "iron", Name: "Iron (Z: 26)", Color: Gray},                                                                  // nolint: lll
	PredefinedMaterialRecord{Value: "cobalt", Name: "Cobalt (Z: 27)", Color: Gray},                                                              // nolint: lll
	PredefinedMaterialRecord{Value: "nickel", Name: "Nickel (Z: 28)", Color: Gray},                                                              // nolint: lll
	PredefinedMaterialRecord{Value: "copper", Name: "Copper (Z: 29)", Color: Gray},                                                              // nolint: lll
	PredefinedMaterialRecord{Value: "zinc", Name: "Zinc (Z: 30)", Color: Gray},                                                                  // nolint: lll
	PredefinedMaterialRecord{Value: "gallium", Name: "Gallium (Z: 31)", Color: Gray},                                                            // nolint: lll
	PredefinedMaterialRecord{Value: "germanium", Name: "Germanium (Z: 32)", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "arsenic", Name: "Arsenic (Z: 33)", Color: Gray},                                                            // nolint: lll
	PredefinedMaterialRecord{Value: "selenium", Name: "Selenium (Z: 34)", Color: Gray},                                                          // nolint: lll
	PredefinedMaterialRecord{Value: "bromine", Name: "Bromine (Z: 35)", Color: Gray},                                                            // nolint: lll
	PredefinedMaterialRecord{Value: "krypton", Name: "Krypton (Z: 36)", Color: Gray},                                                            // nolint: lll
	PredefinedMaterialRecord{Value: "rubidium", Name: "Rubidium (Z: 37)", Color: Gray},                                                          // nolint: lll
	PredefinedMaterialRecord{Value: "strontium", Name: "Strontium (Z: 38)", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "yttrium", Name: "Yttrium (Z: 39)", Color: Gray},                                                            // nolint: lll
	PredefinedMaterialRecord{Value: "zirconium", Name: "Zirconium (Z: 40)", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "niobium", Name: "Niobium (Z: 41)", Color: Gray},                                                            // nolint: lll
	PredefinedMaterialRecord{Value: "molybdenum", Name: "Molybdenum (Z: 42)", Color: Gray},                                                      // nolint: lll
	PredefinedMaterialRecord{Value: "technetium", Name: "Technetium (Z: 43)", Color: Gray},                                                      // nolint: lll
	PredefinedMaterialRecord{Value: "ruthenium", Name: "Ruthenium (Z: 44)", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "rhodium", Name: "Rhodium (Z: 45)", Color: Gray},                                                            // nolint: lll
	PredefinedMaterialRecord{Value: "palladium", Name: "Palladium (Z: 46)", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "silver", Name: "Silver (Z: 47)", Color: Gray},                                                              // nolint: lll
	PredefinedMaterialRecord{Value: "cadmium", Name: "Cadmium (Z: 48)", Color: Gray},                                                            // nolint: lll
	PredefinedMaterialRecord{Value: "indium", Name: "Indium (Z: 49)", Color: Gray},                                                              // nolint: lll
	PredefinedMaterialRecord{Value: "tin", Name: "Tin (Z: 50)", Color: Gray},                                                                    // nolint: lll
	PredefinedMaterialRecord{Value: "antimony", Name: "Antimony (Z: 51)", Color: Gray},                                                          // nolint: lll
	PredefinedMaterialRecord{Value: "tellurium", Name: "Tellurium (Z: 52)", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "iodine", Name: "Iodine (Z: 53)", Color: Gray},                                                              // nolint: lll
	PredefinedMaterialRecord{Value: "xenon", Name: "Xenon (Z: 54)", Color: Gray},                                                                // nolint: lll
	PredefinedMaterialRecord{Value: "cesium", Name: "Cesium (Z: 55)", Color: Gray},                                                              // nolint: lll
	PredefinedMaterialRecord{Value: "barium", Name: "Barium (Z: 56)", Color: Gray},                                                              // nolint: lll
	PredefinedMaterialRecord{Value: "lanthanum", Name: "Lanthanum (Z: 57)", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "cerium", Name: "Cerium (Z: 58)", Color: Gray},                                                              // nolint: lll
	PredefinedMaterialRecord{Value: "praseodymium", Name: "Praseodymium (Z: 59)", Color: Gray},                                                  // nolint: lll
	PredefinedMaterialRecord{Value: "neodymium", Name: "Neodymium (Z: 60)", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "promethium", Name: "Promethium (Z: 61)", Color: Gray},                                                      // nolint: lll
	PredefinedMaterialRecord{Value: "samarium", Name: "Samarium (Z: 62)", Color: Gray},                                                          // nolint: lll
	PredefinedMaterialRecord{Value: "europium", Name: "Europium (Z: 63)", Color: Gray},                                                          // nolint: lll
	PredefinedMaterialRecord{Value: "gadolinium", Name: "Gadolinium (Z: 64)", Color: Gray},                                                      // nolint: lll
	PredefinedMaterialRecord{Value: "terbium", Name: "Terbium (Z: 65)", Color: Gray},                                                            // nolint: lll
	PredefinedMaterialRecord{Value: "dysprosium", Name: "Dysprosium (Z: 66)", Color: Gray},                                                      // nolint: lll
	PredefinedMaterialRecord{Value: "holmium", Name: "Holmium (Z: 67)", Color: Gray},                                                            // nolint: lll
	PredefinedMaterialRecord{Value: "erbium", Name: "Erbium (Z: 68)", Color: Gray},                                                              // nolint: lll
	PredefinedMaterialRecord{Value: "thulium", Name: "Thulium (Z: 69)", Color: Gray},                                                            // nolint: lll
	PredefinedMaterialRecord{Value: "ytterbium", Name: "Ytterbium (Z: 70)", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "lutetium", Name: "Lutetium (Z: 71)", Color: Gray},                                                          // nolint: lll
	PredefinedMaterialRecord{Value: "hafnium", Name: "Hafnium (Z: 72)", Color: Gray},                                                            // nolint: lll
	PredefinedMaterialRecord{Value: "tantalum", Name: "Tantalum (Z: 73)", Color: Gray},                                                          // nolint: lll
	PredefinedMaterialRecord{Value: "tungsten", Name: "Tungsten (Z: 74)", Color: Gray},                                                          // nolint: lll
	PredefinedMaterialRecord{Value: "rhenium", Name: "Rhenium (Z: 75)", Color: Gray},                                                            // nolint: lll
	PredefinedMaterialRecord{Value: "osmium", Name: "Osmium (Z: 76)", Color: Gray},                                                              // nolint: lll
	PredefinedMaterialRecord{Value: "iridium", Name: "Iridium (Z: 77)", Color: Gray},                                                            // nolint: lll
	PredefinedMaterialRecord{Value: "platinum", Name: "Platinum (Z: 78)", Color: Gray},                                                          // nolint: lll
	PredefinedMaterialRecord{Value: "gold", Name: "Gold (Z: 79)", Color: Gray},                                                                  // nolint: lll
	PredefinedMaterialRecord{Value: "mercury", Name: "Mercury (Z: 80)", Color: Gray},                                                            // nolint: lll
	PredefinedMaterialRecord{Value: "thallium", Name: "Thallium (Z: 81)", Color: Gray},                                                          // nolint: lll
	PredefinedMaterialRecord{Value: "lead", Name: "Lead (Z: 82)", Color: Gray},                                                                  // nolint: lll
	PredefinedMaterialRecord{Value: "bismuth", Name: "Bismuth (Z: 83)", Color: Gray},                                                            // nolint: lll
	PredefinedMaterialRecord{Value: "polonium", Name: "Polonium (Z: 84)", Color: Gray},                                                          // nolint: lll
	PredefinedMaterialRecord{Value: "astatine", Name: "Astatine (Z: 85)", Color: Gray},                                                          // nolint: lll
	PredefinedMaterialRecord{Value: "radon", Name: "Radon (Z: 86)", Color: Gray},                                                                // nolint: lll
	PredefinedMaterialRecord{Value: "francium", Name: "Francium (Z: 87)", Color: Gray},                                                          // nolint: lll
	PredefinedMaterialRecord{Value: "radium", Name: "Radium (Z: 88)", Color: Gray},                                                              // nolint: lll
	PredefinedMaterialRecord{Value: "actinium", Name: "Actinium (Z: 89)", Color: Gray},                                                          // nolint: lll
	PredefinedMaterialRecord{Value: "thorium", Name: "Thorium (Z: 90)", Color: Gray},                                                            // nolint: lll
	PredefinedMaterialRecord{Value: "protactinium", Name: "Protactinium (Z: 91)", Color: Gray},                                                  // nolint: lll
	PredefinedMaterialRecord{Value: "uranium", Name: "Uranium (Z: 92)", Color: Gray},                                                            // nolint: lll
	PredefinedMaterialRecord{Value: "neptunium", Name: "Neptunium (Z: 93)", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "plutonium", Name: "Plutonium (Z: 94)", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "americium", Name: "Americium (Z: 95)", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "curium", Name: "Curium (Z: 96)", Color: Gray},                                                              // nolint: lll
	PredefinedMaterialRecord{Value: "berkelium", Name: "Berkelium (Z: 97)", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "californium", Name: "Californium (Z: 98)", Color: Gray},                                                    // nolint: lll
	PredefinedMaterialRecord{Value: "a_150", Name: "A-150 Tissue-Equivalent Plastic", Color: Gray},                                              // nolint: lll
	PredefinedMaterialRecord{Value: "acetone", Name: "Acetone", Color: Gray},                                                                    // nolint: lll
	PredefinedMaterialRecord{Value: "acetylene", Name: "Acetylene", Color: Gray},                                                                // nolint: lll
	PredefinedMaterialRecord{Value: "adenine", Name: "Adenine", Color: Gray},                                                                    // nolint: lll
	PredefinedMaterialRecord{Value: "adipose", Name: "Adipose Tissue (ICRP)", Color: Gray},                                                      // nolint: lll
	PredefinedMaterialRecord{Value: "air_dry", Name: "Air, Dry (near sea level)", Color: Gray},                                                  // nolint: lll
	PredefinedMaterialRecord{Value: "alanine", Name: "Alanine", Color: Gray},                                                                    // nolint: lll
	PredefinedMaterialRecord{Value: "aluminum_oxide", Name: "Aluminum Oxide", Color: Gray},                                                      // nolint: lll
	PredefinedMaterialRecord{Value: "amber", Name: "Amber", Color: Gray},                                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "ammonia", Name: "Ammonia", Color: Gray},                                                                    // nolint: lll
	PredefinedMaterialRecord{Value: "aniline", Name: "Aniline", Color: Gray},                                                                    // nolint: lll
	PredefinedMaterialRecord{Value: "anthracene", Name: "Anthracene", Color: Gray},                                                              // nolint: lll
	PredefinedMaterialRecord{Value: "b_100", Name: "B-100 Bone-Equivalent Plastic", Color: Gray},                                                // nolint: lll
	PredefinedMaterialRecord{Value: "bakelite", Name: "Bakelite", Color: Gray},                                                                  // nolint: lll
	PredefinedMaterialRecord{Value: "barium_fluoride", Name: "Barium Fluoride", Color: Gray},                                                    // nolint: lll
	PredefinedMaterialRecord{Value: "barium_sulfate", Name: "Barium Sulfate", Color: Gray},                                                      // nolint: lll
	PredefinedMaterialRecord{Value: "benzene", Name: "Benzene", Color: Gray},                                                                    // nolint: lll
	PredefinedMaterialRecord{Value: "beryllium_oxide", Name: "Beryllium oxide", Color: Gray},                                                    // nolint: lll
	PredefinedMaterialRecord{Value: "bismuth_germanium_oxide", Name: "Bismuth Germanium oxide", Color: Gray},                                    // nolint: lll
	PredefinedMaterialRecord{Value: "blood", Name: "Blood (ICRP)", Color: Gray},                                                                 // nolint: lll
	PredefinedMaterialRecord{Value: "bone_compact", Name: "Bone, Compact (ICRU)", Color: Gray},                                                  // nolint: lll
	PredefinedMaterialRecord{Value: "bone_cortical", Name: "Bone, Cortical (ICRP)", Color: Gray},                                                // nolint: lll
	PredefinedMaterialRecord{Value: "boron_carbide", Name: "Boron Carbide", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "boron_oxide", Name: "Boron Oxide", Color: Gray},                                                            // nolint: lll
	PredefinedMaterialRecord{Value: "brain", Name: "Brain (ICRP)", Color: Gray},                                                                 // nolint: lll
	PredefinedMaterialRecord{Value: "butane", Name: "Butane", Color: Gray},                                                                      // nolint: lll
	PredefinedMaterialRecord{Value: "n_butyl_alcohol", Name: "N-Butyl Alcohol", Color: Gray},                                                    // nolint: lll
	PredefinedMaterialRecord{Value: "c_552", Name: "C-552 Air-Equivalent Plastic", Color: Gray},                                                 // nolint: lll
	PredefinedMaterialRecord{Value: "cadmium_telluride", Name: "Cadmium Telluride", Color: Gray},                                                // nolint: lll
	PredefinedMaterialRecord{Value: "cadmium_tungstate", Name: "Cadmium Tungstate", Color: Gray},                                                // nolint: lll
	PredefinedMaterialRecord{Value: "calcium_carbonate", Name: "Calcium Carbonate", Color: Gray},                                                // nolint: lll
	PredefinedMaterialRecord{Value: "calcium_fluoride", Name: "Calcium Fluoride", Color: Gray},                                                  // nolint: lll
	PredefinedMaterialRecord{Value: "calcium_oxide", Name: "Calcium Oxide", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "calcium_sulfate", Name: "Calcium Sulfate", Color: Gray},                                                    // nolint: lll
	PredefinedMaterialRecord{Value: "calcium_tungstate", Name: "Calcium Tungstate", Color: Gray},                                                // nolint: lll
	PredefinedMaterialRecord{Value: "carbon_dioxide", Name: "Carbon Dioxide", Color: Gray},                                                      // nolint: lll
	PredefinedMaterialRecord{Value: "carbon_tetrachloride", Name: "Carbon Tetrachloride", Color: Gray},                                          // nolint: lll
	PredefinedMaterialRecord{Value: "cellulose_acetate_cellophane", Name: "Cellulose Acetate, Cellophane", Color: Gray},                         // nolint: lll
	PredefinedMaterialRecord{Value: "cellulose_acetate_butyrate", Name: "Cellulose Acetate Butyrate", Color: Gray},                              // nolint: lll
	PredefinedMaterialRecord{Value: "cellulose_nitrate", Name: "Cellulose Nitrate", Color: Gray},                                                // nolint: lll
	PredefinedMaterialRecord{Value: "ceric_sulfate_dosimeter_solution", Name: "Ceric Sulfate Dosimeter Solution", Color: Gray},                  // nolint: lll
	PredefinedMaterialRecord{Value: "cesium_fluoride", Name: "Cesium Fluoride", Color: Gray},                                                    // nolint: lll
	PredefinedMaterialRecord{Value: "cesium_iodide", Name: "Cesium Iodide", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "chlorobenzene", Name: "Chlorobenzene", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "chloroform", Name: "Chloroform", Color: Gray},                                                              // nolint: lll
	PredefinedMaterialRecord{Value: "concrete_portland", Name: "Concrete, Portland", Color: Gray},                                               // nolint: lll
	PredefinedMaterialRecord{Value: "cyclohexane", Name: "Cyclohexane", Color: Gray},                                                            // nolint: lll
	PredefinedMaterialRecord{Value: "1_2_ddihlorobenzene", Name: "1,2-Ddihlorobenzene", Color: Gray},                                            // nolint: lll
	PredefinedMaterialRecord{Value: "dichlorodiethyl_ether", Name: "Dichlorodiethyl Ether", Color: Gray},                                        // nolint: lll
	PredefinedMaterialRecord{Value: "1_2_dichloroethane", Name: "1,2-Dichloroethane", Color: Gray},                                              // nolint: lll
	PredefinedMaterialRecord{Value: "diethyl_ether", Name: "Diethyl Ether", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "n_n_dimethyl_formamide", Name: "N,N-Dimethyl Formamide", Color: Gray},                                      // nolint: lll
	PredefinedMaterialRecord{Value: "dimethyl_sulfoxide", Name: "Dimethyl Sulfoxide", Color: Gray},                                              // nolint: lll
	PredefinedMaterialRecord{Value: "ethane", Name: "Ethane", Color: Gray},                                                                      // nolint: lll
	PredefinedMaterialRecord{Value: "ethyl_alcohol", Name: "Ethyl Alcohol", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "ethyl_cellulose", Name: "Ethyl Cellulose", Color: Gray},                                                    // nolint: lll
	PredefinedMaterialRecord{Value: "ethylene", Name: "Ethylene", Color: Gray},                                                                  // nolint: lll
	PredefinedMaterialRecord{Value: "eye_lens", Name: "Eye Lens (ICRP)", Color: Gray},                                                           // nolint: lll
	PredefinedMaterialRecord{Value: "ferric_oxide", Name: "Ferric Oxide", Color: Gray},                                                          // nolint: lll
	PredefinedMaterialRecord{Value: "ferroboride", Name: "Ferroboride", Color: Gray},                                                            // nolint: lll
	PredefinedMaterialRecord{Value: "ferrous_oxide", Name: "Ferrous Oxide", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "ferrous_sulfate_dosimeter_solution", Name: "Ferrous Sulfate Dosimeter Solution", Color: Gray},              // nolint: lll
	PredefinedMaterialRecord{Value: "freon_12", Name: "Freon-12", Color: Gray},                                                                  // nolint: lll
	PredefinedMaterialRecord{Value: "freon_12b2", Name: "Freon-12B2", Color: Gray},                                                              // nolint: lll
	PredefinedMaterialRecord{Value: "freon_13", Name: "Freon-13", Color: Gray},                                                                  // nolint: lll
	PredefinedMaterialRecord{Value: "freon_13b1", Name: "Freon-13B1", Color: Gray},                                                              // nolint: lll
	PredefinedMaterialRecord{Value: "freon_13i1", Name: "Freon-13I1", Color: Gray},                                                              // nolint: lll
	PredefinedMaterialRecord{Value: "gadolinium_oxysulfide", Name: "Gadolinium Oxysulfide", Color: Gray},                                        // nolint: lll
	PredefinedMaterialRecord{Value: "gallium_arsenide", Name: "Gallium Arsenide", Color: Gray},                                                  // nolint: lll
	PredefinedMaterialRecord{Value: "gel_in_photographic_emulsion", Name: "Gel in Photographic Emulsion", Color: Gray},                          // nolint: lll
	PredefinedMaterialRecord{Value: "glass_lead", Name: "Glass, Lead", Color: Gray},                                                             // nolint: lll
	PredefinedMaterialRecord{Value: "glass_plate", Name: "Glass, Plate", Color: Gray},                                                           // nolint: lll
	PredefinedMaterialRecord{Value: "glass_pyrex", Name: "Glass, Pyrex", Color: Gray},                                                           // nolint: lll
	PredefinedMaterialRecord{Value: "glucose", Name: "Glucose", Color: Gray},                                                                    // nolint: lll
	PredefinedMaterialRecord{Value: "glutamine", Name: "Glutamine", Color: Gray},                                                                // nolint: lll
	PredefinedMaterialRecord{Value: "glycerol", Name: "Glycerol", Color: Gray},                                                                  // nolint: lll
	PredefinedMaterialRecord{Value: "guanine", Name: "Guanine", Color: Gray},                                                                    // nolint: lll
	PredefinedMaterialRecord{Value: "gypsum", Name: "Gypsum, Plaster of Paris", Color: Gray},                                                    // nolint: lll
	PredefinedMaterialRecord{Value: "n_heptane", Name: "N-Heptane", Color: Gray},                                                                // nolint: lll
	PredefinedMaterialRecord{Value: "n_hexane", Name: "N-Hexane", Color: Gray},                                                                  // nolint: lll
	PredefinedMaterialRecord{Value: "kapton_polyimide_film", Name: "Kapton Polyimide Film", Color: Gray},                                        // nolint: lll
	PredefinedMaterialRecord{Value: "lanthanum_oxybromide", Name: "Lanthanum Oxybromide", Color: Gray},                                          // nolint: lll
	PredefinedMaterialRecord{Value: "lanthanum_oxysulfide", Name: "Lanthanum Oxysulfide", Color: Gray},                                          // nolint: lll
	PredefinedMaterialRecord{Value: "lead_oxide", Name: "Lead Oxide", Color: Gray},                                                              // nolint: lll
	PredefinedMaterialRecord{Value: "lithium_amide", Name: "Lithium Amide", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "lithium_carbonate", Name: "Lithium Carbonate", Color: Gray},                                                // nolint: lll
	PredefinedMaterialRecord{Value: "lithium_fluoride", Name: "Lithium Fluoride", Color: Gray},                                                  // nolint: lll
	PredefinedMaterialRecord{Value: "lithium_hydride", Name: "Lithium Hydride", Color: Gray},                                                    // nolint: lll
	PredefinedMaterialRecord{Value: "lithium_iodide", Name: "Lithium Iodide", Color: Gray},                                                      // nolint: lll
	PredefinedMaterialRecord{Value: "lithium_oxide", Name: "Lithium Oxide", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "lithium_tetraborate", Name: "Lithium Tetraborate", Color: Gray},                                            // nolint: lll
	PredefinedMaterialRecord{Value: "lung", Name: "Lung (ICRP)", Color: Gray},                                                                   // nolint: lll
	PredefinedMaterialRecord{Value: "m3_wax", Name: "M3 Wax", Color: Gray},                                                                      // nolint: lll
	PredefinedMaterialRecord{Value: "magnesium_carbonate", Name: "Magnesium Carbonate", Color: Gray},                                            // nolint: lll
	PredefinedMaterialRecord{Value: "magnesium_fluoride", Name: "Magnesium Fluoride", Color: Gray},                                              // nolint: lll
	PredefinedMaterialRecord{Value: "magnesium_oxide", Name: "Magnesium Oxide", Color: Gray},                                                    // nolint: lll
	PredefinedMaterialRecord{Value: "magnesium_tetraborate", Name: "Magnesium Tetraborate", Color: Gray},                                        // nolint: lll
	PredefinedMaterialRecord{Value: "mercuric_iodide", Name: "Mercuric Iodide", Color: Gray},                                                    // nolint: lll
	PredefinedMaterialRecord{Value: "methane", Name: "Methane", Color: Gray},                                                                    // nolint: lll
	PredefinedMaterialRecord{Value: "methanol", Name: "Methanol", Color: Gray},                                                                  // nolint: lll
	PredefinedMaterialRecord{Value: "mix_d_wax", Name: "Mix D Wax", Color: Gray},                                                                // nolint: lll
	PredefinedMaterialRecord{Value: "ms20_tissue_substitute", Name: "MS20 Tissue Substitute", Color: Gray},                                      // nolint: lll
	PredefinedMaterialRecord{Value: "muscle_skeletal", Name: "Muscle, Skeletal", Color: Gray},                                                   // nolint: lll
	PredefinedMaterialRecord{Value: "muscle_striated", Name: "Muscle, Striated", Color: Gray},                                                   // nolint: lll
	PredefinedMaterialRecord{Value: "muscle_equivalent_liquid_with_sucrose", Name: "Muscle-Equivalent Liquid, with Sucrose", Color: Gray},       // nolint: lll
	PredefinedMaterialRecord{Value: "muscle_equivalent_liquid_without_sucrose", Name: "Muscle-Equivalent Liquid, without Sucrose", Color: Gray}, // nolint: lll
	PredefinedMaterialRecord{Value: "naphthalene", Name: "Naphthalene", Color: Gray},                                                            // nolint: lll
	PredefinedMaterialRecord{Value: "nitrobenzene", Name: "Nitrobenzene", Color: Gray},                                                          // nolint: lll
	PredefinedMaterialRecord{Value: "nitrous_oxide", Name: "Nitrous Oxide", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "nylon_du_pont_elvamide_8062", Name: "Nylon, Du Pont ELVAmide 8062", Color: Gray},                           // nolint: lll
	PredefinedMaterialRecord{Value: "nylon_type_6_and_type_6/6", Name: "Nylon, type 6 and type 6/6", Color: Gray},                               // nolint: lll
	PredefinedMaterialRecord{Value: "nylon_type_6/10", Name: "Nylon, type 6/10", Color: Gray},                                                   // nolint: lll
	PredefinedMaterialRecord{Value: "nylon_type_11_rilsan", Name: "Nylon, type 11 (Rilsan)", Color: Gray},                                       // nolint: lll
	PredefinedMaterialRecord{Value: "octane_liquid", Name: "Octane, Liquid", Color: Gray},                                                       // nolint: lll
	PredefinedMaterialRecord{Value: "paraffin_wax", Name: "Paraffin Wax", Color: Gray},                                                          // nolint: lll
	PredefinedMaterialRecord{Value: "n_pentane", Name: "N-Pentane", Color: Gray},                                                                // nolint: lll
	PredefinedMaterialRecord{Value: "photographic_emulsion", Name: "Photographic Emulsion", Color: Gray},                                        // nolint: lll
	PredefinedMaterialRecord{Value: "plastic_scintillator", Name: "Plastic Scintillator (Vinyltoluene based)", Color: Gray},                     // nolint: lll
	PredefinedMaterialRecord{Value: "plutonium_dioxide", Name: "Plutonium Dioxide", Color: Gray},                                                // nolint: lll
	PredefinedMaterialRecord{Value: "polyacrylonitrile", Name: "Polyacrylonitrile", Color: Gray},                                                // nolint: lll
	PredefinedMaterialRecord{Value: "polycarbonate", Name: "Polycarbonate (Makrolon, Lexan)", Color: Gray},                                      // nolint: lll
	PredefinedMaterialRecord{Value: "polychlorostyrene", Name: "Polychlorostyrene", Color: Gray},                                                // nolint: lll
	PredefinedMaterialRecord{Value: "polyethylene", Name: "Polyethylene", Color: Gray},                                                          // nolint: lll
	PredefinedMaterialRecord{Value: "polyethylene_terephthalate", Name: "Polyethylene Terephthalate (Mylar)", Color: Gray},                      // nolint: lll
	PredefinedMaterialRecord{Value: "polymethyl_methacralate", Name: "Polymethyl Methacralate (Lucite, Perspex)", Color: Gray},                  // nolint: lll
	PredefinedMaterialRecord{Value: "polyoxymethylene", Name: "Polyoxymethylene", Color: Gray},                                                  // nolint: lll
	PredefinedMaterialRecord{Value: "polypropylene", Name: "Polypropylene", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "polystyrene", Name: "Polystyrene", Color: Gray},                                                            // nolint: lll
	PredefinedMaterialRecord{Value: "polytetrafluoroethylene", Name: "Polytetrafluoroethylene (Teflon)", Color: Gray},                           // nolint: lll
	PredefinedMaterialRecord{Value: "polytrifluorochloroethylene", Name: "Polytrifluorochloroethylene", Color: Gray},                            // nolint: lll
	PredefinedMaterialRecord{Value: "polyvinyl_acetate", Name: "Polyvinyl Acetate", Color: Gray},                                                // nolint: lll
	PredefinedMaterialRecord{Value: "polyvinyl_alcohol", Name: "Polyvinyl Alcohol", Color: Gray},                                                // nolint: lll
	PredefinedMaterialRecord{Value: "polyvinyl_butyral", Name: "Polyvinyl Butyral", Color: Gray},                                                // nolint: lll
	PredefinedMaterialRecord{Value: "polyvinyl_chloride", Name: "Polyvinyl Chloride", Color: Gray},                                              // nolint: lll
	PredefinedMaterialRecord{Value: "polyvinylidene_chloride_saran", Name: "Polyvinylidene Chloride, Saran", Color: Gray},                       // nolint: lll
	PredefinedMaterialRecord{Value: "polyvinylidene_fluoride", Name: "Polyvinylidene Fluoride", Color: Gray},                                    // nolint: lll
	PredefinedMaterialRecord{Value: "polyvinyl_pyrrolidone", Name: "Polyvinyl Pyrrolidone", Color: Gray},                                        // nolint: lll
	PredefinedMaterialRecord{Value: "potassium_iodide", Name: "Potassium Iodide", Color: Gray},                                                  // nolint: lll
	PredefinedMaterialRecord{Value: "potassium_oxide", Name: "Potassium Oxide", Color: Gray},                                                    // nolint: lll
	PredefinedMaterialRecord{Value: "propane", Name: "Propane", Color: Gray},                                                                    // nolint: lll
	PredefinedMaterialRecord{Value: "propane_liquid", Name: "Propane, Liquid", Color: Gray},                                                     // nolint: lll
	PredefinedMaterialRecord{Value: "n_propyl_alcohol", Name: "N-Propyl Alcohol", Color: Gray},                                                  // nolint: lll
	PredefinedMaterialRecord{Value: "pyridine", Name: "Pyridine", Color: Gray},                                                                  // nolint: lll
	PredefinedMaterialRecord{Value: "rubber_butyl", Name: "Rubber, Butyl", Color: Gray},                                                         // nolint: lll
	PredefinedMaterialRecord{Value: "rubber_natural", Name: "Rubber, Natural", Color: Gray},                                                     // nolint: lll
	PredefinedMaterialRecord{Value: "rubber_neoprene", Name: "Rubber, Neoprene", Color: Gray},                                                   // nolint: lll
	PredefinedMaterialRecord{Value: "silicon_dioxide", Name: "Silicon Dioxide", Color: Gray},                                                    // nolint: lll
	PredefinedMaterialRecord{Value: "silver_bromide", Name: "Silver Bromide", Color: Gray},                                                      // nolint: lll
	PredefinedMaterialRecord{Value: "silver_chloride", Name: "Silver Chloride", Color: Gray},                                                    // nolint: lll
	PredefinedMaterialRecord{Value: "silver_halides_in_photographic_emulsion", Name: "Silver Halides in Photographic Emulsion", Color: Gray},    // nolint: lll
	PredefinedMaterialRecord{Value: "silver_iodide", Name: "Silver Iodide", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "skin", Name: "Skin (ICRP)", Color: Gray},                                                                   // nolint: lll
	PredefinedMaterialRecord{Value: "sodium_carbonate", Name: "Sodium Carbonate", Color: Gray},                                                  // nolint: lll
	PredefinedMaterialRecord{Value: "sodium_iodide", Name: "Sodium Iodide", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "sodium_monoxide", Name: "Sodium Monoxide", Color: Gray},                                                    // nolint: lll
	PredefinedMaterialRecord{Value: "sodium_nitrate", Name: "Sodium Nitrate", Color: Gray},                                                      // nolint: lll
	PredefinedMaterialRecord{Value: "stilbene", Name: "Stilbene", Color: Gray},                                                                  // nolint: lll
	PredefinedMaterialRecord{Value: "sucrose", Name: "Sucrose", Color: Gray},                                                                    // nolint: lll
	PredefinedMaterialRecord{Value: "terphenyl", Name: "Terphenyl", Color: Gray},                                                                // nolint: lll
	PredefinedMaterialRecord{Value: "testes", Name: "Testes (ICRP)", Color: Gray},                                                               // nolint: lll
	PredefinedMaterialRecord{Value: "tetrachloroethylene", Name: "Tetrachloroethylene", Color: Gray},                                            // nolint: lll
	PredefinedMaterialRecord{Value: "thallium_chloride", Name: "Thallium Chloride", Color: Gray},                                                // nolint: lll
	PredefinedMaterialRecord{Value: "tissue_soft_icrp", Name: "Tissue, Soft (ICRP)", Color: Gray},                                               // nolint: lll
	PredefinedMaterialRecord{Value: "tissue_soft_icru_four_component", Name: "Tissue, Soft (ICRU four-component)", Color: Gray},                 // nolint: lll
	PredefinedMaterialRecord{Value: "tissue_equivalent_gas_methane_based", Name: "Tissue-Equivalent GAS (Methane based)", Color: Gray},          // nolint: lll
	PredefinedMaterialRecord{Value: "tissue_equivalent_gas_propane_based", Name: "Tissue-Equivalent GAS (Propane based)", Color: Gray},          // nolint: lll
	PredefinedMaterialRecord{Value: "titanium_dioxide", Name: "Titanium Dioxide", Color: Gray},                                                  // nolint: lll
	PredefinedMaterialRecord{Value: "toluene", Name: "Toluene", Color: Gray},                                                                    // nolint: lll
	PredefinedMaterialRecord{Value: "trichloroethylene", Name: "Trichloroethylene", Color: Gray},                                                // nolint: lll
	PredefinedMaterialRecord{Value: "triethyl_phosphate", Name: "Triethyl Phosphate", Color: Gray},                                              // nolint: lll
	PredefinedMaterialRecord{Value: "tungsten_hexafluoride", Name: "Tungsten Hexafluoride", Color: Gray},                                        // nolint: lll
	PredefinedMaterialRecord{Value: "uranium_dicarbide", Name: "Uranium Dicarbide", Color: Gray},                                                // nolint: lll
	PredefinedMaterialRecord{Value: "uranium_monocarbide", Name: "Uranium Monocarbide", Color: Gray},                                            // nolint: lll
	PredefinedMaterialRecord{Value: "uranium_oxide", Name: "Uranium Oxide", Color: Gray},                                                        // nolint: lll
	PredefinedMaterialRecord{Value: "urea", Name: "Urea", Color: Gray},                                                                          // nolint: lll
	PredefinedMaterialRecord{Value: "valine", Name: "Valine", Color: Gray},                                                                      // nolint: lll
	PredefinedMaterialRecord{Value: "viton_fluoroelastomer", Name: "Viton Fluoroelastomer", Color: Gray},                                        // nolint: lll
	PredefinedMaterialRecord{Value: "water_liquid", Name: "Water, Liquid", Color: waterColor},                                                   // nolint: lll
	PredefinedMaterialRecord{Value: "water_vapor", Name: "Water Vapor", Color: waterColor},                                                      // nolint: lll
	PredefinedMaterialRecord{Value: "xylene", Name: "Xylene", Color: Gray},                                                                      // nolint: lll
	PredefinedMaterialRecord{Value: "vacuum", Name: "Vacuum", Color: Gray},                                                                      // nolint: lll
}

// IsotopesList ...
var IsotopesList = []IsotopeRecord{
	IsotopeRecord{Value: "h-1 - hydrogen", Name: "H-1 - Hydrogen"},
	IsotopeRecord{Value: "h-2 - deuterium", Name: "H-2 - Deuterium"},
	IsotopeRecord{Value: "h-3 - tritium", Name: "H-3 - Tritium"},
	IsotopeRecord{Value: "he-3", Name: "He-3"},
	IsotopeRecord{Value: "he-4", Name: "He-4"},
	IsotopeRecord{Value: "li-6", Name: "Li-6"},
	IsotopeRecord{Value: "li-7", Name: "Li-7"},
	IsotopeRecord{Value: "be-9", Name: "Be-9"},
	IsotopeRecord{Value: "b-10", Name: "B-10"},
	IsotopeRecord{Value: "b-11", Name: "B-11"},
	IsotopeRecord{Value: "c-*", Name: "C-*"},
	IsotopeRecord{Value: "n-*", Name: "N-*"},
	IsotopeRecord{Value: "o-*", Name: "O-*"},
	IsotopeRecord{Value: "f-19", Name: "F-19"},
	IsotopeRecord{Value: "na-23", Name: "Na-23"},
	IsotopeRecord{Value: "mg-*", Name: "Mg-*"},
	IsotopeRecord{Value: "al-27", Name: "Al-27"},
	IsotopeRecord{Value: "si-*", Name: "Si-*"},
	IsotopeRecord{Value: "p-31", Name: "P-31"},
	IsotopeRecord{Value: "s-*", Name: "S-*"},
	IsotopeRecord{Value: "cl-*", Name: "Cl-*"},
	IsotopeRecord{Value: "ar-*", Name: "Ar-*"},
	IsotopeRecord{Value: "k-*", Name: "K-*"},
	IsotopeRecord{Value: "ca-*", Name: "Ca-*"},
	IsotopeRecord{Value: "ti-*", Name: "Ti-*"},
	IsotopeRecord{Value: "v-51", Name: "V-51"},
	IsotopeRecord{Value: "cr-*", Name: "Cr-*"},
	IsotopeRecord{Value: "mn-55", Name: "Mn-55"},
	IsotopeRecord{Value: "fe-*", Name: "Fe-*"},
	IsotopeRecord{Value: "co-59", Name: "Co-59"},
	IsotopeRecord{Value: "ni-*", Name: "Ni-*"},
	IsotopeRecord{Value: "cu-*", Name: "Cu-*"},
	IsotopeRecord{Value: "zn-*", Name: "Zn-*"},
	IsotopeRecord{Value: "ga-*", Name: "Ga-*"},
	IsotopeRecord{Value: "ge-*", Name: "Ge-*"},
	IsotopeRecord{Value: "as-75", Name: "As-75"},
	IsotopeRecord{Value: "nb-93", Name: "Nb-93"},
	IsotopeRecord{Value: "mo-*", Name: "Mo-*"},
	IsotopeRecord{Value: "ag-*", Name: "Ag-*"},
	IsotopeRecord{Value: "cd-*", Name: "Cd-*"},
	IsotopeRecord{Value: "sn-*", Name: "Sn-*"},
	IsotopeRecord{Value: "eu-*", Name: "Eu-*"},
	IsotopeRecord{Value: "gd-*", Name: "Gd-*"},
	IsotopeRecord{Value: "er-*", Name: "Er-*"},
	IsotopeRecord{Value: "ta-181", Name: "Ta-181"},
	IsotopeRecord{Value: "w-*", Name: "W-*"},
	IsotopeRecord{Value: "re-*", Name: "Re-*"},
	IsotopeRecord{Value: "au-187", Name: "Au-187"},
	IsotopeRecord{Value: "hg-*", Name: "Hg-*"},
	IsotopeRecord{Value: "pb-*", Name: "Pb-*"},
	IsotopeRecord{Value: "bi-209", Name: "Bi-209"},
	IsotopeRecord{Value: "th-232", Name: "Th-232"},
	IsotopeRecord{Value: "u-235", Name: "U-235"},
	IsotopeRecord{Value: "u-238", Name: "U-238"},
	IsotopeRecord{Value: "pu-239", Name: "Pu-239"},
	IsotopeRecord{Value: "pu-240", Name: "Pu-240"},
}
