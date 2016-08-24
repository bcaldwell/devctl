package plugins

// treat as constant
var serviceList = map[string]Service{
	"mongo": Service{
		image:   "mongo",
		port:    "27017/tcp",
		volumes: []string{"/data/db", "/data/configdb"},
		tag:     "latest",
	},
	"mysql": Service{
		image:       "mysql",
		port:        "3306/tcp",
		volumes:     []string{"/var/lib/mysql"},
		environment: []string{"MYSQL_ROOT_PASSWORD=devctl"},
		tag:         "latest",
	},
	"postgres": Service{
		image:   "postgres",
		port:    "5432/tcp",
		volumes: []string{"/var/lib/postgresql/data"},
		tag:     "alpine",
	},
	"rethinkdb": Service{
		image:   "rethinkdb",
		port:    "3306/tcp",
		volumes: []string{"/data"},
		tag:     "alpine",
	},
	"redis": Service{
		image:   "redis",
		port:    "6379/tcp",
		volumes: []string{"/data"},
		tag:     "alpine",
	},
	"memcached": Service{
		image: "memcached",
		port:  "11211/tcp",
		tag:   "alpine",
	},
	"elasticsearch": Service{
		image:   "elasticsearch",
		port:    "9200/tcp",
		volumes: []string{"/usr/share/elasticsearch/data", "/usr/share/elasticsearch/config"},
		tag:     "alpine",
	},
	"rabbitmq": Service{
		image:   "rabbitmq",
		port:    "8000/tcp",
		volumes: []string{"/var/liv/rabbitmq"},
		tag:     "alpine",
	},
	"solr": Service{
		image:   "solr",
		port:    "8983/tcp",
		volumes: []string{"/var/liv/rabbitmq"},
		tag:     "alpine",
	},
	"celery": Service{
		image:   "celery",
		port:    "8983/tcp",
		volumes: []string{"/var/liv/rabbitmq"},
		tag:     "alpine",
	},
}

// celery
// mariadb
