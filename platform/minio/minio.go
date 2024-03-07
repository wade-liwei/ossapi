// ./platform/minio/minio.go

package minioUpload

import (
	"context"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// MinioConnection func for opening minio connection.
func PrivateMinioConnection() (*minio.Client, error) {
	ctx := context.Background()
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESSKEY")
	secretAccessKey := os.Getenv("MINIO_SECRETKEY")
	useSSL := false
	// Initialize minio client object.
	minioClient, errInit := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if errInit != nil {
		log.Fatalln(errInit)
	}

	// Make a new bucket called dev-minio.
	//bucketName := os.Getenv("MINIO_BUCKET")
	//MINIO_PUBLIC_DOWNLOAD_BUCKET="public-down"
	//MINIO_PRIVATE_DOWNLOAD_BUCKET="private-down"

	//publicBucketName := os.Getenv("MINIO_PUBLIC_DOWNLOAD_BUCKET")
	privateBucketName := os.Getenv("MINIO_BUCKET")

	location := "us-east-1"
	//location := "eu-central-1"

	// minioClient.PutObject()
	// minioClient.PutObjectFanOut()
	// minioClient.PutObjectsSnowball()
	// minioClient.FPutObject()

	// if err := minioClient.MakeBucket(ctx, publicBucketName, minio.MakeBucketOptions{Region: location}); err != nil {
	// 	// Check to see if we already own this bucket (which happens if you run this twice)
	// 	exists, errBucketExists := minioClient.BucketExists(ctx, publicBucketName)
	// 	if errBucketExists == nil && exists {
	// 		log.Printf("We already own %s\n", publicBucketName)
	// 	} else {
	// 		log.Fatalln(err)
	// 	}
	// } else {
	// 	log.Printf("Successfully created %s\n", publicBucketName)
	// }

	if err := minioClient.MakeBucket(ctx, privateBucketName, minio.MakeBucketOptions{Region: location}); err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, privateBucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", privateBucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", privateBucketName)
	}

	return minioClient, errInit
}

// func PublicMinioConnection() (*minio.Client, error) {
// 	ctx := context.Background()
// 	endpoint := os.Getenv("MINIO_ENDPOINT")
// 	accessKeyID := os.Getenv("MINIO_ACCESSKEY")
// 	secretAccessKey := os.Getenv("MINIO_SECRETKEY")
// 	useSSL := false
// 	// Initialize minio client object.
// 	minioClient, errInit := minio.New(endpoint, &minio.Options{
// 		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
// 		Secure: useSSL,
// 	})
// 	if errInit != nil {
// 		log.Fatalln(errInit)
// 	}

// 	// Make a new bucket called dev-minio.
// 	//bucketName := os.Getenv("MINIO_BUCKET")
// 	//MINIO_PUBLIC_DOWNLOAD_BUCKET="public-down"
// 	//MINIO_PRIVATE_DOWNLOAD_BUCKET="private-down"

// 	publicBucketName := os.Getenv("MINIO_PUBLIC_DOWNLOAD_BUCKET")
// 	//privateBucketName := os.Getenv("MINIO_PRIVATE_DOWNLOAD_BUCKET")

// 	location := "us-east-1"
// 	//location := "eu-central-1"

// 	// minioClient.PutObject()
// 	// minioClient.PutObjectFanOut()
// 	// minioClient.PutObjectsSnowball()
// 	// minioClient.FPutObject()

// 	if err := minioClient.MakeBucket(ctx, publicBucketName, minio.MakeBucketOptions{Region: location}); err != nil {
// 		// Check to see if we already own this bucket (which happens if you run this twice)
// 		exists, errBucketExists := minioClient.BucketExists(ctx, publicBucketName)
// 		if errBucketExists == nil && exists {
// 			log.Printf("We already own %s  %d\n", publicBucketName, len(publicBucketName))
// 		} else {
// 			log.Fatalln(err)
// 		}
// 	} else {

// 		//policy := `{"Version": "2012-10-17","Statement": [{"Action": ["s3:GetObject"],"Effect": "Allow","Principal": {"AWS": ["*"]},"Resource": ["arn:aws:s3:::my-bucketname/*"],"Sid": ""}]}`

// 		policy := `{"Version": "2012-10-17","Statement": [{"Action": ["s3:GetObject"],"Effect": "Allow","Principal": {"AWS": ["*"]},"Resource": ["arn:aws:s3:::` + publicBucketName + `/*"],"Sid": ""}]}`

// 		err = minioClient.SetBucketPolicy(context.Background(), publicBucketName, policy)
// 		if err != nil {
// 			log.Fatalln(err)
// 		}

// 		log.Printf("Successfully created %s\n", publicBucketName)
// 	}

// 	return minioClient, errInit
// }

// // MinioConnection func for opening minio connection.
// func MinioConnection(bucketName string) (*minio.Client, error) {
// 	ctx := context.Background()
// 	endpoint := os.Getenv("MINIO_ENDPOINT")
// 	accessKeyID := os.Getenv("MINIO_ACCESSKEY")
// 	secretAccessKey := os.Getenv("MINIO_SECRETKEY")
// 	useSSL := false
// 	// Initialize minio client object.
// 	minioClient, errInit := minio.New(endpoint, &minio.Options{
// 		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
// 		Secure: useSSL,
// 	})
// 	if errInit != nil {
// 		log.Fatalln(errInit)
// 	}

// 	// Make a new bucket called dev-minio.
// 	//bucketName := os.Getenv("MINIO_BUCKET")
// 	//MINIO_PUBLIC_DOWNLOAD_BUCKET="public-down"
// 	//MINIO_PRIVATE_DOWNLOAD_BUCKET="private-down"

// 	//publicBucketName := os.Getenv("MINIO_PUBLIC_DOWNLOAD_BUCKET")
// 	//privateBucketName := os.Getenv("MINIO_PRIVATE_DOWNLOAD_BUCKET")
// 	location := "us-east-1"
// 	//location := "eu-central-1"

// 	// minioClient.PutObject()
// 	// minioClient.PutObjectFanOut()
// 	// minioClient.PutObjectsSnowball()
// 	// minioClient.FPutObject()

// 	// if err := minioClient.MakeBucket(ctx, publicBucketName, minio.MakeBucketOptions{Region: location}); err != nil {
// 	// 	// Check to see if we already own this bucket (which happens if you run this twice)
// 	// 	exists, errBucketExists := minioClient.BucketExists(ctx, publicBucketName)
// 	// 	if errBucketExists == nil && exists {
// 	// 		log.Printf("We already own %s\n", publicBucketName)
// 	// 	} else {
// 	// 		log.Fatalln(err)
// 	// 	}
// 	// } else {
// 	// 	log.Printf("Successfully created %s\n", publicBucketName)
// 	// }

// 	if err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location}); err != nil {
// 		// Check to see if we already own this bucket (which happens if you run this twice)
// 		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
// 		if errBucketExists == nil && exists {
// 			log.Printf("We already own %s\n", bucketName)
// 		} else {
// 			log.Fatalln(err)
// 		}
// 	} else {
// 		log.Printf("Successfully created %s\n", bucketName)
// 	}

// 	return minioClient, errInit
// }
