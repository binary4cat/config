"service" = {
  "consul" = "127.0.0.1:8500"
  "manage" = {
    "limit" = true
    "monitor" = "127.0.0.1:8888"
  }
  "name" = "test"
}

"database" = {
  "conn_str" = "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8\u0026parseTime=true"
  "db_source" = "mysql"
}