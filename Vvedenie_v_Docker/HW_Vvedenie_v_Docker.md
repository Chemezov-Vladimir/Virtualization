Задача 1
==
Сценарий выполения задачи:
* создайте свой репозиторий на https://hub.docker.com;
* выберете любой образ, который содержит веб-сервер Nginx;
* создайте свой fork образа;
* реализуйте функциональность: запуск веб-сервера в фоне с индекс-страницей, 
содержащей HTML-код ниже:

```
<html>
<head>
Hey, Netology
    </head>
    <body>
    <h1>I’m DevOps Engineer!</h1>
</body>
</html>
```

Опубликуйте созданный форк в своем репозитории и предоставьте ответ в виде ссылки на
https://hub.docker.com/username_repo.

Решение
--

Dockerfile:
```
FROM nginx
COPY ./index.html /usr/share/nginx/html/
EXPOSE 80
```

Создан образ командой `docker build -t vladimirchemezov/netolgoy:1 .`

При запуске командой `docker run -d -p 80:80 vladimirchemezov/netolgoy:1`
на localhost отображается нужная страница.
Ссылка на репозиторий: https://hub.docker.com/r/vladimirchemezov/netology

Задача 2
==

Посмотрите на сценарий ниже и ответьте на вопрос: "Подходит ли в этом сценарии 
использование Docker контейнеров или лучше подойдет виртуальная машина, 
физическая машина? Может быть возможны разные варианты?"

Детально опишите и обоснуйте свой выбор.

--

Сценарий:
* Высоконагруженное монолитное java веб-приложение;
* Nodejs веб-приложение;
* Мобильное приложение c версиями для Android и iOS;
* Шина данных на базе Apache Kafka;
* Elasticsearch кластер для реализации логирования продуктивного веб-приложения -
три ноды elasticsearch, два logstash и две ноды kibana;
* Мониторинг-стек на базе Prometheus и Grafana;
* MongoDB, как основное хранилище данных для java-приложения;
* Gitlab сервер для реализации CI/CD процессов и приватный (закрытый) Docker Registry.

Решение
--

Для высоконагруженного монолитного java веб-приложения подойдет вариант физического 
сервера, так как нужен прямой доступ к ресурсам из-за высокой нагрузки и, из-за 
монолитности, сложно реализовать контейнерную структуру.

Для Nodejs веб-приложения подойдет Docker, так как это веб-платформа с подключаемыми 
внешними библиотеками, что звучит как прекрасное основание для использования Docker. 
Виртуалка в этом случае не так подвижна, и обновление среды будет занимать больше
времени.

Для мобильного приложения жизненно необходим GUI, так что Docker не подойдет. А 
подойдет виртуалка.

Для Apache Kafka все зависит от данных и вида среды (тестовая или рабочая). 
Если среда рабочая и полнота данных критична – то лучше использовать ВМ; если среда 
тестовая и потеря данных некритична – можно использовать Docker.

Для Elasticsearch кластера:
1. Сам Elasticsearch можно запаковать в Docker для лучшей масштабируемости;
2. Logstash собирает и хранит данные в реальном времени, так что его можно в 
виртуалку или на физический сервер, в зависимости от объема данных;
3. Kibana – веб-интерфейс с возможностью подключения дополнительного функционала.
Для этого подойдет Docker.

Для мониторинг-стека на базе Prometheus и Grafana подойдет Docker, так как данные 
не хранятся, и масштабировать легко.

В случае использования MongoDB все зависит от нагрузки на БД. Если нагрузка большая,
то физический сервер, если нет – ВМ.

Для Gitlab сервера и Docker Registry, в зависимости от нагрузки, подойдет ВМ или 
физический сервер. Так как подразумевается хранение данных, то Docker не подойдет.

Задача 3
==
* Запустите первый контейнер из образа centos c любым тэгом в фоновом режиме, 
подключив папку /data из текущей рабочей директории на хостовой машине в /data 
контейнера;
* Запустите второй контейнер из образа debian в фоновом режиме, подключив папку 
/data из текущей рабочей директории на хостовой машине в /data контейнера;
* Подключитесь к первому контейнеру с помощью docker exec и создайте текстовый 
файл любого содержания в /data;
* Добавьте еще один файл в папку /data на хостовой машине;
* Подключитесь во второй контейнер и отобразите листинг и содержание файлов в /data 
контейнера.

Решение
--
Контейнеры запущены командами:

`docker run -v ~/ex3/data:/data --name first_centos -itd centos /bin/bash`

`docker run -v ~/ex3/data:/data --name second_debian -itd debian /bin/bash`

Подключение к первой машине с Centos:
```
vladimir@vladimir-VirtualBox:~$ docker exec -it first_centos /bin/bash
[root@726b0b743bd0 /]# cd /data/
[root@726b0b743bd0 data]# touch centos_info
[root@726b0b743bd0 data]# ls
centos_info
[root@726b0b743bd0 data]# echo 'test info for centos' > centos_info
[root@726b0b743bd0 data]# cat centos_info
test info for centos
[root@726b0b743bd0 data]#
```
Хостовая машина:
```
vladimir@vladimir-VirtualBox:~/ex3/data$ touch host_info
vladimir@vladimir-VirtualBox:~/ex3/data$ ls
centos_info host_info
vladimir@vladimir-VirtualBox:~/ex3/data$ echo 'test info for host' > host_info
vladimir@vladimir-VirtualBox:~/ex3/data$ cat host_info
test info for host
vladimir@vladimir-VirtualBox:~/ex3/data$
```
Вторая машина с Debian:
```
vladimir@vladimir-VirtualBox:~$ docker exec -it second_debian /bin/bash
root@e8ad19917af6:/# cd /data/
root@e8ad19917af6:/data# ls
centos_info host_info
root@e8ad19917af6:/data# cat centos_info
test info for centos
root@e8ad19917af6:/data# cat host_info
test info for host
root@e8ad19917af6:/data#
```
Задача 4 (*)
==
Воспроизвести практическую часть лекции самостоятельно.
Соберите Docker образ с Ansible, загрузите на Docker Hub и пришлите ссылку вместе 
с остальными ответами к задачам.

Решение
--
https://hub.docker.com/repository/docker/vladimirchemezov/ansible