package initialization

import (
	"github.com/novliang/lurkers/global"
	"github.com/novliang/lurkers/storage"
)

func Storage() {
	s := &storage.Storage{
		global.GIANT_DB,
	}
	global.GIANT_STORAGE = s
}
