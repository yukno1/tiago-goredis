package main

type Store struct {
	data map[string]string
}

func (s *Store) Get(key string) (string, bool) {
	return "", false
}

func (s *Store) Set(key, val string) {

}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func main() {}
