package entities

type TeamEntity struct {
	ID        int64           `gorm:"primary_key"  json:"id"`
	Name      string          `json:"name"`
	Address   string          `json:"address"`
	CreatedAt int64           `json:"created_at"`
	UpdatedAt int64           `json:"updated_at"`
	Players   []*PlayerEntity `gorm:"foreignkey:TeamId;association_foreignkey:ID" json:"player,omitempty"`
}

func (*TeamEntity) TableName() string {
	return "tbl_team"
}
