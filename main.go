package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

func failOnError(err error) {
	if err != nil {
		fmt.Printf("::err ::error checking CLA %q", err)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) < 3 {
		panic("usage: clasee <sheet to check> <range to check>")
	}

	author := os.Getenv("GITHUB_ACTOR")
	if author == "" {
		failOnError(errors.New("no PR author provided"))
	}

	secret := os.Getenv("CLASEE_SECRET")
	if secret == "" {
		failOnError(errors.New("no secret provided"))
	}

	data, err := base64.StdEncoding.DecodeString(secret)
	failOnError(err)

	conf, err := google.JWTConfigFromJSON(data, sheets.SpreadsheetsScope)
	failOnError(err)

	client := conf.Client(context.TODO())

	srv, err := sheets.New(client)
	failOnError(err)

	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	// "1Pnv4Xv7n08tD8frCKhqprj5QxEszy4MtNJInIMS9jgs"
	spreadsheetId := os.Args[1]
	readRange := os.Args[2]
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	failOnError(err)

	for _, row := range resp.Values {
		if len(row) > 0 && row[0] == author {
			os.Exit(0)
		}
	}

	fmt.Println("::warning ::You need to sign the CLA")
	os.Exit(1)
}
