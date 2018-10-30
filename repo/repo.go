package repo

var savedReports = []string{}

func SaveReport(report string) {
	savedReports = append(savedReports, report)
}
