package migrations

var Up_03_create_reports = []string{
	`CREATE TABLE reports (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		created_at DATETIME NOT NULL,
		diff_of_last_report DOUBLE,
		diff_of_last_month DOUBLE
	);`,

	`CREATE INDEX idx_reports_created_at ON reports (created_at);`,
}
