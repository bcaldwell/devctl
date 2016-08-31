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
	"mariadb": Service{
		image:       "mariadb",
		port:        "3306/tcp",
		volumes:     []string{"/var/lib//mysql"},
		environment: []string{"MYSQL_ROOT_PASSWORD=devctl"},
		tag:         "latest",
	},
	"rethinkdb": Service{
		image:   "rethinkdb",
		port:    "8080/tcp",
		volumes: []string{"/data"},
		tag:     "latest",
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
		tag:     "latest",
	},
	"rabbitmq": Service{
		image:   "rabbitmq",
		port:    "5672/tcp",
		volumes: []string{"/var/liv/rabbitmq"},
		tag:     "latest",
	},
	"solr": Service{
		image:   "solr",
		port:    "8983/tcp",
		volumes: []string{"/var/liv/rabbitmq"},
		tag:     "alpine",
	},
}

// mariadb
