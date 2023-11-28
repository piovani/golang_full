package storage

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/piovani/go_full/infra/config"
)

type Storage struct {
	sess *session.Session
}

func NewStorage() *Storage {
	var sess *session.Session
	var err error

	if config.Env.StageAPP == "dev" {
		sess, err = session.NewSession(&aws.Config{
			Credentials:      credentials.NewStaticCredentials(config.Env.AwsAccessKeyID, config.Env.AwsSecretAccessKey, ""),
			Region:           aws.String(config.Env.AwsRegion),
			Endpoint:         aws.String(fmt.Sprintf("http://localhost:%s", config.Env.AwsPort)),
			S3ForcePathStyle: aws.Bool(len(config.Env.AwsPort) > 0),
		})
	} else {
		sess = session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
			Config: aws.Config{
				Region: aws.String(config.Env.AwsRegion),
			},
		}))
	}

	if err != nil {
		log.Panic(err)
	}

	return &Storage{
		sess: sess,
	}
}

func (s *Storage) Donwload(path string) (io.Reader, error) {
	dwParams := &s3.GetObjectInput{
		Bucket: aws.String(config.Env.AwsBucket),
		Key:    aws.String(path),
	}

	resp, err := s3.New(s.sess).GetObject(dwParams)
	if err != nil {
		return nil, err
	}

	body := resp.Body
	defer body.Close()

	return body, nil
}

func (s *Storage) Upload(file *File) error {
	upParams := &s3manager.UploadInput{
		Bucket: aws.String(config.Env.AwsBucket),
		Key:    aws.String(s.getKey()),
		Body:   file.Reader,
	}

	result, err := s3manager.NewUploader(s.sess).Upload(upParams)
	if err != nil {
		return err
	} else {
		file.Path = result.Location
		return nil
	}
}

func (s *Storage) getKey() string {
	return time.Now().UTC().Format("20060102150405")
}
