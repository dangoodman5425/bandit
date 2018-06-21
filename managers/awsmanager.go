package managers

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
	"github.com/aws/aws-sdk-go/service/sqs"
	"fmt"
	"io"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"bytes"
	"net/http"
	"golang.org/x/net/context"
	"time"
)


var bucket = aws.String("bandit-test")

func loadAwsConfig() *session.Session	{
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String("us-east-1")},
		Profile: "bandit",
	}))
	return sess
}


// SQS Functions
func AddToSQSQueue(qUrl string, content []byte)	{
	sess := loadAwsConfig()
	svc := sqs.New(sess)
	result, err := svc.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(content)),
		QueueUrl: aws.String(qUrl),
	})
	if err != nil{
		log.Panic(err)
	}
	fmt.Println("Success", *result.MessageId)
}


// S3 Functions
func ReadFromObject(key string)	*io.ReadCloser {
	sess := loadAwsConfig()
	svc := s3.New(sess)
	ctx := context.Background()
	result, err := svc.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket: bucket,
		Key:	aws.String(key),
	})
	if err != nil 	{
		log.Panic(err)
	}
	return &result.Body
}

func DownloadObject(key string, w http.ResponseWriter)	{
	sess := loadAwsConfig()
	buff := &aws.WriteAtBuffer{}
	downloader := s3manager.NewDownloader(sess)
	_, err := downloader.Download(buff, &s3.GetObjectInput{
		Bucket: bucket,
		Key:	aws.String(key),
	})
	if err != nil {
		log.Panic(err)
	}
	io.Copy(w, bytes.NewReader(buff.Bytes()))
}

func GeneratePresignedUrl(key string) string {
	sess := loadAwsConfig()
	svc := s3.New(sess)
	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: bucket,
		Key:    aws.String(key),
		ACL:	aws.String("public-read"),
	})
	url, _, err := req.PresignRequest(time.Hour * 1)
	if err != nil {
		log.Panic(err)
	}
	return url
}
