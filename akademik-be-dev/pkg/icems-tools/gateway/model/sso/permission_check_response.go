package modelsso

type PermissionDetailResponse struct {
	Color            string             `json:"color"`
	CreatedAt        int64              `json:"created_at"`
	Description      string             `json:"description"`
	GuardName        string             `json:"guard_name"`
	ID               string             `json:"id"`
	Level            int                `json:"level"`
	Name             string             `json:"name"`
	PermissionDetail []PermissionDetail `json:"permission_detail"`
	UpdatedAt        int64              `json:"updated_at"`
}

func (p *PermissionDetailResponse) GetID() string {
	return p.ID
}

type PermissionDetail struct {
	ID         string     `json:"id"`
	Actions    Action     `json:"actions"`
	Permission Permission `json:"permission"`
}

type Action struct {
	ActionName string `json:"action_name"`
	CreatedAt  int64  `json:"created_at"`
	ID         string `json:"id"`
	UpdatedAt  int64  `json:"updated_at"`
}

type Permission struct {
	ApplicationID string `json:"application_id"`
	CreatedAt     int64  `json:"created_at"`
	Group         string `json:"group"`
	GuardName     string `json:"guard_name"`
	ID            string `json:"id"`
	IsReadOnly    bool   `json:"is_read_only"`
	UpdatedAt     int64  `json:"updated_at"`
}
