package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 6
	if got != want {
		t.Errorf("Add(2,3) = %d; want %d", got, want)
	}
}

func TestAddTable(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{0, 5, 5},
	}

	for _, tt := range tests {
		got := Add(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Add(%d,%d) = %d; want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

type DB interface {
	GetUser(id int) string
}

type StubDB struct{}

func (s StubDB) GetUser(id int) string {
	return "stub user"
}

func TestGetUser(t *testing.T) {
	var db DB = StubDB{}
	user := db.GetUser(1)
	if user != "stub user" {
		t.Errorf("got %s; want stub user", user)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Gopher!"))
}

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	HelloHandler(w, req)

	resp := w.Result()
	body := w.Body.String()

	if resp.StatusCode != http.StatusOK || body != "Hello, Gopher!" {
		t.Errorf("unexpected response: %d %s", resp.StatusCode, body)
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}
