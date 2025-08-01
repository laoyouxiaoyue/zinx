package core

import (
	"sync"
)

const (
	AOI_MIN_X  int = 85
	AOI_MAX_X  int = 410
	AOI_CNTS_X int = 10
	AOI_MIN_Y  int = 75
	AOI_MAX_Y  int = 400
	AOI_CNTS_Y int = 20
)

type WorldManager struct {
	AoiMgr *AOIManager
	Player map[int32]*Player
	pLock  sync.RWMutex
}

var WorldManagerObj *WorldManager

func init() {
	WorldManagerObj = &WorldManager{
		Player: make(map[int32]*Player),
		AoiMgr: NewAOIManager(AOI_MIN_X, AOI_MAX_X, AOI_CNTS_X, AOI_MIN_Y, AOI_MAX_Y, AOI_CNTS_Y),
	}
}
func (wm *WorldManager) AddPlayer(player *Player) {
	wm.pLock.Lock()
	wm.Player[player.Pid] = player
	wm.pLock.Unlock()

	wm.AoiMgr.AddToGridByPos(int(player.Pid), player.X, player.Z)
	//slog.Info(fmt.Sprint(len(wm.Player)))
}

func (wm *WorldManager) RemovePlayer(pid int32) {
	wm.pLock.Lock()
	delete(wm.Player, pid)
	wm.pLock.Unlock()

}

func (wm *WorldManager) GetPlayer(pid int32) *Player {
	wm.pLock.RLock()
	defer wm.pLock.RUnlock()
	player, ok := wm.Player[pid]
	if !ok {
		return nil
	}
	return player
}
func (wm *WorldManager) GetAllPlayers() []*Player {
	wm.pLock.RLock()
	defer wm.pLock.RUnlock()
	players := make([]*Player, 0, len(wm.Player))
	for _, player := range wm.Player {
		players = append(players, player)
	}
	return players
}
