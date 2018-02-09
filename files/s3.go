package files

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"fmt"
)

func newSession() (*session.Session, error)  {
	ak := ""
	sk := ""
	cred := credentials.NewStaticCredentials(ak, sk, "")
	endpoint := ""
	disableSSL := false
	config := &aws.Config{
		Region:           aws.String(""),
		Endpoint:         &endpoint,
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      cred,
		DisableSSL:       &disableSSL,
	}
	return session.NewSession(config)
}

func doUpload(uploader *UploadFile,config UploadConfig) {
	sessions, err := newSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}
	svc := s3.New(sessions)

	params := &s3.PutObjectInput{
		Bucket			:   aws.String(config.SaveBucket()), // Required
		Key				:   aws.String(config.GetKey()),  // Required
		ACL          	:	aws.String("public-read"),
		Body         	:	uploader,
		ContentLength	:	aws.Int64(uploader.GetSize()),
		ContentType  	:	aws.String(uploader.GetMimeType()),
		Metadata     	:	map[string]*string{
			"Key":aws.String("MetadataValue"),
		},
	}
	_, err = svc.PutObject(params)
}