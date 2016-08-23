package plugins

// treat as constant
var serviceList = map[string]Service{
	"mongo": Service{
		image:   "mongo",
		port:    "27017/tcp",
		volumes: []string{"/data/db", "/data/configdb"},
	},
	"mysql": Service{
		image:       "mysql",
		port:        "3306/tcp",
		volumes:     []string{"/var/lib/mysql"},
		environment: []string{"MYSQL_ROOT_PASSWORD=devctl"},
	},
	"postgres": Service{
		image:   "postgres",
		port:    "5432/tcp",
		volumes: []string{"/var/lib/postgresql/data"},
	},
	"rethinkdb": Service{
		image:   "rethinkdb",
		port:    "3306/tcp",
		volumes: []string{"/data"},
	},
	"redis": Service{
		image:   "redis",
		port:    "6379/tcp",
		volumes: []string{"/data"},
	},
	"memcached": Service{
		image:   "memcached",
		port:    "11211/tcp",
		volumes: []string{"/data"},
	},
	"elasticsearch": Service{
		image:   "elasticsearch",
		port:    "9200/tcp",
		volumes: []string{"/usr/share/elasticsearch/data", "/usr/share/elasticsearch/config"},
	},
	"rabbitmq": Service{
		image:   "rabbitmq",
		port:    "8000/tcp",
		volumes: []string{"/var/liv/rabbitmq"},
	},
}

// solr
// celery
// mariadb
