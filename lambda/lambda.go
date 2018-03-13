package lambda

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type LambdaServer struct {
	state string
}

type Request events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// stdout and stderr are sent to AWS CloudWatch Logs
	pp, err := json.Marshal(request.PathParameters)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Error1:" + err.Error(),
			StatusCode: 200,
		}, nil
	}
	qp, err := json.Marshal(request.QueryStringParameters)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Error3:" + err.Error(),
			StatusCode: 200,
		}, nil
	}
	sv, err := json.Marshal(request.StageVariables)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Error4:" + err.Error(),
			StatusCode: 200,
		}, nil
	}

	paths := strings.Split(request.PathParameters["proxy"], "/")

	return events.APIGatewayProxyResponse{
		Body: "Body:" + request.Body +
			" PathParams:" + string(pp) +
			" Path:" + request.Path +
			" ResourcePath:" + request.RequestContext.ResourcePath +
			" QueryPath:" + string(qp) +
			" StageVars:" + string(sv) +
			" Paths:" + string(paths[0]),
		StatusCode: 200,
	}, nil
}

func (l *LambdaServer) Start() error {
	lambda.Start(Handler)
}

func (l *LambdaServer) Stop(err error) {
	if err != nil {
		log.Println(err)
	}

	lambda.Stop()
}
