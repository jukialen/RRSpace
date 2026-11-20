package storage

import (
	"context"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"rrspace/backend/internal/config"
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

// UploadFile wysyła plik do R2 i zwraca presigned URL
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
	presignClient := s3.NewPresignClient(r.client)
	req, err := presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(r.bucket),
		Key:    aws.String(key),
	}, func(o *s3.PresignOptions) {
		o.Expires = 365 * 24 * time.Hour
	})
	if err != nil {
		return "", err
	}
	return req.URL, nil
}

// PresignedUploadURL — zwraca URL do bezpośredniego uploadu z frontendu
// PresignedUploadURL — zwraca URL do bezpośredniego uploadu z frontendu
func (r *R2Client) PresignedUploadURL(ctx context.Context, key string) (string, error) {
	presignClient := s3.NewPresignClient(r.client)
	req, err := presignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(r.bucket),
		Key:    aws.String(key),
	}, func(o *s3.PresignOptions) {
		o.Expires = 15 * time.Minute
	})
	if err != nil {
		return "", err
	}
	return req.URL, nil
}