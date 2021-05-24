package entity

import "time"

type Parking struct {
	Id                string    `gorm:"TYPE:uniqueidentifier;PRIMARY_KEY;NOT NULL;COLUMN:id" json:"id"`
	FechaIngreso      time.Time `gorm:"TYPE:datetime;SIZE:100;NOT NULL;COLUMN:fechaIngreso" json:"fechaIngreso"`
	Placa             string    `gorm:"TYPE:varchar;NOT NULL;COLUMN:placa" json:"placa"`
	Valor             float32   `gorm:"TYPE:float;NOT NULL;COLUMN:valor" json:"valor"`
	NombrePropietario string    `gorm:"TYPE:varchar;NOT NULL;COLUMN:nombrePropietario" json:"nombrePropietario"`
	FechaRegistro     time.Time `gorm:"TYPE:datetime;SIZE:100;NOT NULL;COLUMN:fechaRegistro" json:"fechaRegistro"`
}

func (Parking) TableName() string {
	return "parking"
}
