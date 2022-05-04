package repo

var sqlStatements = map[string]string{
	"register": `INSERT INTO users (email, username, password, date_added)
		VALUES ($1, $2, $3, NOW()) RETURNING id;`,

	"login": "select count(*) from users where email = $1 AND password = $2;",

	"reserve": `INSERT INTO reservations (username, date, time, location)
	VALUES ($1, $2, $3, $4 RETURNING id);`,

	"get_locations": "select id, name, lon, lat from locations"}
