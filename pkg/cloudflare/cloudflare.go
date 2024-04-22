package cloudflare

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"os"
)

type S3Service struct {
	s3Client *s3.Client
	bucket   string
	endpoint string // Cloudflare R2 endpoint
}

func NewR2Service() (*S3Service, error) {
	account := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	accessKey := os.Getenv("CLOUDFLARE_ACCESS_KEY_ID")
	secretKey := os.Getenv("CLOUDFLARE_SECRET_ACCESS_KEY")
	bucket := os.Getenv("CLOUDFLARE_BUCKET_NAME")

	endpoint := fmt.Sprintf("https://%s.r2.cloudflarestorage.com", account)

	r2Resolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: endpoint,
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolver(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
		config.WithRegion("apac"),
	)
	if err != nil {
		return nil, err
	}

	s3Client := s3.NewFromConfig(cfg)

	return &S3Service{
		s3Client: s3Client,
		bucket:   bucket,
		endpoint: endpoint,
	}, nil
}

func (s *S3Service) UploadVideoFileToR2(ctx context.Context, key string, videoData []byte) error {
	input := &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(videoData),
		ContentType: aws.String("video/mp4"),
	}

	_, err := s.s3Client.PutObject(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (s *S3Service) GetFileURL(key string) string {
	return fmt.Sprintf("%s/%s/%s", s.endpoint, s.bucket, key)
}
