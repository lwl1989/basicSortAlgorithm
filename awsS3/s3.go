package awsS3

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"fmt"
	"basicSortAlgorithm/files"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

var Con *files.Constant
func newSession() (*session.Session, error)  {
	ak := Con.GetConstant("S3_KEY")
	sk := Con.GetConstant("S3_SECRET")
	cred := credentials.NewStaticCredentials(ak, sk, "")
	endpoint := ""
	disableSSL := false
	config := &aws.Config{
		Region:           aws.String(Con.GetConstant("S3_REGION")),
		Endpoint:         &endpoint,
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      cred,
		DisableSSL:       &disableSSL,
	}
	return session.NewSession(config)
}

func DoUpload(uploader *files.UploadFile) (*s3.PutObjectOutput, error) {
	sessions, err := newSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return nil,err
	}
	svc := s3.New(sessions)

	params := &s3.PutObjectInput{
		Bucket			:   aws.String("dev-smart-app"), // Required
		Key				:   aws.String("test"),  // Required
		ACL          	:	aws.String("public-read"),
		Body         	:	uploader,
		ContentLength	:	aws.Int64(uploader.GetSize()),
		ContentType  	:	aws.String(uploader.GetMimeType()),
		Metadata     	:	map[string]*string{
			"Key":aws.String("MetadataValue"),
		},
	}
	return svc.PutObject(params)
}

func GetFileSteam(file string) (*s3.GetObjectOutput, error) {

	sessions, err := newSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return nil,err
	}

	svc := s3.New(sessions)
	params := &s3.GetObjectInput{
		Bucket			:   aws.String("dev-smart-app"), // Required
		Key				:   aws.String(file),  // Required
	}
	obj,err := svc.GetObject(params)
	if err != nil {
		fmt.Println("no file:", err)
		return nil,err
	}

	return obj,err
}