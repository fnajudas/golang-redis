package migrate

type Redisz struct {
	Id        int    `gorm:"primaryKey;autoIncrement"`
	Value     string `gorm:"column:value"`
	CreatedOn string `gorm:"column:created_on"`
}
