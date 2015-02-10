package parc

func ByteEq(b byte) Predict {
	return func(input []byte) (int, bool) {
		if len(input) > 0 && input[0] == b {
			return 1, true
		}
		return 0, false
	}
}
func ByteIn(bs []byte) Predict {
	return func(input []byte) (int, bool) {
		if len(input) > 0 {
			c := input[0]
			for _, c2 := range bs {
				if c == c2 {
					return 1, true
				}
			}
		}
		return 0, false
	}
}
