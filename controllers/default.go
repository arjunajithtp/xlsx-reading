package controllers

import (
	"github.com/astaxie/beego"
	"time"
	"os"
	"io"
	"fmt"
	"github.com/tealeg/xlsx"
	"log"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "home.html"
}

func (c *MainController) Upload()  {

	r := c.Ctx.Request
	r.ParseMultipartForm(0)
	templateName := r.FormValue("documentName")
	//File upload
	file, header, err := r.FormFile("templateFile")
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	fileName := time.Now().Format("20060102150405")+header.Filename
	uploadedFile := "uploadedfile/"+ fileName
	out, err := os.Create(uploadedFile)
	if err!=nil{
		panic(err)
	}
	//document.FileLocation = uploadedFile
	defer out.Close()
	_, err = io.Copy(out, file)
	if err!=nil{
		panic(err)
	}
	//File upload Ends
	fmt.Println("Name: ", templateName, "/nFile Location: ", uploadedFile)

	xlFile, err := xlsx.OpenFile(uploadedFile)
	if err != nil {
		log.Println(err)
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}
}
