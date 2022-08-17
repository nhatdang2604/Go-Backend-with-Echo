package models

import (
	"github.com/beego/beego/orm"
)

func init() {

	//init the model from database for Beego
	orm.RegisterModel(new(User))
}
