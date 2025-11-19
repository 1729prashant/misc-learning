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


## SQLite Metadata and Introspection Capabilities

| **#** | **Metadata Category**          | **SQLite Support** | **SQL / PRAGMA with Notes**                                                                                     |
|------|--------------------------------|--------------------|------------------------------------------------------------------------------------------------------------------|
| 1    | All Tables                     | ✅ Yes             | -- Notes: User-defined tables only<br>`SELECT name FROM sqlite_master WHERE type = 'table';`                     |
|      | All Views                      | ✅ Yes             | -- Notes: Views stored alongside tables<br>`SELECT name FROM sqlite_master WHERE type = 'view';`                 |
|      | Temp Tables and Views          | ✅ Yes             | -- Notes: Only visible during session<br>`SELECT name FROM sqlite_temp_master WHERE type IN ('table', 'view');`  |
| 2    | Table/View DDL                 | ✅ Yes             | -- Notes: DDL is stored in `sql` column<br>`SELECT type, name, tbl_name, sql FROM sqlite_master WHERE type IN ('table', 'view');` |
| 3    | DML Activity                   | ❌ No              | -- Notes: No DML audit log; use triggers manually if needed                                                      |
| 4    | Index Metadata                 | ✅ Yes             | -- Notes: Lists indexes defined on any table<br>`SELECT name, tbl_name, sql FROM sqlite_master WHERE type = 'index';` |
|      | Index Columns                  | ✅ Yes             | -- Notes: Shows indexed column names and positions<br>`PRAGMA index_info('<index_name>');`                       |
| 5    | Triggers                       | ✅ Yes             | -- Notes: All trigger definitions are introspectable<br>`SELECT name, tbl_name, sql FROM sqlite_master WHERE type = 'trigger';` |
| 6    | Functions (UDFs)               | ❌ No              | -- Notes: Defined only in host language, not visible in SQLite                                                  |
| 7    | Permissions / Grants / Roles  | ❌ No              | -- Notes: No built-in auth model; relies on file permissions                                                    |
| 8    | Constraints: Primary Key       | ✅ Yes             | -- Notes: `pk` column indicates primary key<br>`PRAGMA table_info('<table_name>');`                              |
|      | Constraints: Unique            | ✅ Yes             | -- Notes: Usually declared inline or via unique indexes<br>`SELECT * FROM sqlite_master WHERE sql LIKE '%UNIQUE%';` |
|      | Constraints: Foreign Key       | ✅ Yes             | -- Notes: Shows table, from/to columns, and ON DELETE/UPDATE<br>`PRAGMA foreign_key_list('<table_name>');`       |
|      | Constraints: Check             | ⚠️ Partial         | -- Notes: No structured catalog view, must parse SQL<br>Inspect `sql` in `sqlite_master` manually                |
| 9    | Busy Timeout                   | ✅ Yes             | -- Notes: Set/get time to wait before failing on locked DB (in ms)<br>`PRAGMA busy_timeout;` or `PRAGMA busy_timeout = 5000;` |
| 10   | Attached Databases             | ✅ Yes             | -- Notes: Shows logical name, file path, and origin<br>`PRAGMA database_list;`                                   |
| 11   | Page Size (Storage Unit)       | ✅ Yes             | -- Notes: Shows/sets size (in bytes) of each DB page<br>`PRAGMA page_size;` or `PRAGMA page_size = 4096;`        |
|      | File Size (Pages)              | ✅ Yes             | -- Notes: Multiply `page_size * page_count` for actual size<br>`PRAGMA page_count;`                              |
|      | Auto-Vacuum Status             | ✅ Yes             | -- Notes: Values: 0 (none), 1 (full), 2 (incremental)<br>`PRAGMA auto_vacuum;`                                   |
| 12   | Locking Mode                   | ✅ Yes             | -- Notes: Values: NORMAL (default), EXCLUSIVE<br>`PRAGMA locking_mode;` or `PRAGMA locking_mode = EXCLUSIVE;`    |
|      | Current Lock Status            | ⚠️ Partial         | -- Notes: No direct SQL interface<br>Use `.locks` in the sqlite3 shell                                           |
| 13   | Journaling Mode                | ✅ Yes             | -- Notes: Modes: DELETE, TRUNCATE, PERSIST, MEMORY, WAL, OFF<br>`PRAGMA journal_mode;` or `PRAGMA journal_mode = WAL;` |
|      | Journal Size Limit             | ✅ Yes             | -- Notes: Sets max journal file size<br>`PRAGMA journal_size_limit;`                                             |
| 14   | WAL Checkpoint Info            | ✅ Yes             | -- Notes: Runs a manual checkpoint<br>`PRAGMA wal_checkpoint(FULL);`                                             |
|      | WAL Status Metrics             | ✅ Yes             | -- Notes: Controls/monitors WAL size & frequency<br>`PRAGMA wal_checkpoint(TRUNCATE);` + `PRAGMA wal_autocheckpoint;` |
| 15   | Temp File Storage              | ✅ Yes             | -- Notes: 0 = default, 1 = file, 2 = memory<br>`PRAGMA temp_store;`                                               |
|      | Temp Store Directory           | ✅ Yes             | -- Notes: Deprecated; requires legacy compile option<br>`PRAGMA temp_store_directory;`                           |
| 16   | Cache Size                     | ✅ Yes             | -- Notes: Negative = KB units, Positive = pages<br>`PRAGMA cache_size;`                                          |
|      | Cache Spilling                 | ✅ Yes             | -- Notes: Controls whether cache overflows to temp files<br>`PRAGMA cache_spill;`                                |
| 17   | Database Integrity Check       | ✅ Yes             | -- Notes: Checks for schema and data corruption<br>`PRAGMA integrity_check;` or `PRAGMA quick_check;`            |
| 18   | File Path and Name             | ✅ Yes             | -- Notes: Returns list of attached DBs and file paths<br>`PRAGMA database_list;`                                 |
|      | Journal File Path              | ⚠️ Indirect        | -- Notes: Named by convention (e.g., `mydb.db-journal`, `-wal`, `-shm`)<br>Use filesystem inspection             |
| 19   | Application ID / User Version  | ✅ Yes             | -- Notes: Used by apps to tag DBs or store schema version<br>`PRAGMA application_id;`, `PRAGMA user_version;`    |



## SQLite Sort DD-MON-YYYY text columns

To order by a date stored in DD-MON-YYYY format, convert it to a sortable format. How to order by a date in DD-MON-YYYY format:
```
SELECT
  your_column_name,
  STRFTIME('%Y-%m-%d', SUBSTR(your_column_name, 8, 4) || '-' ||
                        CASE SUBSTR(your_column_name, 4, 3)
                          WHEN 'JAN' THEN '01'
                          WHEN 'FEB' THEN '02'
                          WHEN 'MAR' THEN '03'
                          WHEN 'APR' THEN '04'
                          WHEN 'MAY' THEN '05'
                          WHEN 'JUN' THEN '06'
                          WHEN 'JUL' THEN '07'
                          WHEN 'AUG' THEN '08'
                          WHEN 'SEP' THEN '09'
                          WHEN 'OCT' THEN '10'
                          WHEN 'NOV' THEN '11'
                          WHEN 'DEC' THEN '12'
                        END || '-' ||
                        SUBSTR(your_column_name, 1, 2)) AS sortable_date
FROM
  your_table_name
ORDER BY
  sortable_date;
```
