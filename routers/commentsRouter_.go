package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["booksmanagementSys/controllers:BooksController"] = append(beego.GlobalControllerRouter["booksmanagementSys/controllers:BooksController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["booksmanagementSys/controllers:BooksController"] = append(beego.GlobalControllerRouter["booksmanagementSys/controllers:BooksController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["booksmanagementSys/controllers:BooksController"] = append(beego.GlobalControllerRouter["booksmanagementSys/controllers:BooksController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:bookid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["booksmanagementSys/controllers:BooksController"] = append(beego.GlobalControllerRouter["booksmanagementSys/controllers:BooksController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:bookid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["booksmanagementSys/controllers:BooksController"] = append(beego.GlobalControllerRouter["booksmanagementSys/controllers:BooksController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:bookid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["booksmanagementSys/controllers:BooksController"] = append(beego.GlobalControllerRouter["booksmanagementSys/controllers:BooksController"],
        beego.ControllerComments{
            Method: "BorrowBook",
            Router: `/borrowbook`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["booksmanagementSys/controllers:BooksController"] = append(beego.GlobalControllerRouter["booksmanagementSys/controllers:BooksController"],
        beego.ControllerComments{
            Method: "ReturnBook",
            Router: `/returnbook`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["booksmanagementSys/controllers:LoginController"] = append(beego.GlobalControllerRouter["booksmanagementSys/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["booksmanagementSys/controllers:LogoutController"] = append(beego.GlobalControllerRouter["booksmanagementSys/controllers:LogoutController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/:id`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["booksmanagementSys/controllers:LogsController"] = append(beego.GlobalControllerRouter["booksmanagementSys/controllers:LogsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["booksmanagementSys/controllers:LogsController"] = append(beego.GlobalControllerRouter["booksmanagementSys/controllers:LogsController"],
        beego.ControllerComments{
            Method: "Addlog",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["booksmanagementSys/controllers:LogsController"] = append(beego.GlobalControllerRouter["booksmanagementSys/controllers:LogsController"],
        beego.ControllerComments{
            Method: "GetById",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["booksmanagementSys/controllers:LogsController"] = append(beego.GlobalControllerRouter["booksmanagementSys/controllers:LogsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["booksmanagementSys/controllers:RegistorController"] = append(beego.GlobalControllerRouter["booksmanagementSys/controllers:RegistorController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["booksmanagementSys/controllers:RegistorController"] = append(beego.GlobalControllerRouter["booksmanagementSys/controllers:RegistorController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["booksmanagementSys/controllers:UserController"] = append(beego.GlobalControllerRouter["booksmanagementSys/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/index`,
            AllowHTTPMethods: []string{"Get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["booksmanagementSys/controllers:UserController"] = append(beego.GlobalControllerRouter["booksmanagementSys/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["booksmanagementSys/controllers:UserController"] = append(beego.GlobalControllerRouter["booksmanagementSys/controllers:UserController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
