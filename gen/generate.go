package gen

func GenerateRestartScript() string {
	return `#!/bin/bash
set -ex

if [ -f /var/log/mysql/mysqld-slow.log ]; then
	sudo mv /var/log/mysql/mysqld-slow.log /var/log/mysql/mysqld-slow.log.$(date "+%Y%m%d_%H%M%S")
fi
if [ -f /var/log/nginx/access.log ]; then
	sudo mv /var/log/nginx/access.log /var/log/nginx/access.log.$(date "+%Y%m%d_%H%M%S")
fi
sudo systemctl restart mysql
sudo systemctl restart nginx`
}
