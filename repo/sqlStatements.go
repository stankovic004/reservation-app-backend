package repo

var sqlStatements = map[string]string{
	"register": `INSERT INTO users (email, username, password, role ,date_added)
		VALUES ($1, $2, $3, $4, NOW()) RETURNING id;`,

	"login": "select username, role from users where email = $1 AND password = $2;",

	"reserve": `INSERT INTO reservations (username, location, date)
	VALUES ($1, $2, $3)`,

	"addLocation": `INSERT INTO locations (name, lat, lon) VALUES ($1, $2, $3);`,

	"get_locations": "select id, name, lon, lat from locations",

	"get_reservations": "select id, username, location, date from reservations"}
