module main

go 1.18

// директивой replace указываем положение корня 
// модуля somemodule относительно main/go.mod
replace somemodule => ../somemodule  