package controllers

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"log"
	"qlsv/models"
)

// Operations about GiaoVien
type TeacherController struct {
	beego.Controller
}

// @Title DeleteST
// @Description delete the GiaoVien
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success! models.GiaoVien true
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *TeacherController) Delete() {
	uid := u.GetString(":uid")
	re := models.DeleteGV(uid)
	u.Data["json"] = map[string]string{"result": re}
	u.ServeJSON()
}

// @Title CreateGV
// @Description create ST
// @Param	body		body 	models.GiaoVien	true		"body for teacher content"
// @Success 200 {int} models.GiaoVien.MaGV
// @Failure 403 body is empty
// @router / [post]
func (u *TeacherController) Post() {
	var gv models.GiaoVien
	json.Unmarshal(u.Ctx.Input.RequestBody, &gv)
	uid := models.AddGV(gv)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()

}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.GiaoVien
// @router /list/ [get]
func (u *TeacherController) GetAll() {
	st := models.GetAllGV()
	u.Data["json"] = st
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.GiaoVien
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *TeacherController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		st, err := models.GetGV(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = st
		}
	}
	u.ServeJSON()
}

// @Title FindGV
// @Description get all dsgv
// @Param	key		query  string	true		"The key for staticblock"
// @Success 200 {object} models.GiaoVien
// @router /find_teacher [get]
func (u *TeacherController) FindGV() {
	uid := u.GetString("key")
	if uid != "" {
		log.Println(uid)
		gv := models.FindGV(uid)
		u.Data["json"] = gv
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.GiaoVien	true		"body for user content"
// @Success 200 {object} models.GiaoVien
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *TeacherController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var gv models.GiaoVien
		json.Unmarshal(u.Ctx.Input.RequestBody, &gv)
		uu, err := models.UpdateGV(uid, &gv)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}
