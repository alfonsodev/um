-- db: usermanager
-- mantainer: Alfonso Fernandez <alfonso@isla.io>
-- Simple user managment database inspired on github user, team organization structure.
DROP SCHEMA IF EXISTS usermanager CASCADE;

CREATE SCHEMA usermanager;

DROP TABLE IF EXISTS usermanager.users_orgs;
DROP TABLE IF EXISTS usermanager.users;
DROP TABLE IF EXISTS usermanager.teams;
DROP TABLE IF EXISTS usermanager.orgs;
DROP TABLE IF EXISTS usermanager.envs;

CREATE TABLE usermanager.users (
  id serial UNIQUE NOT NULL,
  username varchar(20),
  fullname varchar(255),
  email varchar(255),
  location varchar(255),
  person json
);

CREATE TABLE usermanager.orgs (
  id serial UNIQUE NOT NULL,
  name varchar(60),
  website varchar(255),
  location varchar(255)
);


CREATE TABLE usermanager.teams (
  id serial UNIQUE NOT NULL,
  orgs_id integer NOT NULL REFERENCES usermanager.orgs(id),
  name varchar(255),
  description varchar(255),
  permission smallint
);

CREATE TABLE usermanager.envs (
  id serial UNIQUE NOT NULL,
  name varchar(255)
);

CREATE TABLE usermanager.users_orgs  (
  users_id integer NOT NULL REFERENCES usermanager.users(id),
  orgs_id integer NOT NULL REFERENCES usermanager.orgs(id),
  teams_id integer NOT NULL REFERENCES usermanager.teams(id),
  PRIMARY KEY (users_id, orgs_id, teams_id)
);
