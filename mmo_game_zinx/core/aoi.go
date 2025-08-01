package core

import (
	"fmt"
	"log/slog"
)

type AOIManager struct {
	MinX, MinY, MaxX, MaxY int
	CntX, CntY             int

	grids map[int]*Grid
}

func NewAOIManager(minX, maxX, cntx, minY, maxY, cnty int) *AOIManager {
	aoiMgr := &AOIManager{
		MinX:  minX,
		MinY:  minY,
		MaxX:  maxX,
		MaxY:  maxY,
		CntX:  cntx,
		CntY:  cnty,
		grids: make(map[int]*Grid),
	}
	idx := 0
	width := aoiMgr.gridWidth()
	height := aoiMgr.gridHeight()
	for y := 0; y < cnty; y++ {
		for x := 0; x < cntx; x++ {
			//slog.Info(strconv.Itoa(idx))
			aoiMgr.grids[idx] = NewGrid(idx,
				minX+x*width,
				minX+(x+1)*width,
				minY+y*height,
				minY+(x+1)*height,
			)

			idx++
		}
	}
	return aoiMgr
}

func (aoiMgr *AOIManager) gridWidth() int {
	return (aoiMgr.MaxX - aoiMgr.MinX) / aoiMgr.CntX
}
func (aoiMgr *AOIManager) gridHeight() int {
	return (aoiMgr.MaxY - aoiMgr.MinY) / aoiMgr.CntY
}

func (aoiMgr *AOIManager) String() string {
	msg := fmt.Sprintf("AOIManager:\n Minx: %d\n MaxX: %d\n CntX:%d\n MinY: %d\n Maxy: %d\n CntY: %d", aoiMgr.MinX, aoiMgr.MaxX, aoiMgr.CntX, aoiMgr.MinY, aoiMgr.MaxY, aoiMgr.CntY)
	for _, ids := range aoiMgr.grids {
		msg += fmt.Sprintln(ids)
	}
	return msg
}

// 根据格子的gID得到当前周边的九宫格信息
func (m *AOIManager) GetSurroundGridsByGid(gID int) (grids []*Grid) {
	//判断gID是否存在
	if _, ok := m.grids[gID]; !ok {
		return
	}

	//将当前gid添加到九宫格中
	grids = append(grids, m.grids[gID])

	//根据gid得到当前格子所在的X轴编号
	idx := gID % m.CntX

	//判断当前idx左边是否还有格子
	if idx > 0 {
		grids = append(grids, m.grids[gID-1])
	}
	//判断当前的idx右边是否还有格子
	if idx < m.CntX-1 {
		grids = append(grids, m.grids[gID+1])
	}

	//将x轴当前的格子都取出，进行遍历，再分别得到每个格子的上下是否有格子

	//得到当前x轴的格子id集合
	gidsX := make([]int, 0, len(grids))
	for _, v := range grids {
		gidsX = append(gidsX, v.GID)
	}

	//遍历x轴格子
	for _, v := range gidsX {
		//计算该格子处于第几列
		idy := v / m.CntX

		//判断当前的idy上边是否还有格子
		if idy > 0 {
			grids = append(grids, m.grids[v-m.CntX])
		}
		//判断当前的idy下边是否还有格子
		if idy < m.CntY-1 {
			grids = append(grids, m.grids[v+m.CntX])
		}
	}

	return
}

func (aoiMgr *AOIManager) GetPidsByPos(x, y float32) (playerIDs []int) {
	gID := aoiMgr.GetGidByPos(x, y)
	grids := aoiMgr.GetSurroundGridsByGid(gID)
	for _, grid := range grids {
		playerIDs = append(playerIDs, grid.GetPlayerIDs()...)
	}
	return playerIDs
}
func (aoiMgr *AOIManager) GetGidByPos(x, y float32) (gID int) {
	idx := (int(x) - aoiMgr.MinX) / aoiMgr.gridWidth()
	idy := (int(y) - aoiMgr.MinY) / aoiMgr.gridHeight()
	return idy*aoiMgr.CntX + idx
}

func (aoiMgr *AOIManager) AddPidToGrid(pID, gID int) {
	slog.Info(fmt.Sprintf("AddPidToGrid: pID:%d gID:%d", pID, gID))
	aoiMgr.grids[gID].Add(pID)
}

func (aoiMgr *AOIManager) RemovePidFromGrid(pID, gID int) {
	aoiMgr.grids[gID].Remove(pID)
}

func (aoiMgr *AOIManager) GetPidsByGid(gID int) (playerIDs []int) {
	return aoiMgr.grids[gID].GetPlayerIDs()
}
func (aoiMgr *AOIManager) AddToGridByPos(pID int, x float32, y float32) {
	gID := aoiMgr.GetGidByPos(x, y)
	aoiMgr.AddPidToGrid(pID, gID)
}

func (aoiMgr *AOIManager) RemoveFromGridByPos(pID int, x, y float32) {
	gID := aoiMgr.GetGidByPos(x, y)
	aoiMgr.RemovePidFromGrid(pID, gID)
}
