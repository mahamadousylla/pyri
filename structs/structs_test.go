package structs

import (
	"fmt"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Test ...
type Test struct {
	ID         uuid.UUID `db:"id" json:"ID"`
	This       string    `db:"this" json:"This"`
	Is         string    `db:"is" json:"is"`
	Used       int32     `db:"used" json:"used"`
	For        float32   `db:"for" json:"for"`
	Testing    int64     `db:"testing" json:"testing"`
	CreateDate time.Time `db:"create_date" json:"create_date"`
	UpdateDate time.Time `db:"update_date" json:"update_date"`
}

func TestGetTags(t *testing.T) {
	got := GetTags(&Test{}, "db", nil)
	expected := []string{"id", "this", "is", "used", "for", "testing", "create_date", "update_date"}

	if len(got) != len(expected) {
		t.Fatalf("GetTags() = %v, want %v", got, expected)
	}

	for i := range got {
		if got[i] != expected[i] {
			t.Fatalf("GetTags() = %v, want %v", got, expected)
		}
	}
}

func TestGetTagsIgnoreFields(t *testing.T) {
	got := GetTags(&Test{}, "json", []string{"update_date", "used"})
	expected := []string{"ID", "This", "is", "for", "testing", "create_date"}

	if len(got) != len(expected) {
		t.Fatalf("GetTags() = %v, want %v", got, expected)
	}

	for i := range got {
		if got[i] != expected[i] {
			t.Fatalf("GetTags() = %v, want %v", got, expected)
		}
	}
}

func TestGetTagsAndValues(t *testing.T) {
	id := uuid.NewV4()
	now := time.Now()

	expectedTags := []string{"ID", "This", "is", "used", "for", "testing", "create_date", "update_date"}
	var expectedValues []interface{}
	expectedValues = append(expectedValues, id, "yo", "are", int32(1), float32(2.0), int64(123456789), now, now)

	gotTags, gotValues := GetTagsAndValues(createStruct(id, now), "json", nil)
	if (len(gotTags) != len(expectedTags)) || (len(gotValues) != len(expectedValues)) {
		t.Fatalf("GetTagsAndValues() = %v, %v, want %v %v", gotTags, gotValues, expectedTags, expectedValues)
	}

	for i := range gotTags {
		if (gotTags[i] != expectedTags[i]) || (gotValues[i] != expectedValues[i]) {
			fmt.Println(gotTags[i], expectedTags[i], gotValues[i], expectedValues[i])
			t.Fatalf("GetTagsAndValues() = %v, %v, want %v %v", gotTags, gotValues, expectedTags, expectedValues)
		}
	}
}

func TestGetTagsAndValuesForJSONIgnoreFields(t *testing.T) {
	id := uuid.NewV4()
	now := time.Now()

	expectedTags := []string{"ID", "This", "is", "used", "for", "testing", "create_date"}
	var expectedValues []interface{}
	expectedValues = append(expectedValues, id, "yo", "are", int32(1), float32(2.0), int64(123456789), now)

	gotTags, gotValues := GetTagsAndValues(createStruct(id, now), "json", []string{"update_date"})
	if (len(gotTags) != len(expectedTags)) || (len(gotValues) != len(expectedValues)) {
		t.Fatalf("GetTagsAndValues() = %v, %v, want %v %v", gotTags, gotValues, expectedTags, expectedValues)
	}

	for i := range gotTags {
		if (gotTags[i] != expectedTags[i]) || (gotValues[i] != expectedValues[i]) {
			fmt.Println(gotTags[i], expectedTags[i], gotValues[i], expectedValues[i])
			t.Fatalf("GetTagsAndValues() = %v, %v, want %v %v", gotTags, gotValues, expectedTags, expectedValues)
		}
	}
}

func createStruct(id uuid.UUID, now time.Time) *Test {
	return &Test{
		ID:         id,
		This:       "yo",
		Is:         "are",
		Used:       1,
		For:        2.0,
		Testing:    123456789,
		CreateDate: now,
		UpdateDate: now,
	}
}
