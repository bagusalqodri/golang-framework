package db

type Users struct {
	Id       int64  `xorm:"'id' pk autoincr" json:"id" schema:"id"`
	First    string `xorm:"first" json:"first" schema:"first"`
	Last     string `xorm:"last" json:"last" schema:"last"`
	Email    string `xorm:"email" json:"email" schema:"email"`
	Password string `xorm:"password" json:"password" schema:"password"`
}
