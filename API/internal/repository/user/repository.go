package user

import (
	"context"
	"database/sql"

	"chalas.com/forum_project/API/internal/models"
)

type UserRepository interface {
	GetUser(ctx context.Context, username string) (*models.UserModel, error)
	CreateUser(ctx context.Context, user *models.UserModel) (int64, error)
	//GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*models.RefreshTokenModel, error)
	GetUserByID(ctx context.Context, userID int64) (*models.UserModel, error)
}
type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) GetUser(ctx context.Context, username string) (*models.UserModel, error) {
	query := `SELECT user_id, username, password, registration_date  FROM USERS WHERE username = ?`
	selectedRow := repo.db.QueryRowContext(ctx, query, username)
	var resultedUser models.UserModel
	err := selectedRow.Scan(&resultedUser.UserID, &resultedUser.Username, &resultedUser.Password, &resultedUser.RegistrationDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &resultedUser, nil
}

func (repo *userRepository) CreateUser(ctx context.Context, user *models.UserModel) (int64, error) {
	query := "INSERT INTO `USERS` (`username`, `password`, `registration_date`) VALUES (?, ?, ?)"

	insertedRow, err := repo.db.ExecContext(ctx, query, user.Username, user.Password, user.RegistrationDate)
	if err != nil {
		return 0, err
	}

	id, _ := insertedRow.LastInsertId()
	return id, nil
}

func (repo *userRepository) GetUserByID(ctx context.Context, userID int64) (*models.UserModel, error) {
	query := "SELECT user_id, username, password, registration_date FROM `USERS` WHERE user_id = ?"

	selectedRow := repo.db.QueryRowContext(ctx, query, userID)
	var resultedUser models.UserModel
	err := selectedRow.Scan(&resultedUser.UserID, &resultedUser.Username, &resultedUser.Password, &resultedUser.RegistrationDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &resultedUser, nil
}
