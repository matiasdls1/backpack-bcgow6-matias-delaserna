package movie

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/matiasdls1/backpack-bcgow6-matias-delaserna/go-sql/api-movies/internal/domain"
	"github.com/matiasdls1/backpack-bcgow6-matias-delaserna/go-sql/api-movies/pkg/sql"
	"github.com/stretchr/testify/assert"
)

var (
	ERRORFORZADO = errors.New("Error forzado")
)

var movie_test = domain.Movie{
	ID:           1,
	Created_at:   time.Now(),
	Updated_at:   time.Now(),
	Title:        "Cars 1",
	Rating:       4,
	Awards:       2,
	Release_date: time.Now(),
	Length:       0,
	Genre_id:     0,
}

func TestSave(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("Store Ok", func(t *testing.T) {

		mock.ExpectPrepare(regexp.QuoteMeta(sql.SAVE_MOVIE))
		mock.ExpectExec(regexp.QuoteMeta(sql.SAVE_MOVIE)).WillReturnResult(sqlmock.NewResult(1, 1))

		columns := []string{"id", "title", "rating", "awards", "length", "genre_id"}
		rows := sqlmock.NewRows(columns)
		rows.AddRow(movie_test.ID, movie_test.Title, movie_test.Rating, movie_test.Awards, movie_test.Length, movie_test.Genre_id)
		mock.ExpectQuery(regexp.QuoteMeta(sql.GET_MOVIE)).WithArgs(1).WillReturnRows(rows)

		repository := NewRepository(db)
		ctx := context.TODO()

		newID, err := repository.Save(ctx, movie_test)
		assert.NoError(t, err)

		movieResult, err := repository.Get(ctx, int(newID))
		assert.NoError(t, err)

		assert.NotNil(t, movieResult)
		assert.Equal(t, movie_test.ID, movieResult.ID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Store Fail", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectPrepare(regexp.QuoteMeta(sql.SAVE_MOVIE))
		mock.ExpectExec(regexp.QuoteMeta(sql.SAVE_MOVIE)).WillReturnError(ERRORFORZADO)

		repository := NewRepository(db)
		ctx := context.TODO()

		id, err := repository.Save(ctx, movie_test)

		assert.EqualError(t, err, ERRORFORZADO.Error())
		assert.Empty(t, id)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func Test_RepositoryGetAllOK(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	// Columns
	columns := []string{"id", "title", "rating", "awards", "length", "genre_id"}
	rows := sqlmock.NewRows(columns)
	movies := []domain.Movie{{ID: 1, Title: "Avatar", Rating: 22, Awards: 99, Length: 0, Genre_id: 1}, {ID: 2, Title: "Simpson", Rating: 33, Awards: 11, Length: 2, Genre_id: 2}}

	for _, movie := range movies {
		rows.AddRow(movie.ID, movie.Title, movie.Rating, movie.Awards, movie.Length, movie.Genre_id)
	}

	mock.ExpectQuery(regexp.QuoteMeta(sql.GET_ALL_MOVIE)).WillReturnRows(rows)

	repo := NewRepository(db)
	resultMovies, err := repo.GetAll(context.TODO())

	assert.NoError(t, err)
	assert.Equal(t, movies, resultMovies)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_RepositoryGetAllFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	// Columns
	columns := []string{"id", "title", "rating", "awards", "length", "genre_id"}
	rows := sqlmock.NewRows(columns)
	movies := []domain.Movie{{ID: 1, Title: "Avatar", Rating: 22, Awards: 99, Length: 0, Genre_id: 1}, {ID: 2, Title: "Simpson", Rating: 33, Awards: 11, Length: 2, Genre_id: 2}}

	for _, movie := range movies {
		rows.AddRow(movie.ID, movie.Title, movie.Rating, movie.Awards, movie.Length, movie.Genre_id)
	}

	mock.ExpectQuery(regexp.QuoteMeta(sql.GET_ALL_MOVIE)).WillReturnError(ERRORFORZADO)

	repo := NewRepository(db)
	resultMovies, err := repo.GetAll(context.TODO())

	assert.EqualError(t, err, ERRORFORZADO.Error())
	assert.Empty(t, resultMovies)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepositoryUpdateOk(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("Update Ok", func(t *testing.T) {
		mock.ExpectPrepare(regexp.QuoteMeta(sql.UPDATE_MOVIE))
		mock.ExpectExec(regexp.QuoteMeta(sql.UPDATE_MOVIE)).WithArgs(movie_test.ID, movie_test.Title, movie_test.Rating, movie_test.Awards, movie_test.Length, movie_test.Genre_id).WillReturnResult(sqlmock.NewResult(0, 1))

		columns := []string{"id", "title", "rating", "awards", "length", "genre_id"}
		rows := sqlmock.NewRows(columns)
		rows.AddRow(movie_test.ID, movie_test.Title, movie_test.Rating, movie_test.Awards, movie_test.Length, movie_test.Genre_id)
		mock.ExpectQuery(regexp.QuoteMeta(sql.GET_MOVIE)).WithArgs(1).WillReturnRows(rows)

		repository := NewRepository(db)
		ctx := context.TODO()

		movie_update := domain.Movie{
			ID:    1,
			Title: "Cars 2",
		}

		err := repository.Update(ctx, movie_update, movie_test.ID)
		assert.NoError(t, err)

		movieResult, err := repository.Get(ctx, movie_test.ID)
		assert.NoError(t, err)
		movie_test.Title = "Cars 2"
		assert.Equal(t, movie_test, movieResult)

	})
}
