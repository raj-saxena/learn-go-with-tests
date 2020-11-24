package maps

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("dictionary", func(t *testing.T) {
		dictionary := Dictionary{"mykey": "myvalue"}

		got, _ := dictionary.Search("mykey")
		want := "myvalue"

		if got != want {
			t.Errorf("got %q, want %q given %q", got, want, dictionary)
		}
	})

	t.Run("Unknown work", func(t *testing.T) {
		dictionary := Dictionary{"mykey": "myvalue"}

		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("We were kinds expecting err")
		}

		if err != ErrNotFound {
			t.Errorf("Got %q, want %q", err, ErrNotFound)
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		d := Dictionary{}
		err := d.Add("mykey", "myvalue")

		if err != nil {
			t.Fatal("Pair should be added")
		}

		got, err := d.Search("mykey")
		want := "myvalue"

		if got != want {
			t.Errorf("Got %q, want %q", got, want)
		}
	})

	t.Run("Err on existing", func(t *testing.T) {
		d := Dictionary{"mykey": "myvalue"}

		err := d.Add("mykey", "myvalue")

		if err == nil {
			t.Fatal("Expecting error")
		}

		if err != ErrKeyExists {
			t.Errorf("Got %q, want %q", err, ErrKeyExists)
		}
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update", func(t *testing.T) {
		word := "someWord"
		d := Dictionary{word: "definition"}

		newDefinition := "new definition"
		d.Update(word, newDefinition)

		got, _ := d.Search(word)

		if got != newDefinition {
			t.Errorf("got %q, want %q", got, newDefinition)
		}
	})

	t.Run("Err on missing", func(t *testing.T) {
		d := Dictionary{}
		err := d.Update("not", "existing")

		if err == nil {
			t.Fatal("Expecting error")
		}

		if err != ErrNotFound {
			t.Errorf("got % q, want %q", err, ErrNotFound)
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete", func(t *testing.T) {
		word := "someWord"
		d := Dictionary{word: "definition"}

		d.Delete(word)

		_, err := d.Search(word)

		if err != ErrNotFound {
			t.Errorf("word %q should have been deleted, given: %v", word, d)
		}
	})
}
