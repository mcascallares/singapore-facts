package controllers

import (
	"github.com/jgraham909/revmgo"
	"github.com/mcascallares/singapore-facts/app/models"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
	revmgo.MongoController
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) RetrievePois(layers []string) revel.Result {
	pois, err := models.GetPoisByLayer(c.MongoSession, layers)
	if err != nil {
		return c.NotFound("Error")
	}
	return c.RenderJson(pois)
}
