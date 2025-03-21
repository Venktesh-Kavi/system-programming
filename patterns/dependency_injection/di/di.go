package di

import (
	"fmt"
	"net/http"
	"time"
)

// Logger interface to make any logger substitutable & make implementations decoupled.
type Logger interface {
	Log(msg string)
}

// A type to implement the LoggerInterface. No state is required, so func type.
type LoggerAdapter func(msg string)

// Log log function should just call the loggerAdapter method
func (l LoggerAdapter) Log(msg string) {
	l(msg)
}

type DataStore interface {
	UserNameById(string) (string, bool)
}

type Logic interface {
	SayHello(msg string) (string, error)
}

type Controller struct {
	l     Logger
	logic Logic
}

func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}

// to adhere to DataStore define a stateful type SimpleDataSource.
type SimpleDataStore struct {
	store map[string]string
}

// LogOutput function which has the same format as logAdapter.
func LogOutput(msg string) {
	fmt.Printf("%v::info::msg -- %s\n", time.Now(), msg)
}

func (sd SimpleDataStore) UserNameById(id string) (string, bool) {
	name, ok := sd.store[id]
	return name, ok
}

func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		store: map[string]string{
			"1": "foo",
			"2": "bar",
			"4": "baz",
		},
	}
}

type SimpleLogic struct {
	ds DataStore
	l  Logger
}

func NewSimpleLogic(ds DataStore, l Logger) SimpleLogic {
	return SimpleLogic{
		ds: ds,
		l:  l,
	}
}

func UserNotFoundError(userID string) error {
	return fmt.Errorf("user not found, user:%s", userID)
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("received request for say hello")
	userId := r.URL.Query().Get("user_id")
	msg, err := c.logic.SayHello(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(msg))
}

func (s SimpleLogic) SayHello(userId string) (string, error) {
	s.l.Log("starting say hello method")
	name, ok := s.ds.UserNameById(userId)
	if !ok {
		return "", UserNotFoundError(userId)
	}
	return fmt.Sprintf("Hello user:%s, id:%s\n", name, userId), nil
}

func (s SimpleLogic) SayBye(userId string) error {
	s.l.Log("starting bye method")
	name, ok := s.ds.UserNameById(userId)
	if !ok {
		return UserNotFoundError(name)
	}
	fmt.Printf("bye to user:%s, id:%s\n", name, userId)
	return nil
}
