package vars

type Variables struct {
	Population     int
	GeneLength     int
	NumGenerations int
}

func GetVariables() Variables {
	return Variables{
		Population:     100,
		GeneLength:     10,
		NumGenerations: 100,
	}
}