package models

type User struct {
	//Id          int64  `gorm:"primaryKey" json:"id"`
	Iduser int64 `gorm:"primaryKey" json:"iduser"`
	//NamaLengkap string `gorm:"varchar(100)" json:"nama_lengkap"`
	Name_user string `gorm:"varchar(100)" json:"name_user"`
	Username  string `gorm:"varchar(100)" json:"username"`
	Password  string `gorm:"varchar(100)" json:"password"`
}
