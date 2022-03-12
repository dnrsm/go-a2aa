package dir

import (
	"github.com/spf13/afero"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	type args struct {
		s   []string
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "normal",
			args: args{s: []string{".mp3", ".aac"}, str: ".mp3"},
			want: true,
		},
		{
			name: "fail",
			args: args{s: []string{".mp3", ".aac"}, str: "xxx"},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Contains(tc.args.s, tc.args.str); got != tc.want {
				t.Errorf("Contains() = %v, want %v", got, tc.want)
			}
		})

	}
}

func TestMatch(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "normal(mp3)",
			args: args{file: ".mp3"},
			want: true,
		},
		{
			name: "normal(aac)",
			args: args{file: ".aac"},
			want: true,
		},
		{
			name: "fail",
			args: args{file: ".xxx"},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Match(tc.args.file); got != tc.want {
				t.Errorf("Match() = %v, want %v", got, tc.want)
			}
		})

	}
}

func TestCurrent(t *testing.T) {
	appFs := afero.NewMemMapFs()
	tmpDir := "./Music"

	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want []string
		mock func()
	}{
		{
			name: "normal",
			args: args{path: tmpDir},
			want: []string{"Music/a.mp3", "Music/b.mp3", "Music/c.mp3"},
			mock: func() {
				_ = afero.WriteFile(appFs, filepath.Join(tmpDir, "a.mp3"), []byte(nil), os.ModePerm)
				_ = afero.WriteFile(appFs, filepath.Join(tmpDir, "b.mp3"), []byte(nil), os.ModePerm)
				_ = afero.WriteFile(appFs, filepath.Join(tmpDir, "c.mp3"), []byte(nil), os.ModePerm)
			},
		},
		{
			name: "normal(empty)",
			args: args{path: tmpDir},
			want: []string{},
			mock: func() {},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Cleanup(func() {
				_ = appFs.RemoveAll(tmpDir)
			})
			tc.mock()
			if got, _ := Current(appFs, tc.args.path); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Current() = %v, want %v", got, tc.want)
			}
		})

	}
}

func TestAll(t *testing.T) {
	appFs := afero.NewMemMapFs()
	fileSystem := afero.NewIOFS(appFs)

	tmpDir1 := "./Music"
	tmpDir2 := "./Music/bar"
	tmpDir3 := "./Music/baz"

	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want []string
		mock func()
	}{
		{
			name: "normal",
			args: args{path: "./"},
			want: []string{"Music/a.mp3", "Music/bar/b.mp3", "Music/baz/c.mp3"},
			mock: func() {
				_ = afero.WriteFile(appFs, filepath.Join(tmpDir1, "a.mp3"), []byte(nil), os.ModePerm)
				_ = afero.WriteFile(appFs, filepath.Join(tmpDir2, "b.mp3"), []byte(nil), os.ModePerm)
				_ = afero.WriteFile(appFs, filepath.Join(tmpDir3, "c.mp3"), []byte(nil), os.ModePerm)
			},
		},
		{
			name: "normal(empty)",
			args: args{path: "./"},
			want: nil,
			mock: func() {},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Cleanup(func() {
				_ = appFs.RemoveAll("./")
			})
			tc.mock()
			if got, _ := All(fileSystem, tc.args.path); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("All() = %v, want %v", got, tc.want)
			}
		})

	}
}
