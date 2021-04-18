package controllers

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"log"
	"qlsv/models"
)

// Operations about SinhVien
type STController struct {
	beego.Controller
}

// @Title DeleteST
// @Description delete the SinhVien
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success! models.SinhVien true
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *STController) DeleteST() {
	uid := u.GetString(":uid")
	re := models.DeleteST(uid)
	u.Data["json"] = map[string]string{"result": re}
	u.ServeJSON()
}

// @Title CreateST
// @Description create ST
// @Param	body		body 	models.SinhVien	true		"body for student content"
// @Success 200 {int} models.SinhVien
// @Failure 403 body is empty
// @router / [post]
func (u *STController) Post() {

	var st models.SinhVien

	json.Unmarshal(u.Ctx.Input.RequestBody, &st)

	uid := models.AddST(st)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()

}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.SinhVien
// @router /list/ [get]
func (u *STController) GetAll() {
	st := models.GetAllST()
	u.Data["json"] = st
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SinhVien
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *STController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		st, err := models.GetST(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = st
		}
	}
	u.ServeJSON()
}

// @Title FindST
// @Description get all dssv
// @Param	key		query  string	true		"The key for staticblock"
// @Success 200 {object} models.SinhVien
// @router /find_st [get]
func (u *STController) FindST() {
	uid := u.GetString("key")
	if uid != "" {
		log.Println(uid)
		st := models.FindStudent(uid)
		u.Data["json"] = st
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.SinhVien	true		"body for user content"
// @Success 200 {object} models.SinhVien
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *STController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var sv models.SinhVien
		json.Unmarshal(u.Ctx.Input.RequestBody, &sv)
		uu, err := models.UpdateST(uid, &sv)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}
