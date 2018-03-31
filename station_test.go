package rail_test

import (
	"testing"
	"time"

	"github.com/go-india/rail"
)

func TestTrainBetweenStations(t *testing.T) {
	c := &rail.Client{
		Auth: rail.NewAuth(getAPIKey()),
	}
	testClient(c, t)

	req := rail.TrainBetweenStationsReq{
		StationFrom: "BE",
		StationTo:   "ADI",
		Date:        time.Now(),
	}

	var resp rail.TrainBetweenStationsResp
	err := c.Do(c.Auth(req), &resp)
	if err != nil {
		t.Fatalf("client Do failed: %+v", err)
	}

	if len(resp.Trains) < 1 {
		t.Fatal("invalid trains length")
	}
}
