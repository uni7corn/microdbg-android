package android

type pthread[P any] struct {
	next     P
	prev     P
	tid      uint32
	_filling [30]P
	_buf     [512]byte
}

type pthread_key_data[P any] struct {
	seq  P
	data P
}

type pthread_mutex32 struct {
	__private [1]int32
}

type pthread_mutex64 struct {
	__private [10]int32
}

type pthread_rwlock32 struct {
	__private [10]int32
}

type pthread_rwlock64 struct {
	__private [14]int32
}
