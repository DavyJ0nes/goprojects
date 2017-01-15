-- Connect to Postgres.
-- psql -U goprojects

-- Create a user named `goprojects` with the password `goprojects`. If you ever
-- do this in production, please use a better password.
--CREATE USER goprojects WITH PASSWORD 'goprojects';

-- Create the database we're going to use.
-- CREATE DATABASE goprojects;

-- Grant all privleges to our user on the DB.
GRANT ALL PRIVILEGES ON DATABASE goprojects to goprojects;

-- Create a new table to store our pages.
CREATE TABLE IF NOT EXISTS PAGES(
  id             SERIAL    PRIMARY KEY,
  title          TEXT      NOT NULL,
  content        TEXT      NOT NULL
);

-- Create a new table to store our posts.
CREATE TABLE IF NOT EXISTS POSTS(
  id             SERIAL    PRIMARY KEY,
  title          TEXT      NOT NULL,
  content        TEXT      NOT NULL,
  date_created   DATE      NOT NULL
);

-- Create a new table to store our comments.
CREATE TABLE IF NOT EXISTS COMMENTS(
  id             SERIAL    PRIMARY KEY,
  author         TEXT      NOT NULL,
  content        TEXT      NOT NULL,
  date_created   DATE      NOT NULL,
  post_id        INT       references POSTS(id)
);
