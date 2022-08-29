go build -o fsbb.exe ./cmd/web/. || exit /b
fsbb.exe -dbname=fsbb -dbuser=postgres -dbpass=solomon7 -cache=false -production=false