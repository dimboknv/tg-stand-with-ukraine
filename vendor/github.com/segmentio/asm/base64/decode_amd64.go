// Code generated by command: go run decode_asm.go -pkg base64 -out ../base64/decode_amd64.s -stubs ../base64/decode_amd64.go. DO NOT EDIT.

//go:build !purego
// +build !purego

package base64

func decodeAVX2(dst []byte, src []byte, lut *int8) (int, int)

func decodeAVX2URI(dst []byte, src []byte, lut *int8) (int, int)
