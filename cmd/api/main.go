package main

import (
	"flag"

	"github.com/epavanello/gorsk/pkg/api"

	"github.com/epavanello/gorsk/pkg/utl/config"
	_ "github.com/joho/godotenv/autoload"
)

/*
func main() {
	iso, _ := v8go.NewIsolate()
	global, _ := v8go.NewObjectTemplate(iso)
	printfn, _ := v8go.NewFunctionTemplate(iso, func(info *v8go.FunctionCallbackInfo) *v8go.Value {
		fmt.Printf("%+v\n", info.Args())
		return nil
	})
	global.Set("print", printfn, v8go.ReadOnly)
	ctx, _ := v8go.NewContext(iso, global)
	val, _ := ctx.RunScript("print('foo', 'bar', 0, 1); let obj = {'ciao':2}; obj;", "")
	s, _ := v8go.JSONStringify(nil, val)
	fmt.Println(s)
}
*/

func main() {
	cfgPath := flag.String("p", "./cmd/api/conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	checkErr(err)

	checkErr(api.Start(cfg))
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
