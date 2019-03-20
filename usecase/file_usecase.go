package usecase

import (
	"collector/domain/entity"
	"collector/domain/repository"
)

type FileUsecase struct {
	repo repository.IFileRepository
}

// initializer
func NewFileUsecase(repo repository.IFileRepository) FileUsecase {
	FileUsecase := FileUsecase{
		repo: repo,
	}
	return FileUsecase
}

// Slackから新規ファイル
func (u *FileUsecase) AddFile(file *entity.File) (*entity.File, error) {
	res, err := u.repo.Create(file)
	if nil != err {
		return nil, err
	}

	return res, nil
}

// Slackからファイル情報更新
func (u *FileUsecase) UpdateFile(file *entity.File) (*entity.File, error) {
	res, err := u.repo.Update(file)
	if nil != err {
		return nil, err
	}

	return res, nil
}

// Slackからファイル削除
func (u *FileUsecase) DeleteFile(file *entity.File) (*entity.File, error) {
	res, err := u.repo.Update(file)
	if nil != err {
		return nil, err
	}

	return res, nil
}
