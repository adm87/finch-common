package camera

import (
	"github.com/adm87/finch-application/config"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/errors"
	"github.com/adm87/finch-core/hash"
	"github.com/adm87/finch-core/transform"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	CameraLateUpdateSystemType   = ecs.SystemType(hash.GetHashFromType[CameraLateUpdate]())
	CameraLateUpdateSystemFilter = []ecs.ComponentType{
		transform.TransformComponentType,
		CameraComponentType,
	}
)

type CameraLateUpdate struct {
	world              *ecs.World
	window             *config.Window
	cameraEntity       *ecs.Entity
	cameraComponent    *CameraComponent
	transformComponent *transform.TransformComponent
}

func NewCameraLateUpdate(world *ecs.World, window *config.Window) *CameraLateUpdate {
	return &CameraLateUpdate{
		world:  world,
		window: window,
	}
}

func (s *CameraLateUpdate) Type() ecs.SystemType {
	return CameraLateUpdateSystemType
}

func (s CameraLateUpdate) Filter() []ecs.ComponentType {
	return CameraLateUpdateSystemFilter
}

func (s *CameraLateUpdate) LateUpdate(entities []*ecs.Entity, deltaSeconds float64) error {
	if len(entities) == 0 {
		return nil
	}
	if len(entities) > 1 {
		return errors.NewAmbiguousError("multiple camera entities found, expected only one")
	}

	if s.cameraEntity == nil || s.cameraEntity.ID() != entities[0].ID() {
		if err := s.CacheOperationComponents(entities[0]); err != nil {
			return err
		}
	}

	if s.cameraComponent == nil || s.transformComponent == nil {
		return errors.NewNotFoundError("cannot operate on camera without camera and transform components")
	}

	// TODO: Added matrix caching here so we don't have to recalculate it every frame there is no change to the camera.

	zoom := s.cameraComponent.Zoom()

	_, yoff := ebiten.Wheel()
	if yoff != 0 {
		zoom += yoff * 0.1
		if zoom < 0.01 {
			zoom = 0.01
		}
		s.cameraComponent.SetZoom(zoom)
	}

	scale := s.transformComponent.Scale()
	origin := s.transformComponent.Origin()

	scale.X = zoom
	scale.Y = zoom

	origin.X = float64(s.window.ScreenWidth) / 2
	origin.Y = float64(s.window.ScreenHeight) / 2

	s.transformComponent.SetScale(scale)
	s.transformComponent.SetOrigin(origin)

	invWorldMatrix := s.transformComponent.WorldMatrix()
	invWorldMatrix.Invert()

	s.world.SetView(invWorldMatrix)

	return nil
}

func (s *CameraLateUpdate) CacheOperationComponents(entity *ecs.Entity) error {
	if entity == nil {
		return errors.NewNilError("camera entity cannot be nil")
	}

	s.cameraEntity = entity

	component, _, err := entity.GetComponent(CameraComponentType)
	if err != nil {
		return err
	}

	cameraComponent, ok := component.(*CameraComponent)
	if !ok {
		return errors.NewNotFoundError("camera component not found in entity")
	}

	s.cameraComponent = cameraComponent

	component, _, err = entity.GetComponent(transform.TransformComponentType)
	if err != nil {
		return err
	}

	transformComponent, ok := component.(*transform.TransformComponent)
	if !ok {
		return errors.NewNotFoundError("transform component not found in entity")
	}

	s.transformComponent = transformComponent

	return nil
}
