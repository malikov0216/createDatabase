package migration

const (
	CreateInteraction = `CREATE TABLE IF NOT EXISTS interactions (
		id serial PRIMARY KEY,
		login text,
		phone text,
		direction text,
		interaction_type text,
		source text,
		reason1 text,
		reason2 text,
		reason3 text,
		reason4 text,
		title text,
		notes text,
		start_date text,
		end_date text,
		last_name text,
		first_name text,
		email text
	)`
	TwoFA = `CREATE TABLE IF NOT EXISTS tfa_users (
		id serial PRIMARY KEY,
		login text,
		email text,
		key text,
		reg_date timestamptz,
		last_success_date timestamptz,
		last_unsuccess_date timestamptz,
		unsuccess_auth_count int,
		last_auth_type text,
		block_date timestamptz
	)`

	Cti = `CREATE TABLE IF NOT EXISTS cti_logs (
		id serial PRIMARY KEY,
		date_time timestamptz,
		method text,
		ctn text,
		login text,
		call_id text,
		call_type text,
		call_origin text,
		ivr_point text,
		agent_id text,
		skill_id text,
		talk_duration int,
		call_end_reason text,
		target_skill_id text,
		target_ivr_point text,
		target_number text
	)`

	Socs = `CREATE TABLE IF NOT EXISTS socs (
		id serial PRIMARY KEY,
		soc_code text,
		soc_desc text
	)`

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
