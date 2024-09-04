package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrNotFound

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Add a word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("word", "A definition")
	
		want := "A definition"
		got, err := dictionary.Search("word")
	
		assertError(t, err, nil)
		assertStrings(t, got, want)
	})
	t.Run("Word already added", func(t *testing.T) {
		dictionary := Dictionary{"This word": "Is already there"}
		err := dictionary.Add("This word", "Doesn't matter")
	
		assertError(t, err, ErrWordExists)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing word", func(t *testing.T) {
		dictionary := Dictionary{"word":"old definition"}
		newDef := "new definition"
		err := dictionary.Update("word", newDef)
		got, _ := dictionary.Search("word")
		assertError(t, err, nil)
		assertStrings(t, got, newDef)
	})
	t.Run("New word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update("word", "definition")
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete word", func(t *testing.T) {
		testWord := "word"
		dictionary := Dictionary{testWord:"definition"}
		err := dictionary.Delete(testWord)
		gotDef, gotErr := dictionary.Search(testWord)

		assertError(t, err, nil)
		assertError(t, gotErr, ErrNotFound)
		assertStrings(t, gotDef, "")
	})
	t.Run("doesn't xist", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Delete("doesn't exist")

		assertError(t, err, ErrWordDoesNotExist)

	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}