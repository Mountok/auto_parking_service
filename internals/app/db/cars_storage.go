package db

import (
	"context"
	"errors"
	"fmt"
	"go_service_parking/example/internals/app/models"
	log "github.com/sirupsen/logrus"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type CarsStorage struct {
	databasePool *pgxpool.Pool
}

type userCar struct {
	UserId int64 `db:"user_id"`
	Name string
	Rank string
	CarId int64 `db:"carid"`
	Brand string
	Colour string
	LicensePlate string
}

func converteJoinedQueryToCar(input userCar) models.Car {
	return models.Car{
		Id: input.CarId,
		Colour: input.Colour,
		Brand: input.Brand,
		LicensePlate: input.LicensePlate,
		Owner: models.User{
			Id: input.UserId,
			Name: input.Name,
			Rank: input.Rank,
		},
	}
}


func NewCarStorage(pool *pgxpool.Pool) *CarsStorage {
	storage := new(CarsStorage)
	storage.databasePool = pool
	return storage
}


func (storage *CarsStorage) GetCarsList(userIdFilter int64, brandFilter,colourFilter,licenseFilter string) []models.Car {
	query := "SELECT user.id AS userid,user.name,user.rank, c.id AS carid, c.brand, c.colour, c.license_plate FROM users JOIN cars c on users.id = c.user_id WHERE 1=1"
	
	placeholderNum := 1
	args := make([]interface{},0)
	if userIdFilter != 0 {
		query += fmt.Sprintf(" AND user.id = $%d",placeholderNum)
		args = append(args, userIdFilter)
		placeholderNum++
	}
	if brandFilter != "" {
		query += fmt.Sprintf(" AND brand LIKE = %d",placeholderNum)
		args = append(args, fmt.Sprintf("%%%s%%", brandFilter))
		placeholderNum++
	}
	if colourFilter != "" {
		query += fmt.Sprintf(" AND colour LIKE %d",placeholderNum)
		args = append(args, fmt.Sprintf("%%%s%%", colourFilter))
		placeholderNum++
	}
	if licenseFilter != "" {
		query += fmt.Sprintf(" AND lecense_plate = %d",placeholderNum)
		args = append(args,  fmt.Sprintf("%%%s%%", licenseFilter))
		placeholderNum++
	}

	var dbResult []userCar

	err := pgxscan.Select(context.Background(),storage.databasePool,&dbResult,query,args...)
	if err != nil {
		log.Errorln(err)
	}

	result := make([]models.Car,len(dbResult))

	for idx, dbEntity := range dbResult {
		result[idx] = converteJoinedQueryToCar(dbEntity)
	}

	return result
}

func (storage *CarsStorage) GetCarById(id int64) models.Car {
	query := "SSELECT users.id AS userid, users.name, users.rank, c.id AS carid, c.brand, c.colour, c.license_plate FROM users JOIN cars c on users.id = c.user_id WHERE c.id = $1"
	var result userCar
	err := pgxscan.Get(context.Background(),storage.databasePool,&result,query,id)
	if err != nil {
		log.Errorln(err)
	}
	return converteJoinedQueryToCar(result)
}


func (storage *CarsStorage) CreateCar(car models.Car) error {
	ctx := context.Background()
	tx, err := storage.databasePool.Begin(ctx) //здесь будем пользоваться транзакцией, чтобы проверка пользователей и вставка нового автомобиля выглядели одним запросом с ее точки зрения
	defer func() {
		err = tx.Rollback(context.Background())
		if err != nil {
			log.Errorln(err)
		}
	}()

	query := "SELECT id FROM users WHERE id = $1"

	id := -1

	err = pgxscan.Get(ctx, tx, &id, query, car.Owner.Id)
	if err != nil {
		log.Errorln(err)
		err = tx.Rollback(context.Background()) //если получили ошибку откатываем транзакцию целиком
		if err != nil {
			log.Errorln(err)
		}
		return err
	}

	if id == -1 {
		return errors.New("user not found")
	}

	insertQuery := "INSERT INTO cars(user_id, colour, brand, license_plate) VALUES ($1,$2,$3,$4)"

	_, err = tx.Exec(context.Background(),insertQuery, car.Owner.Id, car.Colour, car.Brand, car.LicensePlate) //вызываем exec НЕ У соединения а У транзакции

	if err != nil {
		log.Errorln(err)
		err = tx.Rollback(context.Background())
		if err != nil {
			log.Errorln(err)
		}
		return err
	}
	err = tx.Commit(context.Background()) // в конце посылаем транзакцию, база сохранит значения, если до этого ничего не было откачено
	if err != nil {
		log.Errorln(err)
	}

	return err
}

