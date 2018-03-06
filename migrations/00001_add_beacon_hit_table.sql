-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE beacon_hit (
	id SERIAL,
	date timestamp with time zone,
	scheme varchar(10),
	host varchar(255),
	path varchar(255),
	query varchar(255),
	fragment varchar(255)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE beacon_hit;
