package api

func Init() {
	On.Room = make(map[string]*Room)
	On.Client = make(map[string]*Client)
}
