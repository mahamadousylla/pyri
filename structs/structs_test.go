package structs

import (
	"fmt"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
)

// ForTesting ...
type ForTesting struct {
	ID         uuid.UUID `db:"id" json:"ID"`
	This       string    `db:"this" json:"This"`
	Is         string    `db:"is" json:"is"`
	Used       int32     `db:"used" json:"used"`
	For        float32   `db:"for" json:"for"`
	Testing    int64     `db:"testing" json:"testing"`
	CreateDate time.Time `db:"create_date" json:"create_date"`
	UpdateDate time.Time `db:"update_date" json:"update_date"`
}

// NestedStructForTesting ...
type NestedStructForTesting struct {
	ID      uuid.UUID `db:"id" json:"id"`
	This    string    `db:"this" json:"this"`
	Is      string    `db:"is" json:"is"`
	Used    int32     `db:"used" json:"used2"`
	For     float32   `db:"for" json:"for"`
	Testing int64     `db:"testing" json:"testing"`
	ForTesting
	CreateDate time.Time `db:"create_date2" json:"create_date2"`
	UpdateDate time.Time `db:"update_date2" json:"update_date2"`
}

func TestGetTagsForDB(t *testing.T) {
	got := GetTags(&ForTesting{}, "db", nil)
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

func TestGetTagsForDBWithNestedStruct(t *testing.T) {
	got := GetTags(&NestedStructForTesting{}, "db", nil)
	expected := []string{"id", "this", "is", "used", "for",
		"testing", "id", "this", "is", "used", "for", "testing",
		"create_date", "update_date", "create_date2", "update_date2"}

	if len(got) != len(expected) {
		t.Fatalf("GetTags() = %v, want %v", got, expected)
	}

	for i := range got {
		if got[i] != expected[i] {
			t.Fatalf("GetTags() = %v, want %v", got, expected)
		}
	}
}

func TestGetTagsForDBAndIgnoreCertainFields(t *testing.T) {
	got := GetTags(&ForTesting{}, "db", []string{"create_date", "is"})
	expected := []string{"id", "this", "used", "for", "testing", "update_date"}

	if len(got) != len(expected) {
		t.Fatalf("GetTags() = %v, want %v", got, expected)
	}

	for i := range got {
		if got[i] != expected[i] {
			t.Fatalf("GetTags() = %v, want %v", got, expected)
		}
	}
}

func TestGetTagsForJSON(t *testing.T) {
	got := GetTags(&ForTesting{}, "json", nil)
	expected := []string{"ID", "This", "is", "used", "for", "testing", "create_date", "update_date"}

	if len(got) != len(expected) {
		t.Fatalf("GetTags() = %v, want %v", got, expected)
	}

	for i := range got {
		if got[i] != expected[i] {
			t.Fatalf("GetTags() = %v, want %v", got, expected)
		}
	}
}

func TestGetTagsForJSONWithNestedStructAndIgnoreFields(t *testing.T) {
	got := GetTags(&NestedStructForTesting{}, "json", []string{"create_date2", "This"})
	expected := []string{"id", "this", "is", "used2", "for",
		"testing", "ID", "is", "used", "for", "testing",
		"create_date", "update_date", "update_date2"}

	if len(got) != len(expected) {
		t.Fatalf("GetTags() = %v, want %v", got, expected)
	}

	for i := range got {
		if got[i] != expected[i] {
			t.Fatalf("GetTags() = %v, want %v", got, expected)
		}
	}
}

func TestGetTagsForJSONAndIgnoreCertainFields(t *testing.T) {
	got := GetTags(&ForTesting{}, "json", []string{"update_date", "used"})
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

func TestGetTagsAndValuesForJSON(t *testing.T) {
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

func TestGetNestedTagsAndValuesForDB(t *testing.T) {
	id := uuid.NewV4()
	now := time.Now()

	expectedTags := []string{"id", "this", "is", "used", "for", "testing", "create_date", "update_date"}
	var expectedValues []interface{}
	expectedValues = append(expectedValues, id, "yo", "are", int32(1), float32(2.0), int64(123456789), now, now)

	gotTags, gotValues := GetNestedTagsAndValues(createStruct(id, now), "db", nil, nil)
	if (len(gotTags) != len(expectedTags)) || (len(gotValues) != len(expectedValues)) {
		t.Fatalf("GetNestedTagsAndValues() = %v, %v, want %v %v", gotTags, gotValues, expectedTags, expectedValues)
	}

	for i := range gotTags {
		if (gotTags[i] != expectedTags[i]) || (gotValues[i] != expectedValues[i]) {
			fmt.Println(gotTags[i], expectedTags[i], gotValues[i], expectedValues[i])
			t.Fatalf("GetNestedTagsAndValues() = %v, %v, want %v %v", gotTags, gotValues, expectedTags, expectedValues)
		}
	}
}

func TestGetNestedTagsAndValuesForDBAndIgnoreCertainFields(t *testing.T) {
	id := uuid.NewV4()
	now := time.Now()

	expectedTags := []string{"id", "this", "is", "for", "testing", "update_date"}
	var expectedValues []interface{}
	expectedValues = append(expectedValues, id, "yo", "are", float32(2.0), int64(123456789), now)

	gotTags, gotValues := GetNestedTagsAndValues(createStruct(id, now), "db", []string{"used", "create_date"}, nil)
	if (len(gotTags) != len(expectedTags)) || (len(gotValues) != len(expectedValues)) {
		t.Fatalf("GetNestedTagsAndValues() = %v, %v, want %v %v", gotTags, gotValues, expectedTags, expectedValues)
	}

	for i := range gotTags {
		if (gotTags[i] != expectedTags[i]) || (gotValues[i] != expectedValues[i]) {
			fmt.Println(gotTags[i], expectedTags[i], gotValues[i], expectedValues[i])
			t.Fatalf("GetNestedTagsAndValues() = %v, %v, want %v %v", gotTags, gotValues, expectedTags, expectedValues)
		}
	}
}

func TestGetNestedTagsAndValuesForJSON(t *testing.T) {
	id := uuid.NewV4()
	now := time.Now()

	expectedTags := []string{"ID", "This", "is", "used", "for", "testing", "create_date", "update_date"}
	var expectedValues []interface{}
	expectedValues = append(expectedValues, id, "yo", "are", int32(1), float32(2.0), int64(123456789), now, now)

	gotTags, gotValues := GetNestedTagsAndValues(createStruct(id, now), "json", nil, nil)
	if (len(gotTags) != len(expectedTags)) || (len(gotValues) != len(expectedValues)) {
		t.Fatalf("GetNestedTagsAndValues() = %v, %v, want %v %v", gotTags, gotValues, expectedTags, expectedValues)
	}

	for i := range gotTags {
		if (gotTags[i] != expectedTags[i]) || (gotValues[i] != expectedValues[i]) {
			fmt.Println(gotTags[i], expectedTags[i], gotValues[i], expectedValues[i])
			t.Fatalf("GetNestedTagsAndValues() = %v, %v, want %v %v", gotTags, gotValues, expectedTags, expectedValues)
		}
	}
}

func TestGetNestedTagsAndValuesForJSONWithNestedStructAndIgnoreCertainFields(t *testing.T) {
	id := uuid.NewV4()
	now := time.Now()

	expectedTags := []string{"id", "this", "is", "used2", "for", "testing",
		"ID", "This", "is", "for", "testing", "update_date", "update_date2"}
	var expectedValues []interface{}
	expectedValues = append(expectedValues, id, "i", "am", int32(2), float32(3.0), int64(987654321),
		id, "yo", "are", float32(2.0), int64(123456789), now, now)

	gotTags, gotValues := GetNestedTagsAndValues(createNestedStruct(id, now), "json", []string{"used", "create_date", "create_date2"}, []string{"ForTesting"})
	if (len(gotTags) != len(expectedTags)) || (len(gotValues) != len(expectedValues)) {
		t.Fatalf("GetNestedTagsAndValues() = %v, %v, want %v %v", gotTags, gotValues, expectedTags, expectedValues)
	}

	for i := range gotTags {
		if (gotTags[i] != expectedTags[i]) || (gotValues[i] != expectedValues[i]) {
			fmt.Println(gotTags[i], expectedTags[i], gotValues[i], expectedValues[i])
			t.Fatalf("GetNestedTagsAndValues() = %v, %v, want %v %v", gotTags, gotValues, expectedTags, expectedValues)
		}
	}
}

func TestGetNestedTagsAndValuesForJSONWithNestedStruct(t *testing.T) {
	id := uuid.NewV4()
	now := time.Now()

	expectedTags := []string{"id", "this", "is", "used2", "for", "testing",
		"ID", "This", "is", "used", "for", "testing", "create_date", "update_date", "create_date2", "update_date2"}
	var expectedValues []interface{}
	expectedValues = append(expectedValues, id, "i", "am", int32(2), float32(3.0), int64(987654321),
		id, "yo", "are", int32(1), float32(2.0), int64(123456789), now, now, now, now)

	gotTags, gotValues := GetNestedTagsAndValues(createNestedStruct(id, now), "json", nil, []string{"ForTesting"})
	if (len(gotTags) != len(expectedTags)) || (len(gotValues) != len(expectedValues)) {
		t.Fatalf("GetNestedTagsAndValues() = %v, %v, want %v %v", gotTags, gotValues, expectedTags, expectedValues)
	}

	for i := range gotTags {
		if (gotTags[i] != expectedTags[i]) || (gotValues[i] != expectedValues[i]) {
			fmt.Println(gotTags[i], expectedTags[i], gotValues[i], expectedValues[i])
			t.Fatalf("GetNestedTagsAndValues() = %v, %v, want %v %v", gotTags, gotValues, expectedTags, expectedValues)
		}
	}
}

func createStruct(id uuid.UUID, now time.Time) *ForTesting {
	return &ForTesting{
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

func createNestedStruct(id uuid.UUID, now time.Time) *NestedStructForTesting {
	return &NestedStructForTesting{
		ID:         id,
		This:       "i",
		Is:         "am",
		Used:       2,
		For:        3.0,
		Testing:    987654321,
		ForTesting: *createStruct(id, now),
		CreateDate: now,
		UpdateDate: now,
	}
}
