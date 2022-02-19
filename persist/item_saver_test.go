package persist

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/charlsonz/crawler/engine"
	"github.com/charlsonz/crawler/model"
	"github.com/olivere/elastic/v7"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "http://newcar.xcar.com.cn/4007",
		Type: "xcar",
		Id:   "108906739",
		Payload: model.Car{
			Name:         "",
			Price:        0,
			ImageURL:     "",
			Size:         "",
			Fuel:         0,
			Transmission: "",
			Engine:       "",
			Displacement: 0,
			MaxSpeed:     0,
			Acceleration: 0,
		},
	}

	// TODO: Try to start up elastic search
	// here using docker go client.
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	// Save expected item
	err = Save(client, index, expected)
	if err != nil {
		panic(err)
	}

	// Fetch saved item
	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

	var actual engine.Item
	json.Unmarshal(resp.Source, &actual)

	// Verify result
	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
