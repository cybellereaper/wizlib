package wizlib

import (
	"encoding/json"
)

type ItemCardInfobox struct {
	School                string `json:"school"`
	PipCost               string `json:"pipcost"`
	ShadowPipCost         string `json:"shadpipcost"`
	SchoolPipCost         string `json:"schoolpipcost"`
	Accuracy              string `json:"accuracy"`
	Type                  string `json:"type"`
	Description           string `json:"descrip1"`
	Variable              string `json:"var"`
	VariableSchool        string `json:"varschool"`
	VariablePipCost       string `json:"varpipcost"`
	VariableShadPipCost   string `json:"varshadpipcost"`
	VariableSchoolPipCost string `json:"varschoolpipcost"`
	VariableAccuracy      string `json:"varaccuracy"`
	VariableType          string `json:"vartype"`
	VariableDescription   string `json:"vardescrip1"`
}

type TreasureCardInfobox struct {
	School        string `json:"school"`
	PipCost       string `json:"pipcost"`
	ShadowPipCost string `json:"shadpipcost"`
	SchoolPipCost string `json:"schoolpipcost"`
	Accuracy      string `json:"accuracy"`
	Type          string `json:"type"`
	Type2         string `json:"type2"`
	Monstrology   string `json:"monstrology"`
	PvP           string `json:"PvP"`
	PvPLevel      string `json:"PvPlevel"`
	Description   string `json:"descrip1"`
	BuyLevel      string `json:"buylevel"`
	BuyRank       string `json:"buyrank"`
	Enchantable   string `json:"enchantable"`
	FishChestLoc1 string `json:"fishchestloc1"`
	FishChestLoc2 string `json:"fishchestloc2"`
	FishChestLoc3 string `json:"fishchestloc3"`
}

type LocationInfobox struct {
	World           string `json:"world"`
	Location        string `json:"location"`
	Subloc          string `json:"subloc"`
	Instance        string `json:"instance"`
	Classic         string `json:"classic"`
	Description     string `json:"descrip"`
	Connect1        string `json:"connect1"`
	Connect2        string `json:"connect2"`
	Connect3        string `json:"connect3"`
	Connect4        string `json:"connect4"`
	Connect5        string `json:"connect5"`
	Reagent1        string `json:"reagent1"`
	Lev1Creature1   string `json:"lev1creature1"`
	Lev1Creature2   string `json:"lev1creature2"`
	Lev1Description string `json:"lev1descrip"`
	Lev2Creature1   string `json:"lev2creature1"`
	Lev2Creature2   string `json:"lev2creature2"`
	Lev2Description string `json:"lev2descrip"`
}

type FishInfobox struct {
	School        string `json:"school"`
	Rank          string `json:"rank"`
	Aquarium      string `json:"aquarium"`
	MinSize       string `json:"minsize"`
	MaxSize       string `json:"maxsize"`
	InitialXP     string `json:"initialxp"`
	RegularXP     string `json:"regularxp"`
	Rarity        string `json:"rarity"`
	Description   string `json:"descrip"`
	MinSell       string `json:"minsell"`
	MinSellSize   string `json:"minsellsize"`
	MaxSellSF     string `json:"maxsellSF"`
	FishLocation1 string `json:"fishlocation1"`
	FishLocation2 string `json:"fishlocation2"`
	FishLocation3 string `json:"fishlocation3"`
}

type MountInfobox struct {
	School      string `json:"school"`
	PipCost     string `json:"pipcost"`
	Description string `json:"descrip1"`
}

type SpellInfobox struct {
	School        string `json:"school"`
	PipCost       string `json:"pipcost"`
	Accuracy      string `json:"accuracy"`
	Type          string `json:"type"`
	Description   string `json:"descrip1"`
	ObtainedFrom  string `json:"obtainedfrom"`
	ObtainedAt    string `json:"obtainedat"`
	Price         string `json:"price"`
	SellPrice     string `json:"sellprice"`
	Cooldown      string `json:"cooldown"`
	Requirement1  string `json:"req1"`
	Requirement2  string `json:"req2"`
	Requirement3  string `json:"req3"`
	Requirement4  string `json:"req4"`
	Requirement5  string `json:"req5"`
	Requirement6  string `json:"req6"`
	Requirement7  string `json:"req7"`
	Requirement8  string `json:"req8"`
	Requirement9  string `json:"req9"`
	Requirement10 string `json:"req10"`
}

type HousingInfobox struct {
	World          string `json:"world"`
	Location       string `json:"location"`
	SalePrice      string `json:"saleprice"`
	RentGold       string `json:"rentgold"`
	RentCrowns     string `json:"rentcrowns"`
	PurchaseRank   string `json:"purchaserank"`
	PurchaseGold   string `json:"purchasegold"`
	PurchaseCrowns string `json:"purchasecrowns"`
	SellPrice      string `json:"sellprice"`
	MaxAllowed     string `json:"maxallowed"`
	HouseSize      string `json:"housesize"`
	FloorCount     string `json:"floorcount"`
	Indoor         string `json:"indoor"`
	Outdoor        string `json:"outdoor"`
	Patio          string `json:"patio"`
	Description    string `json:"descrip"`
}

type JewelInfobox struct {
	School       string `json:"school"`
	MaxInc       string `json:"maxinc"`
	Description  string `json:"descrip1"`
	ObtainedFrom string `json:"obtainedfrom"`
	Price        string `json:"price"`
	SellPrice    string `json:"sellprice"`
	RecipeCard   string `json:"recipecard"`
	RecipeVendor string `json:"recipevendor"`
}

type QuestInfobox struct {
	QuestGiver  string `json:"questgiver"`
	Goal        string `json:"goal"`
	Reward      string `json:"reward"`
	NextQuest   string `json:"nextquest"`
	GiveText    string `json:"givetext"`
	Dialogue    string `json:"dialogue"`
	Description string `json:"descrip1"`
}

type CreatureInfobox struct {
	School      string `json:"school"`
	Rank        string `json:"rank"`
	Tower       string `json:"tower"`
	Gold        string `json:"gold"`
	Description string `json:"descrip"`
	Spell1      string `json:"spell1"`
	Spell2      string `json:"spell2"`
	Spell3      string `json:"spell3"`
}

type MinionInfobox struct {
	School      string `json:"school"`
	Rank        string `json:"rank"`
	Description string `json:"descrip1"`
}

func Parser(data string) (t interface{}, err error) {
	err = json.Unmarshal([]byte(data), &t)
	return
}
