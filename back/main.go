package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"cloud.google.com/go/bigquery"
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

func handler() ([][]bigquery.Value, error) {
	companyData, err := queryWithNamedParams()
	if err != nil {
		fmt.Println(err)
	}
	return companyData, nil
}

// type CompanyRowData struct {
// 	Company_name bigquery.NullString  `json:"companyName"`
// 	Measure_tag  bigquery.NullString  `json:"measureTag"`
// 	Value        bigquery.NullFloat64 `json:"value"`
// 	Units        bigquery.NullString  `json:"units"`
// 	Fiscal_year  bigquery.NullInt64   `json:"fiscalYear"`
// }

func queryWithNamedParams() ([][]bigquery.Value, error) {
	// projectID := "my-project-id"

	q := client.Query(
		`SELECT company_name, measure_tag, value, units, fiscal_year
		FROM ` + "`bigquery-public-data.sec_quarterly_financials.quick_summary`" + `
		WHERE company_name = @company_name
		AND form = @form
		AND fiscal_year = @fiscal_year
		AND period_end_date = @period_end_date
		ORDER BY measure_tag ASC
		LIMIT 5;`)
	q.Parameters = []bigquery.QueryParameter{
		{
			Name:  "company_name",
			Value: "ORACLE CORP",
		},
		{
			Name:  "form",
			Value: "10-K",
		},
		{
			Name:  "fiscal_year",
			Value: 2017,
		},
		{
			Name:  "period_end_date",
			Value: "20170531",
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
	var rows [][]bigquery.Value
	for {
		var row []bigquery.Value
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
