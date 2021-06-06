package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["qlsv/controllers:ClassController"] = append(beego.GlobalControllerRouter["qlsv/controllers:ClassController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["qlsv/controllers:ClassController"] = append(beego.GlobalControllerRouter["qlsv/controllers:ClassController"],
        beego.ControllerComments{
            Method: "DeleteClass",
            Router: "/:uid",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["qlsv/controllers:ClassController"] = append(beego.GlobalControllerRouter["qlsv/controllers:ClassController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:uid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["qlsv/controllers:ClassController"] = append(beego.GlobalControllerRouter["qlsv/controllers:ClassController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:uid",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["qlsv/controllers:ClassController"] = append(beego.GlobalControllerRouter["qlsv/controllers:ClassController"],
        beego.ControllerComments{
            Method: "AddSTtoClass",
            Router: "/add_st_to_class",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["qlsv/controllers:ClassController"] = append(beego.GlobalControllerRouter["qlsv/controllers:ClassController"],
        beego.ControllerComments{
            Method: "AddGVtoClass",
            Router: "/add_teacher_to_class",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["qlsv/controllers:ClassController"] = append(beego.GlobalControllerRouter["qlsv/controllers:ClassController"],
        beego.ControllerComments{
            Method: "DeleteSTInClass",
            Router: "/delete_st_in_class",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["qlsv/controllers:ClassController"] = append(beego.GlobalControllerRouter["qlsv/controllers:ClassController"],
        beego.ControllerComments{
            Method: "DeleteTeacherInClass",
            Router: "/delete_teacher_in_class",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["qlsv/controllers:ClassController"] = append(beego.GlobalControllerRouter["qlsv/controllers:ClassController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/list/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["qlsv/controllers:STController"] = append(beego.GlobalControllerRouter["qlsv/controllers:STController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["qlsv/controllers:STController"] = append(beego.GlobalControllerRouter["qlsv/controllers:STController"],
        beego.ControllerComments{
            Method: "DeleteST",
            Router: "/:uid",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["qlsv/controllers:STController"] = append(beego.GlobalControllerRouter["qlsv/controllers:STController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:uid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["qlsv/controllers:STController"] = append(beego.GlobalControllerRouter["qlsv/controllers:STController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:uid",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["qlsv/controllers:STController"] = append(beego.GlobalControllerRouter["qlsv/controllers:STController"],
        beego.ControllerComments{
            Method: "FindST",
            Router: "/find_st",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["qlsv/controllers:STController"] = append(beego.GlobalControllerRouter["qlsv/controllers:STController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/list/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["qlsv/controllers:TeacherController"] = append(beego.GlobalControllerRouter["qlsv/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["qlsv/controllers:TeacherController"] = append(beego.GlobalControllerRouter["qlsv/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:uid",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["qlsv/controllers:TeacherController"] = append(beego.GlobalControllerRouter["qlsv/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:uid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["qlsv/controllers:TeacherController"] = append(beego.GlobalControllerRouter["qlsv/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:uid",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["qlsv/controllers:TeacherController"] = append(beego.GlobalControllerRouter["qlsv/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "FindGV",
            Router: "/find_teacher",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["qlsv/controllers:TeacherController"] = append(beego.GlobalControllerRouter["qlsv/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/list/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
