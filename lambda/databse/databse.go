// database layer is the lowest layer in the application, so it doesn't import any other package.

package databse

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDBClient struct{

	databaseStore *dynamodb.DynamoDB
}

func NewDynamoDBClient() DynamoDBClient{
	return DynamoDBClient{
		databaseStore: dynamodb.New(session.Must(session.NewSession())),
		
	}
}