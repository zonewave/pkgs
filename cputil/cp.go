package cputil

import (
	"github.com/jinzhu/copier"
)

// ShallowCopy  not deepCopy
func ShallowCopy(toValue interface{}, fromValue interface{}) error {
	return copier.CopyWithOption(toValue, fromValue, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    false,
	})
}

// DeepCopy deep copies
func DeepCopy(toValue interface{}, fromValue interface{}) error {
	return copier.CopyWithOption(toValue, fromValue, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
	})
}
