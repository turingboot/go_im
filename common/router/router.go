package router

import (
	"fmt"
	"go_im/common/global"
	"go_im/controller"
	"html/template"
	"log"
	"net/http"
)

// 10行代码实现万能注册模版
func registerView() {
	tpl, err := template.ParseGlob("./view/**/*")
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, v := range tpl.Templates() {
		tplName := v.Name()
		http.HandleFunc(tplName, func(writer http.ResponseWriter, request *http.Request) {
			err := tpl.ExecuteTemplate(writer, tplName, nil)
			if err != nil {
				return
			}
		})
	}
}

func Router() {
	//engine := gin.Default()

	// 静态资源请求映射
	//engine.Static("/asset", "./asset")
	//engine.StaticFS("/resource", http.Dir("./resource"))

	// 前台
	//engine.GET("/", controller.Index)
	//engine.GET("/index", controller.Index)

	//提供静态资源目录支持
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	http.Handle("/resource/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/user/login", controller.UserLogin)
	http.HandleFunc("/user/register", controller.UserRegister)
	//渲染模板
	registerView()
	// 启动、监听端口
	//post := fmt.Sprintf(":%s", global.Config.Server.Post)
	//if err := engine.Run(post); err != nil {
	//	fmt.Printf("server start error: %s", err)
	//}

	post := fmt.Sprintf(":%s", global.Config.Server.Post)
	if err := http.ListenAndServe(post, nil); err != nil {
		fmt.Printf("server start error: %s", err)
	}

}
