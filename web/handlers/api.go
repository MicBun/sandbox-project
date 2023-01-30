package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"sandbox/service"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ApiHandlerInterface interface {
	Hello(c *gin.Context)
	GetOrders(c *gin.Context)
}

type apiHandler struct {
	container *service.Container
}

func NewApiHandler(container *service.Container) ApiHandlerInterface {
	return &apiHandler{
		container: container,
	}
}

// Hello godoc
//
//	@Summary		Hello
//	@Description	Hello
//	@Tags			Hello
//	@Produce		json
//	@Success		200	{object}	string
//	@Router			/hello [get]
func (h *apiHandler) Hello(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello"})
	return
}

// orderResponse
type orderResponse struct {
	ID                  uint
	TrackingNumber      string
	ConsigneeAddress    string
	ConsigneeCity       string
	ConsigneeProvince   string
	ConsigneePostalCode string
	ConsigneeCountry    string
	Weight              float32
	Height              float32
	Width               float32
	Length              float32
}

// GetOrders godoc
//
//	@Summary		GetOrders
//	@Description	GetOrders
//	@Tags			GetOrders
//	@Produce		json
//	@Success		200	{object}	orderResponse
//	@Router			/orders [get]
func (h *apiHandler) GetOrders(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 1
	}
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		offset = 0
	}
	orders, err := h.container.OrdersManager.ListOrders(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp []orderResponse
	for _, o := range orders {
		resp = append(resp, orderResponse{
			ID:                  o.ID,
			TrackingNumber:      o.TrackingNumber,
			ConsigneeAddress:    o.ConsigneeAddress,
			ConsigneeCity:       o.ConsigneeCity,
			ConsigneeProvince:   o.ConsigneeProvince,
			ConsigneePostalCode: o.ConsigneePostalCode,
			ConsigneeCountry:    o.ConsigneeCountry,
			Weight:              o.Weight,
			Height:              o.Height,
			Width:               o.Width,
			Length:              o.Length,
		})
	}

	// function untuk max page

	c.JSON(200, gin.H{"data": resp, "count": )})
	return
}

// for pokemon
func GetJSON(url string, target interface{}) error {
	timestamp := time.Now().Unix()
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("timestamp", strconv.FormatInt(timestamp, 10))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if time.Now().Unix()-timestamp > 60 {
		return errors.New("stale request")
	}
	return json.NewDecoder(resp.Body).Decode(target)
}

type Pokemon struct {
	Name      string   `json:"name"`
	Types     []string `json:"types"`
	Abilities []string `json:"abilities"`
}

type PokemonGetterInterface interface {
	GetPokemon(name string) (Pokemon, error)
}

type PokemonGetter struct {
	response struct {
		Types []struct {
			Type struct {
				Name string `json:"name"`
			} `json:"type"`
		} `json:"types"`
		Abilities []struct {
			Ability struct {
				Name string `json:"name"`
			} `json:"ability"`
		} `json:"abilities"`
	}
}

func NewPokemonGetter() PokemonGetterInterface {
	return &PokemonGetter{}
}

// GetPokemon gets a pokemon by name and returns its types and abilities
func (p PokemonGetter) GetPokemon(name string) (Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + name
	if err := GetJSON(url, &p.response); err != nil {
		return Pokemon{}, err
	}

	// Extract types
	var types []string
	for _, t := range p.response.Types {
		types = append(types, t.Type.Name)
	}

	// Extract abilities
	var abilities []Ability
	for _, ability := range p.response.Abilities {
		abilityName := ability.Ability.Name
		abilityRes, err := NewAbilityGetter().GetAbility(abilityName)
		if err != nil {
			return Pokemon{}, err
		}
		abilities = append(abilities, abilityRes)
	}

	// combine abilities name and description into a single string for each ability
	var abilityNamesAndDescriptions []string
	for _, ability := range abilities {
		abilityNamesAndDescriptions = append(abilityNamesAndDescriptions, ability.Name+" ("+ability.Description+")")
	}

	return Pokemon{
		Name:      name,
		Types:     types,
		Abilities: abilityNamesAndDescriptions,
	}, nil
}

type Ability struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type AbilityGetterInterface interface {
	GetAbility(name string) (Ability, error)
}

type AbilityGetter struct {
	response struct {
		EffectEntries []struct {
			Language struct {
				Name string `json:"name"`
			} `json:"language"`
			ShortEffect string `json:"short_effect"`
		} `json:"effect_entries"`
	}
}

func NewAbilityGetter() AbilityGetterInterface {
	return &AbilityGetter{}
}

// GetAbility gets an ability by name
func (a AbilityGetter) GetAbility(name string) (Ability, error) {
	url := "https://pokeapi.co/api/v2/ability/" + name
	if err := GetJSON(url, &a.response); err != nil {
		return Ability{}, err
	}

	// Extract description
	var description string
	for _, effect := range a.response.EffectEntries {
		if effect.Language.Name == "en" {
			description = effect.ShortEffect
		}
	}

	return Ability{
		Name:        name,
		Description: description,
	}, nil
}

// GetPokemon godoc
// @Summary Get a pokemon
// @Description Get a pokemon by name
// @Tags pokemon
// @Produce  json
// @Param q query string true "Pokemon name substitute spaces with dashes"
// @Success 200 {object} Pokemon
// @Failure 400 {object} string
// @Router /pokemon [get]
func GetPokemon(ctx *gin.Context) {
	name := strings.ToLower(ctx.Query("q"))
	pokemon, err := NewPokemonGetter().GetPokemon(name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, pokemon)
}
