module gek_app

require (
	gek_downloader v0.0.0
	gek_exec v0.0.0
	gek_file v0.0.0
	gek_github v0.0.0
	gek_service_systemd v0.0.0
	gek_service_rc v0.0.0
)

require (
	gek_json v0.0.0 // indirect
	gek_math v0.0.0 // indirect
)

replace (
	gek_downloader => ../gek_downloader
	gek_exec => ../gek_exec
	gek_file => ../gek_file
	gek_github => ../gek_github
	gek_json => ../gek_json
	gek_math => ../gek_math
	gek_service_systemd => ../gek_service_systemd
	gek_service_rc => ../gek_service_rc
)
