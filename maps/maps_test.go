package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{
		"test": "this is just a test",
	}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertEquals(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknownKey")
		want := ErrKeyNotFound

		if err == nil {
			t.Fatal("expected to get an error.")
		}
		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{
			"testKey": "testValue",
		}

		//when
		err := dictionary.Add("newKey", "newValue")
		if err != nil {
			t.Fatal(err)
		}

		//Then
		got, err := dictionary.Search("newKey")
		want := "newValue"
		if err != nil {
			t.Fatal(err)
		}

		assertEquals(t, got, want)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		defination := "this is just a test"
		dictionary := Dictionary{word: defination}
		err := dictionary.Add(word, "new test")
		if err == nil {
			t.Fatal("no error found, expected error", err)
		}

		assertError(t, err, ErrKeyExists)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing key", func(t *testing.T) {
		key := "testKey"
		val := "testVal"
		newVal := "newVal"

		dict := Dictionary{
			key: val,
		}

		//When
		err := dict.Update(key, newVal)
		if err != nil {
			t.Fatal(err)
		}
		//Then
		got, err := dict.Search(key)
		if err != nil {
			t.Fatal(err)
		}
		assertEquals(t, got, newVal)
	})

	t.Run("try to update non existing key", func(t *testing.T) {
		key := "testKey"
		val := "testVal"
		newKey := "newNonExistingKey"
		newVal := "newVal"

		dict := Dictionary{
			key: val,
		}

		//When
		err := dict.Update(newKey, newVal)
		if err == nil {
			t.Fatal(err)
		}

		assertError(t, err, ErrKeyNotFound)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete existing word", func(t *testing.T) {
		testKey := "testKey"
		testVal := "testVal"
		dict := Dictionary{
			testKey: testVal,
		}

		//When
		if err := dict.Delete(testKey); err != nil {
			t.Fatal(err)
		}

		//Then
		_, err := dict.Search(testKey)
		assertError(t, err, ErrKeyNotFound)
	})

	t.Run("delete non existing word", func(t *testing.T) {
		testKey := "testKey"
		testVal := "testVal"
		newTestKey := "newTestKey"
		dict := Dictionary{
			testKey: testVal,
		}

		//When
		err := dict.Delete(newTestKey)
		if err == nil {
			t.Fatal(err)
		}

		//Then
		assertError(t, err, ErrKeyNotFound)
	})
}

func assertEquals(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q :: want %q", got.Error(), want.Error())
	}
}
