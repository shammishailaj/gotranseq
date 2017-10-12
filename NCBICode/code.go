// Package NCBICode stores codon <-> AA
// translation.
//
// Relevant documentation:
//
//    https://www.ncbi.nlm.nih.gov/Taxonomy/Utils/wprintgc.cgi?chapter=tgencodes#SG1
//
package NCBICode

var (
	// Standard the standard code
	Standard = map[string]byte{
		"TTT": 'F',
		"TCT": 'S',
		"TAT": 'Y',
		"TGT": 'C',
		"TTC": 'F',
		"TCC": 'S',
		"TAC": 'Y',
		"TGC": 'C',
		"TTA": 'L',
		"TCA": 'S',
		"TAA": '*',
		"TGA": '*',
		"TTG": 'L',
		"TCG": 'S',
		"TAG": '*',
		"TGG": 'W',
		"CTT": 'L',
		"CCT": 'P',
		"CAT": 'H',
		"CGT": 'R',
		"CTC": 'L',
		"CCC": 'P',
		"CAC": 'H',
		"CGC": 'R',
		"CTA": 'L',
		"CCA": 'P',
		"CAA": 'Q',
		"CGA": 'R',
		"CTG": 'L',
		"CCG": 'P',
		"CAG": 'Q',
		"CGG": 'R',
		"ATT": 'I',
		"ACT": 'T',
		"AAT": 'N',
		"AGT": 'S',
		"ATC": 'I',
		"ACC": 'T',
		"AAC": 'N',
		"AGC": 'S',
		"ATA": 'I',
		"ACA": 'T',
		"AAA": 'K',
		"AGA": 'R',
		"ATG": 'M',
		"ACG": 'T',
		"AAG": 'K',
		"AGG": 'R',
		"GTT": 'V',
		"GCT": 'A',
		"GAT": 'D',
		"GGT": 'G',
		"GTC": 'V',
		"GCC": 'A',
		"GAC": 'D',
		"GGC": 'G',
		"GTA": 'V',
		"GCA": 'A',
		"GAA": 'E',
		"GGA": 'G',
		"GTG": 'V',
		"GCG": 'A',
		"GAG": 'E',
		"GGG": 'G',
	}
	// ************************************
	// diff from the standard code
	vertebrateMitochondrialeDiff = map[string]byte{
		"AGA": '*',
		"AGG": '*',
		"AUA": 'M',
		"UGA": 'W',
	}
	// TableDiff store the map of diff from standard code
	TableDiff = map[int]map[string]byte{
		2: vertebrateMitochondrialeDiff,
	}
)