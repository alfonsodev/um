# usermanager
[![Build Status](https://travis-ci.org/alfonsodev/usermanager.svg?branch=master)](https://travis-ci.org/alfonsodev/usermanager)  [![Coverage Status](https://coveralls.io/repos/alfonsodev/usermanager/badge.svg?branch=master)](https://coveralls.io/r/alfonsodev/usermanager?branch=master)  

It's a simple user managment system inspeired on Github's Users and Organizations apis.
*Disclaimer* This is *WIP* help is wanted, and pull request are welcome! 

## Planed features (help wanted!)
- Google signin
- Github sigin 
- User register process
- Organizations
- Teams
- Permissions
- Web gui 
- Command line tool
- REST api 

## Requirements
- Postgresql
- go
## Configuration
Database name is defined in a env variable USRMNG_DBNAME
## Api usage

## Library usage

## Command line Usage
`
  go get github.com/alfonsodev/usermanager
`
Create database structure with 
`
  # Create datbase strucutre (will create usermanager schema on db) usermanager.sql
  usermanager create-database
`


