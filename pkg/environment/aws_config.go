package environment

import (
	"github.com/spf13/pflag"
)

const (
	awsConfigFile = "config/aws.env"
)

type awsConfig struct {
	AccessKey       string
	SecretAccessKey string
	BucketName      string
	Region          string
}

func NewAwsConfig() *awsConfig {
	return &awsConfig{
		AccessKey:       "",
		SecretAccessKey: "",
		BucketName:      "finease-uwaterloo",
		Region:          "us-east-1",
	}
}

func (c *awsConfig) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&c.AccessKey, "aws-access-key", c.AccessKey, "Access key to access the Amazon S3 bucket")
	fs.StringVar(&c.SecretAccessKey, "aws-secret-access-key", c.SecretAccessKey, "Secret access key to access the Amazon S3 bucket")
	fs.StringVar(&c.BucketName, "aws-bucket-name", c.BucketName, "Name of Amazon S3 bucket")
	fs.StringVar(&c.Region, "aws-region", c.Region, "Region from the Amazon S3 bucket is exposed")
}

func (c *awsConfig) ReadFromFile() error {
	return nil
}

func (c *awsConfig) ReadFromEnv() error {
	c.AccessKey = getEnvDefault("AWS_ACCESS_KEY", c.AccessKey)
	c.SecretAccessKey = getEnvDefault("AWS_SECRET_ACCESS_KEY", c.SecretAccessKey)
	c.BucketName = getEnvDefault("AWS_BUCKET_NAME", c.BucketName)
	c.Region = getEnvDefault("AWS_REGION", c.Region)
	return nil
}
