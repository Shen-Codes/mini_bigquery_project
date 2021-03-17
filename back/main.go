package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"cloud.google.com/go/bigquery"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var client *bigquery.Client
var ctx context.Context

func init() {
	svc := s3.New(session.New())
	input := &s3.GetObjectInput{
		Bucket: aws.String("jsonfiles312021"),
		Key:    aws.String("Project-8a8c500b8c6d.json"),
	}
	result, _ := svc.GetObject(input)
	defer result.Body.Close()
	body, _ := ioutil.ReadAll(result.Body)

	ctx = context.Background()

	var err error
	client, err = bigquery.NewClient(ctx, "first-vision-305321", option.WithCredentialsJSON(body))
	if err != nil {
		log.Fatalf("bigquery.NewClient: %v", err)
	}

}

func main() {
	lambda.Start(handler)
	defer client.Close()
}

func handler() {
	companyData, err := queryWithNamedParams()
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()
	fmt.Println(companyData)
}

type CompanyRowData struct {
	Company_name bigquery.NullString  `json:"companyName"`
	Measure_tag  bigquery.NullString  `json:"measureTag"`
	Value        bigquery.NullFloat64 `json:"value"`
	Units        bigquery.NullString  `json:"units"`
	Fiscal_year  bigquery.NullInt64   `json:"fiscalYear"`
}

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
