package strand

func ToRNA(dna string) string {
	var rna string
	for i := 0; i < len(dna); i++ {
		switch dna[i] {
		case 'A':
			rna += "U"
		case 'G':
			rna += "C"
		case 'T':
			rna += "A"
		case 'C':
			rna += "G"
		}
	}
	return rna
}
