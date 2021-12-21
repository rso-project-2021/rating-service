package db

import (
	"context"
	"database/sql"
	"rating-service/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomStation(t *testing.T) Station {
	arg := CreateStationParam{
		Name:      util.RandomString(5),
		Latitude:  float32(util.RandomInt(46, 49)) + util.RandomFloat(),
		Longitude: float32(util.RandomInt(21, 23)) + util.RandomFloat(),
		Provider:  util.RandomString(5),
	}

	result, err := testStore.Create(context.Background(), arg)

	// Check if method executed correctly.
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, arg.Name, result.Name)
	require.Equal(t, arg.Latitude, result.Latitude)
	require.Equal(t, arg.Longitude, result.Longitude)
	require.Equal(t, arg.Provider, result.Provider)

	require.NotZero(t, result.ID)

	return result
}

func TestCreateStation(t *testing.T) {
	createRandomStation(t)
}

func TestGetStation(t *testing.T) {
	station1 := createRandomStation(t)
	station2, err := testStore.GetByID(context.Background(), station1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, station2)

	require.Equal(t, station1.ID, station2.ID)
	require.Equal(t, station1.Name, station2.Name)
	require.Equal(t, station1.Latitude, station2.Latitude)
	require.Equal(t, station1.Longitude, station2.Longitude)
	require.Equal(t, station1.Provider, station2.Provider)
}

func TestListStations(t *testing.T) {

	// Create a list of stations in database.
	var createdStations [10]Station
	for i := 0; i < 10; i++ {
		createdStations[i] = createRandomStation(t)
	}

	arg := ListStationsParam{
		Limit:  10,
		Offset: 0,
	}

	// Retrieve list of stations.
	stations, err := testStore.GetAll(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, stations)

	for _, u := range stations {
		require.NotEmpty(t, u)
	}
}

func TestUpdateStation(t *testing.T) {
	station1 := createRandomStation(t)

	arg := UpdateStationParam{
		Name:      util.RandomString(5),
		Latitude:  float32(util.RandomInt(46, 49)) + util.RandomFloat(),
		Longitude: float32(util.RandomInt(21, 23)) + util.RandomFloat(),
		Provider:  util.RandomString(5),
	}

	station2, err := testStore.Update(context.Background(), arg, station1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, station2)

	require.Equal(t, station1.ID, station2.ID)
	require.Equal(t, arg.Name, station2.Name)
	require.Equal(t, arg.Latitude, station2.Latitude)
	require.Equal(t, arg.Longitude, station2.Longitude)
	require.Equal(t, arg.Provider, station2.Provider)
}

func TestDeleteStation(t *testing.T) {
	station1 := createRandomStation(t)
	err := testStore.Delete(context.Background(), station1.ID)
	require.NoError(t, err)

	station2, err := testStore.GetByID(context.Background(), station1.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, station2)
}
