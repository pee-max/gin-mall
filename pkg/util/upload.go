package util

import (
	"context"
	"fmt"
	"gin_mall/conf"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

func UploadAvatarToLocalStatic(file multipart.File, uId uint, userName string) (string, error) {
	bId := strconv.Itoa(int(uId))
	basePath := "." + conf.AvatarPath + "user" + bId + "/"
	if !DirExists(basePath) {
		CreateDir(basePath)
	}
	avatarPath := basePath + userName + ".jpg"
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return "", err
	}
	return "user" + bId + "/" + userName + ".jpg", nil
}

func DirExists(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func CreateDir(path string) bool {
	err := os.MkdirAll(path, 755)
	if err != nil {
		return false
	}
	return true
}

func UploadToR2(file multipart.File, fileSize int64) (string, error) {
	accessKey := conf.R2Accesskey
	secretKey := conf.R2Secretkey
	bucket := conf.R2Bucket
	accountId := conf.R2AccountId
	//r2Domain := conf.R2Domain

	r2Endpoint := fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		return "", fmt.Errorf("load config failed: %w", err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(r2Endpoint)
	})

	fileName := uuid.New().String() + ".jpg"

	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:        aws.String(bucket),
		Key:           aws.String(fileName),
		Body:          file,
		ContentLength: aws.Int64(fileSize),
		ContentType:   aws.String("image/jpeg"),
	})
	if err != nil {
		return "", fmt.Errorf("upload to r2 failed: %w", err)
	}

	return fileName, nil
}
