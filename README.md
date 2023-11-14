# sqlstring

Simple SQL escape and format 

[![Go](https://github.com/feiin/sqlstring/actions/workflows/go.yml/badge.svg)](https://github.com/feiin/sqlstring/actions/workflows/go.yml)
[![GoDoc](https://godoc.org/github.com/feiin/sqlstring?status.svg)](https://godoc.org/github.com/feiin/sqlstring)

## Escaping sql values

```golang
//Format
sql := sqlstring.Format("select * from users where name=? and age=? limit ?,?", "t'est", 10, 10, 10)

fmt.Printf("sql: %s",sql)

//Escape
sql = "select * from users WHERE name = " + sqlstring.Escape(name);
fmt.Printf("sql: %s",sql)

```

## License

MIT