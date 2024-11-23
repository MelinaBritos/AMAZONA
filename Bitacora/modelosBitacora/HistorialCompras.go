package modelosBitacora

import (
	"time"

	"github.com/MelinaBritos/TP-Principal-AMAZONA/Proveedor/modelosProveedor"
	"gorm.io/gorm"
)

type HistorialCompras struct {
	gorm.Model

	RepuestoCompradoID int
	RepuestoComprado   modelosProveedor.Repuesto
	Cantidad           int
	Costo              float32
	FechaCompra        time.Time `gorm:"not null;type:date"`
}
