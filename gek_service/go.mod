module "gek_service"

require gek_exec v0.0.0
require gek_toolbox v0.0.0

replace gek_exec => ./../gek_exec
replace gek_toolbox => ./../gek_toolbox
replace gek_github => ./../gek_github
replace gek_json => ./../gek_json
replace gek_file => ./../gek_file