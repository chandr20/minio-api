package main

import (
	_ "minio_api_server/routers"
	"github.com/astaxie/beego"
	"github.com/minio/minio-go"
	"log"
	"fmt"
	"github.com/astaxie/beego/orm"

        _ "github.com/go-sql-driver/mysql"

)

func init(){
	orm.RegisterDataBase("default", "mysql", "root:devstack@tcp(127.0.0.1:3306)/upload")
	// Database alias.
	name := "default"

	// Drop table and re-create.
	force := false

	// Print log.
	verbose := true

	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	endpoint := beego.AppConfig.String("minio_endpoint")
	accessKeyID := beego.AppConfig.String("ak")
	secretAccessKey := beego.AppConfig.String("sk")
	useSSL := false

	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)

	}
	bucketName := beego.AppConfig.String("bucketname")
	location := beego.AppConfig.String("region")

	exists := minioClient.BucketExists(bucketName)
	fmt.Println(exists)

	if exists!=nil {
		minioClient.MakeBucket(bucketName, location)

	}

}

func main() {
	beego.Run()
}

