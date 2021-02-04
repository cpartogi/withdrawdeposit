package init

// StartAppInit init all server needs
func StartAppInit() {
	setupLogger()
	setupMainConfig()

	setupAuthHelper()
}
