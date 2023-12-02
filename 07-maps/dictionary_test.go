package maps

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "just a test"}

		got, _ := dictionary.Search("test")
		want := "just a test"

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{}

		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "new"
		def := "new word"
		err := dictionary.Add(key, def)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, def)
	})

	t.Run("existing word", func(t *testing.T) {
		key := "key"
		def := "def 1"
		dictionary := Dictionary{key: def}
		err := dictionary.Add(key, "def 2")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, key, def)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		key := "key"
		def := "def 1"
		dictionary := Dictionary{key: def}
		newDef := "def 2"

		err := dictionary.Update(key, newDef)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, newDef)
	})

	t.Run("new word", func(t *testing.T) {
		key := "key"
		def := "def 2"
		dictionary := Dictionary{}

		err := dictionary.Update(key, def)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	key := "key"
	dictionary := Dictionary{key: "def"}

	dictionary.Delete(key)

	_, err := dictionary.Search(key)
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("Expected %q to be deleted", key)
	}
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if !errors.Is(got, want) {
		t.Errorf("got error %q, want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, key, def string) {
	t.Helper()
	got, err := dictionary.Search(key)

	if err != nil {
		t.Fatal("should find added word", err)
	}

	assertString(t, got, def)
}
