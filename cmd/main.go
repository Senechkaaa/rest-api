package main

import (
	"github.com/Senechkaaa/todo-app"
	"github.com/Senechkaaa/todo-app/pkg/handler"
	"github.com/Senechkaaa/todo-app/pkg/repository"
	"github.com/Senechkaaa/todo-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error init configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading .env file: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	//fmt.Printf("host: %s, port: %s, username: %s, password: %s, dbname: %s, sslmode: %s", viper.GetString("db.host"), viper.GetString("db.port"), viper.GetString("db.username"), os.Getenv("DB_PASSWORD"), viper.GetString("db.dbname"), viper.GetString("db.sslmode"))

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	srv := new(todo.Server)
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("err occured while running server: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

/*
-- not null - не может быть пустым
-- unique - должно быть уникальным, без дублирований
-- varchar(255) - тип данных - строка с наиб кол-вом символов 255
-- serial - тип данных - целое число, к которому автоматически добавляется последовательность для автоинкрементации.
-- default false - значение по умолчанию false
-- references todo_items(id) - Внешний ключ, ссылающийся на столбец id таблицы todo_items
-- on delete cascade - Удаление связанных записей из этой таблицы при удалении записи из todo_items
*/

/*
		Работа с docker:
	1. docker pull postgres

	2. docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
	* docker run - запускает контейнер
	* --name=todo-db - задает имя контейнеру
 	* -e POSTGRES_PASSWORD='sigma' - устанавливает переменную окружения, задается пароль
	* -p 5436:5433 -  Пробрасывает порт. Внешний порт 5436 (на хост-машине) будет связан с внутренним портом 5433 (в контейнере)
	* -d - Запускает контейнер в фоновом режиме (detached mode).
	* --rm Удаляет контейнер после его остановки.

	3. migrate create -ext sql -dir ./schema -seq init
	* migrate create - создание новой миграции
	* -ext sql - указывает расширение файлов миграции
	* -dir ./schema - указывает папку в которой будет миграция
	* -seq - миграция будет пронумерована
	* init - имя миграции

	4. migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
	* -path ./schema - путь
	* postgres:qwerty@localhost:5436 - учетные данные и адрес бд
	* postgres - имя usera, qwerty - password, localhost - локальная машина + port
	* /postgres - имя бд
	* ?sslmode=disable: Параметр отключения SSL для соединения.
	* up - запустить все миграций

	5. docker exec -it b8f8c0dea60e /bin/bash
	* docker exec - Используется для выполнения команды в запущенном контейнере Docker.
	* it - подключаемся к терминалу

	6. psql -U postgres

	7. \d


	Др команды:
	select * from schema_migrations;
	update schema_migrations set version='000001',  dirty=false;


*/
