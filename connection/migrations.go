package connection

var migrations = []string{
	`CREATE TABLE IF NOT EXISTS short_urls(
		key 			TEXT NOT NULL PRIMARY KEY, 
		original_url 	TEXT NOT NULL UNIQUE,
		expire_date		INTEGER NOT NULL
	)`,
}
