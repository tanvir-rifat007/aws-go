// api layer talks to the database layer, so it needs to import the database package.

package api

import "awsgo.tanvirrifat.io/databse"


type ApiHandler struct{
	dbStore databse.DynamoDBClient
}

func NewApiHandler(dbStore databse.DynamoDBClient) ApiHandler{
	return ApiHandler{
		dbStore:dbStore,
	}
}