package main

import (
	"fmt"
	"log"
	"path"
	"time"

	"github.com/andygrunwald/go-jira"
	"github.com/gramoflava/jira-go-timesheet/internal/pkg/appconfig"
	"github.com/gramoflava/jira-go-timesheet/internal/pkg/writer/csv"
)

const secondsPerHour = 3600
const configFilePath = "~/.config/jira-go-timesheet/config.json"

func main() {
	cfg := appconfig.GetDummy()
	log.Println("Got configuration.")

	tp := jira.BasicAuthTransport{
		Username: cfg.Servers[0].Login,
		Password: cfg.Servers[0].Password,
	}

	client, _ := jira.NewClient(tp.Client(), cfg.Servers[0].URL)
	log.Println("Connection established.")
	log.Printf("Running query: '%s'.\n", cfg.Servers[0].BaseJQL)

	issues, _, _ := client.Issue.Search(cfg.Servers[0].BaseJQL, nil)
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

	filename := path.Join(cfg.TargetDir, "Timesheet.csv")
	log.Printf("Dumping CSV to '%s'.\n", filename)

	csv.Write(filename, timesheets)
}
