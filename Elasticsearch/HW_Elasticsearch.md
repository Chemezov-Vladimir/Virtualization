Задача 1
=

В этом задании вы потренируетесь в:

* установке elasticsearch
* первоначальном конфигурировании elastcisearch
* запуске elasticsearch в docker

Используя докер образ elasticsearch:7 как базовый:

* составьте Dockerfile-манифест для elasticsearch
* соберите docker-образ и сделайте `push` в ваш docker.io репозиторий
* запустите контейнер из получившегося образа и выполните запрос пути `/` c хост-машины

Требования к `elasticsearch.yml`:

* данные path должны сохраняться в `/var/lib`
* имя ноды должно быть `netology_test`

В ответе приведите:

* текст Dockerfile манифеста
* ссылку на образ в репозитории dockerhub
* ответ `elasticsearch` на запрос пути `/` в json виде

Подсказки:

* при сетевых проблемах внимательно изучите кластерные и сетевые настройки в elasticsearch.yml
* при некоторых проблемах вам поможет docker директива ulimit
* elasticsearch в логах обычно описывает проблему и пути ее решения
* обратите внимание на настройки безопасности такие как `xpack.security.enabled`
* если докер образ не запускается и падает с ошибкой 137 в этом случае может помочь настройка `-e ES_HEAP_SIZE`
* при настройке `path` возможно потребуется настройка прав доступа на директорию

Далее мы будем работать с данным экземпляром elasticsearch.

Решение
-

Dockerfile:
```
FROM elasticsearch:7.17.3
LABEL elastic search
EXPOSE 9200:9200
EXPOSE 9300:9300
ADD elasticsearch.yml /usr/share/elasticsearch/config
RUN mkdir /var/lib/data && \
    mkdir /var/lib/logs && \
    chmod 777 /var/lib/data && \
    chmod 777 /var/lib/logs
```
Ссылка на образ: https://hub.docker.com/r/vladimirchemezov/elastic

Ответ elasticsearch:
```
{
  "name" : "00de11279896",
  "cluster_name" : "netology_test",
  "cluster_uuid" : "117dEK6QTZadzNicxjoYnA",
  "version" : {
    "number" : "7.17.3",
    "build_flavor" : "default",
    "build_type" : "docker",
    "build_hash" : "5ad023604c8d7416c9eb6c0eadb62b14e766caff",
    "build_date" : "2022-04-19T08:11:19.070913226Z",
    "build_snapshot" : false,
    "lucene_version" : "8.11.1",
    "minimum_wire_compatibility_version" : "6.8.0",
    "minimum_index_compatibility_version" : "6.0.0-beta1"
  },
  "tagline" : "You Know, for Search"
}
```

Задача 2
=

В этом задании вы научитесь:

* создавать и удалять индексы
* изучать состояние кластера
* обосновывать причину деградации доступности данных

Ознакомтесь с документацией и добавьте в `elasticsearch` 3 индекса, в соответствии со таблицей:

| Имя |	Количество реплик | Количество шард |
| ----------|--------|-------|
| ind-1 | 0 | 1 |
| ind-2 | 1 | 2 |
| ind-3 | 2 | 4 |

Получите список индексов и их статусов, используя API и **приведите в ответе** на задание.

Получите состояние кластера `elasticsearch`, используя API.

Как вы думаете, почему часть индексов и кластер находится в состоянии yellow?

Удалите все индексы.

**Важно**

При проектировании кластера elasticsearch нужно корректно рассчитывать количество реплик и шард, иначе возможна потеря данных индексов, вплоть до полной, при деградации системы.

Решение
-

Список индексов и их статусов:
```
health status index            uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green  open   .geoip_databases Ifg-DnxjQ9qT1UF-2F8JRA   1   0         41            0     38.7mb         38.7mb
green  open   ind-1            WVk2y2nVQ0GLH-hshE89wg   1   0          0            0       226b           226b
yellow open   ind-3            TtDX31zZRhypbMOYHuOZXg   4   2          0            0       904b           904b
yellow open   ind-2            s-x1P4qjS_SLQaQQMC63zw   2   1          0            0       452b           452b
```

Состояние кластера:
```
{
  "cluster_name" : "netology_test",
  "status" : "yellow",
  "timed_out" : false,
  "number_of_nodes" : 1,
  "number_of_data_nodes" : 1,
  "active_primary_shards" : 10,
  "active_shards" : 10,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 10,
  "delayed_unassigned_shards" : 0,
  "number_of_pending_tasks" : 0,
  "number_of_in_flight_fetch" : 0,
  "task_max_waiting_in_queue_millis" : 0,
  "active_shards_percent_as_number" : 50.0
}
```
Статус индексов и кластера yellow потому, что были заданы реплики для индексов, но реплицировать их некуда, так как нет других серверов.

Задача 3
=

В данном задании вы научитесь:

* создавать бэкапы данных
* восстанавливать индексы из бэкапов

Создайте директорию `{путь до корневой директории с elasticsearch в образе}/snapshots`.

Используя API зарегистрируйте данную директорию как `snapshot repository` c именем `netology_backup`.

**Приведите в ответе** запрос API и результат вызова API для создания репозитория.

Создайте индекс `test` с 0 реплик и 1 шардом и **приведите в ответе** список индексов.

Создайте `snapshot` состояния кластера `elasticsearch`.

**Приведите в ответе** список файлов в директории со `snapshot`ами.

Удалите индекс `test` и создайте индекс `test-2`. **Приведите в ответе** список индексов.

Восстановите состояние кластера `elasticsearch` из `snapshot`, созданного ранее.

**Приведите в ответе** запрос к API восстановления и итоговый список индексов.

Подсказки:

* возможно вам понадобится доработать `elasticsearch.yml` в части директивы `path.repo` и перезапустить `elasticsearch`

Решение
-
Была добавлена следующая строка в `elasticsearch.yml`:
```
path.repo: /usr/share/elasticsearch/snapshots
```
Запрос на регистрацию репозитория: 
```
curl -X PUT "localhost:9200/_snapshot/netology_backup?pretty" -H 'Content-Type: application/json' -d'
{
  "type": "fs",
  "settings": {
    "location": "snapshots"
  }
}
'
```
Вывод:
```
{
  "acknowledged" : true
}
```
Первый список индексов:
```
health status index            uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green  open   .geoip_databases UezzioBPRyWAffl2C8USqA   1   0         41            0     38.7mb         38.7mb
green  open   test             QCBgqUKKSdyAn413BTVz2g   1   0          0            0       226b           226b
```
Список файлов в директории со снапшотами:
```
root@2927cbe14b1b:/usr/share/elasticsearch/snapshots/snapshots# ls -la
total 56
drwxrwxr-x 3 elasticsearch root  4096 May 26 15:00 .
drwxrwxr-x 3 elasticsearch root  4096 May 26 14:28 ..
-rw-rw-r-- 1 elasticsearch root  1425 May 26 15:00 index-0
-rw-rw-r-- 1 elasticsearch root     8 May 26 15:00 index.latest
drwxrwxr-x 6 elasticsearch root  4096 May 26 15:00 indices
-rw-rw-r-- 1 elasticsearch root 29222 May 26 15:00 meta-gCAOr9wYSfS1U5Vs3r8N8A.dat
-rw-rw-r-- 1 elasticsearch root   712 May 26 15:00 snap-gCAOr9wYSfS1U5Vs3r8N8A.dat

```
Второй список индексов:
```
health status index            uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green  open   .geoip_databases UezzioBPRyWAffl2C8USqA   1   0         41            0     38.7mb         38.7mb
green  open   test-2           U6PlB7rBSQWeXvIvwBDFDw   1   0          0            0       226b           226b
```
Запросы для восстановления состояния кластера из снапшота:
1. Остановка следующих служб - 
* GeoIP database downloader
```
curl -X PUT "localhost:9200/_cluster/settings?pretty" -H 'Content-Type: application/json' -d'
{
  "persistent": {
    "ingest.geoip.downloader.enabled": false
  }
}
'
```
* ILM
```
curl -X POST "localhost:9200/_ilm/stop?pretty"
```
* Machine Learning
```
curl -X POST "localhost:9200/_ml/set_upgrade_mode?enabled=true&pretty"
```
* Monitoring
```
curl -X PUT "localhost:9200/_cluster/settings?pretty" -H 'Content-Type: application/json' -d'
{
  "persistent": {
    "xpack.monitoring.collection.enabled": false
  }
}
'
```
* Watcher
```
curl -X POST "localhost:9200/_watcher/_stop?pretty"
```
2. Настройка API для удаления потоков данных и индексов -
```
curl -X PUT "localhost:9200/_cluster/settings?pretty" -H 'Content-Type: application/json' -d'
{
  "persistent": {
    "action.destructive_requires_name": false
  }
}
'
```
3. Удаление потоков данных - 
```
curl -X DELETE "localhost:9200/_data_stream/*?expand_wildcards=all&pretty"
```
4. Удаление индексов -
```
curl -X DELETE "localhost:9200/*?expand_wildcards=all&pretty"
```
5. Восстановление из снапшота - 
```
curl -X POST "localhost:9200/_snapshot/netology_backup/elasticsearch/_restore?pretty" -H 'Content-Type: application/json' -d'
{
  "indices": "*",
  "include_global_state": true
}
'
```
Результат:
```
{
  "accepted" : true
}
```
Итоговый список индексов:
```
health status index            uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green  open   .geoip_databases qZ1kwNLNSSav3z9jmfqovA   1   0         41            0     38.7mb         38.7mb
green  open   test             GibrJdE6Q6yL9lhts1Ho9A   1   0          0            0       226b           226b
```