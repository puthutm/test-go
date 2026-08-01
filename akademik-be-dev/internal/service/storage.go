package service

import (
	"context"
	"errors"
	"mime/multipart"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
	"unsia.ac.id/akademic_be/internal/config"
	"unsia.ac.id/akademic_be/pkg/utils"
	"unsia.ac.id/akademic_be/pkg/validation"
)

type StorageService struct {
	cnf         *config.Config
	minioClient *minio.Client
	log         *logrus.Logger
}

func NewStorageService(cnf *config.Config, minioClient *minio.Client, log *logrus.Logger) *StorageService {
	return &StorageService{
		cnf:         cnf,
		minioClient: minioClient,
		log:         log,
	}
}

func (s *StorageService) UploadFile(ctx context.Context, file *multipart.FileHeader, subfolder string, validateExtension []string, maxSizeMb int64) (string, error) {
	bucketName := s.cnf.Minio.BucketName
	if maxSizeMb == 0 {
		maxSizeMb = 5
	}
	size := int64(1024*1024) * maxSizeMb
	if file.Size > size {
		s.log.Warnln("file size to large")
		return "", fiber.NewError(fiber.StatusBadRequest, "file size is to large")
	}
	// Validasi ekstensi file
	if !validation.IsValidFileExtension(file.Filename, validateExtension) {
		s.log.Warn("invalid extension")
		return "", fiber.NewError(fiber.StatusBadRequest, "invalid file extension")
	}
	exists, err := s.minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		s.log.Errorf("error gettting bucket %+v", err)
		return "", fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if !exists {
		err := s.minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			s.log.Errorf("error Create bucket %+v", err)
			return "", fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
	}

	src, err := file.Open()
	if err != nil {
		s.log.Errorf("error open file %+v", err)
		return "", fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer src.Close()

	var objectName string

	if subfolder != "" {
		objectName = subfolder + "/"
	}

	// Mengunggah file ke MinIO
	objectName += utils.GeneratorRandomString(20) + "_fn_" + strings.ReplaceAll(strings.TrimSpace(file.Filename), " ", "")
	_, err = s.minioClient.PutObject(ctx, bucketName, objectName, src, file.Size, minio.PutObjectOptions{})
	if err != nil {
		s.log.Errorf("error upload image or file in %s %+v", bucketName, err)
		return "", fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	stringUrl := "/api/objects/" + bucketName + "/" + objectName
	return stringUrl, nil
}

func (s *StorageService) UploadFileV2(
	ctx context.Context,
	file *multipart.FileHeader,
	subfolder string, validateExtension []string,
	maxSizeMb int64,
) (string, error) {
	bucketName := s.cnf.Minio.BucketName
	if maxSizeMb == 0 {
		maxSizeMb = int64(s.cnf.Minio.MaxSizeFile)
	}
	size := int64(1024*1024) * maxSizeMb
	if file.Size > size {
		s.log.Warnln("file size to large")
		return "", fiber.NewError(fiber.StatusBadRequest, "file size is to large")
	}
	// Validasi ekstensi file
	if !validation.IsValidFileExtension(file.Filename, validateExtension) {
		s.log.Warn("invalid extension")
		return "", fiber.NewError(fiber.StatusBadRequest, "invalid file extension")
	}
	exists, err := s.minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		s.log.Errorf("error gettting bucket %+v", err)
		return "", fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if !exists {
		err := s.minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			s.log.Errorf("error Create bucket %+v", err)
			return "", fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
	}

	src, err := file.Open()
	if err != nil {
		s.log.Errorf("error open file %+v", err)
		return "", fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer src.Close()

	var objectName string

	if subfolder != "" {
		objectName = subfolder + "/"
	}

	// Mengunggah file ke MinIO
	objectName += utils.GeneratorRandomString(20) + "_fn_" + strings.ReplaceAll(strings.TrimSpace(file.Filename), " ", "")
	_, err = s.minioClient.PutObject(ctx, bucketName, objectName, src, file.Size, minio.PutObjectOptions{})
	if err != nil {
		s.log.Errorf("error upload image or file in %s %+v", bucketName, err)
		return "", fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return bucketName + "/" + objectName, nil
}

func (s *StorageService) UploadFileV3(
	ctx context.Context,
	file *multipart.FileHeader,
	visibility bool,
	subfolder, alias string,
	validateExtension []string,
	maxSizeMb int64,
) (string, error) {
	bucketName := s.cnf.Minio.BucketName
	if maxSizeMb == 0 {
		maxSizeMb = int64(s.cnf.Minio.MaxSizeFile)
	}
	size := int64(1024*1024) * maxSizeMb
	if file.Size > size {
		s.log.Warnln("file size to large")
		return "", fiber.NewError(fiber.StatusBadRequest, "file size is to large")
	}
	// Validasi ekstensi file
	if !validation.IsValidFileExtension(file.Filename, validateExtension) {
		s.log.Warn("invalid extension")
		return "", fiber.NewError(fiber.StatusBadRequest, "invalid file extension")
	}
	exists, err := s.minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		s.log.Errorf("error gettting bucket %+v", err)
		return "", fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if !exists {
		err := s.minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			s.log.Errorf("error Create bucket %+v", err)
			return "", fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
	}

	src, err := file.Open()
	if err != nil {
		s.log.Errorf("error open file %+v", err)
		return "", fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer src.Close()

	var objectName string

	if visibility {
		objectName = "private/"
	} else {
		objectName = "public/"
	}

	if subfolder != "" {
		objectName += subfolder + "/"
	}

	// Mengunggah file ke MinIO
	objectName += utils.GeneratorRandomString(20) + "_fn_" + strings.ReplaceAll(strings.TrimSpace(file.Filename), " ", "")
	if alias != "" {
		objectName += "_as_" + alias
	}
	_, err = s.minioClient.PutObject(ctx, bucketName, objectName, src, file.Size, minio.PutObjectOptions{})
	if err != nil {
		s.log.Errorf("error upload image or file in %s %+v", bucketName, err)
		return "", fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return bucketName + "/" + objectName, nil
	// bucketname / visibility / subfolder / nama file / alias (opsional)
}

func (s *StorageService) GetObject(ctx context.Context, bucketName, objectName string) (*minio.Object, error) {
	exists, err := s.minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("bucket does not exist")
	}

	minioObject, err := s.minioClient.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return minioObject, nil
}

func GetOriginalFileName(uniqueFileName string) string {
	parts := strings.Split(uniqueFileName, "_fn_")

	if len(parts) > 1 {
		return parts[len(parts)-1]
	}

	return uniqueFileName
}
