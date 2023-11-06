package task

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
)

type Action struct {
	Human
}

type Human struct {
	HealthPoints    int
	HappinessPoints int
	Expeirience     []interface{}
}

func (h *Human) ReadBook(content string) {
	contentRows := strings.Fields(content)
	pureContent := strings.Join(contentRows, " ")
	h.Expeirience = append(h.Expeirience, pureContent)
}

func (h *Human) GoSkiing() {
	if isSuccess(90) {
		h.Expeirience = append(h.Expeirience, struct {
			RidingPoints interface{}
			Track        interface{}
			Landscape    interface{}
		}{
			RidingPoints: 10,
			Track:        "track_path.map",
			Landscape:    "mountains.png",
		})
		h.HappinessPoints++
		h.HealthPoints++

		return
	}

	h.HealthPoints -= 5
	h.HappinessPoints -= 5

	h.Expeirience = append(h.Expeirience, struct {
		RidingPoints   interface{}
		ImportantNotes interface{}
	}{
		RidingPoints:   20,
		ImportantNotes: "hire a trainer",
	})
}

func (h Human) Print() {
	dHumanJSON, _ := json.MarshalIndent(h, "", "	")
	fmt.Println(string(dHumanJSON))
}

func isSuccess(probability int) bool {
	r := rand.Intn(100)

	return r < probability
}

func HumanActionEmbeded() {
	action := Action{
		Human{
			HealthPoints:    1000,
			HappinessPoints: 1000,
			Expeirience:     []interface{}{},
		},
	}

	action.Print()

	action.ReadBook(`
		My uncle's goodness is extreme, 
	If seriously he hath disease; 
	He hath acquired the world's esteem 
	And nothing more important sees...`)

	action.GoSkiing()

	action.Print()
}
