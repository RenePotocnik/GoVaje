# Docker basics

### Installation

Instalirajte si docker od tod: [Docker]

### Starting MySQL DB

```sh
docker run --name mariadb -e MYSQL_ROOT_PASSWORD=superSecretPass -e MYSQL_DATABASE=test -p 3306:3306 -d --restart always mariadb
```
MySQL server ostane vklopljen tudi po restartu. Če to ne želite izbrišete --restart always iz komande

MySQL server se zažene z naslednjimi podatki  
IP: 127.0.0.1  
Username: root  
Password: superSecretPass  
DB: test  
Port: 3306  

   [Docker]: <https://hub.docker.com/search?q=&type=edition&offering=community&sort=updated_at&order=desc>
