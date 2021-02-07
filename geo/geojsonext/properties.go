package geojsonext

import (
	"github.com/cockroachdb/errors"
)

// PropertyBool ...
func PropertyBool(props map[string]interface{}, key string, def ...bool) (bool, error) {
	v := props[key]
	if b, ok := v.(bool); ok {
		return b, nil
	}
	if v != nil {
		return false, errors.Newf("not a bool, but a %T: %v", v, v)
	}
	if len(def) > 0 {
		return def[0], nil
	}
	return false, errors.Errorf("property not found")
}

// PropertyInt ...
func PropertyInt(props map[string]interface{}, key string, def ...int) (int, error) {
	v := props[key]
	if i, ok := v.(int); ok {
		return i, nil
	}
	if f, ok := v.(float64); ok {
		return int(f), nil
	}
	if v != nil {
		return 0, errors.Newf("not a number, but a %T: %v", v, v)
	}
	if len(def) > 0 {
		return def[0], nil
	}
	return 0, errors.Errorf("property not found")
}

// PropertyFloat64 ...
func PropertyFloat64(props map[string]interface{}, key string, def ...float64) (float64, error) {
	v := props[key]
	if f, ok := v.(float64); ok {
		return f, nil
	}
	if i, ok := v.(int); ok {
		return float64(i), nil
	}
	if v != nil {
		return 0, errors.Newf("not a number, but a %T: %v", v, v)
	}
	if len(def) > 0 {
		return def[0], nil
	}
	return 0, errors.Errorf("property not found")
}

// PropertyString ...
func PropertyString(props map[string]interface{}, key string, def ...string) (string, error) {
	v := props[key]
	if s, ok := v.(string); ok {
		return s, nil
	}
	if v != nil {
		return "", errors.Newf("not a string, but a %T: %v", v, v)
	}
	if len(def) > 0 {
		return def[0], nil
	}
	return "", errors.Errorf("property not found")
}
