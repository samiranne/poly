package crispr

import (
	"testing"
)

func TestValidateGCContent(t *testing.T) {
	validSequence := "AAAATTTTGGGGCCCC"
	if !validateGCContent(validSequence) {
		t.Errorf("Sequence with between 40%% and 80%% GC content should be valid")
	}

	sequenceTooFewGC := "AAAATTTTAAAATTTT"
	if validateGCContent(sequenceTooFewGC) {
		t.Errorf("Sequence with less than 40%% GC content should not be valid")
	}

	sequenceTooManyGC := "GGGGGGGGCCCCCCCC"
	if validateGCContent(sequenceTooManyGC) {
		t.Errorf("Sequence with more 80%% GC content should not be valid")
	}
}
