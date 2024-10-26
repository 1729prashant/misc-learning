## SQLite and Oracle SQL Syntax Differences

Syntax and functionality differences between Oracle SQL and SQLite for DDL, DML, and metadata queries.

| **Feature/Operation**           | **Oracle SQL**                                                                                                       | **SQLite**                                                                        |
|---------------------------------|----------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------|
| **Insert**                      | `INSERT ALL` for multiple rows; `RETURNING INTO` to capture values                                                  | `INSERT INTO` for basic insertions; no `RETURNING` clause                         |
| **Insert Example**              | `INSERT INTO employees VALUES (1, 'John');`                                                                         | `INSERT INTO employees VALUES (1, 'John');`                                       |
| **Update**                      | `UPDATE ... RETURNING INTO` to capture updated values                                                                | Basic `UPDATE`; no `RETURNING` clause                                             |
| **Update Example**              | `UPDATE employees SET salary = 5000 WHERE id = 1 RETURNING id INTO :id_var;`                                        | `UPDATE employees SET salary = 5000 WHERE id = 1;`                                |
| **Delete**                      | `DELETE ... RETURNING INTO` to capture deleted values                                                                | Basic `DELETE`; no `RETURNING` clause                                             |
| **Delete Example**              | `DELETE FROM employees WHERE id = 1 RETURNING id INTO :id_var;`                                                     | `DELETE FROM employees WHERE id = 1;`                                             |
| **Create Table**                | Complex syntax with advanced data types, constraints, `ON DELETE CASCADE` support                                   | Basic `CREATE TABLE`; `ON DELETE CASCADE` limited, fewer data types               |
| **Create Table Example**        | `CREATE TABLE emp (id NUMBER, name VARCHAR2(50));`                                                                  | `CREATE TABLE emp (id INTEGER, name TEXT);`                                       |
| **Primary Key**                 | `PRIMARY KEY` defined as a column or table constraint                                                                | `INTEGER PRIMARY KEY AUTOINCREMENT` for auto-incrementing primary key             |
| **Primary Key Example**         | `id NUMBER PRIMARY KEY`                                                                                             | `id INTEGER PRIMARY KEY AUTOINCREMENT`                                            |
| **Alter Table**                 | Supports adding/dropping columns, adding constraints, renaming tables, modifying columns                            | Limited `ALTER TABLE`; can only add columns                                      |
| **Add Column Example**          | `ALTER TABLE employees ADD (birthdate DATE);`                                                                       | `ALTER TABLE employees ADD COLUMN birthdate TEXT;`                                |
| **Fetch Metadata**              | Uses metadata views like `ALL_TABLES`, `USER_TABLES`, `DBA_TABLES`                                                  | Uses `sqlite_master` table or `PRAGMA` statements                                 |
| **Fetch Metadata Example**      | `SELECT * FROM ALL_TABLES;`                                                                                         | `SELECT name FROM sqlite_master WHERE type='table';`                              |
| **Transaction Control**         | Full transaction control, with `SAVEPOINT` and `ROLLBACK TO SAVEPOINT`                                              | Basic transaction control with `BEGIN`, `COMMIT`, `ROLLBACK`                      |
| **Start Transaction Example**   | `BEGIN;`                                                                                                            | `BEGIN TRANSACTION;`                                                              |
| **End Transaction Example**     | `COMMIT;`                                                                                                           | `COMMIT;`                                                                         |
| **Rollback Example**            | `ROLLBACK;`                                                                                                         | `ROLLBACK;`                                                                       |
| **Sequences**                   | Supported with `CREATE SEQUENCE` for generating unique numbers                                                      | Not supported; use `AUTOINCREMENT` with `INTEGER PRIMARY KEY`                     |
| **Sequence Example**            | `CREATE SEQUENCE emp_seq START WITH 1 INCREMENT BY 1;`                                                              | Not directly supported; use `INTEGER PRIMARY KEY AUTOINCREMENT`                   |
| **Data Types**                  | `VARCHAR2`, `NUMBER`, `DATE`, `TIMESTAMP`, `BLOB`, `CLOB`                                                           | `TEXT`, `INTEGER`, `REAL`, `BLOB`; all columns are `NULL` by default              |
| **Triggers**                    | Supports `BEFORE` and `AFTER` triggers, row-level and statement-level triggers                                      | Supports basic `AFTER` and `BEFORE` triggers; no statement-level triggers         |
| **User-defined Functions**      | Can create PL/SQL functions and procedures                                                                          | Limited to built-in functions; user-defined functions are available in extensions |
| **Stored Procedures**           | Full PL/SQL support for stored procedures                                                                           | Not supported directly; requires extensions or workarounds                        |
| **Joins in Update/Delete**      | Can perform complex `UPDATE` or `DELETE` statements with joins                                                      | Basic joins supported in `UPDATE`; limited functionality                          |
| **Subquery Support**            | Extensive subquery support, including correlated subqueries                                                         | Supports basic subqueries; correlated subqueries may be limited                   |

## SQLite Date and Time Formats

SQLite doesn’t have a dedicated `DATE` type; instead, it stores dates and times as **TEXT**, **REAL**, or **INTEGER**. It supports a variety of date and time formats by using these three storage classes to represent dates in different ways:

1. **TEXT**: ISO-8601 date strings (`"YYYY-MM-DD HH:MM:SS.SSS"`).
2. **REAL**: Julian day numbers.
3. **INTEGER**: Unix timestamps.

Here’s a summary of how SQLite handles these formats with example usages.

| **Storage Type** | **Format**                        | **Description**                                                         | **Example Value**             |
|------------------|-----------------------------------|-------------------------------------------------------------------------|-------------------------------|
| **TEXT**         | `YYYY-MM-DD HH:MM:SS.SSS`        | ISO-8601 format string for date and time, including optional fractions. | `'2024-10-26 15:35:30.123'`   |
|                  | `YYYY-MM-DD`                     | ISO-8601 format string for just the date.                               | `'2024-10-26'`                |
|                  | `YYYY-MM-DDTHH:MM`               | Compact ISO-8601 format, using 'T' separator instead of space.          | `'2024-10-26T15:35'`          |
| **REAL**         | Julian day                        | Julian day number, including optional fractional time of day.           | `2460367.1475694444`          |
| **INTEGER**      | Unix timestamp                   | Number of seconds since 1970-01-01 00:00:00 UTC.                        | `1724951730`                  |

SQLite provides several **date and time functions** to work with these formats, converting between them as needed. Here’s how you can work with each format.

### Examples of Using SQLite Date and Time Functions

| **Expression**                                | **Description**                                              | **Example Output**                      |
|-----------------------------------------------|--------------------------------------------------------------|-----------------------------------------|
| `DATE('now')`                                 | Returns the current date in `YYYY-MM-DD` format.             | `'2024-10-26'`                          |
| `TIME('now')`                                 | Returns the current time in `HH:MM:SS` format.               | `'15:35:30'`                            |
| `DATETIME('now')`                             | Returns the current date and time in `YYYY-MM-DD HH:MM:SS`.  | `'2024-10-26 15:35:30'`                 |
| `JULIANDAY('now')`                            | Returns the current Julian day as a floating-point number.   | `2460367.1475694444`                    |
| `STRFTIME('%Y-%m-%d', 'now')`                 | Formats the current date with custom format specifications.  | `'2024-10-26'`                          |
| `STRFTIME('%s', 'now')`                       | Returns the current Unix timestamp.                          | `1724951730`                            |
| `DATETIME(1724951730, 'unixepoch')`           | Converts Unix timestamp to `YYYY-MM-DD HH:MM:SS` format.     | `'2024-10-26 15:35:30'`                 |
| `DATE(2460367.1475694444, 'julian')`          | Converts Julian day to `YYYY-MM-DD`.                         | `'2024-10-26'`                          |
| `JULIANDAY('2024-10-26')`                     | Converts a `TEXT` date to Julian day.                        | `2460367.0`                             |
| `STRFTIME('%Y-%m-%d %H:%M:%S', 1724951730, 'unixepoch')` | Converts Unix timestamp to a formatted string. | `'2024-10-26 15:35:30'`                 |

### Permutations with `STRFTIME` Function

The `STRFTIME` function in SQLite can be used to create custom date and time formats. Below are some examples of common permutations:

| **Expression**                                | **Description**                                 | **Example Output**      |
|-----------------------------------------------|-------------------------------------------------|-------------------------|
| `STRFTIME('%Y-%m-%d', 'now')`                 | Year, month, day                                | `'2024-10-26'`          |
| `STRFTIME('%d-%m-%Y', 'now')`                 | Day, month, year                                | `'26-10-2024'`          |
| `STRFTIME('%m/%d/%Y', 'now')`                 | Month, day, year                                | `'10/26/2024'`          |
| `STRFTIME('%H:%M', 'now')`                    | Hour and minute                                 | `'15:35'`               |
| `STRFTIME('%Y-%m-%d %H:%M:%S', 'now')`        | Full date and time                              | `'2024-10-26 15:35:30'` |
| `STRFTIME('%Y', 'now')`                       | Year only                                       | `'2024'`                |
| `STRFTIME('%m', 'now')`                       | Month only                                      | `'10'`                  |
| `STRFTIME('%d', 'now')`                       | Day only                                        | `'26'`                  |
| `STRFTIME('%H:%M:%S', 'now')`                 | Time in hours, minutes, seconds                 | `'15:35:30'`            |


## Useful Built-in Functions in SQLite

Here’s a rundown of some of the most useful built-in functions in SQLite, categorized by their primary function type:

### 1. String Functions

| **Function**                 | **Description**                                                                           | **Example**                                 | **Result**                   |
|------------------------------|-------------------------------------------------------------------------------------------|---------------------------------------------|-------------------------------|
| `LENGTH(X)`                  | Returns the number of characters in a string.                                             | `LENGTH('SQLite')`                          | `6`                           |
| `UPPER(X)`                   | Converts all characters to uppercase.                                                     | `UPPER('sqlite')`                           | `'SQLITE'`                    |
| `LOWER(X)`                   | Converts all characters to lowercase.                                                     | `LOWER('SQLite')`                           | `'sqlite'`                    |
| `TRIM(X)`                    | Removes spaces from both ends of the string.                                              | `TRIM('  SQLite  ')`                        | `'SQLite'`                    |
| `SUBSTR(X, Y, Z)`            | Extracts a substring of `X` starting at `Y` with `Z` length (or until the end if omitted).| `SUBSTR('SQLite', 1, 3)`                    | `'SQL'`                       |
| `REPLACE(X, Y, Z)`           | Replaces occurrences of `Y` with `Z` in `X`.                                              | `REPLACE('SQLite', 'i', 'y')`               | `'Sqlyte'`                    |
| `INSTR(X, Y)`                | Returns the index of the first occurrence of `Y` in `X`.                                  | `INSTR('SQLite', 'l')`                      | `3`                           |

### 2. Date and Time Functions

| **Function**                     | **Description**                                                                             | **Example**                                  | **Result**                   |
|----------------------------------|---------------------------------------------------------------------------------------------|----------------------------------------------|-------------------------------|
| `DATE('now')`                    | Returns the current date.                                                                   | `DATE('now')`                                | `'2024-10-26'`               |
| `TIME('now')`                    | Returns the current time.                                                                   | `TIME('now')`                                | `'15:35:30'`                 |
| `DATETIME('now')`                | Returns the current date and time.                                                          | `DATETIME('now')`                            | `'2024-10-26 15:35:30'`      |
| `JULIANDAY('now')`               | Returns the Julian day number for the current date.                                         | `JULIANDAY('now')`                           | `2460367.1475694444`         |
| `STRFTIME(format, 'now')`        | Returns a formatted date string based on the specified format.                              | `STRFTIME('%Y-%m-%d', 'now')`                | `'2024-10-26'`               |
| `DATE('now', '+7 days')`         | Adds or subtracts time intervals to/from the current date.                                  | `DATE('now', '+7 days')`                     | `'2024-11-02'`               |

### 3. Aggregate Functions

| **Function**                     | **Description**                                                                             | **Example**                                  | **Result**                   |
|----------------------------------|---------------------------------------------------------------------------------------------|----------------------------------------------|-------------------------------|
| `COUNT(X)`                       | Returns the number of non-null values in a column.                                          | `COUNT(id)`                                  | `10` (depending on data)      |
| `SUM(X)`                         | Returns the sum of all non-null values in a column.                                         | `SUM(price)`                                 | Sum of `price` column         |
| `AVG(X)`                         | Returns the average of all non-null values in a column.                                     | `AVG(age)`                                   | Average age                   |
| `MIN(X)`                         | Returns the minimum value in a column.                                                      | `MIN(salary)`                                | Minimum salary                |
| `MAX(X)`                         | Returns the maximum value in a column.                                                      | `MAX(salary)`                                | Maximum salary                |
| `GROUP_CONCAT(X, Y)`             | Concatenates all values in a column, separated by `Y`.                                      | `GROUP_CONCAT(name, ', ')`                   | `'Alice, Bob, Carol'`         |

### 4. Mathematical Functions

| **Function**                     | **Description**                                                                             | **Example**                                  | **Result**                   |
|----------------------------------|---------------------------------------------------------------------------------------------|----------------------------------------------|-------------------------------|
| `ABS(X)`                         | Returns the absolute value of `X`.                                                          | `ABS(-5)`                                    | `5`                           |
| `ROUND(X, Y)`                    | Rounds `X` to `Y` decimal places.                                                           | `ROUND(3.14159, 2)`                          | `3.14`                        |
| `CEIL(X)`                        | Returns the smallest integer greater than or equal to `X`.                                  | `CEIL(3.5)`                                  | `4`                           |
| `FLOOR(X)`                       | Returns the largest integer less than or equal to `X`.                                     | `FLOOR(3.5)`                                 | `3`                           |
| `RANDOM()`                       | Returns a random integer between -2^63 and 2^63 - 1.                                       | `RANDOM()`                                   | Random integer                |

### 5. Conditional Expressions

| **Function**                     | **Description**                                                                             | **Example**                                  | **Result**                   |
|----------------------------------|---------------------------------------------------------------------------------------------|----------------------------------------------|-------------------------------|
| `IFNULL(X, Y)`                   | Returns `Y` if `X` is `NULL`; otherwise, returns `X`.                                       | `IFNULL(NULL, 'default')`                    | `'default'`                   |
| `NULLIF(X, Y)`                   | Returns `NULL` if `X` equals `Y`; otherwise, returns `X`.                                   | `NULLIF(5, 5)`                               | `NULL`                        |
| `COALESCE(X, Y, ...)`            | Returns the first non-null value in the argument list.                                      | `COALESCE(NULL, 'hello', 'world')`           | `'hello'`                     |
| `CASE`                           | Allows for conditional logic in queries.                                                    | `CASE WHEN age > 18 THEN 'adult' ELSE 'minor' END` | `'adult'` or `'minor'` |

