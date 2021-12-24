package db

import (
	"context"
	"database/sql"
	"rating-service/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomRating(t *testing.T) Rating {
	arg := CreateRatingParam{
		Station_id: util.RandomInt(1261, 654561),
		User_id:    util.RandomInt(1261, 654561),
		Rating:     util.RandomInt(1, 5),
		Comment:    util.RandomString(5),
	}

	result, err := testStore.Create(context.Background(), arg)

	// Check if method executed correctly.
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, arg.Station_id, result.Station_id)
	require.Equal(t, arg.User_id, result.User_id)
	require.Equal(t, arg.Rating, result.Rating)
	require.Equal(t, arg.Comment, result.Comment)

	require.NotZero(t, result.ID)
	require.NotZero(t, result.CreatedAt)

	return result
}

func TestCreateRating(t *testing.T) {
	createRandomRating(t)
}

func TestGetRating(t *testing.T) {
	rating1 := createRandomRating(t)
	rating2, err := testStore.GetByID(context.Background(), rating1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, rating2)

	require.Equal(t, rating1.ID, rating2.ID)
	require.Equal(t, rating1.Station_id, rating2.Station_id)
	require.Equal(t, rating1.User_id, rating2.User_id)
	require.Equal(t, rating1.Rating, rating2.Rating)
	require.Equal(t, rating1.Comment, rating2.Comment)
	require.Equal(t, rating1.CreatedAt, rating2.CreatedAt)
}

func TestListRatings(t *testing.T) {

	// Create a list of ratings in database.
	var createdRatings [10]Rating
	for i := 0; i < 10; i++ {
		createdRatings[i] = createRandomRating(t)
	}

	arg := ListRatingParam{
		Limit:  10,
		Offset: 0,
	}

	// Retrieve list of ratings.
	ratings, err := testStore.GetAll(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, ratings)

	for _, u := range ratings {
		require.NotEmpty(t, u)
	}
}

func TestUpdateRating(t *testing.T) {

	rating1 := createRandomRating(t)

	arg := UpdateRatingParam{
		Station_id: util.RandomInt(1261, 654561),
		User_id:    util.RandomInt(1261, 654561),
		Rating:     util.RandomInt(1, 5),
		Comment:    util.RandomString(5),
	}

	rating2, err := testStore.Update(context.Background(), arg, rating1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, rating2)

	require.Equal(t, rating1.ID, rating2.ID)
	require.Equal(t, arg.Station_id, rating2.Station_id)
	require.Equal(t, arg.User_id, rating2.User_id)
	require.Equal(t, arg.Rating, rating2.Rating)
	require.Equal(t, arg.Comment, rating2.Comment)
	require.Equal(t, rating1.CreatedAt, rating2.CreatedAt)

}

func TestDeleteRating(t *testing.T) {
	rating1 := createRandomRating(t)
	err := testStore.Delete(context.Background(), rating1.ID)
	require.NoError(t, err)

	rating2, err := testStore.GetByID(context.Background(), rating1.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, rating2)
}
