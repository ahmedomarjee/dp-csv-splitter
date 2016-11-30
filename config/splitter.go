package config

import (
	"github.com/ONSdigital/go-ns/log"
	"os"
)

const bindAddrKey = "BIND_ADDR"
const kafkaAddrKey = "KAFKA_ADDR"
const s3BucketKey = "S3_BUCKET"
const awsRegionKey = "AWS_REGION"
const topicNameKey = "TOPIC_NAME"

// BindAddr the address to bind to.
var BindAddr = ":21000"

// KafkaAddr the Kafka address to send messages to.
var KafkaAddr = "localhost:9092"

// S3Bucket the name of the AWS s3 bucket to get the CSV files from.
var S3Bucket = "dp-csv-splitter-1"

// AWSRegion the AWS region to use.
var AWSRegion = "eu-west-1"

// TopicName the name of the Kafka topic to send messages to.
var TopicName = "test"

func init() {
	if bindAddrEnv := os.Getenv(bindAddrKey); len(bindAddrEnv) > 0 {
		BindAddr = bindAddrEnv
	}

	if kafkaAddrEnv := os.Getenv(kafkaAddrKey); len(kafkaAddrEnv) > 0 {
		KafkaAddr = kafkaAddrEnv
	}

	if s3BucketEnv := os.Getenv(s3BucketKey); len(s3BucketEnv) > 0 {
		S3Bucket = s3BucketEnv
	}

	if awsRegionEnv := os.Getenv(awsRegionKey); len(awsRegionEnv) > 0 {
		AWSRegion = awsRegionEnv
	}

	if topicNameEnv := os.Getenv(topicNameKey); len(topicNameEnv) > 0 {
		TopicName = topicNameEnv
	}

	log.Debug("dp-csv-splitter Configuration", log.Data{
		bindAddrKey:  BindAddr,
		kafkaAddrKey: KafkaAddr,
		s3BucketKey:  S3Bucket,
		awsRegionKey: AWSRegion,
		topicNameKey: TopicName,
	})
}