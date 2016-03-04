package persistence

import (
	"errors"
	"github.com/deluan/gosonic/domain"
	"sort"
)

type mediaFileRepository struct {
	ledisRepository
}

func NewMediaFileRepository() domain.MediaFileRepository {
	r := &mediaFileRepository{}
	r.init("mediafile", &domain.MediaFile{})
	return r
}

func (r *mediaFileRepository) Put(m *domain.MediaFile) error {
	if m.Id == "" {
		return errors.New("MediaFile Id is not set")
	}
	return r.saveOrUpdate(m.Id, m)
}

func (r *mediaFileRepository) Get(id string) (*domain.MediaFile, error) {
	m, err := r.readEntity(id)
	if err != nil {
		return nil, err
	}
	mf := m.(*domain.MediaFile)
	if mf.Id != id {
		return nil, nil
	}
	return mf, nil
}

func (r *mediaFileRepository) FindByAlbum(albumId string) (domain.MediaFiles, error) {
	var mfs = make(domain.MediaFiles, 0)
	err := r.loadChildren("album", albumId, &mfs)
	sort.Sort(mfs)
	return mfs, err
}

var _ domain.MediaFileRepository = (*mediaFileRepository)(nil)
