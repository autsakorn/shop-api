package models

import (
	"testing"
	"time"
)

func TestCategory_StatusRes(t *testing.T) {
	type fields struct {
		ID        int64
		Name      string
		Detail    string
		Status    int32
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Base case", fields{1, "Name", "Detail", 1, time.Now(), time.Now()}, "Active"},
		{"Invalid case", fields{1, "Name", "Detail", -1, time.Now(), time.Now()}, "Invalid"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			category := &Category{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				Detail:    tt.fields.Detail,
				Status:    tt.fields.Status,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if got := category.StatusRes(); got != tt.want {
				t.Errorf("Category.StatusRes() = %v, want %v", got, tt.want)
			}
		})
	}
}
