package controllers

import (
	"github.com/astaxie/beego"
	//"encoding/json"
	//"net/http"
	//"fmt"
	//"fmt"
	//"encoding/json"
	//"fmt"
	//"fmt"
	//"fmt"
	"fmt"
	"os"
	"io"
	"log"
	"github.com/minio/minio-go"
	"minio_api_server/models"
	"encoding/json"
	"strconv"
)

type MainController struct {
	beego.Controller
}

type super struct {
	name string `json:"name"`

}

func (c *MainController) Post() {
	// var s superle
	 //err := json.Unmarshal(c.Ctx.Input.RequestBody,&s)

	file, header, er := c.GetFile("file") // where <<this>> is the controller and <<file>> the id of your form field
	fmt.Println("er",er)
	if file != nil {
		// get the filename
		fileName := header.Filename
		fmt.Println(fileName)

		// save to server

		os.Mkdir("/tmp/chandra",0777)
		pathfile:= "/tmp/chandra/" + fileName
		files,err:= os.Create(pathfile)
		fmt.Println("files",files)
		fmt.Println("err",err)


		// save to server
		_,err5:=io.Copy(files,file)

		fmt.Println(err5)

		objectName := file
		contentType := "application/txt"
		endpoint := beego.AppConfig.String("minio_endpoint")
		accessKeyID := beego.AppConfig.String("ak")
		secretAccessKey := beego.AppConfig.String("sk")
		bucketName := beego.AppConfig.String("bucketname")
		useSSL := false
		m := models.Uploads{1,fileName,""}
		b,erb:=json.Marshal(m)
		fmt.Println(b,"b")
		fmt.Println(erb,"erb")
		var v models.Uploads
		json.Unmarshal(b,&v)
		s,ers:=models.AddStatus(&v)
		fmt.Println(s)
		fmt.Println(ers,"ers")
		c.Data["json"] = v



		// Upload the zip file with FPutObject
		minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
		n, err := minioClient.FPutObject(bucketName,fileName,pathfile,contentType)
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("Successfully uploaded %s of size %d\n", objectName, n)

	}


	 //if err!=nil{
		// beego.Info(err)
	// }

	// c.Data["json"] = string(c.Ctx.Input.RequestBody)

         c.ServeJSON()




}

func (c *MainController)Get(){
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetStatesById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()


}