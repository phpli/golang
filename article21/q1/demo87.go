package main

type StringSet map[string]struct{}

func (ss StringSet) Has(key string) bool {
	_, present := ss[key]
	return present
}

type Age int

func (age *Age) IsNil() bool  {
	return age == nil
}

func (age *Age) Increase() {
	*age++ // 如果age是一个空指针，则此行将产生一个恐慌。
}

func main()  {
	_ = (StringSet(nil)).Has
	_ = ((*Age)(nil)).IsNil
	_ = ((*Age)(nil)).Increase

	_ = (StringSet{}).Has("key")
	_ = (StringSet(nil)).Has("key")
	_ = ((*Age)(nil)).IsNil()
	// 下面这行将产生一个恐慌，但是此恐慌不是在调用方法的时
	// 候产生的，而是在此方法体内解引用空指针的时候产生的。
	((*Age)(nil)).Increase()
}