// app talks to the api layer

package app

import (
	"awsgo.tanvirrifat.io/api"
	"awsgo.tanvirrifat.io/databse"
)

type App struct{
	ApiHandler api.ApiHandler
}

func NewApp() App{
	db:= databse.NewDynamoDBClient()
	return App{
		ApiHandler: api.NewApiHandler(db),
	}
}