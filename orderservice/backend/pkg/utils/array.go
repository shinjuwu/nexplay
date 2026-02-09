package utils

func ArrayIntersection[T comparable](arr1, arr2 []T) []T {
	resp := make([]T, 0)
	hash := make(map[T]struct{}, 0)
	for _, item := range arr1 {
		hash[item] = struct{}{}
	}
	for _, item := range arr2 {
		if _, find := hash[item]; find {
			resp = append(resp, item)
		}
	}
	return resp
}
