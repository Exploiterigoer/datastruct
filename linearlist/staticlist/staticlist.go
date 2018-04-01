package staticlist

// MAXSIZE length of StaticList
const MAXSIZE = 10

// StaticList  struct of static list
type StaticList struct {
	data   interface{}
	cursur int
}

// InitStaticList Inits static list
func InitStaticList() *[]StaticList {
	sl := make([]StaticList, MAXSIZE)

	for i := 0; i < MAXSIZE; i++ {
		sl[i].cursur = i + 1
	}

	sl[MAXSIZE-1].cursur = 0

	return &sl
}
