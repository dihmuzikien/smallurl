package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/dihmuzikien/smallurl/domain"
	"time"
)

type urlUsecase struct {
	articleRepo domain.UrlRepository
}

func NewUrlUseCase(r domain.UrlRepository) domain.UrlUseCase {
	return &urlUsecase{
		articleRepo: r,
	}
}

func (u *urlUsecase) Create(ctx context.Context, destination string) (domain.Url, error) {
	id, iErr := makeId(destination)
	if iErr != nil {
		return domain.Url{}, iErr
	}
	url := domain.Url{
		ID:          id,
		Destination: destination,
		Created:     time.Now(),
	}
	err := u.articleRepo.Put(ctx, url)
	if err != nil {
		return domain.Url{}, err
	}
	return url, nil
}

func (u *urlUsecase) CreateWithAlias(ctx context.Context, id, destination string) (domain.Url, error) {
	url := domain.Url{
		ID:          id,
		Destination: destination,
		Created:     time.Now(),
	}
	err := u.articleRepo.Put(ctx, url)
	if err != nil {
		return domain.Url{}, err
	}
	return url, nil
}

func (u *urlUsecase) GetById(ctx context.Context, id string) (domain.Url, error) {
	res, err := u.articleRepo.Get(ctx, id)
	if err != nil {
		return domain.Url{}, err
	}
	return res, nil
}

func (u *urlUsecase) List(ctx context.Context) ([]domain.Url, error) {
	result, err := u.articleRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *urlUsecase) Delete(ctx context.Context, id string) error {
	err := u.articleRepo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete %w", err)
	}
	return nil
}

func makeId(url string) (string, error) {
	h := sha1.New()
	_, err := h.Write([]byte(url))
	if err != nil {
		return "", err
	}
	hashed := h.Sum(nil)
	return string(hashed), nil
}
