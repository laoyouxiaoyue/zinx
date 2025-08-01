package core

import (
	"reflect"
	"testing"
)

func TestAOIManager_GetSurroundGridsByGid(t *testing.T) {
	type fields struct {
		MinX  int
		MinY  int
		MaxX  int
		MaxY  int
		CntX  int
		CntY  int
		grids map[int]*Grid
	}
	type args struct {
		gID int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantGrids []*Grid
	}{
		{
			name: "test1",
			fields: fields{
				MinX: 0,
				MinY: 0,
				MaxX: 150,
				MaxY: 250,
				CntX: 3,
				CntY: 5,
			},
			args: args{
				gID: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aoiMgr := NewAOIManager(tt.fields.MinX, tt.fields.MaxX, tt.fields.CntX, tt.fields.MinY, tt.fields.MaxY, tt.fields.CntY)

			if gotGrids := aoiMgr.GetSurroundGridsByGid(tt.args.gID); !reflect.DeepEqual(gotGrids, tt.wantGrids) {
				t.Errorf("GetSurroundGridsByGid() = %v, want %v", gotGrids, tt.wantGrids)
			}
		})
	}
}

func TestAOIManager_String(t *testing.T) {
	type fields struct {
		MinX  int
		MinY  int
		MaxX  int
		MaxY  int
		CntX  int
		CntY  int
		grids map[int]*Grid
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aoiMgr := &AOIManager{
				MinX:  tt.fields.MinX,
				MinY:  tt.fields.MinY,
				MaxX:  tt.fields.MaxX,
				MaxY:  tt.fields.MaxY,
				CntX:  tt.fields.CntX,
				CntY:  tt.fields.CntY,
				grids: tt.fields.grids,
			}
			if got := aoiMgr.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAOIManager_gridHeight(t *testing.T) {
	type fields struct {
		MinX  int
		MinY  int
		MaxX  int
		MaxY  int
		CntX  int
		CntY  int
		grids map[int]*Grid
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aoiMgr := &AOIManager{
				MinX:  tt.fields.MinX,
				MinY:  tt.fields.MinY,
				MaxX:  tt.fields.MaxX,
				MaxY:  tt.fields.MaxY,
				CntX:  tt.fields.CntX,
				CntY:  tt.fields.CntY,
				grids: tt.fields.grids,
			}
			if got := aoiMgr.gridHeight(); got != tt.want {
				t.Errorf("gridHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAOIManager_gridWidth(t *testing.T) {
	type fields struct {
		MinX  int
		MinY  int
		MaxX  int
		MaxY  int
		CntX  int
		CntY  int
		grids map[int]*Grid
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aoiMgr := &AOIManager{
				MinX:  tt.fields.MinX,
				MinY:  tt.fields.MinY,
				MaxX:  tt.fields.MaxX,
				MaxY:  tt.fields.MaxY,
				CntX:  tt.fields.CntX,
				CntY:  tt.fields.CntY,
				grids: tt.fields.grids,
			}
			if got := aoiMgr.gridWidth(); got != tt.want {
				t.Errorf("gridWidth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAOIManager(t *testing.T) {
	type fields struct {
		MinX  int
		MinY  int
		MaxX  int
		MaxY  int
		CntX  int
		CntY  int
		grids map[int]*Grid
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test1",
			fields: fields{
				MinX: 0,
				MinY: 0,
				MaxX: 150,
				MaxY: 250,
				CntX: 3,
				CntY: 5,
			},
			want: "AOIManager:\n Minx: 0\n MaxX: 150\n CntX:3\n MinY: 0\n Maxy: 250\n CntY: 5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aoiMgr := &AOIManager{
				MinX:  tt.fields.MinX,
				MinY:  tt.fields.MinY,
				MaxX:  tt.fields.MaxX,
				MaxY:  tt.fields.MaxY,
				CntX:  tt.fields.CntX,
				CntY:  tt.fields.CntY,
				grids: tt.fields.grids,
			}
			if got := aoiMgr.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
