package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	content, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	const resultSize = 470
	expectedUrls := []string{
		"",
		"",
		"",
	}
	expectedCities := []string{
		"",
		"",
		"",
	}

	result := ParseCityList(content)
	if len(result.DataItem) != resultSize {
		t.Errorf("expected: %d, but got: %d", resultSize, len(result.DataItem))
	}
	if len(result.Requests) != resultSize {
		t.Errorf("expected: %d, but got: %d", resultSize, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected: %s, got: %s", url, result.Requests[i])
		}
	}
	for i, city := range expectedCities {
		if result.DataItem[i].(string) != city {
			t.Errorf("expected: %s, got: %s", city, result.DataItem[i].(string))
		}
	}

}
