package server

import (
	"net/http"
	"rating-service/db"

	"github.com/gin-gonic/gin"
)

type StationController struct{}

type getStationRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type getStationListRequest struct {
	Offset int32 `form:"offset"`
	Limit  int32 `form:"limit" binding:"required,min=1,max=20"`
}

type createStationRequest struct {
	Name      string  `json:"name" db:"name"`
	Latitude  float32 `json:"lat" db:"lat"`
	Longitude float32 `json:"lng" db:"lng"`
	Provider  string  `json:"provider" db:"provider"`
}

type updateStationRequest struct {
	Name      string  `json:"name" db:"name"`
	Latitude  float32 `json:"lat" db:"lat"`
	Longitude float32 `json:"lng" db:"lng"`
	Provider  string  `json:"provider" db:"provider"`
}

func (server *Server) GetByID(ctx *gin.Context) {

	// Check if request has ID field in URI.
	var req getStationRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	// Execute query.
	result, err := server.store.GetByID(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) GetAll(ctx *gin.Context) {

	// Check if request has parameters offset and limit for pagination.
	var req getStationListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	arg := db.ListStationsParam{
		Offset: req.Offset,
		Limit:  req.Limit,
	}

	// Execute query.
	result, err := server.store.GetAll(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) Create(ctx *gin.Context) {

	// Check if request has all required fields in json body.
	var req createStationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	arg := db.CreateStationParam{
		Name:      req.Name,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Provider:  req.Provider,
	}

	// Execute query.
	result, err := server.store.Create(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (server *Server) Update(ctx *gin.Context) {

	// Check if request has ID field in URI.
	var reqID getStationRequest
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	// Check if request has all required fields in json body.
	var req updateStationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	arg := db.UpdateStationParam{
		Name:      req.Name,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Provider:  req.Provider,
	}

	// Execute query.
	result, err := server.store.Update(ctx, arg, reqID.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (server *Server) Delete(ctx *gin.Context) {

	// Check if request has ID field in URI.
	var req getStationRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	// Execute query.
	if err := server.store.Delete(ctx, req.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
