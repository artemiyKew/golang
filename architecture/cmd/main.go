package main

import (
	"architecture/cmd/search"
	"errors"

	"time"
)

func main() {
	// repo := &InMemoryUserRepository{}
	// user := &User{
	// 	Name:     "John",
	// 	Email:    "john@example.com",
	// 	Password: "secret",
	// }
	// repo.Create(user)
	// fmt.Println(repo.GetAll())
	// user.Name = "Jane"
	// repo.Update(user)
	// fmt.Println(repo.GetAll())
	// repo.Delete(user.ID)
	// fmt.Println(repo.GetAll())
	search.Search()
}

// repository.go
type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type UserInterface interface {
	GetByID(id int) (*User, error)
	GetAll() ([]*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id int) error
}

// in_memory.go

type InMemoryUserRepository struct {
	users []*User
}

func (r *InMemoryUserRepository) GetByID(id int) (*User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, errors.New("users not found")
}

func (r *InMemoryUserRepository) GetAll() ([]*User, error) {
	return r.users, nil
}

func (r *InMemoryUserRepository) Create(user *User) error {
	user.ID = len(r.users) + 1
	r.users = append(r.users, user)
	return nil
}

func (r *InMemoryUserRepository) Update(user *User) error {
	for i, u := range r.users {
		if u.ID == user.ID {
			r.users[i] = user
			return nil
		}
	}
	return errors.New("user not found")
}

func (r *InMemoryUserRepository) Delete(id int) error {
	for i, user := range r.users {
		if user.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func NewClient(host string, port uint16, setters ...Option) *Client {
	options := Options{
		timeout: 1 * time.Second,
		retry:   2,
	}

	for _, setter := range setters {
		setter(&options)
	}

	return &Client{
		host: host,
		port: port,
	}
}

type Client struct {
	host    string
	port    uint16
	timeOut time.Duration
	retry   int
}

type Options struct {
	timeout time.Duration
	retry   uint
}

type Option func(*Options)

func Timeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.timeout = timeout
	}
}

func Retry(retry uint) Option {
	return func(o *Options) {
		o.retry = retry
	}
}
