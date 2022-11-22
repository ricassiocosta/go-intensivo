package entity

import (
	"reflect"
	"testing"
)

func TestNewOrder(t *testing.T) {
	type args struct {
		id    string
		price float64
		tax   float64
	}
	tests := []struct {
		name    string
		args    args
		want    *Order
		wantErr bool
	}{
		{
			name: "should be able to create an order",
			args: args{
				id:    "fake-uuid",
				price: 1,
				tax:   1,
			},
			wantErr: false,
			want: &Order{
				ID:         "fake-uuid",
				Price:      1,
				Tax:        1,
				FinalPrice: 2,
			},
		},

		{
			name: "should fail if args are invalid",
			args: args{
				id:    " ",
				price: 1,
				tax:   1,
			},
			wantErr: true,
			want:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewOrder(tt.args.id, tt.args.price, tt.args.tax)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_IsValid(t *testing.T) {
	type fields struct {
		ID         string
		Price      float64
		Tax        float64
		FinalPrice float64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "should be able to validated order",
			fields: fields{
				ID:         "fake-uuid",
				Price:      1,
				Tax:        1,
				FinalPrice: 2,
			},
			wantErr: false,
		},

		{
			name: "should fail if ID is missing",
			fields: fields{
				ID:    " ",
				Price: 1,
				Tax:   1,
			},
			wantErr: true,
		},

		{
			name: "should fail if price is invalid",
			fields: fields{
				ID:    "fake-uuid",
				Price: -1,
				Tax:   1,
			},
			wantErr: true,
		},

		{
			name: "should fail if tax is invalid",
			fields: fields{
				ID:    "fake-uuid",
				Price: 1,
				Tax:   -1,
			},
			wantErr: true,
		},

		{
			name: "should fail if final price is invalid",
			fields: fields{
				ID:         "fake-uuid",
				Price:      1,
				Tax:        1,
				FinalPrice: -1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Order{
				ID:         tt.fields.ID,
				Price:      tt.fields.Price,
				Tax:        tt.fields.Tax,
				FinalPrice: tt.fields.FinalPrice,
			}
			if err := o.IsValid(); (err != nil) != tt.wantErr {
				t.Errorf("Order.IsValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
