package main

var Operate func(f func(xs ...interface{}) interface{}, i ...interface{}) interface{} = func(f func(xs ...interface{}) interface{}, i ...interface{}) interface{} {
	return f(i...)
}
var Concat func(xs ...interface{}) interface{} = func(xs ...interface{}) interface{} {
	var res string
	for _, val := range xs {
		if str, ok := val.(string); ok {
			res += str
		} else {
			continue
		}
	}
	return res
}
var Sum func(xs ...interface{}) interface{} = func(xs ...interface{}) interface{} {
	var res float64
	for _, val := range xs {
		switch v := val.(type) {
		case int:
			res += float64(v)

		case float64:
			res += v
		default:
			continue
		}
	}
	switch xs[0].(type) {
	case int:
		return int(res)

	default:
		return res
	}

}
