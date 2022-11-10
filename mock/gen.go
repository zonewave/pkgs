package mock

//go:generate mockgen --build_flags=--mod=mod -destination aferomock/fs_mock.go -package=aferomock github.com/spf13/afero Fs
//go:generate mockgen --build_flags=--mod=mod -destination aferomock/file_mock.go -package=aferomock github.com/spf13/afero File
//go:generate mockgen -destination aferomock/afero_mock.go -package=aferomock github.com/zonewave/pkgs/util/fileutil Afero
