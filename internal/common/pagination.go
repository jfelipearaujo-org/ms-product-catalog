package common

type Pagination struct {
	Page int64 `json:"page"`
	Size int64 `json:"size"`
}

func (p *Pagination) SetDefaults() {
	if p.Page < 0 {
		p.Page = 0
	}

	if p.Size < 10 {
		p.Size = 10
	}

	if p.Size > 100 {
		p.Size = 100
	}
}
