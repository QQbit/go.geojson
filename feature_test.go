package geojson

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestNewFeature(t *testing.T) {
	f := NewFeature(NewPointGeometry([]float64{1, 2}))

	if f.Type != "Feature" {
		t.Errorf("should have type of Feature, got %v", f.Type)
	}
}

func TestFeatureMarshalJSON(t *testing.T) {
	f := NewFeature(NewPointGeometry([]float64{1, 2}))
	blob, err := f.MarshalJSON()

	if err != nil {
		t.Fatalf("should marshal to json just fine but got %v", err)
	}

	if !bytes.Contains(blob, []byte(`"properties":{}`)) {
		t.Errorf("json should set properties to empty object if there are none")
	}
}

func TestFeatureMarshal(t *testing.T) {
	f := NewFeature(NewPointGeometry([]float64{1, 2}))
	blob, err := json.Marshal(f)

	if err != nil {
		t.Fatalf("should marshal to json just fine but got %v", err)
	}

	if !bytes.Contains(blob, []byte(`"properties":{}`)) {
		t.Errorf("json should set properties to empty object if there are none")
	}
}

func TestFeatureMarshalValue(t *testing.T) {
	f := NewFeature(NewPointGeometry([]float64{1, 2}))
	blob, err := json.Marshal(*f)

	if err != nil {
		t.Fatalf("should marshal to json just fine but got %v", err)
	}

	if !bytes.Contains(blob, []byte(`"properties":{}`)) {
		t.Errorf("json should set properties to empty object if there are none")
	}
}

func TestUnmarshalFeature(t *testing.T) {
	rawJSON := `
	  { "type": "Feature",
	    "bbox": [1, 2, 3, 4],
	    "geometry": {"type": "Point", "coordinates": [102.0, 0.5]},
	    "properties": {"prop0": "value0"}
	  }`

	f, err := UnmarshalFeature([]byte(rawJSON))
	if err != nil {
		t.Fatalf("should unmarshal feature without issue, err %v", err)
	}

	if f.Type != "Feature" {
		t.Errorf("should have type of Feature, got %v", f.Type)
	}

	if len(f.Properties) != 1 {
		t.Errorf("should have 1 property but got %d", len(f.Properties))
	}

	if len(f.BoundingBox) != 4 {
		t.Errorf("should have unmarshalled bounding box")
	}
}
