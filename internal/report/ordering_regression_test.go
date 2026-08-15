package report

import (
	"encoding/json"
	"testing"
	"time"

	"LogPilot/internal/model"
)

func TestEventJSONPreservesCallerOrdering(t *testing.T) {
	base := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	older := model.NewEvent(base, model.LevelInfo, "api", "older", nil)
	newer := model.NewEvent(base.Add(time.Minute), model.LevelInfo, "api", "newer", nil)
	data, err := EventJSON([]model.Event{newer, older})
	if err != nil {
		t.Fatal(err)
	}
	var got []model.Event
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[0].ID != newer.ID || got[1].ID != older.ID {
		t.Fatalf("JSON output changed requested descending order: %v", []string{got[0].ID, got[1].ID})
	}
}
