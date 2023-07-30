package user

import (
	"testing"

	"github.com/mohammadshabab/go-unit/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryStub struct {
	mock.Mock
}
type BadWordsRepositoryStub struct {
	mock.Mock
}

func (b *BadWordsRepositoryStub) FindAll() ([]string, error) {
	args := b.Called()

	return args.Get(0).([]string), args.Error(1)
}

func (r *UserRepositoryStub) Add(user entity.User) error {
	args := r.Called(user)
	return args.Error(0)
}
func TestShouldSuccessfullyRegisterAnUser(t *testing.T) {

	user := entity.User{
		Name:        "Shabab",
		Email:       "abc@xyz.com",
		Description: "Software Developer",
	}

	userRepository := &UserRepositoryStub{}
	userRepository.On("Add", user).Return(nil)
	badWordRepository := &BadWordsRepositoryStub{}
	badWordRepository.On("FindAll").Return([]string{"chicken", "egg"}, nil)
	userService := NewUserService(userRepository, badWordRepository)

	err := userService.Register(user)
	//in this we don't have back word so assert message should be called
	userRepository.AssertCalled(t, "Add", user)
	assert.Nil(t, err)

}

func TestShouldNotRegisterAnUserWhenBadWordFound(t *testing.T) {

	user := entity.User{
		Name:        "Shabab",
		Email:       "abc@xyz.com",
		Description: "Software chicken Developer",
	}

	userRepository := &UserRepositoryStub{}
	userRepository.On("Add", user).Return(nil)
	badWordRepository := &BadWordsRepositoryStub{}
	badWordRepository.On("FindAll").Return([]string{"chicken", "egg"}, nil)
	userService := NewUserService(userRepository, badWordRepository)

	err := userService.Register(user)
	//We need to use assertNotcalled because if someone changes the code and bring add method call in
	//begining then our test is not able to detect that error so we are ensuring that method not called before
	userRepository.AssertNotCalled(t, "Add", user)
	assert.Error(t, err)

}
