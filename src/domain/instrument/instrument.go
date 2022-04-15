// Package instrument holds entities of instrument.
package instrument

// Instument contain instument data.
type Instument struct {
	InstumentID   int32  `json:"instument_id"`
	InstumentName string `json:"instument_name"`
}
