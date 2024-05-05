package database

import (
	"fmt"
	"github.com/ahmos0/DyanamodbConnectMobile/pkg/models"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetAllItems() ([]models.Item, error) {
	svc, err := createDynamoDBService()
	if err != nil {
		return nil, fmt.Errorf("Error creating DynamoDB service: %v", err)
	}

	input := &dynamodb.ScanInput{
		TableName: aws.String("TegeiGallery"),
	}

	result, err := svc.Scan(input)
	if err != nil {
		return nil, fmt.Errorf("Failed to scan DynamoDB: %v", err)
	}

	fmt.Printf("Found %d items in the table.\n", len(result.Items))
	var items []models.Item

	for _, item := range result.Items {
		var newItem models.Item
		err := dynamodbattribute.UnmarshalMap(item, &newItem)
		if err != nil {
			return nil, fmt.Errorf("Error unmarshalling item: %v", err)
		}
		items = append(items, newItem)
	}

	return items, nil
}

func SaveItem(uuid string, works string, author string, url string, other string) (*models.Item, error) {
	svc, err := createDynamoDBService()
	if err != nil {
		log.Fatalf("Error creating DynamoDB service: %v", err)
	}

	log.Printf("UUID: %s", uuid)
	log.Printf("Works: %s", works)
	log.Printf("Author: %s", author)
	log.Printf("URL: %s", url)
	log.Printf("Other: %s", other)

	input := &dynamodb.PutItemInput{
		TableName: aws.String("TegeiGallery"),
		Item: map[string]*dynamodb.AttributeValue{
			"uuid":   {S: aws.String(uuid)},
			"works":  {S: aws.String(works)},
			"author": {S: aws.String(author)},
			"url":    {S: aws.String(url)},
			"other":  {S: aws.String(other)},
		},
	}
	_, err = svc.PutItem(input)
	if err != nil {
		log.Printf("Error putting item: %v", err)
		return nil, err
	}

	item := &models.Item{
		UUID: uuid,
	}
	return item, nil
}

func createDynamoDBService() (*dynamodb.DynamoDB, error) {
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		log.Printf("Error creating AWS session: %v", err)
		return nil, err
	}

	svc := dynamodb.New(sess)
	return svc, nil
}
