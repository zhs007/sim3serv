package sim3core

import "errors"

var (
	// ErrInvalidTileType - invalid TileType
	ErrInvalidTileType = errors.New("invalid TileType")
	// ErrInvalidRandResult - invalid RandResult
	ErrInvalidRandResult = errors.New("invalid RandResult")
	// ErrInvalidMapTileID - invalid MapTileID
	ErrInvalidMapTileID = errors.New("invalid MapTileID")
	// ErrInvalidGenMapParams - invalid GenMapParams
	ErrInvalidGenMapParams = errors.New("invalid GenMapParams")

	// ErrInvalidFuncGetTargetDir - invalid FuncGetTargetDir
	ErrInvalidFuncGetTargetDir = errors.New("invalid FuncGetTargetDir")
)
