package categories

import (
	"context"
	"reflect"
	"testing"

	"github.com/LexusEgorov/items-accounting/internal/models"
	"github.com/LexusEgorov/items-accounting/internal/services/categories/mocks"
)

func TestCategories_Add(t *testing.T) {
	mockRepository := mocks.NewCategoryRepository(t)
	categoriesManager := New(mockRepository)

	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name        string
		c           Categories
		args        args
		mockPrepare func()
		want        models.CategoryDTO
		wantErr     bool
	}{
		{
			name: "basic test",
			c:    *categoriesManager,
			mockPrepare: func() {
				mockRepository.On("Add", context.Background(), "test").Return(1, nil)
			},
			args: args{
				ctx:  context.Background(),
				name: "test",
			},
			want: models.CategoryDTO{
				ID:   1,
				Name: "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockPrepare()
			got, err := tt.c.Add(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Categories.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Categories.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
