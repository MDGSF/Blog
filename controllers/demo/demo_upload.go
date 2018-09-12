package demo

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

// UploadController upload controller
type UploadController struct {
	beego.Controller
}

// Get upload controller get
func (c *UploadController) Get() {
	if c.Ctx.Request.Form == nil {
		c.Ctx.Request.ParseForm()
	}

	beego.Info("UploadController get, c.Data =", c.Data)
	//beego.Info("UploadController get, c.Ctx.Request.Form =", c.Ctx.Request.Form)

	curTime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(curTime, 10))
	token := fmt.Sprintf("%s", h.Sum(nil))

	t, err := template.ParseFiles("views/demo/upload.tpl")
	if err != nil {
		beego.Error("template parse failed, err =", err)
		return
	}

	t.Execute(c.Ctx.ResponseWriter, token)
}

// Post upload controller post
func (c *UploadController) Post() {

	c.Ctx.Request.ParseMultipartForm(32 << 20)

	beego.Info("UploadController post, c.Data =", c.Data)
	beego.Info("UploadController post, c.Ctx.Request.Form =", c.Ctx.Request.Form)

	file, handler, err := c.Ctx.Request.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(c.Ctx.ResponseWriter, "%v", handler.Header)
	f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}
