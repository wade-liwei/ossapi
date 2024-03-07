// ./app/controllers.upload_controller.go
package app

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	minioUpload "github.com/wade-liwei/ossapi/platform/minio"
)

func PrivateDownloadFile(c *fiber.Ctx) error {
	ctx := context.Background()
	bucketName := os.Getenv("MINIO_BUCKET")
	//MINIO_PRIVATE_DOWNLOAD_BUCKET

	//bucketName := os.Getenv("MINIO_PRIVATE_DOWNLOAD_BUCKET")

	//file, err := c.FormFile("fileUpload")

	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get Buffer from file
	buffer, err := file.Open()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	defer buffer.Close()

	// Create minio connection.
	minioClient, err := minioUpload.MinioConnection()
	if err != nil {
		// Return status 500 and minio connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	objectName := file.Filename
	fileBuffer := buffer
	contentType := file.Header["Content-Type"][0]
	fileSize := file.Size

	// policy := `{"Version": "2012-10-17","Statement": [{"Action": ["s3:GetObject"],"Effect": "Allow","Principal": {"AWS": ["*"]},"Resource": ["arn:aws:s3:::my-bucketname/*"],"Sid": ""}]}`

	// err = minioClient.SetBucketPolicy(context.Background(), "my-bucketname", policy)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// Upload the zip file with PutObject
	info, err := minioClient.PutObject(ctx, bucketName, objectName, fileBuffer, fileSize, minio.PutObjectOptions{ContentType: contentType})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	policy, err := minioClient.GetBucketPolicy(context.Background(), bucketName)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("policy-------")
	fmt.Println(policy)
	fmt.Println("policy-------")

	// Set request parameters
	reqParams := make(url.Values)
	//reqParams.Set("response-content-disposition", "attachment; filename=\"your-filename.txt\"")
	//reqParams.Set("response-content-disposition", "attachment; filename=\"your-filename.txt\"")

	//reqParams.Set("response-content-disposition", "attachment; filename=\""+objectName+"\"")

	reqParams.Set("response-content-disposition", "attachment; filename="+objectName)

	// Gernerate presigned get object url.
	//presignedURL, err := minioClient.PresignedGetObject(context.Background(), bucketName, objectName, time.Duration(1000)*time.Second, reqParams)
	presignedURL, err := minioClient.PresignedGetObject(context.Background(), bucketName, objectName, time.Duration(10)*time.Second, reqParams)

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(presignedURL)
	log.Println(presignedURL.String())

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)

	return c.SendString(presignedURL.String())

	// return c.JSON(fiber.Map{
	// 	"error": false,
	// 	"msg":   nil,
	// 	"info":  info,
	// })
}
