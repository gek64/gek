module "gek_toolbox"

require gek_exec v0.0.0
require gek_github v0.0.0

replace gek_exec => ../gek_exec
replace gek_github => ../gek_github
replace gek_json => ../gek_json
replace gek_file => ../gek_file