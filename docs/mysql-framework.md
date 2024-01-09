# mysql-framework
1. "database/sql" ：This is a built-in database access interface for Go that can interact with a variety of databases, including MySQL. It provides a common set of interfaces and methods that can be adapted to different database drivers to achieve access to MySQL。
2. "https://github.com/go-sql-driver/mysql"：This is a popular MySQL driver that implements the database/sql interface and can be used directly in Go. It provides the ability to connect, query, and trade operations to the MySQL database。
3. "https://github.com/go-mysql-org/go-mysql"：This is a pure Go-based MySQL client library that provides low-level access to the MySQL protocol. It has a lower-level API that can be used to implement custom MySQL client functionality // with some limitations. Since it is the underlying MySQL client library, using it requires some knowledge of the MySQL protocol and may require more code writing to handle database operations。

| Built-in SQL database interface |  go-sql-driver |                go-mysql | 
| --- | --- | --- |
| Standardized interface: database/sql is a built-in database access interface in Go, which provides a set of standardized APIs that can interact with a variety of databases. This allows you to switch between different database drivers without changing the code | Performance and compatibility: Since it is specifically designed for MySQL databases, "go-sql-driver/mysql" can provide better performance and compatibility. It provides more efficient database access by communicating directly with the MySQL server, leveraging MySQL's specific features and optimizations. |   Pure Go implementation: As a MySQL client library implemented in pure Go language, "go-mysql/mysql" does not depend on external C libraries and can be easily built and deployed on different platforms. This makes it easier to use and maintain. |  |
| Extensive database support: The database/sql interface supports many popular databases, such as MySQL, PostgreSQL, SQLite, etc. By using the appropriate drivers, you can easily connect and manipulate these databases. | Community support: As an active open source project, "go-sql-driver/mysql" is widely supported and maintained by the community. It has an active community with rich documentation, examples, and solutions that can help you better use and understand the driver. |  Direct access to the MySQL protocol: "go-mysql/mysql" provides low-level access to the MySQL protocol, allowing you to communicate directly with the MySQL server. This allows for better control and optimization of interactions with MySQL, providing more efficient database access. |  |
| Mature and stable: As part of the Go language, the database/sql interface has been extensively tested and used, and is considered a relatively mature and stable solution. It has good documentation and community support to help you solve problems and learn Xi use. |  |  Advanced function support: go-mysql/mysql provides some advanced functions, such as binary protocol support, replication protocol support, etc. These features allow you to implement more complex database operations, such as subscribing to MySQL binary logs, real-time data synchronization, and more |  


If we want to use a standardized interface, support multiple databases, and are more concerned about stability and consistency, then Go's built-in database/sql interface may be a good choice. If we're more concerned about performance, compatibility, and broad community support, we might consider using "go-sql-driver/mysql". If we need more advanced features, a pure Go implementation, and the ability to access the MySQL protocol directly, then "go-mysql/mysql" may be a better fit for our needs.


# database/sql
database/sql has the following characteristics:
- Unified programming interfaces: The database/sql library provides a unified set of interfaces that allow developers to operate different databases in the same way without having to Xi the API for a specific database.
- Driver support: By importing third-party database drivers, database/sql can interact with a variety of common relational database systems, such as MySQL, PostgreSQL, SQLite, etc.
- Prevent SQL injection: The database/sql library effectively prevents SQL injection attacks by using technologies such as precompiled statements and parameterized queries.
- Support for transactions: Transactions are a must-have feature for a good SQL package.


# go-sql-driver
A popular MySQL driver that implements the database/sql interface
Compared with Go's built-in database access interface (database/sql), go-sql-driver/mysql has the following advantages:
1. MySQL Dedicated Driver: go-sql-driver/mysql is a driver specially designed for MySQL databases, so it can provide better performance and compatibility. It provides more efficient database access by communicating directly with the MySQL server, leveraging MySQL's specific features and optimizations.
2. Support extra features: go-sql-driver/mysql provides some extra features and options to better meet the needs of MySQL databases. For example, it supports features such as connection pooling, transaction isolation levels, bulk inserts, and more, which can provide more flexible and advanced database operations.
3. Update and Maintenance: go-sql-driver/mysql is an active open-source project that is widely supported and maintained by the community. It is frequently updated, bug fixes and new features are added to adapt to the latest changes and needs of MySQL databases.
4. Ecosystem support: go-sql-driver/mysql is one of the most widely used MySQL drivers in the Go language, so it has a rich ecosystem and community support. You can find many examples, documentation, and solutions in the community to help you better use and understand the driver.
Although go-sql-driver/mysql is superior to the built-in database/sql interface in many ways, in some cases it may also be appropriate to use the built-in interface. For example, if you want to keep your code generic and can easily switch to a different database if needed, then using the built-in interface may be more appropriate. Ultimately, you can choose the right database access based on your project needs and preferences.

# go-mysql/mysql
Compared with Go's built-in database access interface (database/sql), go-mysql has the following advantages:
1. Pure Go implementation: go-mysql/mysql is a MySQL client library implemented in pure Go language, which does not depend on external C libraries. This means that you can easily build and deploy your applications on different platforms without worrying about the dependencies of the C library.
2. Direct access to the MySQL protocol: go-mysql/mysql provides low-level access to the MySQL protocol, allowing you to communicate directly with the MySQL server. This allows for better control and optimization of interactions with MySQL, providing more efficient database access.
3. Advanced function support: go-mysql/mysql provides some advanced functions, such as binary protocol support, replication protocol support, etc. These features allow you to implement more complex database operations, such as subscribing to MySQL binary logs, real-time data synchronization, and more.
4. Performance and flexibility: Due to direct access to the MySQL protocol, go-mysql/mysql can provide high-performance database access. It also has a flexible API to cater to a variety of different usage scenarios and needs.
Although go-mysql/mysql is superior to the built-in database/sql interface in many ways, it also has some limitations. Since it is the underlying MySQL client library, using it requires some knowledge of the MySQL protocol and may require more code writing to handle database operations
