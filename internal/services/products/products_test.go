package products

import (
	"context"
	"errors"
	"testing"

	"github.com/LexusEgorov/items-accounting/internal/models"
	"github.com/LexusEgorov/items-accounting/internal/services/products/mocks"
	"github.com/stretchr/testify/assert"
)

func TestProducts_Add(t *testing.T) {
	mockRepository := mocks.NewProductRepository(t)
	productManager := New(mockRepository)

	type args struct {
		ctx     context.Context
		product models.ProductDTO
	}
	tests := []struct {
		name        string
		prepareMock func()
		c           Products
		args        args
		want        models.ProductDTO
		wantErr     bool
	}{
		{
			name: "basic test",
			prepareMock: func() {
				mockRepository.On("Add", context.Background(), models.ProductDTO{
					CatID: 1,
					Name:  "test",
					Price: 10,
					Count: 100,
				}).Return(1, nil)
			},
			c: *productManager,
			args: args{
				ctx: context.Background(),
				product: models.ProductDTO{
					CatID: 1,
					Name:  "test",
					Price: 10,
					Count: 100,
				},
			},
			want: models.ProductDTO{
				ID:    1,
				CatID: 1,
				Name:  "test",
				Price: 10,
				Count: 100,
			},
			wantErr: false,
		},
		{
			name:        "empty catId",
			prepareMock: func() {},
			c:           *productManager,
			args: args{
				ctx: context.Background(),
				product: models.ProductDTO{
					Name:  "test",
					Price: 10,
					Count: 100,
				},
			},
			want:    models.ProductDTO{},
			wantErr: true,
		},
		{
			name:        "empty name",
			prepareMock: func() {},
			c:           *productManager,
			args: args{
				ctx: context.Background(),
				product: models.ProductDTO{
					CatID: 1,
					Price: 10,
					Count: 100,
				},
			},
			want:    models.ProductDTO{},
			wantErr: true,
		},
		{
			name: "storage err",
			prepareMock: func() {
				mockRepository.On("Add", context.Background(), models.ProductDTO{
					CatID: 2,
					Name:  "test",
					Price: 10,
					Count: 100,
				}).Return(0, errors.New("testErr"))
			},
			c: *productManager,
			args: args{
				ctx: context.Background(),
				product: models.ProductDTO{
					CatID: 2,
					Name:  "test",
					Price: 10,
					Count: 100,
				},
			},
			want:    models.ProductDTO{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMock()
			got, err := tt.c.Add(tt.args.ctx, tt.args.product)

			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestProducts_Set(t *testing.T) {
	mockRepository := mocks.NewProductRepository(t)
	productManager := New(mockRepository)

	type args struct {
		ctx     context.Context
		product models.ProductDTO
	}
	tests := []struct {
		name        string
		prepareMock func()
		c           Products
		args        args
		want        models.ProductDTO
		wantErr     bool
	}{
		{
			name: "basic test",
			prepareMock: func() {
				mockRepository.On("Set", context.Background(), models.ProductDTO{
					ID:    1,
					CatID: 1,
					Name:  "test",
					Price: 10,
					Count: 100,
				}).Return(nil)
			},
			c: *productManager,
			args: args{
				ctx: context.Background(),
				product: models.ProductDTO{
					ID:    1,
					CatID: 1,
					Name:  "test",
					Price: 10,
					Count: 100,
				},
			},
			want: models.ProductDTO{
				ID:    1,
				CatID: 1,
				Name:  "test",
				Price: 10,
				Count: 100,
			},
			wantErr: false,
		},
		{
			name:        "empty id",
			prepareMock: func() {},
			c:           *productManager,
			args: args{
				ctx: context.Background(),
				product: models.ProductDTO{
					Name:  "test",
					Price: 10,
					Count: 100,
				},
			},
			want:    models.ProductDTO{},
			wantErr: true,
		},
		{
			name:        "empty catId",
			prepareMock: func() {},
			c:           *productManager,
			args: args{
				ctx: context.Background(),
				product: models.ProductDTO{
					ID:    1,
					Name:  "test",
					Price: 10,
					Count: 100,
				},
			},
			want:    models.ProductDTO{},
			wantErr: true,
		},
		{
			name:        "empty name",
			prepareMock: func() {},
			c:           *productManager,
			args: args{
				ctx: context.Background(),
				product: models.ProductDTO{
					CatID: 1,
					Price: 10,
					Count: 100,
				},
			},
			want:    models.ProductDTO{},
			wantErr: true,
		},
		{
			name: "storage err",
			prepareMock: func() {
				mockRepository.On("Set", context.Background(), models.ProductDTO{
					ID:    2,
					CatID: 1,
					Name:  "test",
					Price: 10,
					Count: 100,
				}).Return(errors.New("testErr"))
			},
			c: *productManager,
			args: args{
				ctx: context.Background(),
				product: models.ProductDTO{
					ID:    2,
					CatID: 1,
					Name:  "test",
					Price: 10,
					Count: 100,
				},
			},
			want:    models.ProductDTO{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMock()
			got, err := tt.c.Set(tt.args.ctx, tt.args.product)

			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestProducts_Get(t *testing.T) {
	mockRepository := mocks.NewProductRepository(t)
	productManager := New(mockRepository)

	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name        string
		prepareMock func()
		c           Products
		args        args
		want        models.ProductDTO
		wantErr     bool
	}{
		{
			name: "basic test",
			prepareMock: func() {
				mockRepository.On("Get", context.Background(), 1).Return(models.Product{
					ID:    1,
					CatID: 1,
					Name:  "test",
					Price: 10,
					Count: 100,
				}, nil)
			},
			c: *productManager,
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			want: models.ProductDTO{
				ID:    1,
				CatID: 1,
				Name:  "test",
				Price: 10,
				Count: 100,
			},
			wantErr: false,
		},
		{
			name: "storage err",
			prepareMock: func() {
				mockRepository.On("Get", context.Background(), 2).Return(models.Product{}, errors.New("testErr"))
			},
			c: *productManager,
			args: args{
				ctx: context.Background(),
				id:  2,
			},
			want:    models.ProductDTO{},
			wantErr: true,
		},
		{
			name:        "empty id",
			prepareMock: func() {},
			c:           *productManager,
			args: args{
				ctx: context.Background(),
				id:  0,
			},
			want:    models.ProductDTO{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMock()
			got, err := tt.c.Get(tt.args.ctx, tt.args.id)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestProducts_Delete(t *testing.T) {
	mockRepository := mocks.NewProductRepository(t)
	productManager := New(mockRepository)

	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name        string
		prepareMock func()
		c           Products
		args        args
		wantErr     bool
	}{
		{
			name: "basic test",
			prepareMock: func() {
				mockRepository.On("Delete", context.Background(), 1).Return(nil)
			},
			c: *productManager,
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantErr: false,
		},
		{
			name: "storage err",
			prepareMock: func() {
				mockRepository.On("Delete", context.Background(), 2).Return(errors.New("testErr"))
			},
			c: *productManager,
			args: args{
				ctx: context.Background(),
				id:  2,
			},
			wantErr: true,
		},
		{
			name:        "empty id",
			prepareMock: func() {},
			c:           *productManager,
			args: args{
				ctx: context.Background(),
				id:  0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMock()
			err := tt.c.Delete(tt.args.ctx, tt.args.id)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
