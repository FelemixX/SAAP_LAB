package lab6_test

import (
	"os"
	"testing"

	"lab6"
)

func TestTask1(t *testing.T) {
	dir := "test_dir"
	expectedFile := "test_dir/file2.txt"
	expectedSize := int64(500)

	err := os.Mkdir(dir, 0755)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	file1, err := os.Create(dir + "/file1.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file1.Close()

	file2, err := os.Create(expectedFile)
	if err != nil {
		t.Fatal(err)
	}
	defer file2.Close()

	file3, err := os.Create(dir + "/file3.txt")
	if err != nil {
		t.Fatal(err)
	}

	defer file3.Close()

	file1.Write(make([]byte, 100))
	file2.Write(make([]byte, 500))
	file3.Write(make([]byte, 2))

	maxFile, maxSize := lab6.Task1(dir)

	if maxFile != expectedFile {
		t.Errorf("Expected maxFile to be %s, but got %s", expectedFile, maxFile)
	}

	if maxSize != expectedSize {
		t.Errorf("Expected maxSize to be %d, but got %d", expectedSize, maxSize)
	}
}

func TestTask8(t *testing.T) {
	dir := "test_dir"
	expectedResult := []string{"file3.kek"}
	extension := "kek"

	err := os.Mkdir(dir, 0755)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	file1, err := os.Create(dir + "/file1.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer file1.Close()

	file2, err := os.Create(dir + "/file2.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file2.Close()

	file3, err := os.Create(dir + "/file3.kek")
	if err != nil {
		t.Fatal(err)
	}

	defer file3.Close()

	result := lab6.Task8(dir, extension)

	if result[0] != expectedResult[0] {
		t.Errorf("Expected result to be %s, but got %s", expectedResult, result)
	}
}
