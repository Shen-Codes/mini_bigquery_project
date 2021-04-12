package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"cloud.google.com/go/bigquery"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var client *bigquery.Client
var ctx context.Context

func init() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Region: aws.String("us-east-1"),
		},
	}))

	svc := ssm.New(sess)

	paramsFromAWS := paramsByPath(svc)
	paramsByte, _ := json.Marshal(paramsFromAWS)
	ctx = context.Background()

	var err error
	client, err = bigquery.NewClient(ctx, "first-vision-305321", option.WithCredentialsJSON(paramsByte))
	if err != nil {
		log.Fatalf("bigquery.NewClient: %v", err)
	}

}

func paramsByPath(svc *ssm.SSM) map[string]string {
	pathInput := &ssm.GetParametersByPathInput{
		Path: aws.String("/bqconfig"),
	}

	res, err := svc.GetParametersByPath(pathInput)
	if err != nil {
		log.Println(err)
	}

	params := make(map[string]string)

	for _, param := range res.Parameters {
		name := strings.Replace(*param.Name, "/bqconfig/", "", -1)
		value := *param.Value
		params[name] = value
	}

	return params
}

func main() {
	lambda.Start(handler)
	defer client.Close()
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	queryParams := req.QueryStringParameters
	companyData, err := queryWithNamedParams(queryParams)
	if err != nil {
		serverError(err)
	}
	body, _ := json.Marshal(companyData)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(body),
	}, nil
}

type BqStruct struct {
	Company_name    string  `json:"companyName"`
	Measure_Tag     string  `json:"measureTag"`
	Value           float64 `json:"value"`
	Units           string  `json:"units"`
	Period_end_date string  `json:"periodEndDate"`
}

func queryWithNamedParams(queryParams map[string]string) ([]BqStruct, error) {
	fiscalYear, _ := strconv.Atoi(queryParams["fiscal_year"])
	min, _ := strconv.ParseFloat(queryParams["min"], 32)
	max, _ := strconv.ParseFloat(queryParams["max"], 32)
	measureTag := queryParams["measure_tag"]

	fmt.Println(queryParams)

	q := client.Query(
		`SELECT company_name, measure_tag, value, units, period_end_date
		FROM ` + "`bigquery-public-data.sec_quarterly_financials.quick_summary`" + `
		WHERE fiscal_year = @fiscal_year
		AND form = "10-K"
		AND value BETWEEN @min and @max
		AND measure_tag IN (@measure_tag)
		ORDER BY company_name ASC
		LIMIT 50;`)
	q.Parameters = []bigquery.QueryParameter{
		{
			Name:  "fiscal_year",
			Value: fiscalYear,
		},
		{
			Name:  "min",
			Value: min,
		},
		{
			Name:  "max",
			Value: max,
		},
		{
			Name:  "measure_tag",
			Value: measureTag,
		},
	}
	// Run the query and print results when the query job is completed.
	job, err := q.Run(ctx)
	if err != nil {
		return nil, err
	}
	status, err := job.Wait(ctx)
	if err != nil {
		return nil, err
	}
	if err := status.Err(); err != nil {
		return nil, err
	}
	it, err := job.Read(ctx)
	var rows []BqStruct
	for {
		var row BqStruct
		err := it.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		rows = append(rows, row)
	}
	return rows, nil
}

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	errByte, _ := json.Marshal(err)
	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       string(errByte),
	}, nil
}
