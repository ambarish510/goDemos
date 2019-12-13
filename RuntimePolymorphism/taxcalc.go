package taxsystem

//import "fmt"
import (
	"fmt"
	ForeignTax "github.com/ambsflip/goDemos/RuntimePolymorphism/ForeignTax"
	"github.com/ambsflip/goDemos/RuntimePolymorphism/ForeignTax/UKTaxation"
	"github.com/ambsflip/goDemos/RuntimePolymorphism/ForeignTax/USTaxation"

	//indianTax "github.com/ambsflip/goDemos/RuntimePolymorphism/IndianTax"
	//singaporeTax "github.com/ambsflip/goDemos/RuntimePolymorphism/SingaporeTax"
)





func GetTaxSystem(country string) (int){


	var ustx *USTaxation.UsTax
	ustx = &USTaxation.UsTax{
		TaxPercentage: 25,
	}

	uktx := &UKTaxation.UkTax{
		TaxPercentage: 15,
	}

	taxSystems := []ForeignTax.TaxSystem{ustx, uktx}
	fmt.Println("US TAX :" , taxSystems[0].GetTaxPercent())
	fmt.Println("UK TAX :" , taxSystems[1].GetTaxPercent())

	var t2 ForeignTax.TaxSystem
	if country == "US" {
		t2 = &USTaxation.UsTax{
			TaxPercentage: 25,
		}
	}else {
		t2 = &UKTaxation.UkTax{
			TaxPercentage: 15,
		}
	}

	return t2.GetTaxPercent()
}
