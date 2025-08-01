package core

import (
	"fmt"
	"log/slog"
	"sync"
)

type Grid struct {
	GID                    int
	MinX, MinY, MaxX, MaxY int
	playerIDs              map[int]bool
	pIDLock                sync.RWMutex
}

func NewGrid(gid, minX, maxX, minY, maxY int) *Grid {
	return &Grid{
		GID:       gid,
		MinX:      minX,
		MinY:      minY,
		MaxX:      maxX,
		MaxY:      maxY,
		playerIDs: make(map[int]bool),
		pIDLock:   sync.RWMutex{},
	}
}
func (g *Grid) Add(playerID int) {
	slog.Info(fmt.Sprintf("Add player: %d", playerID))
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()
	g.playerIDs[playerID] = true
}
func (g *Grid) Remove(playerID int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()
	delete(g.playerIDs, playerID)
}

func (g *Grid) GetPlayerIDs() []int {
	g.pIDLock.RLock()
	defer g.pIDLock.RUnlock()
	ids := make([]int, len(g.playerIDs))
	i := 0
	for id := range g.playerIDs {
		ids[i] = id
	}
	return ids
}

func (g *Grid) String() string {
	return fmt.Sprintf("Grud id:%d,minx:%d,miny:%d,maxx:%d,maxy:%d,playerids:%v", g.GID, g.MinX, g.MinY, g.MaxX, g.MaxY, g.GetPlayerIDs())
}
