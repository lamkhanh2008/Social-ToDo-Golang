package mysql

import (
	"context"
	"log"
	core "social_todo/internal/common"
	"social_todo/internal/services/items/entity"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type mysqlRepo struct {
// 	db *gorm.DB
// }

// func (repo *mysqlRepo) CreateItem(ctx context.Context, data *entity.TodoItemCreation) error {

// 	if err := repo.db.Create(data).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }

func TestCreateItem(t *testing.T) {
	// Mocking DB với go-sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	// Khởi tạo repo với mock DB
	repo := &mysqlRepo{
		db: gormDB,
	}

	// Dữ liệu cần thêm
	data := &entity.TodoItemCreation{
		Title:       "Test Item",
		Description: "Description of test item",
		UserID:      10,
		Image: &core.Images{
			core.Image{
				ID: 12,
			},
		},
	}

	// Xác định mock của chúng ta
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `todo_items`").
		WithArgs(data.UserID, data.Title, data.Description, data.Image).
		WillReturnResult(sqlmock.NewResult(1, 1)) // ID = 1, affected rows = 1
	mock.ExpectCommit()

	// Test CreateItem
	err = repo.CreateItem(context.Background(), data)

	// Kiểm tra lỗi
	assert.NoError(t, err)

	// Kiểm tra rằng ID đã được gán sau khi create
	assert.Equal(t, int(1), data.ID, "Expected ID to be 1")

	// Kiểm tra xem mock có bị gọi đúng không
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
