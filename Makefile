DBNAME=foo
structure:
	psql ${DBNAME} -f usermanager.sql
fixtures:
	psql $(DBNAME) -f fixtures.sql
test: structure fixtures

.PHONY: structure, fixtures, test
