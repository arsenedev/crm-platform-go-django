package audit

type AuditResponse struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	Action     string `json:"action"`
	Resource   string `json:"resource"`
	ResourceID string `json:"resource_id"`
}
