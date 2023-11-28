package usecase_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	usecase "github.com/piovani/go_full/domain/use_case"
	"github.com/piovani/go_full/dto"
	"github.com/piovani/go_full/infra/storage"
	mock_infra "github.com/piovani/go_full/infra/test/mock/infra"
	mock_repositories "github.com/piovani/go_full/infra/test/mock/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type CreateStudentSuite struct {
	suite.Suite
	*require.Assertions
	ctrl *gomock.Controller

	storage           mock_infra.MockStorageContract
	studentRepository mock_repositories.MockStudentRepository
	fileRepository    mock_repositories.MockFileRepository

	Usecase usecase.CreateStudentContract
}

func TestCreateStudentSuite(t *testing.T) {
	suite.Run(t, new(CreateStudentSuite))
}

func (c *CreateStudentSuite) TearDownTest() {
	defer c.ctrl.Finish()
}

func (c *CreateStudentSuite) SetupTest() {
	c.ctrl = gomock.NewController(c.T())
	storage := mock_infra.NewMockStorageContract(c.ctrl)
	studentRepository := mock_repositories.NewMockStudentRepository(c.ctrl)
	fileRepository := mock_repositories.NewMockFileRepository(c.ctrl)

	c.Usecase = usecase.NewCreateStudent(storage, studentRepository, fileRepository)
}

func (c *CreateStudentSuite) TestShoudErrUploadFile() {
	dto := c.getStudentInputDTO()
	errExpected := c.getErr()

	c.fileRepository.EXPECT().Save(&dto.Document).Return(errExpected).Times(1)

	student, err := c.Usecase.Execute(dto)

	assert.Nil(c.T(), student)
	assert.Error(c.T(), err, errExpected)
}

// AUX
func (c *CreateStudentSuite) getErr() error {
	return fmt.Errorf("error expected")
}

func (c *CreateStudentSuite) getStudentInputDTO() dto.StudentInput {
	file, _ := os.Open("./../../docker/s3/bucket-policy.json")
	defer file.Close()

	return dto.StudentInput{
		Document: *storage.NewFie("file", "application/pdf", file),
	}
}
