package raster3d

import "ndsemu/emu/fixed"

// Swap buffers (marker of end-of-frame, with double-buffering)
type Primitive_SwapBuffers struct {
	AlphaYSort bool
	WBuffering bool
}

// New viewport, in pixel coordinates (0-255 / 0-191)
type Primitive_SetViewport struct {
	VX0, VY0, VX1, VY1 int
}

// New vertex to be pushed in Vertex RAM, with coordinates in
// clip space (after model-view-proj)
type Primitive_Vertex struct {
	X, Y, Z, W fixed.F12 // coordinates in clip-space
	S, T       fixed.F12 // texture coordinates
	C          [3]uint8  // vertex color (RGB 555)
}

// New polygon to be pushed in Polygon RAM
type Primitive_Polygon struct {
	Vtx  [4]int  // indices of vertices in Vertex RAM
	Attr uint32  // misc flags
	Tex  Texture // texture for this polygon
}
