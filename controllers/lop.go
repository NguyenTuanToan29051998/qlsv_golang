package controllers

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"log"
	"qlsv/models"
)

// Operations about SinhVien
type ClassController struct {
	beego.Controller
}

// @Title DeleteClass
// @Description delete the lop
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success! models.Lop true
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *ClassController) DeleteClass() {
	lp := u.GetString(":uid")
	re := models.DeleteClass(lp)
	u.Data["json"] = map[string]string{"result": re}
	u.ServeJSON()
}

// @Title CreateClass
// @Description create Lop
// @Param	body		body 	models.Lop		true		"body for class content"
// @Success 200 {int} models.Lop
// @Failure 403 body is empty
// @router / [post]
func (u *ClassController) Post() {
	var lp models.Lop
	//var sv models.SinhVien
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &lp)
	if err != nil {
		log.Println(err)
	}
	log.Println(u.Ctx.Input.RequestBody)
	log.Println(string(u.Ctx.Input.RequestBody))

	//re:=models.AddSVToClass(lp.MaLop,sv.MaSV)

	re := models.AddClass(lp)
	u.Data["json"] = map[string]string{"result": re}
	u.ServeJSON()

}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.Lop
// @router /list/ [get]
func (u *ClassController) GetAll() {
	lp := models.GetAllClass()
	u.Data["json"] = lp
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Lop
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *ClassController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		st, err := models.GetClass(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = st
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.Lop	true		"body for user content"
// @Success 200 {lop} models.Lop
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *ClassController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var lp models.Lop
		json.Unmarshal(u.Ctx.Input.RequestBody, &lp)
		re := models.UpdateClass(uid, &lp)
		u.Data["json"] = map[string]string{"result": re}
	}
	u.ServeJSON()
}

// @Title AddSTtoClass
// @Description create Lop
// @Param	body		body 	models.CrudST		true		"body for user content"
// @Success 200 {int} models.CrudST
// @Failure 403 body is empty
// @router /add_st_to_class [post]
func (u *ClassController) AddSTtoClass() {
	var lp models.CrudST
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &lp)
	if err != nil {
		log.Println(err)
	}
	re := models.AddSTToClass(lp.MaLop, lp.MaSV)

	u.Data["json"] = map[string]string{"result": re}
	u.ServeJSON()

}

// @Title AddGVtoClass
// @Description create Lop
// @Param	body		body 	models.CrudGV		true		"body for user content"
// @Success 200 {int} models.CrudGV
// @Failure 403 body is empty
// @router /add_teacher_to_class [post]
func (u *ClassController) AddGVtoClass() {
	var lp models.CrudGV
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &lp)
	if err != nil {
		log.Println(err)
	}
	re := models.AddGVToClass(lp.MaLop, lp.MaGV)
	u.Data["json"] = map[string]string{"result": re}
	u.ServeJSON()

}

// @Title DeleteSTInClass
// @Description create Lop
// @Param	body		body 	models.CrudST		true		"body for user content"
// @Success 200 {int} models.CrudST
// @Failure 403 body is empty
// @router /delete_st_in_class [delete]
func (u *ClassController) DeleteSTInClass() {
	var lp models.CrudST
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &lp)
	if err != nil {
		log.Println(err)
	}
	re := models.DeleteSTInClass(lp.MaLop, lp.MaSV)

	u.Data["json"] = map[string]string{"result": re}
	u.ServeJSON()

}

// @Title DeleteTeacherInClass
// @Description create Lop
// @Param	body		body 	models.CrudGV		true		"body for user content"
// @Success 200 {int} models.CrudGV
//@Failure 403 body is empty
// @router /delete_teacher_in_class [delete]
func (u *ClassController) DeleteTeacherInClass() {
	var lp models.CrudGV
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &lp)
	if err != nil {
		log.Println(err)
	}
	re := models.DeleteTeacherInClass(lp.MaLop, lp.MaGV)

	u.Data["json"] = map[string]string{"result": re}
	u.ServeJSON()

}

//// @Title Update
//// @Description update the user
//// @Param	uid		path 	string	true		"The uid you want to update"
//// @Param	body		body 	models.CrudST	true		"body for user content"
//// @Success 200 {object} models.CrudST
//// @Failure 403 :uid is not int
//// @router /UpdateSTInClass:uid [put]
//func (u *ClassController) UpdateSTinClass() {
//	var lp models.Lop
//	//u.Ctx.Input.p
//	err := json.Unmarshal(u.Ctx.Input.RequestBody, &lp)
//	if err != nil{
//		log.Println(err)
//	}
//	re:=models.UpdateIClass(lp.MaLop,&lp)
//
//	u.Data["json"]= map[string]int{"result" : re}
//	u.ServeJSON()
//}
