package database

import (
	"fmt"
	"log"
	"github.com/Wojbeg/go-gorilla-crm/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	Close() error
	AutoMigrate(model interface{}) error
	DropTable(model interface{}) error
	Model(value interface{}) *gorm.DB

	Create(model interface{}) *gorm.DB
	Save(model interface{}) *gorm.DB
	Update(model interface{}) *gorm.DB
	Delete(model interface{}, conds ...interface{}) *gorm.DB

	Select(query interface{}, args ...interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	Find(dest interface{}, conds ...interface{}) *gorm.DB
	Exec(sql string, values ...interface{}) *gorm.DB
	Raw(sql string, values ...interface{}) *gorm.DB
}

type infosRepository struct {
	*repository
}

func CreateInfosRepository(config *config.Config) Repository {
	db, err := connectToDatabase(config.DBUser, config.DBPassword, config.DBPort, config.DBHost, config.DBName)

	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	} else {
		fmt.Println("Successfully Connected to the database")
	}

	return &infosRepository{&repository{db: db}}
}

func connectToDatabase(User, Password, Port, Host, DBName string) (*gorm.DB, error) {
	fmt.Println("Connecting to databse:")
	fmt.Println("User: ", User)
	fmt.Println("Password: ", Password)
	fmt.Println("Port: ", Port)
	fmt.Println("Host: ", Host)
	fmt.Println("DBName: ", DBName)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", User, Password, Host, Port, DBName)

	db, err := gorm.Open(mysql.Open(dsn))
	return db, err
}


func (rep *repository) Close() error {
	db, _ := rep.db.DB()
	return db.Close()
}

func (rep *repository) Model(model interface{}) *gorm.DB {
	return rep.db.Model(model)
}

func (rep *repository) AutoMigrate(model interface{}) error {
	return rep.db.AutoMigrate(model)
}

func (rep *repository) DropTable(model interface{}) error {
	return rep.db.Migrator().DropTable(model)
}

func (rep *repository) Create(model interface{}) *gorm.DB {
	return rep.db.Create(model)
}

func (rep *repository) Save(model interface{}) *gorm.DB {
	return rep.db.Save(model)
}

func (rep *repository) Update(model interface{}) *gorm.DB {
	return rep.db.Updates(model)
}

func (rep *repository) Delete(model interface{}, conds ...interface{}) *gorm.DB {
	return rep.db.Delete(model, conds...)
}

func (rep *repository) Select(query interface{}, args ...interface{}) *gorm.DB {
	return rep.db.Select(query, args...)
}

func (rep *repository) First(dest interface{}, conds ...interface{}) *gorm.DB {
	return rep.db.First(dest, conds...)
}

func (rep *repository) Where(query interface{}, args ...interface{}) *gorm.DB {
	return rep.db.Where(query, args...)
}

func (rep *repository) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	return rep.db.Find(dest, conds...)
}

func (rep *repository) Exec(sql string, values ...interface{}) *gorm.DB {
	return rep.db.Exec(sql, values...)
}

func (rep *repository) Raw(sql string, values ...interface{}) *gorm.DB {
	return rep.db.Raw(sql, values...)
}
