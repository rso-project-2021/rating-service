package db

import (
	"context"
)

type Station struct {
	ID        int64   `json:"station_id" db:"station_id"`
	Name      string  `json:"name" db:"name"`
	Latitude  float32 `json:"lat" db:"lat"`
	Longitude float32 `json:"lng" db:"lng"`
	Provider  string  `json:"provider" db:"provider"`
}

type CreateStationParam struct {
	Name      string
	Latitude  float32
	Longitude float32
	Provider  string
}

type UpdateStationParam struct {
	Name      string
	Latitude  float32
	Longitude float32
	Provider  string
}

type ListStationsParam struct {
	Offset int32
	Limit  int32
}

func (store *Store) GetByID(ctx context.Context, id int64) (station Station, err error) {
	const query = `SELECT * FROM "stations" WHERE "station_id" = $1`
	err = store.db.GetContext(ctx, &station, query, id)

	return
}

func (store *Store) GetAll(ctx context.Context, arg ListStationsParam) (stations []Station, err error) {
	const query = `SELECT * FROM "stations" OFFSET $1 LIMIT $2`
	stations = []Station{}
	err = store.db.SelectContext(ctx, &stations, query, arg.Offset, arg.Limit)

	return
}

func (store *Store) Create(ctx context.Context, arg CreateStationParam) (Station, error) {
	const query = `
	INSERT INTO "stations"("name", "lat", "lng", "provider") 
	VALUES ($1, $2, $3, $4)
	RETURNING "station_id", "name", "lat", "lng", "provider"
	`
	row := store.db.QueryRowContext(ctx, query, arg.Name, arg.Latitude, arg.Longitude, arg.Provider)

	var station Station

	err := row.Scan(
		&station.ID,
		&station.Name,
		&station.Latitude,
		&station.Longitude,
		&station.Provider,
	)

	return station, err
}

func (store *Store) Update(ctx context.Context, arg UpdateStationParam, id int64) (Station, error) {
	const query = `
	UPDATE "stations"
	SET "name" = $2,
		"lat" = $3,
		"lng" = $4,
		"provider" = $5
	WHERE "station_id" = $1
	RETURNING "station_id", "name", "lat", "lng", "provider"
	`
	row := store.db.QueryRowContext(ctx, query, id, arg.Name, arg.Latitude, arg.Longitude, arg.Provider)

	var station Station
	err := row.Scan(
		&station.ID,
		&station.Name,
		&station.Latitude,
		&station.Longitude,
		&station.Provider,
	)

	return station, err
}

func (store *Store) Delete(ctx context.Context, id int64) error {
	const query = `
	DELETE FROM stations
	WHERE "station_id" = $1
	`
	_, err := store.db.ExecContext(ctx, query, id)

	return err
}
