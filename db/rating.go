package db

import (
	"context"
	"time"
)

type Rating struct {
	ID         int64     `json:"rating_id" db:"rating_id"`
	Station_id int64     `json:"station_id" db:"station_id"`
	User_id    int64     `json:"user_id" db:"user_id"`
	Rating     int64     `json:"rating" db:"rating"`
	Comment    string    `json:"comment" db:"comment"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

type CreateRatingParam struct {
	Station_id int64
	User_id    int64
	Rating     int64
	Comment    string
}

type UpdateRatingParam struct {
	Station_id int64
	User_id    int64
	Rating     int64
	Comment    string
}

type ListRatingParam struct {
	Offset int32
	Limit  int32
}

// HTTPError types

type Empty struct {
}

type HTTPError400 struct {
	Message []Empty `json:"message" example:[Empty]`
}

type HTTPError404 struct {
	Message Empty `json:"message" example:Empty`
}

type HTTPError500 struct {
	Message Empty `json:"message" example:Empty`
}

/// GetByID godoc
// @Summary      Get a rating by its ID
// @Description  get rating by ID
// @ID           get-rating-by-int
// @Tags         ratings
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Rating ID"
// @Success      200  {object}  Rating
// @Failure      404  {object}  HTTPError404
// @Failure      500  {object}  HTTPError500
// @Router       /ratings/{id} [get]
func (store *Store) GetByID(ctx context.Context, id int64) (rating Rating, err error) {
	const query = `SELECT * FROM "ratings" WHERE "rating_id" = $1`
	err = store.db.GetContext(ctx, &rating, query, id)

	return
}

/// GetAll godoc
// @Summary      Get all ratings and comments
// @Description  get all ratings
// @ID           get-all-ratings
// @Tags         ratings
// @Accept 		 mpfd
// @Produce      json
// @Param        offset   query      int  true  "Offset"
// @Param        limit   query      int  true  "Limit"
// @Success      200  {object}  []Rating
// @Header		 200 {object}	BasicHeader
// @Failure      400  {object}  HTTPError400
// @Failure      404  {object}  HTTPError404
// @Failure      500  {object}  HTTPError500
// @Router       /ratings [get]
func (store *Store) GetAll(ctx context.Context, arg ListRatingParam) (ratings []Rating, err error) {
	const query = `SELECT * FROM "ratings" OFFSET $1 LIMIT $2`
	ratings = []Rating{}
	err = store.db.SelectContext(ctx, &ratings, query, arg.Offset, arg.Limit)

	return
}

/// Create godoc
// @Summary      Create a new rating
// @Description  create rating
// @ID           create-rating
// @Tags         ratings
// @Accept       json
// @Produce      json
// @Param        message  body  CreateRatingParam  true  "Rating parametres"
// @Success      201  {object}  Rating
// @Failure      404  {object}  HTTPError404
// @Failure      500  {object}  HTTPError500
// @Router       /ratings [post]
func (store *Store) Create(ctx context.Context, arg CreateRatingParam) (Rating, error) {
	const query = `
	INSERT INTO "ratings"("station_id", "user_id", "rating", "comment") 
	VALUES ($1, $2, $3, $4)
	RETURNING "rating_id", "station_id", "user_id", "rating", "comment", "created_at"
	`
	row := store.db.QueryRowContext(ctx, query, arg.Station_id, arg.User_id, arg.Rating, arg.Comment)

	var rating Rating

	err := row.Scan(
		&rating.ID,
		&rating.Station_id,
		&rating.User_id,
		&rating.Rating,
		&rating.Comment,
		&rating.CreatedAt,
	)

	return rating, err
}

/// Update godoc
// @Summary      Update a rating
// @Description  update rating
// @ID           update-rating
// @Tags         ratings
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Rating ID"
// @Param        message  body  UpdateRatingParam  true  "Rating parametres"
// @Success      201  {object}  Rating
// @Failure 	 400  {object}  HTTPError400
// @Failure      404  {object}  HTTPError404
// @Failure      500  {object}  HTTPError500
// @Router       /ratings/{id} [put]
func (store *Store) Update(ctx context.Context, arg UpdateRatingParam, id int64) (Rating, error) {
	const query = `
	UPDATE "ratings"
	SET "station_id" = $2,
		"user_id" = $3,
		"rating" = $4,
		"comment" = $5
	WHERE "rating_id" = $1
	RETURNING "rating_id", "station_id", "user_id", "rating", "comment", "created_at"
	`
	row := store.db.QueryRowContext(ctx, query, id, arg.Station_id, arg.User_id, arg.Rating, arg.Comment)

	var rating Rating

	err := row.Scan(
		&rating.ID,
		&rating.Station_id,
		&rating.User_id,
		&rating.Rating,
		&rating.Comment,
		&rating.CreatedAt,
	)

	return rating, err
}

/// Delete godoc
// @Summary      Delete a rating
// @Description  delete rating
// @ID           delete-rating
// @Tags         ratings
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Rating ID"
// @Success      204
// @Failure      404  {object}  HTTPError404
// @Failure      500  {object}  HTTPError500
// @Router       /ratings/{id} [delete]
func (store *Store) Delete(ctx context.Context, id int64) error {
	const query = `
	DELETE FROM ratings
	WHERE "rating_id" = $1
	`
	_, err := store.db.ExecContext(ctx, query, id)

	return err
}

/// GetAllByStation godoc
// @Summary      Get all ratings of a single station by its ID
// @Description  get rating by station
// @ID           get-rating-by-station
// @Tags         ratings
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID of station"
// @Success      200  {object}  []Rating
// @Failure      400  {object}  HTTPError400
// @Failure      404  {object}  HTTPError404
// @Failure      500  {object}  HTTPError500
// @Router       /ratings/station/{id} [get]
func (store *Store) GetAllByStation(ctx context.Context, stationID int64) (ratings []Rating, err error) {
	const query = `SELECT * FROM "ratings" WHERE "station_id" = $1`
	ratings = []Rating{}
	err = store.db.SelectContext(ctx, &ratings, query, stationID)

	return
}
