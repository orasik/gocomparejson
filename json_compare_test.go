package jsoncompare

import "testing"

func TestCompareJSONWithTwoValidEqualStrings(t *testing.T) {
	json1 := []byte(`{"hello":"world"}`)
	json2 := []byte(`{
"hello":"world"
}`)

	val, err := CompareJSON(string(json1), string(json2))

	if err != nil {
		t.Errorf("Error received in comparing two valid json strings. %s", err)
	}

	if val == false {
		t.Error("CompareJSON should have returned true but it returned false")
	}
}

func TestCompareJSONWithTwoValidNonEqualStrings(t *testing.T) {
	json1 := []byte(`{"hello":"world"}`)
	json2 := []byte(`{
"hello":"world",
"second": "test"
}`)

	val, err := CompareJSON(string(json1), string(json2))

	if err != nil {
		t.Errorf("Error received in comparing two valid json strings. %s", err)
	}

	if val == true {
		t.Error("CompareJSON should have returned fakse but it returned true")
	}
}

func TestCompareJSONWithTwoValidEqualStringsWithRandomKeyOrder(t *testing.T) {
	json1 := []byte(`{"key1":"value1", "key2":"value2", "key3":"value3", "key4":"value4"}`)
	json2 := []byte(`{
"key4":"value4",
"key1":"value1",
"key3":"value3",
"key2":"value2"
}`)

	val, err := CompareJSON(string(json1), string(json2))

	if err != nil {
		t.Errorf("Error received in comparing two valid json strings. %s", err)
	}

	if val == false {
		t.Error("CompareJSON should have returned fakse but it returned true")
	}
}

func TestCompareJSONWithFirstStringInvalidJSON(t *testing.T) {
	json1 := []byte(`blablaInvalidJson{"key1":"value1", "key2":"value2", "key3":"value3" "key4":"value4"}`)
	json2 := []byte(`{
"key4":"value4",
"key1":"value1",
"key3":"value3",
"key2":"value2"
}`)

	val, err := CompareJSON(string(json1), string(json2))

	if err == nil {
		t.Error("There should be an error for Invalid JSON")
	}

	if val == true {
		t.Error("CompareJSON should have returned fakse but it returned true")
	}
}

func TestCompareJSONWithSecondStringInvalidJSON(t *testing.T) {
	json1 := []byte(`{"key1":"value1", "key2":"value2", "key3":"value3", "key4":"value4"}`)
	json2 := []byte(`
"key4":"value4",
"key1":"value1",
"key3":"value3",
"key2":"value2"
}`)

	val, err := CompareJSON(string(json1), string(json2))

	if err == nil {
		t.Error("There should be an error for Invalid JSON")
	}

	if val == true {
		t.Error("CompareJSON should have returned fakse but it returned true")
	}
}

func TestCompareJSONWithBothStringsInvalidJSON(t *testing.T) {
	json1 := []byte(`4354545{"key1":"value1", "key2":"value2", "key3":"value3", "key4":"value4"}`)
	json2 := []byte(`errorError
"key4":"value4",
"key1":"value1",
"key3":"value3",
"key2":"value2"
}`)

	val, err := CompareJSON(string(json1), string(json2))

	if err == nil {
		t.Error("There should be an error for Invalid JSON")
	}

	if val == true {
		t.Error("CompareJSON should have returned fakse but it returned true")
	}
}
