package camera

import (
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/hash"
)

var CameraComponentType = ecs.ComponentType(hash.GetHashFromType[CameraComponent]())

type CameraComponent struct {
	zoom float64
}

func NewCameraComponent() *CameraComponent {
	return NewCameraComponentWith(
		1.0,
	)
}

func NewCameraComponentWith(zoom float64) *CameraComponent {
	return &CameraComponent{
		zoom: zoom,
	}
}

func (c *CameraComponent) Type() ecs.ComponentType {
	return CameraComponentType
}

func (c *CameraComponent) Zoom() float64 {
	return c.zoom
}

func (c *CameraComponent) SetZoom(zoom float64) {
	if zoom <= 0 {
		panic("zoom must be greater than 0")
	}
	c.zoom = zoom
}
