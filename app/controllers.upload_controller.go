// ./app/controllers.upload_controller.go
package app

import (
	"context"
	"log"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	minioUpload "github.com/wade-liwei/ossapi/platform/minio"
)

func PrivateDownloadFile(c *fiber.Ctx) error {
	ctx := context.Background()
	//bucketName := os.Getenv("MINIO_BUCKET")
	//MINIO_PRIVATE_DOWNLOAD_BUCKET

	//bucketName := c.Query("bucket")

	var minioClient *minio.Client
	var err error

	//fmt.Println("bucketName", bucketName, len(bucketName))

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

	objectName := file.Filename
	fileBuffer := buffer
	contentType := file.Header["Content-Type"][0]
	fileSize := file.Size

	bucketName := os.Getenv("MINIO_BUCKET")
	// Create minio connection.
	minioClient, err = minioUpload.PrivateMinioConnection()
	if err != nil {
		// Return status 500 and minio connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// if len(bucketName) == 0 {
	// 	bucketName = os.Getenv("MINIO_PRIVATE_DOWNLOAD_BUCKET")
	// 	// Create minio connection.
	// 	minioClient, err = minioUpload.PrivateMinioConnection()
	// 	if err != nil {
	// 		// Return status 500 and minio connection error.
	// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 			"error": true,
	// 			"msg":   err.Error(),
	// 		})
	// 	}
	// } else {
	// 	minioClient, err = minioUpload.MinioConnection(bucketName)
	// 	if err != nil {
	// 		// Return status 500 and minio connection error.
	// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 			"error": true,
	// 			"msg":   err.Error(),
	// 		})
	// 	}
	// }

	expiresAsStr := c.Query("expires", "1800")

	expiresNum, err := strconv.Atoi(expiresAsStr)
	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "query param expires is not valid, err: " + err.Error(),
		})
	}

	if len(objectName) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "must provide  query param 'object' .",
		})
	}

	// Upload the zip file with PutObject
	info, err := minioClient.PutObject(ctx, bucketName, objectName, fileBuffer, fileSize, minio.PutObjectOptions{ContentType: contentType})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set request parameters
	reqParams := make(url.Values)

	reqParams.Set("response-content-disposition", "attachment; filename="+objectName)

	// Gernerate presigned get object url.
	//presignedURL, err := minioClient.PresignedGetObject(context.Background(), bucketName, objectName, time.Duration(1000)*time.Second, reqParams)
	presignedURL, err := minioClient.PresignedGetObject(context.Background(), bucketName, objectName, time.Duration(expiresNum)*time.Second, reqParams)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
		//log.Fatalln(err)
	}
	log.Println(presignedURL)
	log.Println(presignedURL.String())

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)

	return c.SendString(presignedURL.String())
}

// func PublicDownload(c *fiber.Ctx) error {
// 	ctx := context.Background()
// 	bucketName := os.Getenv("MINIO_PUBLIC_DOWNLOAD_BUCKET")
// 	//file, err := c.FormFile("fileUpload")

// 	fmt.Println("bucketName", bucketName, len(bucketName))

// 	file, err := c.FormFile("file")

// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   err.Error(),
// 		})
// 	}

// 	// Get Buffer from file
// 	buffer, err := file.Open()

// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   err.Error(),
// 		})
// 	}
// 	defer buffer.Close()

// 	// Create minio connection.
// 	minioClient, err := minioUpload.PublicMinioConnection()
// 	if err != nil {
// 		// Return status 500 and minio connection error.
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   err.Error(),
// 		})
// 	}

// 	objectName := file.Filename
// 	fileBuffer := buffer
// 	contentType := file.Header["Content-Type"][0]
// 	fileSize := file.Size

// 	// Upload the zip file with PutObject
// 	info, err := minioClient.PutObject(ctx, bucketName, objectName, fileBuffer, fileSize, minio.PutObjectOptions{ContentType: contentType})

// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   err.Error(),
// 		})
// 	}

// 	_ = info

// 	fmt.Println("minioClient.EndpointURL()--------", minioClient.EndpointURL())

// 	return c.SendString(minioClient.EndpointURL().String() + "/" + bucketName + "/" + objectName)
// }

// http://192.168.1.21:13000/api/v1/presignedGetObject?bucket=Peter&object=123&expires=123
// func PresignedGetObject(c *fiber.Ctx) error {

// 	//c.Params()
// 	bucketName := c.Query("bucket")
// 	objectName := c.Query("object")
// 	expiresAsStr := c.Query("expires")

// 	expiresNum, err := strconv.Atoi(expiresAsStr)
// 	if err != nil {

// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   "query param expires is not valid, err: " + err.Error(),
// 		})
// 	}

// 	if len(objectName) == 0 {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   "must provide  query param 'object' .",
// 		})
// 	}

// 	var minioClient *minio.Client

// 	if len(bucketName) == 0 {

// 		bucketName = os.Getenv("MINIO_PRIVATE_DOWNLOAD_BUCKET")

// 		// Create minio connection.
// 		minioClient, err = minioUpload.PrivateMinioConnection()
// 		if err != nil {
// 			// Return status 500 and minio connection error.
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"error": true,
// 				"msg":   err.Error(),
// 			})
// 		}
// 	} else {
// 		minioClient, err = minioUpload.MinioConnection(bucketName)
// 		if err != nil {
// 			// Return status 500 and minio connection error.
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"error": true,
// 				"msg":   err.Error(),
// 			})
// 		}
// 	}

// 	// Set request parameters
// 	reqParams := make(url.Values)

// 	reqParams.Set("response-content-disposition", "attachment; filename="+objectName)

// 	presignedURL, err := minioClient.PresignedGetObject(context.Background(), bucketName, objectName, time.Duration(expiresNum)*time.Second, reqParams)

// 	if err != nil {
// 		//log.Fatalln(err)
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   err.Error(),
// 		})
// 	}
// 	log.Println(presignedURL)
// 	log.Println(presignedURL.String())

// 	//log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)

// 	return c.SendString(presignedURL.String())

// }
