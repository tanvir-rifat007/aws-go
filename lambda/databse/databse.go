// database layer is the lowest layer in the application, so it doesn't import any other package.

package databse

import (
	"awsgo.tanvirrifat.io/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const TABLE_NAME = "userTable"

type DynamoDBClient struct{

	databaseStore *dynamodb.DynamoDB
}

func NewDynamoDBClient() DynamoDBClient{
	return DynamoDBClient{
		databaseStore: dynamodb.New(session.Must(session.NewSession())),
		
	}
}


func (u DynamoDBClient) IsUserExists(username string) (bool,error){

	result,err:= u.databaseStore.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"username":{
				 S: aws.String(username),
			},
		},

	})

	// if user exist
	// but some aws error occured

	if err!=nil{
		return true,err
	}

	// if user doesn't exist

	if result.Item==nil{
		return false,nil
	}

	// if user exist

	return true,nil

}


func (u DynamoDBClient) AddUser(user types.Register) error{
	
	 item:= &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item: map[string]*dynamodb.AttributeValue{
			"username":{
				S: aws.String(user.Username),
			},
			"password":{
				S: aws.String(user.Password),
			},
		},
	 }

	_, err:= u.databaseStore.PutItem(item)


	if err!=nil{
		return err
	}

	return nil


}