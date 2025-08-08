package renderers

import (
	"github.com/adm87/finch-core/geometry"
	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteRenderer struct {
	image  *ebiten.Image
	anchor geometry.Point64
}

func NewSpriteRenderer(image *ebiten.Image, anchorX, anchorY float64) *SpriteRenderer {
	return &SpriteRenderer{
		image: image,
		anchor: geometry.Point64{
			X: anchorX * float64(image.Bounds().Dx()),
			Y: anchorY * float64(image.Bounds().Dy()),
		},
	}
}

func (r *SpriteRenderer) Dispose() {
	r.image = nil
}

func (r *SpriteRenderer) Render(buffer *ebiten.Image, view, transform ebiten.GeoM) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	op.GeoM.Translate(-r.anchor.X, -r.anchor.Y)
	op.GeoM.Concat(transform)
	op.GeoM.Concat(view)
	buffer.DrawImage(r.image, op)
	return nil
}
