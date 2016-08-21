package plugins

// treat as constant
var serviceList = map[string]Service{
	"mongo": Service{
		image:   "mongo",
		port:    "27017/tcp",
		volumes: []string{"/data/db"},
	},
}
