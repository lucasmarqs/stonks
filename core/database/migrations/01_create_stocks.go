package migrations

var Up_01_create_stocks = []string{
	`CREATE TABLE stocks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(20) NOT NULL UNIQUE COLLATE NOCASE,
		last_quote DOUBLE,
		last_quote_date DATETIME
	);`,
}
