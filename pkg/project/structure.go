package project

import (
	"time"
)

type HarborProject []struct {
	CreationTime       time.Time              `json:"creation_time"`
	CurrentUserRoleID  int                    `json:"current_user_role_id"`
	CurrentUserRoleIds []int                  `json:"current_user_role_ids"`
	CveAllowlist       CveState               `json:"cve_allowlist"`
	Name               string                 `json:"name"`
	OwnerID            int                    `json:"owner_id"`
	OwnerName          string                 `json:"owner_name"`
	ProjectID          int                    `json:"project_id"`
	RepoCount          int                    `json:"repo_count"`
	UpdateTime         time.Time              `json:"update_time"`
	Metadata           map[string]interface{} `json:"metadata,omitempty"`
}

type CveState struct {
	CreationTime time.Time     `json:"creation_time"`
	ID           int           `json:"id"`
	Items        []interface{} `json:"items"`
	ProjectID    int           `json:"project_id"`
	UpdateTime   time.Time     `json:"update_time"`
}
