Задача 1
===
Используя docker поднимите инстанс MySQL (версию 8). Данные БД сохраните в volume.

Изучите бэкап БД и восстановитесь из него.

Перейдите в управляющую консоль `mysql` внутри контейнера.

Используя команду `\h` получите список управляющих команд.

Найдите команду для выдачи статуса БД и **приведите в ответе** из ее вывода версию сервера БД.

Подключитесь к восстановленной БД и получите список таблиц из этой БД.

**Приведите в ответе** количество записей с `price` > 300.

В следующих заданиях мы будем продолжать работу с данным контейнером.

Решение
---
Версия сервера БД:

```
mysql> \s
--------------
mysql  Ver 8.0.27 for Linux on x86_64 (MySQL Community Server - GPL)

Connection id:		12
Current database:	
Current user:		root@localhost
SSL:			Not in use
Current pager:		stdout
Using outfile:		''
Using delimiter:	;
Server version:		8.0.27 MySQL Community Server - GPL
Protocol version:	10
Connection:		Localhost via UNIX socket
Server characterset:	utf8mb4
Db     characterset:	utf8mb4
Client characterset:	latin1
Conn.  characterset:	latin1
UNIX socket:		/var/run/mysqld/mysqld.sock
Binary data as:		Hexadecimal
Uptime:			25 min 12 sec

Threads: 2  Questions: 43  Slow queries: 0  Opens: 156  Flush tables: 3  Open tables: 74  Queries per second avg: 0.028
--------------
```

Количество записей с price > 300:

```
mysql> select count(*) from orders where price > 300;
+----------+
| count(*) |
+----------+
|        1 |
+----------+
1 row in set (0.00 sec)
```

Задача 2
===
Создайте пользователя test в БД c паролем test-pass, используя:

* плагин авторизации mysql_native_password
* срок истечения пароля - 180 дней
* количество попыток авторизации - 3
* максимальное количество запросов в час - 100
* аттрибуты пользователя:
    * Фамилия "Pretty"
    * Имя "James"

Предоставьте привелегии пользователю `test` на операции SELECT базы `test_db`.

Используя таблицу INFORMATION_SCHEMA.USER_ATTRIBUTES получите данные по пользователю `test` и **приведите в ответе к задаче**.

Решение
---
Создание пользователя и изменение его настроек и атрибутов:

```
CREATE USER 'test'
IDENTIFIED WITH mysql_native_password BY 'test-pass';
```
```
ALTER USER 'test'
ATTRIBUTE '{"name":"James", "sname":"Pretty"}';
```
```
ALTER USER 'test'
WITH
MAX_QUERIES_PER_HOUR 100
PASSWORD EXPIRE INTERVAL 180 DAY
FAILED_LOGIN_ATTEMPTS 3;
```
```
GRANT SELECT ON test_db.orders TO 'test';
```
Данные по пользователю `test` из INFORMATION_SCHEMA.USER_ATTRIBUTES:
```
mysql> select * from INFORMATION_SCHEMA.USER_ATTRIBUTES where user='test';
+------+------+--------------------------------------+
| USER | HOST | ATTRIBUTE                            |
+------+------+--------------------------------------+
| test | %    | {"name": "James", "sname": "Pretty"} |
+------+------+--------------------------------------+
1 row in set (0.00 sec)
```

Задача 3
===
Установите профилирование `SET profiling = 1`. Изучите вывод профилирования команд
`SHOW PROFILES;`.

Исследуйте, какой `engine` используется в таблице БД `test_db` и **приведите в ответе**.

Измените `engine` и **приведите время выполнения и запрос на изменения из профайлера 
в ответе**:

* на `MyISAM`
* на `InnoDB`

Решение
---
```
mysql> SELECT TABLE_SCHEMA, TABLE_NAME, ENGINE FROM information_schema.TABLES WHERE TABLE_SCHEMA='test_db';
+--------------+------------+--------+
| TABLE_SCHEMA | TABLE_NAME | ENGINE |
+--------------+------------+--------+
| test_db      | orders     | InnoDB |
+--------------+------------+--------+
1 row in set (0.00 sec)
```
MyISAM:
```
|       12 | 0.01910000 | ALTER TABLE orders ENGINE = MyISAM
```
InnoDB:
```
|       13 | 0.02235500 | ALTER TABLE orders ENGINE = InnoDB
```

Задача 4
===
Изучите файл `my.cnf` в директории /etc/mysql.

Измените его согласно ТЗ (движок InnoDB):

* Скорость IO важнее сохранности данных
* Нужна компрессия таблиц для экономии места на диске
* Размер буффера с незакомиченными транзакциями 1 Мб
* Буффер кеширования 30% от ОЗУ
* Размер файла логов операций 100 Мб

Приведите в ответе измененный файл `my.cnf`.

Решение
---
```
[mysqld]
pid-file        = /var/run/mysqld/mysqld.pid
socket          = /var/run/mysqld/mysqld.sock
datadir         = /var/lib/mysql
secure-file-priv= NULL

#Внесенные изменения

innodb_flush_log_at_trx_commit = 0

innodb_file_format=Barracuda
#Вроде бы этот формат включен по умолчанию...

innodb_log_buffer_size = 1M

innodb_buffer_pool_size = 307M
#При 1 Гб оперативной памяти

innodb_log_file_size = 100M

# Custom config should go here
!includedir /etc/mysql/conf.d/
```