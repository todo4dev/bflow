package healthcheck

type Result struct {
	Uptime   string           `json:"uptime"`
	Status   bool             `json:"status"`
	Services map[string]error `json:"services"`
}
