package generator

import (
	"testing"

	"github.com/muka/go-bluetooth/gen"
	"github.com/muka/go-bluetooth/gen/util"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {

	TplPath = "../../gen/generator/tpl/%s.go.tpl"

	bluezApi, err := gen.Parse("../../src/bluez/doc", []string{}, false)
	if err != nil {
		t.Fatal(err)
	}

	outdir := "../../test/out"

	err = util.Mkdir("../../test")
	if err != nil {
		t.Fatal(err)
	}
	err = util.Mkdir(outdir)
	if err != nil {
		t.Fatal(err)
	}

	err = Generate(bluezApi, outdir, true)
	if err != nil {
		t.Fatal(err)
	}

	assert.DirExists(t, outdir)
	assert.DirExists(t, outdir+"/profile/adapter")

}
