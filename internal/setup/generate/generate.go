package generate

import (
	"os"

	"github.com/pkg/errors"
)

const perm os.FileMode = 0755

func RestartScript() error {
	script := `#!/bin/bash
set -ex

if [ -f /var/log/mysql/mysqld-slow.log ]; then
	sudo mv /var/log/mysql/mysqld-slow.log /var/log/mysql/mysqld-slow.log.$(date "+%Y%m%d_%H%M%S")
fi
if [ -f /var/log/nginx/access.log ]; then
	sudo mv /var/log/nginx/access.log /var/log/nginx/access.log.$(date "+%Y%m%d_%H%M%S")
fi
sudo systemctl restart mysql
sudo systemctl restart nginx`

	f, err := os.Create("restart.sh")
	if err != nil {
		return errors.Wrap(err, "Failed to create file")
	}
	f.WriteString(script)
	f.Chmod(perm)
	f.Close()

	return nil
}
