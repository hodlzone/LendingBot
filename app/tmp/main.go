// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	_ "github.com/DistributedSolutions/LendingBot/app"
	controllers "github.com/DistributedSolutions/LendingBot/app/controllers"
	_ "github.com/DistributedSolutions/LendingBot/app/lender"
	_ "github.com/DistributedSolutions/LendingBot/app/scraper/scraperGRPC"
	tests "github.com/DistributedSolutions/LendingBot/tests"
	controllers0 "github.com/revel/modules/static/app/controllers"
	_ "github.com/revel/modules/testrunner/app"
	controllers1 "github.com/revel/modules/testrunner/app/controllers"
	"github.com/revel/revel"
	"github.com/revel/revel/testing"
	"reflect"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.INFO.Println("Running revel server")

	revel.RegisterController((*controllers.App)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name:           "Sandbox",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{},
				RenderArgNames: map[int][]string{
					41: []string{},
				},
			},
			&revel.MethodType{
				Name:           "Login",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name:           "Register",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
		})

	revel.RegisterController((*controllers.AppAuthRequired)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Dashboard",
				Args: []*revel.MethodArg{},
				RenderArgNames: map[int][]string{
					113: []string{},
				},
			},
			&revel.MethodType{
				Name:           "Logout",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name:           "InfoDashboard",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name:           "InfoAdvancedDashboard",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name:           "SettingsDashboard",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name:           "SysAdminDashboard",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name:           "AdminDashboard",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name:           "AuthUser",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
		})

	revel.RegisterController((*controllers0.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
		})

	revel.RegisterController((*controllers1.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{},
				RenderArgNames: map[int][]string{
					76: []string{
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Suite",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{
					129: []string{},
				},
			},
			&revel.MethodType{
				Name:           "List",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
		})

	revel.DefaultValidationKeys = map[string]map[int]string{}
	testing.TestSuites = []interface{}{
		(*tests.AppTest)(nil),
	}

	revel.Run(*port)
}
