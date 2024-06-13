package tempconv

func CToK(c Celsius) Kelvins {
	return Kelvins(c - AbsoluteZeroC)
}

func KToC(k Kelvins) Celsius {
	return Celsius(k + Kelvins(AbsoluteZeroC))
}
