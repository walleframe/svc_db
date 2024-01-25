package svc_db

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/walleframe/walle/util"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func RawToInt8(val sql.RawBytes) (int8, error) {
	if len(val) == 0 {
		return 0, nil
	}
	v, err := strconv.ParseInt(util.BytesToString(val), 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(v), nil
}

func RawToInt16(val sql.RawBytes) (int16, error) {
	if len(val) == 0 {
		return 0, nil
	}
	v, err := strconv.ParseInt(util.BytesToString(val), 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(v), nil
}

func RawToInt32(val sql.RawBytes) (int32, error) {
	if len(val) == 0 {
		return 0, nil
	}
	v, err := strconv.ParseInt(util.BytesToString(val), 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(v), nil
}

func RawToInt64(val sql.RawBytes) (int64, error) {
	if len(val) == 0 {
		return 0, nil
	}
	v, err := strconv.ParseInt(util.BytesToString(val), 10, 64)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func RawToUint8(val sql.RawBytes) (uint8, error) {
	if len(val) == 0 {
		return 0, nil
	}
	v, err := strconv.ParseUint(util.BytesToString(val), 10, 8)
	if err != nil {
		return 0, err
	}
	return uint8(v), nil
}

func RawToUint16(val sql.RawBytes) (uint16, error) {
	if len(val) == 0 {
		return 0, nil
	}
	v, err := strconv.ParseUint(util.BytesToString(val), 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(v), nil
}

func RawToUint32(val sql.RawBytes) (uint32, error) {
	if len(val) == 0 {
		return 0, nil
	}
	v, err := strconv.ParseUint(util.BytesToString(val), 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(v), nil
}

func RawToUint64(val sql.RawBytes) (uint64, error) {
	if len(val) == 0 {
		return 0, nil
	}
	v, err := strconv.ParseUint(util.BytesToString(val), 10, 64)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func RawToBool(val sql.RawBytes) (bool, error) {
	if len(val) == 0 {
		return false, nil
	}
	return strconv.ParseBool(util.BytesToString(val))
}

func RawToBinary(val sql.RawBytes) ([]byte, error) {
	if len(val) == 0 {
		return nil, nil
	}
	// 此处需要拷贝内存
	data := make([]byte, len(val))
	copy(data, val)
	return data, nil
}

func RawToFloat32(val sql.RawBytes) (float32, error) {
	if len(val) == 0 {
		return 0, nil
	}
	v, err := strconv.ParseFloat(util.BytesToString(val), 32)
	if err != nil {
		return 0, err
	}
	return float32(v), nil
}

func RawToFloat64(val sql.RawBytes) (float64, error) {
	if len(val) == 0 {
		return 0, nil
	}
	v, err := strconv.ParseFloat(util.BytesToString(val), 64)
	if err != nil {
		return 0, err
	}
	return float64(v), nil
}

func RawToString(val sql.RawBytes) (string, error) {
	if len(val) == 0 {
		return "", nil
	}
	return string(val), nil
}

////////////////////////////////////////////////////////////////////////////////
// timestamp

func RawToStampInt64(val sql.RawBytes) (int64, error) {
	if len(val) == 0 {
		return 0, nil
	}
	t, err := time.Parse(time.DateTime, util.BytesToString(val))
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

func AnyFromStampInt64(val int64) any {
	return time.Unix(val, 0).Format(time.DateTime)
}

////////////////////////////////////////////////////////////////////////////////
//

// use for map
func RawToSlice[T any](val sql.RawBytes) ([]T, error) {
	obj := make([]T, 0, 4)
	err := json.Unmarshal(val, &obj)
	return obj, err
}

// use for map
func RawToMap[K comparable, V any](val sql.RawBytes) (map[K]V, error) {
	obj := make(map[K]V)
	err := json.Unmarshal(val, &obj)
	return obj, err
}

// use for struct
func RawToObject[T any](val sql.RawBytes) (*T, error) {
	var obj T
	err := json.Unmarshal(val, &obj)
	return &obj, err
}

// use for map
func AnyFromSlice[T any](val []T) any {
	data, _ := json.MarshalToString(&val)
	return data
}

// use for map
func AnyFromMap[K comparable, V any](val map[K]V) any {
	data, _ := json.MarshalToString(&val)
	return data
}

// use for struct
func AnyFromObject[T any](val *T) any {
	data, _ := json.MarshalToString(&val)
	return data
}

func AnyFromAny[T any](val T) any {
	return val
}
