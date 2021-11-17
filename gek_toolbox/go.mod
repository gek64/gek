module "gek_toolbox"

require (
	gek_exec v0.0.0
)

replace (
	gek_exec => ../gek_exec
)
