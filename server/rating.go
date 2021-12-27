package server

import (
	"net/http"
	"rating-service/db"

	"github.com/gin-gonic/gin"
)

type RatingController struct{}

type getRatingRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type getRatingListRequest struct {
	Offset int32 `form:"offset"`
	Limit  int32 `form:"limit" binding:"required,min=1,max=20"`
}

type createRatingRequest struct {
	Station_id int64  `json:"station_id" db:"station_id"`
	User_id    int64  `json:"user_id" db:"user_id"`
	Rating     int64  `json:"rating" db:"rating"`
	Comment    string `json:"comment" db:"comment"`
}

type updateRatingRequest struct {
	Station_id int64  `json:"station_id" db:"station_id"`
	User_id    int64  `json:"user_id" db:"user_id"`
	Rating     int64  `json:"rating" db:"rating"`
	Comment    string `json:"comment" db:"comment"`
}

func (server *Server) GetByID(ctx *gin.Context) {

	// Check if request has ID field in URI.
	var req getRatingRequest
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
	var req getRatingListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	arg := db.ListRatingParam{
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
	var req createRatingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	arg := db.CreateRatingParam{
		Station_id: req.Station_id,
		User_id:    req.User_id,
		Rating:     req.Rating,
		Comment:    req.Comment,
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
	var reqID getRatingRequest
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	// Check if request has all required fields in json body.
	var req updateRatingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	arg := db.UpdateRatingParam{
		Station_id: req.Station_id,
		User_id:    req.User_id,
		Rating:     req.Rating,
		Comment:    req.Comment,
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
	var req getRatingRequest
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

func (server *Server) GetAllByStation(ctx *gin.Context) {

	// Check if request has ID field in URI.
	var req getRatingRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	// Execute query.
	result, err := server.store.GetAllByStation(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, result)
}
