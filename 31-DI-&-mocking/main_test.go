package main

import (
	"errors"
	"testing"
)

type User struct {
	Name string
}

type MockUserRepository struct {
	Err       error
	WasCalled bool
}

func (m *MockUserRepository) InsertUser(user User) error {
	m.WasCalled = true
	return m.Err
}

type UserRepository interface {
	InsertUser(user User) error
}

func CreateUser(
	user User,
	repo UserRepository,
) error {

	if user.Name == "" {
		return errors.New("name required")
	}

	return repo.InsertUser(user)
}

func TestCreateUser(t *testing.T) {

	repo := &MockUserRepository{
		Err: nil,
	}

	err := CreateUser(
		User{Name: "Vatsal"},
		repo,
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCreateUser_DatabaseError(t *testing.T) {

	repo := &MockUserRepository{
		Err: errors.New("Database Down"),
	}

	err := CreateUser(User{Name: "Vatsal"}, repo)

	if err == nil {
		t.Fatalf("expected an error but got nil")
	}
}

func TestCreateUser_InvalidName(t *testing.T) {

	repo := &MockUserRepository{}

	err := CreateUser(
		User{Name: ""},
		repo,
	)

	if err == nil {
		t.Fatal("expected validation error")
	}

	if repo.WasCalled {
		t.Fatal(
			"repository should not be called",
		)
	}
}
