package storage

import (
	"context"
	"reflect"
	"testing"

	"github.com/LexusEgorov/items-accounting/internal/models"
)

func TestCategories_Get(t *testing.T) {
	mockDB := &DB{
		DB: mockDB{},
	}

	testCategories, _ := NewCategories(mockDB)

	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name         string
		c            *Categories
		args         args
		wantCategory models.Category
		wantErr      bool
	}{
		{
			name: "mock test",
			c:    testCategories,
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCategory, err := tt.c.Get(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Categories.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCategory, tt.wantCategory) {
				t.Errorf("Categories.Get() = %v, want %v", gotCategory, tt.wantCategory)
			}
		})
	}
}
