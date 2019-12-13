package UKTaxation

type UkTax struct {
	TaxPercentage int
}
func (t *UkTax) GetTaxPercent() int {
	return t.TaxPercentage
}

//func New(taxPercent int) ukTax {
//	tax := ukTax {taxPercent}
//	return tax
//}