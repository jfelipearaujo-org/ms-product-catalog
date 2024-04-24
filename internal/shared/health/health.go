package health

type HealthCheck interface {
	Health() *HealthStatus
}

type HealthStatus struct {
	Status string `json:"status"`
	Err    string `json:"err,omitempty"`
}

func (h *HealthStatus) HasError() bool {
	return h.Err != ""
}
