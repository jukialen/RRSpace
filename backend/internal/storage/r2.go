package storage

import (
	"context"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/studiobooking/go-api/internal/config"
)

type R2Client struct {
	client *s3.Client
	bucket string
}

func NewR2Client(cfg *config.Config) *R2Client {
	creds := credentials.NewStaticCredentialsProvider(cfg.R2AccessKey, cfg.R2SecretKey, "")

	client := s3.New(s3.Options{
		Region:       "auto",
		Credentials:  creds,
		BaseEndpoint: aws.String(cfg.R2Endpoint),
		UsePathStyle: true,
	})

	return &R2Client{client: client, bucket: cfg.R2Bucket}
}

// UploadFile wysyła plik do R2 i zwraca publiczny URL
func (r *R2Client) UploadFile(ctx context.Context, key string, body io.Reader, contentType string) (string, error) {
	_, err := r.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(r.bucket),
		Key:         aws.String(key),
		Body:        body,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", err
	}

	url := r.client.URLForPresigned(ctx, &s3.GetObjectInput{
		Bucket: aws.String(r.bucket),
		Key:    aws.String(key),
	}, func(o *s3.PresignOptions) {
		o.Expires = 365 * 24 * time.Hour
	})

	return url, nil
}

// PresignedUploadURL — zwraca URL do bezpośredniego uploadu z frontendu
func (r *R2Client) PresignedUploadURL(ctx context.Context, key string) (string, error) {
	uploader := s3.NewUploadManager(r.client)

	request, err := s3.NewPutObjectRequest(ctx, &s3.PutObjectInput{
		Bucket: aws.String(r.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return "", err
	}

	// W praktyce użyj s3.PresignClient do wygenerowania presigned URL
	_ = uploader
	_ = request
	return "", nil // placeholder — uzupełnij wg SDK
}