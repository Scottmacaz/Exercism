
codon_dict = {
    "AUG"	: "Methionine",
    "UUU": "Phenylalanine",
    "UUC"	: "Phenylalanine",
    "UUA": "Leucine",
    "UUG":	"Leucine",
    "UCU": "Serine",
    "UCC": "Serine",
    "UCA": "Serine",
    "UCG"	: "Serine",
    "UAU": "Tyrosine",
    "UAC"	: "Tyrosine",
    "UGU": "Cysteine",
    "UGC":	"Cysteine",
    "UGG"	: "Tryptophan",
    "UAA": "STOP",
    "UAG": "STOP",
    "UGA": 	"STOP"
}


def proteins(strand):
    amino_acids = []
    n = 3
    codons = [strand[i:i+n] for i in range(0, len(strand), n)]
    for codon in codons:
        protein = codon_dict[codon]
        if protein == "STOP":
            return amino_acids
        amino_acids.append(protein)
    return amino_acids
