package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/andygrunwald/go-jira"
)

const secondsPerHour = 3600
const configFilePath = "~/.config/jira-go-timesheet/config.json"

type config struct {
	URL       string
	BaseJQL   string
	Login     string
	Password  string
	DateFrom  string
	DateUntil string
	CSVOut    string
}

func getConfig() config {
	//TODO: Add reading from the actual config file
	//https://www.thepolyglotdeveloper.com/2017/04/load-json-configuration-file-golang-application/

	var cfg config
	cfg.URL = "https://tandbergdata.atlassian.net"
	cfg.BaseJQL = "PROJECT in (RDX, VTX2U, VTX1U) AND worklogAuthor in (v.redzhepov, vshakhov, a.hrytsevich) AND timespent is not EMPTY"
	cfg.Login = "e.lavnikevich@sam-solutions.com"
	cfg.Password = "Js6Us47UtB78qcsY9[qP"
	cfg.DateFrom = "2018-09-01"
	cfg.DateUntil = "2018-09-30"
	cfg.CSVOut = "/Users/lava/Timesheet.csv"
	return cfg
}

func writeCSV(path string, records [][]string) {
	file, _ := os.Create(path)
	defer file.Close()

	w := csv.NewWriter(file)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("Error writing record to csv: ", err)
		}
	}

	defer w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	cfg := getConfig()
	log.Println("Got configuration.")

	tp := jira.BasicAuthTransport{
		Username: cfg.Login,
		Password: cfg.Password,
	}

	client, _ := jira.NewClient(tp.Client(), cfg.URL)
	log.Println("Connection established.")
	log.Printf("Running query: '%s'.\n", cfg.BaseJQL)

	issues, _, _ := client.Issue.Search(cfg.BaseJQL, nil)
	log.Printf("%d issues found.\n", len(issues))

	timesheets := make([][]string, 0)

	i := 0

	log.Print("Extracting worklogs:")
	for _, issue := range issues {

		worklogs, _, _ := client.Issue.GetWorklogs(issue.ID)
		log.Printf(" %d", len(worklogs.Worklogs))
		for _, record := range worklogs.Worklogs {
			started := time.Time(*record.Started)

			timesheets = append(timesheets, make([]string, 0))

			timesheets[i] = append(timesheets[i], fmt.Sprintf("%s", issue.Fields.Project.Key))
			timesheets[i] = append(timesheets[i], fmt.Sprintf("%s", issue.Key))
			if issue.Fields.Parent == nil {
				timesheets[i] = append(timesheets[i], fmt.Sprintf("%s", issue.Fields.Summary))
			} else {
				parent, _, _ := client.Issue.Get(issue.Fields.Parent.Key, nil)
				timesheets[i] = append(timesheets[i], fmt.Sprintf("%s> %s", parent.Key, issue.Fields.Summary))
			}
			timesheets[i] = append(timesheets[i], fmt.Sprintf("%s", record.Author.DisplayName))
			timesheets[i] = append(timesheets[i], fmt.Sprintf("%d.%d.%d", started.Day(), started.Month(), started.Year()))
			timesheets[i] = append(timesheets[i], fmt.Sprintf("%.1f", float64(record.TimeSpentSeconds)/secondsPerHour))
			timesheets[i] = append(timesheets[i], fmt.Sprintf("%s", record.Comment))

			i++
		}
	}
	log.Println("Done.")
	log.Println("Dumping CSV.")

	writeCSV(cfg.CSVOut, timesheets)
}
