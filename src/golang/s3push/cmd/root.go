package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/spf13/cobra"
)

var bucket string
var target string
var filename string

var svc *s3manager.Uploader

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "s3push",
	Short: "Push files to s3 bucket.",
	Long:  `Push files to s3 bucket.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()

		key := target
		if key == "" {
			key = filepath.Base(filename)
		}
		fmt.Println("Uploading file to S3...")
		result, err := svc.Upload(&s3manager.UploadInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
			Body:   file,
		})

		if err != nil {
			fmt.Println("error", err)
			os.Exit(1)
		}

		fmt.Printf("Successfully uploaded %s to %s\n", filename, result.Location)
		return nil
	},
}

// Execute called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&filename, "filename", "f", "", "File to upload. (required)")
	rootCmd.PersistentFlags().StringVarP(&bucket, "bucket", "b", "", "Bucket to upload to. (required)")
	rootCmd.PersistentFlags().StringVarP(&target, "target", "t", "", "Object key to upload to. If left off, original filename will be used.")

	rootCmd.MarkPersistentFlagRequired("bucket")
	rootCmd.MarkPersistentFlagRequired("upload")

	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("Error creating session ", err)
		os.Exit(1)
	}
	svc = s3manager.NewUploader(sess)
}
