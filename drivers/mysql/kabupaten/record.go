package kabupaten

import (
	"bansosman/drivers/mysql/daerah"

	"gorm.io/gorm"
)

type Kabs struct {
	gorm.Model
	Kabupaten string
	Daerah    daerah.Daerahs `gorm:"foreignKey:IdKabupaten"`
}
