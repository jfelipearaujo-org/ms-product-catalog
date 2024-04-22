package time_provider

import (
	"time"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/provider"
)

type TimeProvider struct {
	funcTime provider.FuncTime
}

func NewTimeProvider(funcTime provider.FuncTime) *TimeProvider {
	return &TimeProvider{
		funcTime: funcTime,
	}
}

func (p *TimeProvider) GetTime() time.Time {
	return p.funcTime()
}
