package coverage

import (
	"errors"
	"os"
	"reflect"
	"strconv"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestLess(t *testing.T) {
	var people People

	people = append(people, Person{"Tom", "Tom", time.Time{}})
	people = append(people, Person{"Tom", "Tom", time.Time{}.Add(5 * time.Minute)})
	people = append(people, Person{"Jack", "Tom", time.Time{}})
	people = append(people, Person{"Jack", "Jack", time.Time{}})
	people = append(people, Person{"Jack", "Jack", time.Time{}})

	if people.Less(0, 1) {
		t.Errorf("Error. Less by Birthday")
	}

	if people.Less(0, 2) {
		t.Errorf("Error. Less by FirstName")
	}

	if people.Less(2, 3) {
		t.Errorf("Error. Less by LastName")
	}

	if people.Less(3, 4) {
		t.Errorf("Error. Less Eq")
	}
}

func TestSwap(t *testing.T) {
	var people People

	peopleOne := Person{"1", "1", time.Time{}}
	peopleTwo := Person{"1", "1", time.Time{}}

	people = append(people, peopleOne)
	people = append(people, peopleTwo)

	people.Swap(0, 1)

	if people[0] != peopleTwo || people[1] != peopleOne {
		t.Errorf("Error. Swap")
	}
}

func TestNew(t *testing.T) {
	actual, err := New("---")
	if actual != nil || !errors.Is(err, strconv.ErrSyntax) {
		t.Errorf("Error. String Error")
	}

	actual, err = New("1 1 \n 2")
	if actual != nil || err.Error() != "Rows need to be the same length" {
		t.Errorf("Error. Matrix Error")
	}

	actual, err = New("1 1 \n 2 3")
	expects := &Matrix{2, 2, []int{1, 1, 2, 3}}

	if actual.cols != expects.cols || actual.rows != expects.rows || !reflect.DeepEqual(actual.data, expects.data) {
		t.Errorf("Error. Empty Matrix")
	}
}

func TestRows(t *testing.T) {
	matrix := &Matrix{2, 2, []int{4, 5, 6, 7}}
	expects := [][]int{{4, 5}, {6, 7}}

	actual := matrix.Rows()

	if !reflect.DeepEqual(actual, expects) {
		t.Errorf("Error. Rows Matrix")
	}
}

func TestCols(t *testing.T) {
	matrix := &Matrix{3, 3, []int{1, 10, 100, 2, 20, 200, 3, 30, 300}}
	expects := [][]int{{1, 2, 3}, {10, 20, 30}, {100, 200, 300}}

	actual := matrix.Cols()

	if !reflect.DeepEqual(actual, expects) {
		t.Errorf("Error. Cols Matrix")
	}
}

func TestSet(t *testing.T) {
	matrix := &Matrix{3, 3, []int{1, 10, 100, 2, 20, 200, 3, 30, 300}}
	expectsData := []int{1, 10, 100, 2, 20, 200, 3, 30, 9999999999}

	actual := matrix.Set(2, 2, 9999999999)

	if !actual || !reflect.DeepEqual(matrix.data, expectsData) {
		t.Errorf("Error Set Matrix")
	}

	actual = matrix.Set(-1, 2, 9999999999)
	if actual {
		t.Errorf("Error. Set Error")
	}
}
