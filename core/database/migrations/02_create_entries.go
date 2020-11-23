package migrations

var Up_02_create_entries = []string{
	`CREATE TABLE entries (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		stock_id INTEGER NOT NULL,
		unit_price DOUBLE NOT NULL,
		quantity INTEGER NOT NULL,
		created_at DATETIME NOT NULL,
		FOREIGN KEY (stock_id) REFERENCES stocks(id)
	);`,
}
