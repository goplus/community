# mysql-framework
1. "database/sql": This is a built-in database access interface for Go that can interact with a variety of databases, including MySQL. It provides a common set of interfaces and methods that can be adapted to different database drivers to achieve access to MySQL
2. "https://github.com/go-sql-driver/mysql": This is a popular MySQL driver that implements the database/sql interface and can be used directly in Go. It provides the ability to connect, query, and trade operations to the MySQL database
3. "https://github.com/go-mysql-org/go-mysql": This is a pure Go-based MySQL client library that provides low-level access to the MySQL protocol. It has a lower-level API that can be used to implement custom MySQL client functionality // with some limitations. Since it is the underlying MySQL client library, using it requires some knowledge of the MySQL protocol and may require more code writing to handle database operations

| Built-in SQL database interface |  go-sql-driver |                go-mysql | 
| --- | --- | --- |
| Standardized interface: database/sql is a built-in database access interface in Go, which provides a set of standardized APIs that can interact with a variety of databases. This allows you to switch between different database drivers without changing the code | Performance and compatibility: Since it is specifically designed for MySQL databases, "go-sql-driver/mysql" can provide better performance and compatibility. It provides more efficient database access by communicating directly with the MySQL server, leveraging MySQL's specific features and optimizations. |   Pure Go implementation: As a MySQL client library implemented in pure Go language, "go-mysql/mysql" does not depend on external C libraries and can be easily built and deployed on different platforms. This makes it easier to use and maintain. |  |
| Extensive database support: The database/sql interface supports many popular databases, such as MySQL, PostgreSQL, SQLite, etc. By using the appropriate drivers, you can easily connect and manipulate these databases. | Community support: As an active open source project, "go-sql-driver/mysql" is widely supported and maintained by the community. It has an active community with rich documentation, examples, and solutions that can help you better use and understand the driver. |  Direct access to the MySQL protocol: "go-mysql/mysql" provides low-level access to the MySQL protocol, allowing you to communicate directly with the MySQL server. This allows for better control and optimization of interactions with MySQL, providing more efficient database access. |  |
| Mature and stable: As part of the Go language, the database/sql interface has been extensively tested and used, and is considered a relatively mature and stable solution. It has good documentation and community support to help you solve problems and learn Xi use. |  |  Advanced function support: go-mysql/mysql provides some advanced functions, such as binary protocol support, replication protocol support, etc. These features allow you to implement more complex database operations, such as subscribing to MySQL binary logs, real-time data synchronization, and more |  

## conclusion
We want this project to have good performance and a low barrier to entry, so we will use go-sql-driver



Here's the code for each framework to connect to MySQL We can see that the second one is the most concise

```
database/sql
db, err := sql.Open("mysql",
    "root:123123@tcp(127.0.0.1:3306)/books?charset=utf8mb4&parseTime=true&loc=Local")
if err != nil {
    log.Fatal(err)
}
defer func() { _ = db.Close() }()
```

```
go-sql-driver
db, err := sql.Open("mysql", "root:123123@/books")
if err != nil {
    panic(err)
}
stmt, err := db.Prepare()
defer stmt.Close()
```

```
cfg := replication.BinlogSyncerConfig {
        ServerID: 100,
        Flavor:   "mysql",
        Host:     "127.0.0.1",
        Port:     3306,
        User:     "root",
        Password: "",
}
syncer := replication.NewBinlogSyncer(cfg)
```
