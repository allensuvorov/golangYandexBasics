module main

go 1.18

// директивой replace указываем положение корня
// модуля somemodule относительно main/go.mod
replace somemodule => ../somemodule

require somemodule v0.0.0-00010101000000-000000000000 // indirect
