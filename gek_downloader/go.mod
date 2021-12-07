module "gek_downloader"

require (
	gek_exec v0.0.0
	gek_file v0.0.0
	gek_math v0.0.0
)

replace (
	gek_exec => ../gek_exec
	gek_file => ../gek_file
	gek_math => ../gek_math
)