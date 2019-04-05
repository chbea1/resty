Resty
--------

Resty is a command line tool to automatically create rest services
given sql database tables. Resty reads the schema and generates a strict
CRUD service which is configurable either by command line flags or the resty
config file.

Resty generates openAPI compliant yaml and ships with swagger already setup. 
Five endpoints will be generated for each table:

* /health GET - An EVC health check for load balancing.
* /create POST - Creates a record in the database table.
* /read GET - Retrieves a record from the database table.
* /update PUT - Updates a record from the database table.
* /delete DELETE - Deletes a record from the database table.


Resty also works in two seperate modes, interactive and code-gen. In interactive 
mode a rest service will be created and immediately start running. In code-gen
mode resty will output a Go rest service using gin to the output directory specified.
That rest service can then be run seperatly.



Status
-------------

Resty is not currently usable. Currently I am trying to map a database schema
to an openAPI yaml file. 
