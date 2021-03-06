// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

import (
	"fmt"

	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

type shaderUniform struct {
	name        string
	size        int
	ty          ShaderDataType
	location    int32
	textureUnit int
}

func (u *shaderUniform) Bind(context *Context, v interface{}) {
	transpose := false
	switch u.ty {
	case FLOAT_MAT2x3:
		gl.UniformMatrix2x3fv(u.location, 1, transpose, &v.([]float32)[0])
	case FLOAT_MAT2x4:
		gl.UniformMatrix2x4fv(u.location, 1, transpose, &v.([]float32)[0])
	case FLOAT_MAT2:
		gl.UniformMatrix2fv(u.location, 1, transpose, &v.([]float32)[0])
	case FLOAT_MAT3x2:
		gl.UniformMatrix3x2fv(u.location, 1, transpose, &v.([]float32)[0])
	case FLOAT_MAT3x4:
		gl.UniformMatrix3x4fv(u.location, 1, transpose, &v.([]float32)[0])
	case FLOAT_MAT3:
		switch m := v.(type) {
		case math.Mat3:
			gl.UniformMatrix3fv(u.location, 1, transpose, &m[0])
		case []float32:
			gl.UniformMatrix3fv(u.location, 1, transpose, &m[0])
		}
	case FLOAT_MAT4x2:
		gl.UniformMatrix4x2fv(u.location, 1, transpose, &v.([]float32)[0])
	case FLOAT_MAT4x3:
		gl.UniformMatrix4x3fv(u.location, 1, transpose, &v.([]float32)[0])
	case FLOAT_MAT4:
		gl.UniformMatrix4fv(u.location, 1, transpose, &v.([]float32)[0])
	case FLOAT_VEC1:
		switch v := v.(type) {
		case float32:
			gl.Uniform1f(u.location, v)
		case []float32:
			gl.Uniform1fv(u.location, int32(len(v)), &v[0])
		}
	case FLOAT_VEC2:
		switch v := v.(type) {
		case math.Vec2:
			gl.Uniform2fv(u.location, 1, &[]float32{v.X, v.Y}[0])
		case []float32:
			if len(v)%2 != 0 {
				panic(fmt.Errorf("Uniform '%s' of type vec2 should be an float32 array with a multiple of two length", u.name))
			}
			gl.Uniform2fv(u.location, int32(len(v)/2), &v[0])
		}
	case FLOAT_VEC3:
		switch v := v.(type) {
		case math.Vec3:
			gl.Uniform3fv(u.location, 1, &[]float32{v.X, v.Y, v.Z}[0])
		case []float32:
			if len(v)%3 != 0 {
				panic(fmt.Errorf("Uniform '%s' of type vec3 should be an float32 array with a multiple of three length", u.name))
			}
			gl.Uniform3fv(u.location, int32(len(v)/3), &v[0])
		}
	case FLOAT_VEC4:
		switch v := v.(type) {
		case math.Vec4:
			gl.Uniform4fv(u.location, 1, &[]float32{v.X, v.Y, v.Z, v.W}[0])
		case gxui.Color:
			gl.Uniform4fv(u.location, 1, &[]float32{v.R, v.G, v.B, v.A}[0])
		case []float32:
			if len(v)%4 != 0 {
				panic(fmt.Errorf("Uniform '%s' of type vec4 should be an float32 array with a multiple of four length", u.name))
			}
			gl.Uniform4fv(u.location, int32(len(v)/4), &v[0])
		}
	case SAMPLER_2D:
		ss := v.(SamplerSource)
		gl.ActiveTexture(gl.TEXTURE0 + uint32(u.textureUnit))
		gl.BindTexture(gl.TEXTURE_2D, ss.Texture())
		gl.Uniform1i(u.location, int32(u.textureUnit))
	default:
		panic(fmt.Errorf("Uniform of unsupported type %s", u.ty))
	}
}
