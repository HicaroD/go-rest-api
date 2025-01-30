# RDM - Relational Database Manager

## Drivers

In order to add a new driver support, go to `drivers.go` and paste the code in
the corresponding driver method.

### SQLite

<details>

1. `import "gorm.io/driver/sqlite"`

2. Replace the placeholder method with the following method:

   ```go
   // Connect to SQLite database
   func connectSQLite(config DatabaseConfig) (*gorm.DB, error) {
       dsn := config.DBName
       return gorm.Open(sqlite.Open(dsn), &gorm.Config{})
   }
   ```

</details>

### MySQL

<details>

1. `import "gorm.io/driver/mysql"`

2. Replace the placeholder method with the following method:

   ```go
   // Connect to MySQL database
   func connectMySQL(config DatabaseConfig) (*gorm.DB, error) {
       dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
           config.User, config.Password, config.Host, config.Port, config.DBName)

       return gorm.Open(mysql.Open(dsn), &gorm.Config{})
   }
   ```

</details>

### PostgreSQL

<details>

1. `import "gorm.io/driver/postgres"`

2. Replace the placeholder method with the following method:

   ```go
    // Connect to PostgreSQL database
    func connectPostgres(config DatabaseConfig) (*gorm.DB, error) {
        dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
            config.Host, config.Port, config.User, config.Password, config.DBName)
        return gorm.Open(postgres.Open(dsn), &gorm.Config{})
    }
   ```

</details>

## Good practices

1. If you don't need a driver anymore, remove it! Don't install unnecessary
   packages.
