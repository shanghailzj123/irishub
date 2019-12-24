package types

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// AssetSource
type AssetSource byte

const (
	NATIVE   AssetSource = 0x00 //
	EXTERNAL AssetSource = 0x01 //
)

var (
	// AssetSourceToStringMap
	AssetSourceToStringMap = map[AssetSource]string{
		NATIVE:   "native",
		EXTERNAL: "external",
	}
	// StringToAssetSourceMap
	StringToAssetSourceMap = map[string]AssetSource{
		"native":   NATIVE,
		"external": EXTERNAL,
	}
)

// AssetSourceFromString
func AssetSourceFromString(str string) (AssetSource, error) {
	if source, ok := StringToAssetSourceMap[strings.ToLower(str)]; ok {
		return source, nil
	}
	return AssetSource(0xff), errors.Errorf("'%s' is not a valid token source", str)
}

// IsValidAssetSource
func IsValidAssetSource(source AssetSource) bool {
	_, ok := AssetSourceToStringMap[source]
	return ok
}

// Format
func (source AssetSource) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		s.Write([]byte(fmt.Sprintf("%s", source.String())))
	default:
		s.Write([]byte(fmt.Sprintf("%v", byte(source))))
	}
}

// String
func (source AssetSource) String() string {
	return AssetSourceToStringMap[source]
}

// Marshal needed for protobuf compatibility
func (source AssetSource) Marshal() ([]byte, error) {
	return []byte{byte(source)}, nil
}

// Unmarshal needed for protobuf compatibility
func (source *AssetSource) Unmarshal(data []byte) error {
	*source = AssetSource(data[0])
	return nil
}

// Marshals to JSON using string
func (source AssetSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(source.String())
}

// Unmarshals from JSON assuming Bech32 encoding
func (source *AssetSource) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return nil
	}
	bz2, err := AssetSourceFromString(s)
	if err != nil {
		return err
	}
	*source = bz2
	return nil
}
