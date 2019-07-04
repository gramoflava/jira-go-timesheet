# jira-go-timeshit

The sole purpose of the project is to bring back timesheets to the JIRA Cloud.

## History

Atlassian never was too user-friendly and always was greedy. For some time now they've been removing functions from the JIRA Cloud platform as "not essential". Some time in 2016 integrated Timesheet reports plugin also became the victim of their marketing. Currently, only paid plugins are available for integration with JIRA Cloud, and there is no way to extract timesheets without a lot of pain for everyone who doesn't want to pay.

## Vision

JIRA Cloud provides some API, which can be used to request worklogs. **jira-go-timeshit** will connect to the JIRA Cloud, get worklogs for the required period of time and format them as a CSV-file.

## Technical solution

The project is implemented in Go and only has CLI-interface planned.

Note: Only tested on a small project. I do not know well the API but have seen some paging involved.
