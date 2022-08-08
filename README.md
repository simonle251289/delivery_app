# delivery_app
### Learning Golang language
* This is the basic delivery app. I write it while leaning Golang.
  - Technique using:
    - Backend: Golang language (https://go.dev/doc/)
    - Database: Mysql
    - Manage database connect with GORM (https://gorm.io/docs/)
    - Manage api router with Gin (https://gin-gonic.com/docs/)
  - IDE using:
    - Coding: GoLand (https://www.jetbrains.com/go/)
    - View database: TablePlus (https://tableplus.com/)
    - Use docker for start mysql container (https://docs.docker.com/)
    - Test request api: Postman (https://www.postman.com/)
### Startup server
  - Open GoLand IDE
    - Add configuration
    - Add environment: `DBConnection=db_user:db_pw@tcp(127.0.0.1:3306)/db_name?charset=utf8mb4&parseTime=True`
    - Press button run