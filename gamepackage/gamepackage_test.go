package gamepackage

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGamePackage_Get(t *testing.T) {

	jsonData, err := os.ReadFile("testdata/gamepackage.json")
	assert.NoError(t, err)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}))
	defer ts.Close()

	t.Run("Get Valid Answer Mock", func(t *testing.T) {
		var response Response
		launcher := Launcher{
			Url:        ts.URL,
			LauncherID: "VYTpXlbWo8",
			GameID:     "4ziysqXOQ8",
		}
		err = Get(launcher, &response)
		assert.NoError(t, err)
		assert.Equal(t, "4ziysqXOQ8", response.Data.GamePackages[0].Game.ID)
	})
	t.Run("Get Empty Url Mock", func(t *testing.T) {
		var response Response
		launcher := Launcher{
			Url:        "",
			LauncherID: "",
			GameID:     "4ziysqXOQ8",
		}
		err = Get(launcher, &response)
		assert.Error(t, err)
	})
	t.Run("Get Empty Launcher Mock", func(t *testing.T) {
		var response Response
		launcher := Launcher{
			Url:        ts.URL,
			LauncherID: "",
			GameID:     "4ziysqXOQ8",
		}
		err = Get(launcher, &response)
		assert.Error(t, err)
	})

}
