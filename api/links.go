package api

import (
	"net/http"

	"github.com/juliankoehn/mchurl/stores/shared"
	"github.com/labstack/echo/v4"
)

// LinkDelete deletes a single entry
func (a *API) LinkDelete(c echo.Context) error {
	l := new(shared.Entry)
	if err := c.Bind(l); err != nil {
		return err
	}
	if l.ID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing ID in Request")
	}

	if err := a.store.DeleteEntry(l.ID); err != nil {
		return err
	}
	return c.NoContent(204)
}

// LinkCreate creates a new Link
func (a *API) LinkCreate(c echo.Context) error {
	l := new(shared.Entry)
	if err := c.Bind(l); err != nil {
		return err
	}

	link, err := a.store.CreateEntry(shared.Entry{
		URL: l.URL,
	}, l.ID)
	if err != nil {
		return err
	}
	return c.JSON(200, link)
}

// LinksList lists all links
func (a *API) LinksList(c echo.Context) error {
	links, err := a.store.LinksList()
	if err != nil {
		// handle error
	}

	return c.JSON(200, links)
}

// LinkUpdate updates link url
func (a *API) LinkUpdate(c echo.Context) error {
	l := new(shared.Entry)
	if err := c.Bind(l); err != nil {
		return err
	}
	link, err := a.store.LinkUpdate(l)
	if err != nil {
		return err
	}
	return c.JSON(200, link)
}