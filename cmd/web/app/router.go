package app

func (a *Application) Routes() {
	api := a.Router.Group("/api/v1/")
	api.GET("/health", HealthHandler)
	api.GET("/slow", SlowHandler)
}
