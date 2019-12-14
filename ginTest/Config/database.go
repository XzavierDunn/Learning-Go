package Config

import (
    "fmt"
    "github.com/jinzhu/gorm"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
    Host string
    Port int
    User string
    DBName string
    Password string
}

func BuildDBConfig() *DBConfig {
    dbConfig := DBConfig{
        Host: "127.0.0.1",
        Port: 5432,
        User: "xd",
        DBName: "gintest",
        Password: "3510",
    }
    return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
    return fmt.Sprintf(
        "postgres://%s:%s@%s:%d/%s",
        dbConfig.User,
        dbConfig.Password,
        dbConfig.Host,
        dbConfig.Port,
        dbConfig.DBName,
    )
}

