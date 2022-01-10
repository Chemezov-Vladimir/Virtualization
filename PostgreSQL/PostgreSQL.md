Задача 1
=
Используя docker поднимите инстанс PostgreSQL (версию 12) c 2 volume, в который 
будут складываться данные БД и бэкапы.

Приведите получившуюся команду или docker-compose манифест.

Решение
-
Команда Docker:
`docker run --name my_postgres -v /hw/postgres-bu:/postgres-bu -v 
/hw/postgres-db:/postgres-db -p 5432:5432 -e POSTGRES_PASSWORD=12345678 -itd
postgres:12-alpine`

Задача 2
=
В БД из задачи 1:

* создайте пользователя test-admin-user и БД test_db
* в БД test_db создайте таблицу orders и clients (спeцификация таблиц ниже)
* предоставьте привилегии на все операции пользователю test-admin-user на таблицы 
БД test_db
* создайте пользователя test-simple-user
* предоставьте пользователю test-simple-user права на SELECT/INSERT/UPDATE/DELETE
данных таблиц БД test_db

Таблица orders:

* id (serial primary key)
* наименование (string)
* цена (integer)

Таблица clients:

* id (serial primary key)
* фамилия (string)
* страна проживания (string, index)
* заказ (foreign key orders)

Приведите:
* итоговый список БД после выполнения пунктов выше;
* описание таблиц (describe);
* SQL-запрос для выдачи списка пользователей с правами над таблицами test_db;
* список пользователей с правами над таблицами test_db.

Решение
-

Список БД:
```
test_db=# \l
                                 List of databases
   Name    |  Owner   | Encoding |  Collate   |   Ctype    |   Access privileges   
-----------+----------+----------+------------+------------+-----------------------
 postgres  | postgres | UTF8     | en_US.utf8 | en_US.utf8 | 
 template0 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +
           |          |          |            |            | postgres=CTc/postgres
 template1 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +
           |          |          |            |            | postgres=CTc/postgres
 test_db   | postgres | UTF8     | en_US.utf8 | en_US.utf8 | 
(4 rows)
```

Описание таблиц:
```
test_db=# \d orders
                            Table "public.orders"
 Column |  Type   | Collation | Nullable |              Default               
--------+---------+-----------+----------+------------------------------------
 id     | integer |           | not null | nextval('orders_id_seq'::regclass)
 name   | text    |           |          | 
 price  | integer |           |          | 
Indexes:
    "orders_pkey" PRIMARY KEY, btree (id)
Referenced by:
    TABLE "clients" CONSTRAINT "clients_ordered_fkey" FOREIGN KEY (ordered) REFERENCES orders(id)
```
```
test_db=# \d clients
                             Table "public.clients"
 Column  |  Type   | Collation | Nullable |               Default               
---------+---------+-----------+----------+-------------------------------------
 id      | integer |           | not null | nextval('clients_id_seq'::regclass)
 surname | text    |           |          | 
 country | text    |           |          | 
 ordered | integer |           |          | 
Indexes:
    "clients_pkey" PRIMARY KEY, btree (id)
Foreign-key constraints:
    "clients_ordered_fkey" FOREIGN KEY (ordered) REFERENCES orders(id)
```
Список пользователей с правами:
```
test_db=# SELECT grantee, table_name, privilege_type
test_db-# FROM information_schema.table_privileges
test_db-# WHERE grantee IN ('test-admin-user','test-simple-user');
     grantee      | table_name | privilege_type 
------------------+------------+----------------
 test-admin-user  | orders     | INSERT
 test-admin-user  | orders     | SELECT
 test-admin-user  | orders     | UPDATE
 test-admin-user  | orders     | DELETE
 test-admin-user  | orders     | TRUNCATE
 test-admin-user  | orders     | REFERENCES
 test-admin-user  | orders     | TRIGGER
 test-simple-user | orders     | INSERT
 test-simple-user | orders     | SELECT
 test-simple-user | orders     | UPDATE
 test-simple-user | orders     | DELETE
 test-admin-user  | clients    | INSERT
 test-admin-user  | clients    | SELECT
 test-admin-user  | clients    | UPDATE
 test-admin-user  | clients    | DELETE
 test-admin-user  | clients    | TRUNCATE
 test-admin-user  | clients    | REFERENCES
 test-admin-user  | clients    | TRIGGER
 test-simple-user | clients    | INSERT
 test-simple-user | clients    | SELECT
 test-simple-user | clients    | UPDATE
 test-simple-user | clients    | DELETE
(22 rows)
```
Задача 3
=
Используя SQL синтаксис - наполните таблицы следующими тестовыми данными:

Таблица orders

| Наименование | цена |
| ----------|--------|
| Шоколад |	10 |
| Принтер |	3000 |
| Книга |	500 |
| Монитор |	7000 |
| Гитара |	4000 |

Таблица clients

| ФИО |	Страна проживания |
|--------|------|
| Иванов Иван Иванович |	USA |
| Петров Петр Петрович |	Canada |
| Иоганн Себастьян Бах |	Japan |
| Ронни Джеймс Дио |	Russia |
| Ritchie Blackmore | Russia |

Используя SQL синтаксис:

* вычислите количество записей для каждой таблицы
* приведите в ответе:
  * запросы
  * результаты их выполнения.

Решение
-
Заполнил таблицы следующими командами:

`insert into orders values (1, 'Шоколад', 10), (2, 'Принтер', 3000), (3, 'Книга', 
500), (4, 'Монитор', 7000), (5, 'Гитара', 4000);`

`insert into clients values (1, 'Иванов Иван Иванович', 'USA'), (2, 'Петров Петр 
Петрович', 'Canada'), (3, 'Иоганн Себастьян Бах', 'Japan'), (4, 'Ронни Джеймс Дио',
'Russia'), (5, 'Ritchie Blackmore', 'Russia');`

Запросы на количество записей:
```
test_db=# select count(*) from clients;
 count 
-------
     5
(1 row)
```
```
test_db=# select count(*) from orders;
 count 
-------
     5
(1 row)
```
Задача 4
=
Часть пользователей из таблицы clients решили оформить заказы из таблицы orders.

Используя foreign keys свяжите записи из таблиц, согласно таблице:

| ФИО |	Заказ |
|-------|-------|
| Иванов Иван Иванович |	Книга |
| Петров Петр Петрович |	Монитор |
| Иоганн Себастьян Бах |	Гитара |

Приведите SQL-запросы для выполнения данных операций.

Приведите SQL-запрос для выдачи всех пользователей, которые совершили заказ, а 
также вывод данного запроса.

Подсказк - используйте директиву `UPDATE`.

Решение
-
`update clients set ordered = 3 where id = 1;`

`update clients set ordered = 4 where id = 2;`

`update clients set ordered = 5 where id = 3;`

Запрос и вывод:
```
test_db=# select c.surname, o.n_name from clients c inner join orders o on (c.ordered=o.id);
       surname        | n_name  
----------------------+---------
 Иванов Иван Иванович | Книга
 Петров Петр Петрович | Монитор
 Иоганн Себастьян Бах | Гитара
(3 rows)
```
Задача 5
=
Получите полную информацию по выполнению запроса выдачи всех пользователей из задачи 
4 (используя директиву EXPLAIN).

Приведите получившийся результат и объясните что значат полученные значения.

Решение
-
```
test_db=# explain select c.surname, o.n_name from clients c inner join orders o on (c.ordered=o.id);
                               QUERY PLAN                                
-------------------------------------------------------------------------
 Hash Join  (cost=37.00..57.24 rows=810 width=64)
   Hash Cond: (c.ordered = o.id)
   ->  Seq Scan on clients c  (cost=0.00..18.10 rows=810 width=36)
   ->  Hash  (cost=22.00..22.00 rows=1200 width=36)
         ->  Seq Scan on orders o  (cost=0.00..22.00 rows=1200 width=36)
(5 rows)
```
Планировщик выбирает соединение по хешу, при котором строки одной таблицы записываются
в хеш-таблицу в памяти, после чего сканируется другая таблица, и для каждой её строки
проверяется соответствие по хеш-таблице.

Первое значение cost означает приблизительную "стоимость" запуска, т. е. это время,
которое проходит, прежде чем начнётся этап вывода данных.

Второе значение cost - приблизительная общая "стоимость". Она вычисляется в 
предположении, что узел плана выполняется до конца, т. е. возвращает все доступные 
строки.

rows - ожидаемое число строк, которое должен вывести этот узел плана. При этом так
же предполагается, что узел выполняется до конца. На практике чтение строк может 
прерваться досрочно. 

width - ожидаемый средний размер строк, выводимых этим узлом плана (в байтах).

Задача 6
=
Создайте бэкап БД test_db и поместите его в volume, предназначенный для бэкапов 
(см. Задачу 1).

Остановите контейнер с PostgreSQL (но не удаляйте volumes).

Поднимите новый пустой контейнер с PostgreSQL.

Восстановите БД test_db в новом контейнере.

Приведите список операций, который вы применяли для бэкапа данных и восстановления.

Решение
-
Для бэкапа подключаемся к контейнеру и выполняем команду `pg_dump test_db > 
postgres-bu/bu-1 -U postgres`. 

Останавливаем старый контейнер `docker stop my_postgres`и запускам новый `docker run 
--name my_postgres2 -v ~/hw/postgres-bu:/postgres-bu -v ~/hw/postgres-db:/postgres-db
-p 5432:5432 -e POSTGRES_PASSWORD=12345678 -itd postgres:12-alpine`. 

Подключаемся ко второму контейнеру `docker exec -it my_postgres2 /bin/bash`.

Подключаемся к дефолтной базе `psql -U postgres`.

Создаем базу `create database test_db;`. 

Выходим из текущей базы и подключаемся к вновь созданной `psql -b test_db -U postgres`. 

Создаем пользователей `create role "test-admin-user";`,
`create role "test-simple-user";`.

Выходим из базы и выполняем команду `psql test_db < postgres-bu/bu-1 -U postgres`. 

База восстановлена.