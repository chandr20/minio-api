

package models

import "github.com/astaxie/beego/orm"

type Uploads struct {

	Id         int    `orm:"column(id);auto"`
	Uploadfile string `orm:"column(uploadfile);size(25);null"`
	Status     string `orm:"column(status);size(25);null"`


}

func init(){
	orm.RegisterModel(new(Uploads))
}

// last inserted Id on success.
func AddStatus(m *Uploads) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}


func GetStatesById(id int) (v *Uploads, err error) {
	o := orm.NewOrm()
	v = &Uploads{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}
