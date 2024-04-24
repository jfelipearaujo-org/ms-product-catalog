package common

type Pagination struct {
	Page int64 `query:"page"`
	Size int64 `query:"size"`
}

func (p *Pagination) SetDefaults() {
	if p.Page < 1 {
		p.Page = 1
	}

	if p.Size < 10 {
		p.Size = 10
	}

	if p.Size > 100 {
		p.Size = 100
	}
}
