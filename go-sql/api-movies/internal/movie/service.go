package movie

import (
	"context"
	"errors"

	"github.com/matiasdls1/backpack-bcgow6-matias-delaserna/go-sql/api-movies/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Movie, error)
	Get(ctx context.Context, id int) (domain.Movie, error)
	GetByTitle(ctx context.Context, title string) (domain.Movie, error)
	Save(ctx context.Context, b domain.Movie) (domain.Movie, error)
	Update(ctx context.Context, b domain.Movie, id int) (domain.Movie, error)
	Delete(ctx context.Context, id int64) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Movie, error) {
	movies, err := s.repo.GetAll(ctx)
	if err != nil {
		return []domain.Movie{}, err
	}
	return movies, err
}

func (s *service) Get(ctx context.Context, id int) (movie domain.Movie, err error) {
	movie, err = s.repo.Get(ctx, id)
	if err != nil {
		return domain.Movie{}, err
	}
	return movie, nil
}

func (s *service) GetByTitle(ctx context.Context, title string) (movie domain.Movie, err error) {
	movie, err = s.repo.GetByTitle(ctx, title)
	if err != nil {
		return domain.Movie{}, err
	}
	return movie, nil
}

func (s *service) Save(ctx context.Context, m domain.Movie) (domain.Movie, error) {
	if s.repo.Exists(ctx, m.ID) {
		return domain.Movie{}, errors.New("error: movie id already exists")
	}
	movie_id, err := s.repo.Save(ctx, m)
	if err != nil {
		return domain.Movie{}, err
	}

	m.ID = int(movie_id)
	return m, nil
}

func (s *service) Update(ctx context.Context, b domain.Movie, id int) (domain.Movie, error) {

	err := s.repo.Update(ctx, b, id)
	if err != nil {
		return domain.Movie{}, err
	}
	updated, err := s.repo.Get(ctx, id)
	if err != nil {
		return b, err
	}
	return updated, nil
}

func (s *service) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
