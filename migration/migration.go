package migration

const (
	Offers = `CREATE TABLE IF NOT EXISTS offer_logs (
		id serial PRIMARY KEY,
		date_time timestamptz,
		user_firstname text,
		user_lastname text,
		user_position text,
		user_department text,
		user_ip text,
		pc_name text,
		user_activity text,
		abonent_ctn text,
		sms_text text,
		login text,
		current_tariff text,
		pre_tariff text,
		offer text,
		status text,
		system text,
		errcode text,
		offers text,
		request_type text,
		offer_reward int
	)`

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
