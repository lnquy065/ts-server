package configs

import "fmt"

var Dialect = "postgres"
var host = "ec2-107-22-238-217.compute-1.amazonaws.com"
var port = "5432"
var user = "tqgywnaaowjduz"
var password = "f156dfac12896acde4a13200ff09b74cf05f1b97bad83989335f62ab93a393d5"
var dbname = "d3n94qjae8lnna"

//sslmode=disable
var ConnectionString = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
	host, port, user, dbname, password)
