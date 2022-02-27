package crispr

type CasType struct {
	Name          string
	GenusOfOrigin string
	PamSequence   string // three-letter, use N for any
}

type CasTarget struct {
	TargetSequence string
	Locus          string
	Organism       string
	CasProtein     CasType
}

// Given a candidate guide RNA sequence, validate its GC content.
// Returns true if valid, false if not
func validateGCContent(grnaSequence string) bool {
	gcCount := 0
	bases := []rune(grnaSequence)
	for _, base := range bases {
		if base == 'G' || base == 'C' {
			gcCount++
		}
	}

	if gcContent := (float64(gcCount) / float64(len(bases))); gcContent > 0.8 || gcContent < 0.4 {
		return false
	}
	return true
}

// Given a candidate guide RNA sequence, validate its length
// Returns true if valid, false if not
func validateLength(grnaSequence string) bool {
	length := len(grnaSequence)
	return (length >= 17 && length <= 24)
}

func (t CasTarget) validateTarget() bool {
	// 1. Create candidate guide RNA based on t.targetSequence and t.CasProtein.PamSequence
	// 2. Validate GC content
	// 3. Validate length
	// 4. Check off-target effects (query BLAST for the organism?)
	return false
}

// func main() {
// 	sequence := "AAAATTTTCCCCGGGGCCCCGGGG"
// 	gcContentValid := validateGCContent(sequence)
// 	lengthValid := validateLength(sequence)

// 	fmt.Println(gcContentValid)
// 	fmt.Println(lengthValid)
// }

// Ideal workflow: you have already selected a TargetSequence based on
// how far away it is from
