package server

import (
	"testing"

	"github.com/bwolf1/gin-postgres-rest-service/pkg/service"
	"github.com/gin-gonic/gin"
)

func TestRest_listProducts(t *testing.T) {
	type fields struct {
		product *service.Product
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"TestRest_listProducts",
			fields{
				product: &service.Product{},
			},
			args{
				c: &gin.Context{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rest{
				products: tt.fields.product,
			}
			r.listProducts(tt.args.c)
		})
	}
}

func TestRest_getProduct(t *testing.T) {
	type fields struct {
		product *service.Product
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"TestRest_getProduct",
			fields{
				product: &service.Product{},
			},
			args{
				c: &gin.Context{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rest{
				products: tt.fields.product,
			}
			r.getProduct(tt.args.c)
		})
	}
}

func TestRest_getProductVersions(t *testing.T) {
	type fields struct {
		product *service.Product
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{	
			"TestRest_getProductVersions",
			fields{
				product: &service.Product{},
			},
			args{
				c: &gin.Context{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rest{
				products: tt.fields.product,
			}
			r.getProductVersions(tt.args.c)
		})
	}
}
