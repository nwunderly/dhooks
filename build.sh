os_list=("windows" "linux" "darwin")
arch_list=("amd64")

for os in "${os_list[@]}"
do
  if [ "$os" == "windows" ]; then
    file="dhooks.exe"
  else
    file="dhooks"
  fi

  for arch in "${arch_list[@]}"
  do
    GOOS=$os GOARCH=$arch go build -o "$file" .
    zip "./build/dhooks_$1_${os}_${arch}.zip" "$file"
    rm "$file"
  done
done
