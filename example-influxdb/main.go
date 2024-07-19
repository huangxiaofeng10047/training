package main

import (
	"context"
	"fmt"
	"github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	// Create a client
	// You can generate an API Token from the "API Tokens Tab" in the UI
	client := influxdb2.NewClient("http://localhost:8086", "L0_ZTbhAHBhSqmKf85vH46byjBuh1ISipwgxZIk1_tXh5jXNiFql_uk4iajpJ8Y6BpHF8pO-8lsrSoyHF_L1Dg==")
	// always close client at the end
	// get non-blocking write client
	writeAPI := client.WriteAPI("bcs", "dataupload")

	// write line protocol
	writeAPI.WriteRecord(fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0))
	// Flush writes
	writeAPI.Flush()
	// Get query client
	queryAPI := client.QueryAPI("bcs")

	query := `from(bucket:"dataupload")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "stat")`

	// get QueryTableResult
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		panic(err)
	}

	// Iterate over query response
	for result.Next() {
		// Notice when group key has changed
		if result.TableChanged() {
			fmt.Printf("table: %s\n", result.TableMetadata().String())
		}
		// Access data
		fmt.Printf("value: %v\n", result.Record().Value())
	}
	// check for an error
	if result.Err() != nil {
		fmt.Printf("query parsing error: %\n", result.Err().Error())
	}

	defer client.Close()
}
