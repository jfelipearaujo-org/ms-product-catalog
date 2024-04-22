package rule

import (
	"fmt"

	"github.com/leekchan/accounting"
)

var ErrMinAmount = "amount must be at least %s"

type CurrencyMinAmount struct {
	MinAmount  float64
	Accounting accounting.Accounting
}

func NewCurrencyMinAmount(minAmount float64, currencyId string) *CurrencyMinAmount {
	lc := accounting.LocaleInfo[currencyId]

	return &CurrencyMinAmount{
		MinAmount: minAmount,
		Accounting: accounting.Accounting{
			Symbol:    lc.ComSymbol,
			Precision: 2,
			Thousand:  lc.ThouSep,
			Decimal:   lc.DecSep,
		},
	}
}

func (r *CurrencyMinAmount) Validate(value any) error {
	if _, ok := value.(float64); !ok {
		panic("value must be a float64")
	}

	if value.(float64) < r.MinAmount {
		return fmt.Errorf(ErrMinAmount, r.Accounting.FormatMoney(r.MinAmount))
	}

	return nil
}
