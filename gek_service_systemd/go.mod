module gek_service_systemd

require (
	gek_exec v0.0.0
    gek_file v0.0.0
)

replace (
	gek_exec => ../gek_exec
    gek_file => ../gek_file
)
