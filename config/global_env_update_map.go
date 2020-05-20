package config

func GlobalEnvUpdateMap() map[string]string {
	m := make(map[string]string)

	m["AppName"] = "APP_NAME"
	m["AppVersion"] = "APP_VERSION"
	m["ServerHost"] = "SERVER_HOST"
	m["ServerPort"] = "SERVER_PORT"
	m["LogConfFile"] = "LOG_CONF_FILE"
	return m
}
