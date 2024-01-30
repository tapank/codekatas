package protein

import "errors"

var ErrStop = errors.New("error stop")
var ErrInvalidBase = errors.New("error invalid base")

func FromRNA(rna string) ([]string, error) {
	proteins := []string{}
	for i := 0; i <= len(rna)-3; i += 3 {
		if i+3 > len(rna) {
			return nil, ErrInvalidBase
		}
		p, err := FromCodon(rna[i : i+3])
		if err == ErrStop {
			break
		}
		if err != nil {
			return nil, err
		}
		proteins = append(proteins, p)
	}
	return proteins, nil
}

func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	}
	return "", ErrInvalidBase
}
