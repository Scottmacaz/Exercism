package strand
var dnaMap = map[byte]byte{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA replaces DNA nucleotids with the complement RNA nucleotides
func ToRNA(dna string) string {
	rna := ""
		for _, di := range dna {
		rna += string(dnaMap[byte(di)])
	}
	return rna
}
