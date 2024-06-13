package config

type ConfigSetting struct {
	// If true, each incoming request will be logged in terminal.
	// Be default - true.
	LogIncomingRequests bool
}

var Settings *ConfigSetting = &ConfigSetting{
	LogIncomingRequests: true,
}
