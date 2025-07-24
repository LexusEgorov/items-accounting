package categories

import (
	"context"
	"errors"
	"testing"

	"github.com/LexusEgorov/items-accounting/internal/models"
	"github.com/LexusEgorov/items-accounting/internal/services/categories/mocks"
	"github.com/stretchr/testify/assert"
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
				Name: "test",
			},
			wantErr: false,
		},
		{
			name:        "add without name",
			c:           *categoriesManager,
			mockPrepare: func() {},
			args: args{
				ctx: context.Background(),
			},
			want:    models.CategoryDTO{},
			wantErr: true,
		},
		{
			name: "error from storage",
			c:    *categoriesManager,
			mockPrepare: func() {
				mockRepository.On("Add", context.Background(), "testErr").Return(0, errors.New("testErr"))
			},
			args: args{
				ctx:  context.Background(),
				name: "testErr",
			},
			want:    models.CategoryDTO{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockPrepare()
			got, err := tt.c.Add(tt.args.ctx, tt.args.name)

			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCategories_Set(t *testing.T) {
	mockRepository := mocks.NewCategoryRepository(t)
	categoriesManager := New(mockRepository)

	type args struct {
		ctx      context.Context
		category models.CategoryDTO
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
				mockRepository.On("Set", context.Background(), 1, "test1").Return(nil)
			},
			args: args{
				ctx: context.Background(),
				category: models.CategoryDTO{
					ID:   1,
					Name: "test1",
				},
			},
			want: models.CategoryDTO{
				ID:   1,
				Name: "test1",
			},
			wantErr: false,
		},
		{
			name:        "empty name",
			c:           *categoriesManager,
			mockPrepare: func() {},
			args: args{
				ctx: context.Background(),
				category: models.CategoryDTO{
					ID: 1,
				},
			},
			want:    models.CategoryDTO{},
			wantErr: true,
		},
		{
			name:        "empty id",
			c:           *categoriesManager,
			mockPrepare: func() {},
			args: args{
				ctx: context.Background(),
				category: models.CategoryDTO{
					Name: "test",
				},
			},
			want:    models.CategoryDTO{},
			wantErr: true,
		},
		{
			name: "storage err",
			c:    *categoriesManager,
			mockPrepare: func() {
				mockRepository.On("Set", context.Background(), 1, "test").Return(errors.New("testErr"))
			},
			args: args{
				ctx: context.Background(),
				category: models.CategoryDTO{
					ID:   1,
					Name: "test",
				},
			},
			want:    models.CategoryDTO{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockPrepare()
			got, err := tt.c.Set(tt.args.ctx, tt.args.category)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCategories_Get(t *testing.T) {
	mockRepository := mocks.NewCategoryRepository(t)
	categoriesManager := New(mockRepository)

	type args struct {
		ctx context.Context
		ID  int
	}
	tests := []struct {
		name        string
		c           Categories
		mockPrepare func()
		args        args
		want        models.CategoryDTO
		wantErr     bool
	}{
		{
			name: "basic test",
			c:    *categoriesManager,
			mockPrepare: func() {
				mockRepository.On("Get", context.Background(), 1).Return(models.Category{
					ID:   1,
					Name: "test",
				}, nil)
			},
			args: args{
				ctx: context.Background(),
				ID:  1,
			},
			want: models.CategoryDTO{
				ID:   1,
				Name: "test",
			},
			wantErr: false,
		},
		{
			name: "storage err",
			c:    *categoriesManager,
			mockPrepare: func() {
				mockRepository.On("Get", context.Background(), 2).Return(models.Category{}, errors.New("testErr"))
			},
			args: args{
				ctx: context.Background(),
				ID:  2,
			},
			want:    models.CategoryDTO{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockPrepare()
			got, err := tt.c.Get(tt.args.ctx, tt.args.ID)

			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

			assert.Equal(t, tt.want, got, "not equal")
		})
	}
}

func TestCategories_Delete(t *testing.T) {
	mockRepository := mocks.NewCategoryRepository(t)
	categoriesManager := New(mockRepository)

	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name        string
		mockPrepare func()
		c           Categories
		args        args
		wantErr     bool
	}{
		{
			name: "basic test",
			mockPrepare: func() {
				mockRepository.On("Delete", context.Background(), 1).Return(nil)
			},
			c: *categoriesManager,
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantErr: false,
		},
		{
			name:        "bad id",
			mockPrepare: func() {},
			c:           *categoriesManager,
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
		},
		{
			name: "storage err",
			mockPrepare: func() {
				mockRepository.On("Delete", context.Background(), 2).Return(errors.New("testErr"))
			},
			c: *categoriesManager,
			args: args{
				ctx: context.Background(),
				id:  2,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockPrepare()
			err := tt.c.Delete(tt.args.ctx, tt.args.id)

			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
