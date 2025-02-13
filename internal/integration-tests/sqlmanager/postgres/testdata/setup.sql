CREATE SCHEMA IF NOT EXISTS "sqlmanagerpostgres@special";

SET search_path TO "sqlmanagerpostgres@special";

CREATE TABLE users (
    id TEXT NOT NULL,
    age int NOT NULL,
    current_salary float NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    fullname TEXT GENERATED ALWAYS AS (first_name || ' ' || last_name) STORED
);

CREATE TABLE users_with_identity (
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    age int NOT NULL
);

CREATE TABLE unique_emails (
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    email text not null UNIQUE
);
CREATE TABLE unique_emails_and_usernames (
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    email text not null,
    username text not null,
    CONSTRAINT unique_email_username UNIQUE (email, username)
);

CREATE TABLE parent1 (id uuid NOT NULL DEFAULT gen_random_uuid(), CONSTRAINT pk_parent1_id PRIMARY KEY (id));
CREATE TABLE child1 (
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    parent_id uuid NULL,
    CONSTRAINT pk_child1_id PRIMARY KEY (id),
    CONSTRAINT fk_child1_parent_id_parent1_id FOREIGN KEY (parent_id) REFERENCES parent1(id) ON
    DELETE
        CASCADE
);

-- testing basic circular deps
CREATE TABLE t1 (
    a int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

    b int NULL,
    CONSTRAINT t1_b_fkey FOREIGN KEY (b) REFERENCES t1(a)
);
CREATE TABLE t2 (
    a int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

    b int NULL
);
CREATE TABLE t3 (
    a int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

    b int NULL
);
ALTER TABLE t2
ADD CONSTRAINT t2_b_fkey FOREIGN KEY (b) REFERENCES t3(a);
ALTER TABLE t3
ADD CONSTRAINT t3_b_fkey FOREIGN KEY (b) REFERENCES t2(a);

-- Testing composite keys
CREATE TABLE t4 (
    a int NOT NULL,
    b int NOT NULL,
    c int NULL,
    PRIMARY KEY (a, b)
);
CREATE TABLE t5 (
    x int NOT NULL,
    y int NOT NULL,
    z int NULL,
    CONSTRAINT t5_t4_fkey FOREIGN KEY (x, y) REFERENCES t4 (a, b)
);

CREATE SEQUENCE custom_seq
  START WITH 1
  INCREMENT BY 1
  NO MINVALUE
  NO MAXVALUE
  CACHE 1;
CREATE TYPE custom_type AS (
  part1 INTEGER,
  part2 TEXT
);

CREATE TYPE custom_type2 AS (
  part1 INTEGER,
  part2 TEXT
);

CREATE FUNCTION custom_function() RETURNS TRIGGER AS $$
BEGIN
  NEW.id := nextval('custom_seq');
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TYPE custom_enum AS ENUM (
  'value1',
  'value2',
  'value3'
);
CREATE DOMAIN custom_domain AS TEXT
  CHECK (VALUE ~ '^[a-zA-Z]+$'); -- Only allows alphabetic characters

CREATE DOMAIN custom_domain2 AS TEXT
  CHECK (VALUE ~ '^[a-zA-Z]+$'); -- Only allows alphabetic characters

CREATE TYPE custom_enum2 AS ENUM (
    'value4',
    'value5',
    'value6'
);

CREATE TABLE custom_table (
  id INTEGER NOT NULL DEFAULT nextval('custom_seq'),
  name custom_domain NOT NULL,
  names custom_domain2[] NOT NULL, -- testing custom domain as array
  data custom_type,
  datas custom_type2[] NOT NULL, -- testing custom composite type as array
  status custom_enum NOT NULL,
  statuses custom_enum2[] NOT NULL, -- testing custom data type as array
  created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);
-- Adding a trigger to use the custom function for setting the 'id' field
CREATE TRIGGER set_custom_id
  BEFORE INSERT ON custom_table
  FOR EACH ROW
  EXECUTE FUNCTION custom_function();
CREATE INDEX idx_name ON custom_table(name);

CREATE TABLE tablewithcount (
    id TEXT NOT NULL
);
INSERT INTO tablewithcount(id) VALUES ('1'), ('2');


CREATE TABLE defaults_table (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    description TEXT,
    age INT DEFAULT 18,
    is_active BOOLEAN DEFAULT true,
    registration_date DATE DEFAULT CURRENT_DATE,
    last_login TIMESTAMP,
    score NUMERIC(10,2) DEFAULT 0.00,
    status VARCHAR(20) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    sequence_number INT GENERATED BY DEFAULT AS IDENTITY,
    uuid UUID DEFAULT gen_random_uuid(),
    serial_number SERIAL
);


CREATE SCHEMA IF NOT EXISTS "CaPiTaL";
CREATE TABLE IF NOT EXISTS "CaPiTaL"."BadName" (
    "ID" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "NAME" text UNIQUE
);

INSERT INTO "CaPiTaL"."BadName" ("NAME")
VALUES
    ('Xk7pQ9nM3v'),
    ('Rt5wLjH2yB'),
    ('Zc8fAe4dN6'),
    ('Ym9gKu3sW5'),
    ('Vb4nPx7tJ2');

CREATE TABLE "CaPiTaL"."Bad Name 123!@#" (
    "ID" SERIAL PRIMARY KEY,
    "NAME" text REFERENCES "CaPiTaL"."BadName" ("NAME")
);


INSERT INTO "CaPiTaL"."Bad Name 123!@#" ("NAME")
VALUES
    ('Xk7pQ9nM3v'),
    ('Rt5wLjH2yB'),
    ('Zc8fAe4dN6'),
    ('Ym9gKu3sW5'),
    ('Vb4nPx7tJ2');



CREATE SCHEMA IF NOT EXISTS "bookings";
SET search_path TO "bookings";


CREATE TABLE aircrafts_data (
    aircraft_code character(3) NOT NULL,
    model jsonb NOT NULL,
    range integer NOT NULL,
    CONSTRAINT aircrafts_range_check CHECK ((range > 0))
);



CREATE TABLE airports_data (
    airport_code character(3) NOT NULL,
    airport_name jsonb NOT NULL,
    city jsonb NOT NULL,
    coordinates point NOT NULL,
    timezone text NOT NULL
);


CREATE TABLE boarding_passes (
    ticket_no character(13) NOT NULL,
    flight_id integer NOT NULL,
    boarding_no integer NOT NULL,
    seat_no character varying(4) NOT NULL
);


CREATE TABLE bookings (
    book_ref character(6) NOT NULL,
    book_date timestamp with time zone NOT NULL,
    total_amount numeric(10,2) NOT NULL
);


CREATE TABLE flights (
    flight_id integer NOT NULL,
    flight_no character(6) NOT NULL,
    scheduled_departure timestamp with time zone NOT NULL,
    scheduled_arrival timestamp with time zone NOT NULL,
    departure_airport character(3) NOT NULL,
    arrival_airport character(3) NOT NULL,
    status character varying(20) NOT NULL,
    aircraft_code character(3) NOT NULL,
    actual_departure timestamp with time zone,
    actual_arrival timestamp with time zone,
    CONSTRAINT flights_check CHECK ((scheduled_arrival > scheduled_departure)),
    CONSTRAINT flights_check1 CHECK (((actual_arrival IS NULL) OR ((actual_departure IS NOT NULL) AND (actual_arrival IS NOT NULL) AND (actual_arrival > actual_departure)))),
    CONSTRAINT flights_status_check CHECK (((status)::text = ANY (ARRAY[('On Time'::character varying)::text, ('Delayed'::character varying)::text, ('Departed'::character varying)::text, ('Arrived'::character varying)::text, ('Scheduled'::character varying)::text, ('Cancelled'::character varying)::text])))
);



CREATE SEQUENCE flights_flight_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE flights_flight_id_seq OWNED BY flights.flight_id;

CREATE TABLE seats (
    aircraft_code character(3) NOT NULL,
    seat_no character varying(4) NOT NULL,
    fare_conditions character varying(10) NOT NULL,
    CONSTRAINT seats_fare_conditions_check CHECK (((fare_conditions)::text = ANY (ARRAY[('Economy'::character varying)::text, ('Comfort'::character varying)::text, ('Business'::character varying)::text])))
);



CREATE TABLE ticket_flights (
    ticket_no character(13) NOT NULL,
    flight_id integer NOT NULL,
    fare_conditions character varying(10) NOT NULL,
    amount numeric(10,2) NOT NULL,
    CONSTRAINT ticket_flights_amount_check CHECK ((amount >= (0)::numeric)),
    CONSTRAINT ticket_flights_fare_conditions_check CHECK (((fare_conditions)::text = ANY (ARRAY[('Economy'::character varying)::text, ('Comfort'::character varying)::text, ('Business'::character varying)::text])))
);


CREATE TABLE tickets (
    ticket_no character(13) NOT NULL,
    book_ref character(6) NOT NULL,
    passenger_id character varying(20) NOT NULL,
    passenger_name text NOT NULL,
    contact_data jsonb
);


ALTER TABLE ONLY flights ALTER COLUMN flight_id SET DEFAULT nextval('flights_flight_id_seq'::regclass);

ALTER TABLE ONLY aircrafts_data
    ADD CONSTRAINT aircrafts_pkey PRIMARY KEY (aircraft_code);


ALTER TABLE ONLY airports_data
    ADD CONSTRAINT airports_data_pkey PRIMARY KEY (airport_code);


ALTER TABLE ONLY boarding_passes
    ADD CONSTRAINT boarding_passes_flight_id_boarding_no_key UNIQUE (flight_id, boarding_no);


ALTER TABLE ONLY boarding_passes
    ADD CONSTRAINT boarding_passes_flight_id_seat_no_key UNIQUE (flight_id, seat_no);


ALTER TABLE ONLY boarding_passes
    ADD CONSTRAINT boarding_passes_pkey PRIMARY KEY (ticket_no, flight_id);


ALTER TABLE ONLY bookings
    ADD CONSTRAINT bookings_pkey PRIMARY KEY (book_ref);


ALTER TABLE ONLY flights
    ADD CONSTRAINT flights_flight_no_scheduled_departure_key UNIQUE (flight_no, scheduled_departure);


ALTER TABLE ONLY flights
    ADD CONSTRAINT flights_pkey PRIMARY KEY (flight_id);


ALTER TABLE ONLY seats
    ADD CONSTRAINT seats_pkey PRIMARY KEY (aircraft_code, seat_no);


ALTER TABLE ONLY ticket_flights
    ADD CONSTRAINT ticket_flights_pkey PRIMARY KEY (ticket_no, flight_id);


ALTER TABLE ONLY tickets
    ADD CONSTRAINT tickets_pkey PRIMARY KEY (ticket_no);


ALTER TABLE ONLY boarding_passes
    ADD CONSTRAINT boarding_passes_ticket_no_fkey FOREIGN KEY (ticket_no, flight_id) REFERENCES ticket_flights(ticket_no, flight_id);


ALTER TABLE ONLY flights
    ADD CONSTRAINT flights_aircraft_code_fkey FOREIGN KEY (aircraft_code) REFERENCES aircrafts_data(aircraft_code);


ALTER TABLE ONLY flights
    ADD CONSTRAINT flights_arrival_airport_fkey FOREIGN KEY (arrival_airport) REFERENCES airports_data(airport_code);


ALTER TABLE ONLY flights
    ADD CONSTRAINT flights_departure_airport_fkey FOREIGN KEY (departure_airport) REFERENCES airports_data(airport_code);


ALTER TABLE ONLY seats
    ADD CONSTRAINT seats_aircraft_code_fkey FOREIGN KEY (aircraft_code) REFERENCES aircrafts_data(aircraft_code) ON DELETE CASCADE;


ALTER TABLE ONLY ticket_flights
    ADD CONSTRAINT ticket_flights_flight_id_fkey FOREIGN KEY (flight_id) REFERENCES flights(flight_id);


ALTER TABLE ONLY ticket_flights
    ADD CONSTRAINT ticket_flights_ticket_no_fkey FOREIGN KEY (ticket_no) REFERENCES tickets(ticket_no);


ALTER TABLE ONLY tickets
    ADD CONSTRAINT tickets_book_ref_fkey FOREIGN KEY (book_ref) REFERENCES bookings(book_ref);


