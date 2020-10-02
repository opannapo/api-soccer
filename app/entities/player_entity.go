package entities

type PlayerEntity struct {
	ID        int64       `gorm:"primary_key"  json:"id"`
	Name      string      `json:"name"`
	Age       int64       `json:"age"`
	JoinDate  int64       `json:"join_date"`
	CreatedAt int64       `json:"created_at"`
	UpdatedAt int64       `json:"updated_at"`
	TeamId    int64       `json:"team_id"`
	Team      *TeamEntity `gorm:"foreignkey:TeamId;association_foreignkey:ID" json:"team,omitempty"`
}

func (*PlayerEntity) TableName() string {
	return "tbl_player"
}
