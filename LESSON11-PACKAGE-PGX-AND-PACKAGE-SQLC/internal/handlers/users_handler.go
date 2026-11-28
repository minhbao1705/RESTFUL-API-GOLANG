package handlers

import (
	"errors"
	"hoc-golang/internal/db/sqlc"
	"hoc-golang/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type UserResponse struct {
	UserId    int32     `json:"user_id"`
	Uuid      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Create_at string    `json:"created_at"`
}

type UserHandler struct {
	repo repository.UserRepository
}

func NewUserHander(repo repository.UserRepository) *UserHandler {
	return &UserHandler{

		repo: repo,
	}
}

func (uh *UserHandler) GetUserById(ctx *gin.Context) {
	uuidParam := ctx.Param("uuid")

	userUUID, err := uuid.Parse(uuidParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := uh.repo.FindByUuid(ctx, userUUID)

	if errors.Is(err, pgx.ErrNoRows) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	} else if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	response := UserResponse{
		UserId:    user.UserID,
		Uuid:      user.Uuid,
		Name:      user.Name,
		Email:     user.Email,
		Create_at: user.CreatedAt.Format("2006-01-02"),
	}

	ctx.JSON(http.StatusOK, gin.H{"data": response})
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var input sqlc.CreateUserParams
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uh.repo.Create(ctx, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := UserResponse{
		UserId:    user.UserID,
		Uuid:      user.Uuid,
		Name:      user.Name,
		Email:     user.Email,
		Create_at: user.CreatedAt.Format("2006-01-02"),
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": response})
}
