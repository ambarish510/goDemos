package USTaxation

type UsTax struct {
	TaxPercentage int
}
func (t *UsTax) GetTaxPercent() int {
	return t.TaxPercentage
}
//func New(taxPercent int) usTax {
//	tax := usTax {taxPercent}
//	return tax
//}

