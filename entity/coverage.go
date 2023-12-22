package entity

type Coverage struct {
	Id    int64  `gorm:"primaryKey;column:id;autoIncrement;type:int"`
	Kdcab string `gorm:"index;column:kdcab;type:varchar(4)"`
	Nik   string `gorm:"column:nik"`
}

func (Coverage) TableName() string {
	return "m_coverage"
}
