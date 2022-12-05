path=$1

gooses=(windows linux darwin)
distsList=$(go tool dist list)

dists=($distsList)

IFS='/'
for dist in "${dists[@]}"; do
  read -ra D <<<"$dist"
  if [[ ! " ${gooses[*]} " =~ ${D[0]} ]]; then
    continue
  fi
  (cd "${path}" && GOOS="${D[0]}" GOARCH="${D[1]}" go build -o "../AusfuehrbaresProgramm/${D[0]}-${D[1]}")
done
