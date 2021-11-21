Задача 1
=
* Опишите своими словами основные преимущества применения на практике IaaC паттернов.
* Какой из принципов IaaC является основополагающим?

Решение
-
Преимущества непрерывной интеграции (CI):
1.	Все проблемы и ошибки продукта оперативно выявляются и могут быть оперативно устранены;
2.	Немедленное тестирование продукта на тестовых стендах;
3.	Ускорено время получения обновлений.

Непрерывная доставка (CD):
1.	Быстрое обновление за счет автоматизации процессов;
2.	Обновление небольшими пакетами позволяет лучше сохранить стабильность системы, в отличие от больших обновлений;
3.	Полная готовность продукта к релизу, так как на этом этапе уже, обычно, протестировано все.

Непрерывное развертывание (CD):
1.	Снижение влияния человеческого фактора на процесс развертывания за счет автоматизации;
2.	Ускорение процесса разработки;
3.	Постоянное улучшение качества продукта.

Основополагающий принцип IaaC – автоматизация всех рутинных задач, чтоб оставались ресурсы для действительно важных вещей.

Задача 2
=
* Чем Ansible выгодно отличается от других систем управление конфигурациями?
* Какой, на ваш взгляд, метод работы систем конфигурации более надёжный push или pull?

Решение
-
Главное отличие Ansible от других систем управления конфигурациями заключается в том, что на управляемые сервера не надо устанавливать дополнительное ПО, достаточно SSH-подключения.

На мой взгляд, pull-система выглядит надежнее, так как каждый сервер отвечает за себя и не ждет команды от центрального сервера, который может вообще не работать. Главное, чтоб был доступ к репозиторию с конфигурацией. 

Задача 3
=
Установить на личный компьютер:

* VirtualBox
* Vagrant
* Ansible

*Приложить вывод команд установленных версий каждой из программ, оформленный в markdown.*

Решение
-
VirtualBox:

    vladimir@vladimir-VirtualBox:~$ vboxmanage --version
    6.1.26_Ubuntur145957

Vagrant:

    vladimir@vladimir-VirtualBox:~$ vagrant --version
    Vagrant 2.2.6

Ansible:

    vladimir@vladimir-VirtualBox:~$ ansible --version
    ansible 2.9.6
      config file = /etc/ansible/ansible.cfg
      configured module search path = ['/home/vladimir/.ansible/plugins/modules', '/usr/share/ansible/plugins/modules']
      ansible python module location = /usr/lib/python3/dist-packages/ansible
      executable location = /usr/bin/ansible
      python version = 3.8.10 (default, Sep 28 2021, 16:10:42) [GCC 9.3.0]

Задача 4 (*)
=
Воспроизвести практическую часть лекции самостоятельно. 

* Создать виртуальную машину. 

* Зайти внутрь ВМ, убедиться, что Docker установлен с помощью команды 


    docker ps

Решение
-
    vagrant@server1:~$ docker ps
    CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES