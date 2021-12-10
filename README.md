# sqlstring

Simple SQL escape and format 


## Escaping sql values

```
//Format
sql := sqlstring.Format("select * from users where name=? and age=? limit ?,?", "t'est", 10, 10, 10)

fmt.Printf("sql: %s",sql)

//Escape
sql = "select * from users WHERE name = " + sqlstring.Escape(name);
fmt.Printf("sql: %s",sql)

```

## License

MIT