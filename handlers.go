package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *APIServer) handleTodo(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		createTodoReq := new(CreateTodoRequest)
		if err := json.NewDecoder(r.Body).Decode(createTodoReq); err != nil {
			return err
		}

		create := NewTodo(createTodoReq.Title)

		if err := s.store.Create(create); err != nil {
			return err
		}
		return WriteJSON(w, http.StatusCreated, create)
	}

	if r.Method == "GET" {
		todos, err := s.store.List()
		if err != nil {
			return err
		}
		return WriteJSON(w, http.StatusOK, todos)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleSingleTodo(w http.ResponseWriter, r *http.Request) error {
	stringID := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(stringID)
	if r.Method == "GET" {
		task, err := s.store.Get(id)
		if err != nil {
			return err
		}

		if task.ID != id {
			return WriteJSON(w, http.StatusNotFound, apiError{Error: "Not found"})
		}

		return WriteJSON(w, http.StatusOK, task)
	}

	if r.Method == "DELETE" {
		task, err := s.store.Get(id)
		if err != nil {
			return err
		}

		if task.ID != id {
			return WriteJSON(w, http.StatusNotFound, apiError{Error: "Not found"})
		}

		if err := s.store.Delete(id); err != nil {
			return err
		}
		return WriteJSON(w, http.StatusOK, map[string]string{"message": "Account deleted"})
	}

	if r.Method == "PUT" {
		task, err := s.store.Get(id)
		if err != nil {
			return err
		}

		updateTodoReq := new(UpdateTodoRequest)
		if err := json.NewDecoder(r.Body).Decode(updateTodoReq); err != nil {
			return err
		}

		if task.ID != id {
			return WriteJSON(w, http.StatusNotFound, apiError{Error: "Not found"})
		}

		if err := s.store.Update(id, updateTodoReq); err != nil {
			return err
		}

		return WriteJSON(w, http.StatusOK, map[string]string{"message": "Todo updated successfully"})

	}

	return fmt.Errorf("method not allowed %s", r.Method)
}
