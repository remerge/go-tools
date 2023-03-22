package typconv

type ConvertableNumber interface {
	int | int32 | int64 | uint32 | uint64 | float32 | float64
}

func SliceToSlice[T, U ConvertableNumber](v []T) []U {
	if len(v) == 0 {
		return nil
	}

	out := make([]U, len(v))
	for i, vv := range v {
		out[i] = U(vv)
	}
	return out
}
