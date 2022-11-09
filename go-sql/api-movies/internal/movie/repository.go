package movie

import (
	"backpack-bcgow6-matias-delaserna/go-sql/api-movies/internal/domain"
	isql "backpack-bcgow6-matias-delaserna/go-sql/api-movies/pkg/sql"
	"context"
	"database/sql"
	"errors"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Movie, error)
	Get(ctx context.Context, id int) (domain.Movie, error)
	GetByTitle(ctx context.Context, title string) (domain.Movie, error)
	Save(ctx context.Context, b domain.Movie) (int64, error)
	Exists(ctx context.Context, id int) bool
	Update(ctx context.Context, b domain.Movie, id int) error
	Delete(ctx context.Context, id int64) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Exists(ctx context.Context, id int) bool {
	rows := r.db.QueryRow(isql.EXIST_MOVIE, id)
	err := rows.Scan(&id)
	return err == nil
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Movie, error) {
	var movies []domain.Movie
	rows, err := r.db.Query(isql.GET_ALL_MOVIE)
	if err != nil {
		return []domain.Movie{}, err
	}

	for rows.Next() {
		var movie domain.Movie
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Rating, &movie.Awards, &movie.Length, &movie.Genre_id)
		if err != nil {
			return []domain.Movie{}, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func (r *repository) GetMovieWithContext(ctx context.Context, id int) (domain.Movie, error) {
	row := r.db.QueryRowContext(ctx, isql.GET_MOVIE, id)

	var movie domain.Movie
	err := row.Scan(&movie.ID, &movie.Title, &movie.Rating, &movie.Awards, &movie.Length, &movie.Genre_id)
	if err != nil {
		return domain.Movie{}, err
	}
	return movie, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Movie, error) {
	row := r.db.QueryRow(isql.GET_MOVIE, id)
	var movie domain.Movie
	if err := row.Scan(&movie.ID, &movie.Title, &movie.Rating, &movie.Awards, &movie.Length, &movie.Genre_id); err != nil {
		return domain.Movie{}, err
	}
	return movie, nil
}

func (r *repository) GetByTitle(ctx context.Context, title string) (domain.Movie, error) {
	row := r.db.QueryRow(isql.GET_BY_TITLE, title)
	var movie domain.Movie
	if err := row.Scan(&movie.ID, &movie.Title, &movie.Rating, &movie.Awards, &movie.Length, &movie.Genre_id); err != nil {
		return domain.Movie{}, err
	}
	return movie, nil
}

func (r *repository) Save(ctx context.Context, m domain.Movie) (int64, error) {
	stm, err := r.db.Prepare(isql.SAVE_MOVIE)
	if err != nil {
		return 0, err
	}

	result, err := stm.Exec(m.Title, m.Rating, m.Awards, m.Length, m.Genre_id)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) Update(ctx context.Context, m domain.Movie, id int) error {
	stm, err := r.db.Prepare(isql.UPDATE_MOVIE)
	if err != nil {
		return err
	}
	defer stm.Close() //cerramos para no perder memoria

	result, err := stm.Exec(m.Title, m.Rating, m.Awards, m.Length, m.Genre_id, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected < 1 {
		return errors.New("error: no affected rows")
	}
	return nil
}

func (r *repository) Delete(ctx context.Context, id int64) error {
	stm, err := r.db.Prepare(isql.DELETE_MOVIE)
	if err != nil {
		return err
	}

	defer stm.Close()
	result, err := stm.Exec(id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected < 1 {
		return errors.New("error: no affected rows")
	}

	return nil
}
