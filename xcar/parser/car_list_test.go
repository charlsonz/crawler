package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCarList(t *testing.T) {
	contents, err := ioutil.ReadFile("car_list_test_data.html")
	if err != nil {
		panic(err)
	}

	got := ParseCarList(contents, "")

	const (
		resultSize   = 30
		carModelSize = 20
	)
	expectedCarModelUrls := []string{
		"http://newcar.xcar.com.cn/4007/",
		"http://newcar.xcar.com.cn/52/",
		"http://newcar.xcar.com.cn/3875/",
	}
	expectedCarListUrls := []string{
		"http://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-1/",
		"http://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-2-1/",
		"http://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-3-1/",
	}

	if len(got.Requests) != resultSize {
		t.Errorf("got should have %d requests; but had %d",
			resultSize, len(got.Requests))
	}
	for i, url := range expectedCarModelUrls {
		if got.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s",
				i, url, got.Requests[i].Url)
		}
	}
	for i, url := range expectedCarListUrls {
		if got.Requests[carModelSize+i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s",
				carModelSize+i, url, got.Requests[i].Url)
		}
	}

}
