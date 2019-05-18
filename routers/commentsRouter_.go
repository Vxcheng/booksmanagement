package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:BooksController"] = append(beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:BooksController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:BooksController"] = append(beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:BooksController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:BooksController"] = append(beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:BooksController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:bookid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:BooksController"] = append(beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:BooksController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:bookid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:BooksController"] = append(beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:BooksController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:bookid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:BooksController"] = append(beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:BooksController"],
        beego.ControllerComments{
            Method: "BorrowBook",
            Router: `/borrowbook`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:BooksController"] = append(beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:BooksController"],
        beego.ControllerComments{
            Method: "ReturnBook",
            Router: `/returnbook`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:LoginController"] = append(beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:LogoutController"] = append(beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:LogoutController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/:id`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:LogsController"] = append(beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:LogsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:LogsController"] = append(beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:LogsController"],
        beego.ControllerComments{
            Method: "Addlog",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:LogsController"] = append(beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:LogsController"],
        beego.ControllerComments{
            Method: "GetById",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:LogsController"] = append(beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:LogsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:RegistorController"] = append(beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:RegistorController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:RegistorController"] = append(beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:RegistorController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:UserController"] = append(beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/index`,
            AllowHTTPMethods: []string{"Get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:UserController"] = append(beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:UserController"] = append(beego.GlobalControllerRouter["ownergit/booksmanagement/controllers:UserController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
