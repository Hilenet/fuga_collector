package controller

import (
	"collector/domain/entity"
	"collector/domain/repository"
	"collector/usecase"
)

// DAOをいい感じにしてusecaseに投げる
type FileController struct {
	Usecase usecase.IFileUsecase
}

// initializer
func NewFileController(repo repository.IFileRepository) FileController {
	usecase := usecase.NewFileUsecase(repo)
	FileController := FileController{
		Usecase: &usecase,
	}
	return FileController
}

// plain File

// FileChanged
func (c *FileController) HandleFileChange(File entity.File) error {
	_, err := c.Usecase.UpdateFile(&File)
	if err != nil {
		return err
	}

	return nil
}

func (c *FileController) HandleFileDelete(File entity.File) error {
	_, err := c.Usecase.DeleteFile(&File)
	if err != nil {
		return err
	}

	return nil
}
