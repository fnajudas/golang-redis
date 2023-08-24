package migrate

import "time"

type Redisz struct {
	Id        int        `gorm:"primaryKey;autoIncrement"`
	Value     string     `gorm:"column:value"`
	CreatedOn *time.Time `gorm:"column:created_on"`
}
