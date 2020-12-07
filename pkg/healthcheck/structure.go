package healthcheck

type HarborSVC struct {
	Status     string     `json:"status"`
	Components []Services `json:"components"`
}

type Services struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}
