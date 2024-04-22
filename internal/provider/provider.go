package provider

import "time"

type FuncTime func() time.Time

type TimeProvider interface {
	GetTime() time.Time
}
