project_dir=$(cd $(dirname $0)/.. && pwd)
target_dir="$project_dir/target"
echo "Target dir: $target_dir"

if [ -d "$target_dir" ]; then
    echo "removing old target dir ($target_dir)"
    rm -rf $target_dir
fi
echo "creating new target dir ($target_dir)"
mkdir -p $target_dir

build() {
    local os="$1"
    local arch="$2"
    local name="$3"
    echo "building $os/$arch distributable $name"
    os_dir="$target_dir/$os-$arch"
    mkdir "$os_dir"
    GOOS="$os" GOARCH="$arch" go build -o "$os_dir/$name" main.go
    (cd "$os_dir" && tar -zcf "$os-$arch.tar.gz" "$name")
}

build "linux" "amd64" "preftime"
build "windows" "amd64" "preftime.exe"
build "darwin" "amd64" "preftime"
