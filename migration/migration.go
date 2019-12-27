package migration

const (
	Users = `CREATE TABLE IF NOT EXISTS users (
		id serial PRIMARY KEY, 
		first_name text NOT NULL, 
		last_name text, 
		username text NOT NULL, 
		password text NOT NULL,
		email text,
		phone text NOT NULL,
		role text NOT NULL,
		flat_id int
	)`

	Flats = `CREATE TABLE IF NOT EXISTS flats (
		id serial PRIMARY KEY, 
		name text NOT NULL, 
		price int NOT NULL, 
		resident_id int, 
		is_free BOOLEAN NOT NULL
	)`

	Resident = `CREATE TABLE IF NOT EXISTS residents (
		id serial PRIMARY KEY, 
		name text NOT NULL, 
		contact text, 
		checkIn_date date NOT NULL, 
		checkOut_date date
	)`
	Payements = `CREATE TABLE IF NOT EXISTS payements (
		id serial PRIMARY KEY,
		reason text NOT NULL,
		accepted_person text NOT NULL,
		amount int NOT NULL,
		electric int NOT NULL,
		resident_id int NOT NULL,
		date date NOT NULL
	)`
)
