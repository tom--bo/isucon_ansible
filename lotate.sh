#!/bin/sh

service mysqld stop

mv /var/log/mysql/slow-query.log /var/log/mysql/slow mysql_slow_`date ":%Y%m%d_%H%M%S"`.log
touch /var/log/mysql/slow-query.log

service mysqld start


