package models

type AuditLog struct {
	Id       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId   uint   `json:"user_id"`
	Action   string `json:"action"`
	Entity   string `json:"entity"`
	Details  string `json:"details"`
	EntityId uint   `json:"entity_id"`
}
