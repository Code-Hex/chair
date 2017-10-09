package config

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

func SQLSlowLog() string {
	return `[mysqld]
slow_query_log = 1
slow_query_log_file = /var/log/mysql/mysqld-slow.log
long_query_time = 0`
}

func SQLDefault() string {
	return `[mysqld]
datadir = /var/lib/mysql
socket = /var/lib/mysql/mysql.sock
symbolic-links = 0

max_allowed_packet = 300M

[mysqld_safe]
log-error = /var/log/mysql/mysqld.log
pid-file = /var/run/mysqld/mysqld.pid`
}

func SQLMaybeNice() (string, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}
	totalGB := v.Total / (1024 * 1024 * 1024)
	PerGB := int(float64(totalGB) * float64(0.6))
	return fmt.Sprintf(`[mysqld]
innodb_buffer_pool_size = %dGB
innodb_flush_log_at_trx_commit = 2
innodb_flush_method = O_DIRECT`, PerGB), nil
}

func SQLCache() string {
	return `[mysqld]
query_cache_limit = 128M

# memory size for caching
query_cache_size = 1024M

# 0: off
# 1: ON (Other than "SELECT SQL_NO_CACHE")
# 2: DEMAND (Only "SELECT SQL_CACHE")
query_cache_type = 1`
}

func SQLFix57GroupByProblem() string {
	return `[mysqld]
sql_mode=NO_ENGINE_SUBSTITUTION,STRICT_TRANS_TABLES`
}

func NginxAccessLog() string {
	return `http {
	log_format ltsv 'domain:$host\t'
		'host:$remote_addr\t'
		'user:$remote_user\t'
		'time:$time_local\t'
		'method:$request_method\t'
		'path:$request_uri\t'
		'protocol:$server_protocol\t'
		'status:$status\t'
		'size:$body_bytes_sent\t'
		'referer:$http_referer\t'
		'agent:$http_user_agent\t'
		'response_time:$request_time\t'
		'cookie:$http_cookie\t'
		'set_cookie:$sent_http_set_cookie\t'
		'upstream_addr:$upstream_addr\t'
		'upstream_cache_status:$upstream_cache_status\t'
		'upstream_response_time:$upstream_response_time';


	access_log /var/log/nginx/access.log ltsv;
}`
}

func NginxEvent() string {
	return `events {
	worker_connections  1024;
	# accept_mutex_delay 100ms;
	multi_accept on;
	use epoll;
}`
}

func NginxOutsideMaybeNice() string {
	return `worker_processes  1;
worker_rlimit_nofile 10000;`
}

func NginxStaticLocation() string {
	return `location / {
	gzip_static always;
	gunzip on;
	open_file_cache max=100 inactive=20s;
	expires 1d;
}`
}
