module github.com/MouseHatGames/hat/cmd/client

go 1.15

replace github.com/MouseHatGames/hat/pkg/client => ../../pkg/client

replace github.com/MouseHatGames/hat => ../../

require (
	github.com/MouseHatGames/hat/pkg/client v0.0.0-00010101000000-000000000000
	github.com/alecthomas/kong v0.2.12
)
