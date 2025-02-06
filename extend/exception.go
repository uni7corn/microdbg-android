package extend

func (ex *extend) defineException() {
	AndroidException := ex.cf.DefineClass("android.util.AndroidException")
	_ = AndroidException
}
