The config component is encapsulated in viper, which simplifies the use and adds changes to the monitoring configuration file. Currently, it does not support the loading of other configuration sources except local files. In theory, any configuration source can be supported, and an extension interface needs to be added.

The config package supports tracking file modifications and real-time modification of the values ​​of configuration-related objects.

For example, there is such a configuration file (config.yaml):

```yaml
service:
  name: test
  consul: 127.0.0.1:8500
  manage:
    limit: true
    monitor: 127.0.0.1:8888
database:
  db_source: mysql
  conn_str: root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true
```

The corresponding configuration object is defined as (tag name is the type of configuration file, for example, yaml type configuration file is `yaml:"xxx"`, json configuration file is `json:"xxx"`):

```golang
type Config struct {
 Service  Service  `yaml:"service"`
 Database Database `yaml:"database"`
}

type Service struct {
 Name   string `yaml:"name"`
 Consul string `yaml:"consul"`
 Manage Manage `yaml:"manage"`
}

type Manage struct {
 Limit   bool   `yaml:"limit"`
 Monitor string `yaml:"monitor"`
}

type Database struct {
 DBSource string `yaml:"db_source"`
 ConnStr  string `yaml:"conn_str"`
}
```

Complete analysis method:

```golang
package main

import "github.com/hjdo/config"

func main() {
  // Parse the content of the configuration file into the Config structure, 
  // and correspond to the configuration file changes in real time
  var c Config
  _ = config.Init("./config.yaml", "", &c)
  for {
    // If you modify the contents of the configuration file, 
    // the values ​​of the corresponding fields in the structure will be modified at the same time
    fmt.Printf("config is %#v\n", c)
    time.Sleep(10 * time.Second)
  }
}
```

Configuration item independent resolution mode:

```golang
package main

import "github.com/hjdo/config"

func main() {
  var s Service
  var d Database
  var m Manage
  // The service part configuration is treated as an independent object. The second parameter needs to be filled in for independent analysis, 
  // which indicates the name of the configuration item corresponding to the independent object. 
  // The "." connection is used in the middle of the multi-level configuration.
  _ = config.Init("./config.yaml", "service", &s)
  _ = config.Init("./config.yaml", "database", &d)
  _ = config.Init("./config.yaml", "service.manage", &m)
  for {
   // Each individual object of the same will also be monitored and automatically modify the value when changes occur
   fmt.Printf("config is %#v\n", s)
   fmt.Printf("config is %#v\n", d)
   fmt.Printf("config is %#v\n", m)
   time.Sleep(10 * time.Second)
 }
}
```