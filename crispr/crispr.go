package crispr

// Help users pick a particular genus of origin based on what PAM
// sequence would be best for this target
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

// Ideal workflow: you have already selected a TargetSequence based on
// how far away it is from

func (t CasTarget) findOffTargetEffects(genome string) string {

}
