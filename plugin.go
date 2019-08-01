package plugin

func Run() {
	initRouter()

	quit := make(chan bool)
	<-quit
}
