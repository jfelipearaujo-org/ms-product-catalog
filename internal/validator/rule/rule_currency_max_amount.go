package rule

import (
	"fmt"

	"github.com/leekchan/accounting"
)

var ErrMaxAmount = "amount must be at most %s"

type CurrencyMaxAmount struct {
	MaxAmount  float64
	Accounting accounting.Accounting
}

func NewCurrencyMaxAmount(maxAmount float64, currencyId string) *CurrencyMaxAmount {
	lc := accounting.LocaleInfo[currencyId]

	return &CurrencyMaxAmount{
		MaxAmount: maxAmount,
		Accounting: accounting.Accounting{
			Symbol:    lc.ComSymbol,
			Precision: 2,
			Thousand:  lc.ThouSep,
			Decimal:   lc.DecSep,
		},
	}
}

func (r *CurrencyMaxAmount) Validate(value any) error {
	if _, ok := value.(float64); !ok {
		panic("value must be a float64")
	}

	if value.(float64) > r.MaxAmount {
		return fmt.Errorf(ErrMaxAmount, r.Accounting.FormatMoney(r.MaxAmount))
	}

	return nil
}
