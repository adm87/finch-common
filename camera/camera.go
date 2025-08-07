package camera

import (
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/geometry"
	"github.com/adm87/finch-core/hash"
)

var CameraComponentType = ecs.ComponentType(hash.GetHashFromType[CameraComponent]())

type CameraComponent struct {
	size geometry.Point64
	zoom float64
}

func NewCameraComponent() *CameraComponent {
	return NewCameraComponentWith(
		geometry.Point64{X: 0, Y: 0},
		1.0,
	)
}

func NewCameraComponentWith(size geometry.Point64, zoom float64) *CameraComponent {
	return &CameraComponent{
		size: size,
		zoom: zoom,
	}
}

func (c *CameraComponent) Type() ecs.ComponentType {
	return CameraComponentType
}

func (c *CameraComponent) Size() geometry.Point64 {
	return c.size
}

func (c *CameraComponent) SetSize(size geometry.Point64) {
	if c.size.X == size.X && c.size.Y == size.Y {
		return
	}
	c.size = size
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
