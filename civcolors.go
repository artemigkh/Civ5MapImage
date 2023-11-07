package main

import (
	"encoding/xml"
	"fmt"
	"image/color"
	"io/ioutil"
	"math"
	"os"

	"github.com/samuelyuan/Civ5MapImage/fileio"
)

type CivColor struct {
	OuterColor color.RGBA
	InnerColor color.RGBA
	TextColor  color.RGBA
}

type ColorXMLFile struct {
	XMLName xml.Name `xml:"GameData"`
	Colors  Colors   `xml:"Colors"`
}

type Colors struct {
	XMLName xml.Name   `xml:"Colors"`
	Rows    []ColorRow `xml:"Row"`
}

type ColorRow struct {
	XMLName xml.Name `xml:"Row"`
	Type    string   `xml:"Type"`
	Red     float64  `xml:"Red"`
	Green   float64  `xml:"Green"`
	Blue    float64  `xml:"Blue"`
	Alpha   float64  `xml:"Alpha"`
}

type PlayerColorXMLFile struct {
	XMLName      xml.Name     `xml:"GameData"`
	PlayerColors PlayerColors `xml:"PlayerColors"`
}

type PlayerColors struct {
	XMLName         xml.Name         `xml:"PlayerColors"`
	PlayerColorRows []PlayerColorRow `xml:"Row"`
}

type PlayerColorRow struct {
	XMLName        xml.Name `xml:"Row"`
	Type           string   `xml:"Type"`
	PrimaryColor   string   `xml:"PrimaryColor"`
	SecondaryColor string   `xml:"SecondaryColor"`
	TextColor      string   `xml:"TextColor"`
}

var (
	colorMap    = initColorMap()
	civColorMap = initCivColorMap()
)

func dumpColors(filename string) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var colorXmlFile ColorXMLFile
	xml.Unmarshal(byteValue, &colorXmlFile)
	for _, row := range colorXmlFile.Colors.Rows {
		fmt.Println(fmt.Sprintf("colorMap[\"%v\"] = color.RGBA{%v, %v, %v, %v}", row.Type,
			math.Round(row.Red*255), math.Round(row.Green*255), math.Round(row.Blue*255), math.Round(row.Alpha*255)))
	}
}

func dumpCivColors(filename string) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var colorXmlFile PlayerColorXMLFile
	xml.Unmarshal(byteValue, &colorXmlFile)
	for _, row := range colorXmlFile.PlayerColors.PlayerColorRows {
		fmt.Println(fmt.Sprintf("civColorMap[\"%v\"] = CivColor{", row.Type))
		fmt.Println(fmt.Sprintf("OuterColor: colorMap[\"%v\"],", row.SecondaryColor))
		fmt.Println(fmt.Sprintf("InnerColor: colorMap[\"%v\"],", row.PrimaryColor))
		fmt.Println(fmt.Sprintf("TextColor: colorMap[\"%v\"],", row.TextColor))
		fmt.Println("}")
	}
}

func initColorMap() map[string]color.RGBA {
	colorMap := make(map[string]color.RGBA)

	colorMap["COLOR_CLEAR"] = color.RGBA{255, 255, 255, 0}
	colorMap["COLOR_ALPHA_GREY"] = color.RGBA{26, 26, 26, 115}
	colorMap["COLOR_WHITE"] = color.RGBA{255, 255, 255, 255}
	colorMap["COLOR_BLACK"] = color.RGBA{0, 0, 0, 255}
	colorMap["COLOR_DARK_GREY"] = color.RGBA{64, 64, 64, 255}
	colorMap["COLOR_GREY"] = color.RGBA{128, 128, 128, 255}
	colorMap["COLOR_LIGHT_GREY"] = color.RGBA{191, 191, 191, 255}
	colorMap["COLOR_GREEN"] = color.RGBA{0, 255, 0, 255}
	colorMap["COLOR_BLUE"] = color.RGBA{0, 0, 255, 255}
	colorMap["COLOR_XP_BLUE"] = color.RGBA{0, 120, 252, 255}
	colorMap["COLOR_CYAN"] = color.RGBA{0, 255, 255, 255}
	colorMap["COLOR_YELLOW"] = color.RGBA{255, 255, 0, 255}
	colorMap["COLOR_MAGENTA"] = color.RGBA{255, 0, 255, 255}
	colorMap["COLOR_YIELD_FOOD"] = color.RGBA{252, 148, 41, 255}
	colorMap["COLOR_YIELD_PRODUCTION"] = color.RGBA{112, 143, 189, 255}
	colorMap["COLOR_YIELD_GOLD"] = color.RGBA{255, 240, 20, 255}
	colorMap["COLOR_CITY_BLUE"] = color.RGBA{18, 117, 204, 255}
	colorMap["COLOR_CITY_GREY"] = color.RGBA{89, 64, 64, 255}
	colorMap["COLOR_CITY_BROWN"] = color.RGBA{179, 115, 0, 255}
	colorMap["COLOR_CITY_GREEN"] = color.RGBA{46, 135, 107, 255}
	colorMap["COLOR_FONT_RED"] = color.RGBA{255, 77, 38, 255}
	colorMap["COLOR_FONT_GREEN"] = color.RGBA{26, 242, 0, 255}
	colorMap["COLOR_RESEARCH_STORED"] = color.RGBA{0, 255, 255, 255}
	colorMap["COLOR_RESEARCH_RATE"] = color.RGBA{0, 255, 255, 153}
	colorMap["COLOR_CULTURE_STORED"] = color.RGBA{153, 0, 255, 255}
	colorMap["COLOR_CULTURE_RATE"] = color.RGBA{153, 0, 255, 153}
	colorMap["COLOR_GREAT_PEOPLE_STORED"] = color.RGBA{255, 255, 0, 255}
	colorMap["COLOR_GREAT_PEOPLE_RATE"] = color.RGBA{255, 255, 0, 153}
	colorMap["COLOR_NEGATIVE_RATE"] = color.RGBA{255, 0, 0, 166}
	colorMap["COLOR_EMPTY"] = color.RGBA{0, 0, 0, 102}
	colorMap["COLOR_POPUP_TEXT"] = color.RGBA{255, 255, 255, 255}
	colorMap["COLOR_POPUP_SELECTED"] = color.RGBA{255, 255, 0, 191}
	colorMap["COLOR_TECH_TEXT"] = color.RGBA{128, 255, 26, 255}
	colorMap["COLOR_UNIT_TEXT"] = color.RGBA{255, 255, 0, 255}
	colorMap["COLOR_BUILDING_TEXT"] = color.RGBA{204, 204, 217, 255}
	colorMap["COLOR_PROJECT_TEXT"] = color.RGBA{204, 204, 217, 255}
	colorMap["COLOR_HIGHLIGHT_TEXT"] = color.RGBA{102, 230, 255, 255}
	colorMap["COLOR_ALT_HIGHLIGHT_TEXT"] = color.RGBA{128, 255, 26, 255}
	colorMap["COLOR_WARNING_TEXT"] = color.RGBA{255, 77, 77, 255}
	colorMap["COLOR_POSITIVE_TEXT"] = color.RGBA{128, 255, 26, 255}
	colorMap["COLOR_NEGATIVE_TEXT"] = color.RGBA{255, 77, 77, 255}
	colorMap["COLOR_BROWN_TEXT"] = color.RGBA{102, 61, 41, 255}
	colorMap["COLOR_SELECTED_TEXT"] = color.RGBA{255, 209, 125, 255}
	colorMap["COLOR_WATER_TEXT"] = color.RGBA{179, 179, 255, 255}
	colorMap["COLOR_MENU_BLUE"] = color.RGBA{71, 212, 242, 255}
	colorMap["COLOR_DAWN_OF_MAN_TEXT"] = color.RGBA{56, 23, 8, 255}
	colorMap["COLOR_ADVISOR_HIGHLIGHT_TEXT"] = color.RGBA{255, 255, 0, 255}
	colorMap["COLOR_TECH_GREEN"] = color.RGBA{41, 179, 69, 128}
	colorMap["COLOR_TECH_BLUE"] = color.RGBA{54, 59, 173, 128}
	colorMap["COLOR_TECH_WORKING"] = color.RGBA{54, 59, 173, 128}
	colorMap["COLOR_TECH_BLACK"] = color.RGBA{0, 0, 0, 128}
	colorMap["COLOR_TECH_RED"] = color.RGBA{255, 0, 0, 128}
	colorMap["COLOR_RED"] = color.RGBA{255, 0, 0, 255}
	colorMap["COLOR_PLAYER_BLACK"] = color.RGBA{33, 33, 33, 255}
	colorMap["COLOR_PLAYER_BLACK_TEXT"] = color.RGBA{204, 206, 217, 255}
	colorMap["COLOR_PLAYER_BLUE"] = color.RGBA{54, 102, 255, 255}
	colorMap["COLOR_PLAYER_LIGHT_BLUE_TEXT"] = color.RGBA{179, 204, 255, 255}
	colorMap["COLOR_PLAYER_BROWN"] = color.RGBA{99, 61, 0, 255}
	colorMap["COLOR_PLAYER_BROWN_TEXT"] = color.RGBA{230, 166, 77, 255}
	colorMap["COLOR_PLAYER_CYAN"] = color.RGBA{18, 204, 245, 255}
	colorMap["COLOR_PLAYER_CYAN_TEXT"] = color.RGBA{153, 255, 248, 255}
	colorMap["COLOR_PLAYER_DARK_BLUE"] = color.RGBA{41, 0, 163, 255}
	colorMap["COLOR_PLAYER_DARK_BLUE_TEXT"] = color.RGBA{166, 140, 230, 255}
	colorMap["COLOR_PLAYER_DARK_CYAN"] = color.RGBA{0, 138, 140, 255}
	colorMap["COLOR_PLAYER_DARK_CYAN_TEXT"] = color.RGBA{0, 212, 201, 255}
	colorMap["COLOR_PLAYER_DARK_GREEN"] = color.RGBA{0, 99, 0, 255}
	colorMap["COLOR_PLAYER_DARK_DARK_GREEN"] = color.RGBA{0, 69, 0, 255}
	colorMap["COLOR_PLAYER_DARK_GREEN_TEXT"] = color.RGBA{143, 204, 143, 255}
	colorMap["COLOR_PLAYER_DARK_PINK"] = color.RGBA{176, 0, 97, 255}
	colorMap["COLOR_PLAYER_DARK_PINK_TEXT"] = color.RGBA{255, 0, 255, 255}
	colorMap["COLOR_PLAYER_DARK_PURPLE"] = color.RGBA{115, 0, 125, 255}
	colorMap["COLOR_PLAYER_DARK_PURPLE_TEXT"] = color.RGBA{204, 89, 217, 255}
	colorMap["COLOR_PLAYER_DARK_RED"] = color.RGBA{158, 0, 0, 255}
	colorMap["COLOR_PLAYER_DARK_RED_TEXT"] = color.RGBA{255, 56, 56, 255}
	colorMap["COLOR_PLAYER_DARK_YELLOW"] = color.RGBA{247, 191, 0, 255}
	colorMap["COLOR_PLAYER_DARK_YELLOW_TEXT"] = color.RGBA{255, 204, 0, 255}
	colorMap["COLOR_PLAYER_GRAY"] = color.RGBA{179, 179, 179, 255}
	colorMap["COLOR_PLAYER_GRAY_TEXT"] = color.RGBA{204, 204, 204, 255}
	colorMap["COLOR_PLAYER_GREEN"] = color.RGBA{125, 224, 0, 255}
	colorMap["COLOR_PLAYER_GREEN_TEXT"] = color.RGBA{124, 225, 0, 255}
	colorMap["COLOR_PLAYER_ORANGE"] = color.RGBA{252, 89, 0, 255}
	colorMap["COLOR_PLAYER_ORANGE_TEXT"] = color.RGBA{254, 117, 0, 255}
	colorMap["COLOR_PLAYER_PEACH"] = color.RGBA{255, 217, 143, 255}
	colorMap["COLOR_PLAYER_PEACH_TEXT"] = color.RGBA{194, 178, 101, 255}
	colorMap["COLOR_PLAYER_PINK"] = color.RGBA{250, 171, 125, 255}
	colorMap["COLOR_PLAYER_PINK_TEXT"] = color.RGBA{250, 184, 145, 255}
	colorMap["COLOR_PLAYER_PURPLE"] = color.RGBA{196, 87, 255, 255}
	colorMap["COLOR_PLAYER_PURPLE_TEXT"] = color.RGBA{217, 166, 255, 255}
	colorMap["COLOR_PLAYER_RED"] = color.RGBA{219, 5, 5, 255}
	colorMap["COLOR_PLAYER_RED_TEXT"] = color.RGBA{255, 76, 106, 255}
	colorMap["COLOR_PLAYER_WHITE"] = color.RGBA{230, 230, 230, 255}
	colorMap["COLOR_PLAYER_WHITE_TEXT"] = color.RGBA{255, 242, 242, 255}
	colorMap["COLOR_PLAYER_YELLOW"] = color.RGBA{255, 255, 43, 255}
	colorMap["COLOR_PLAYER_YELLOW_TEXT"] = color.RGBA{254, 255, 44, 255}
	colorMap["COLOR_PLAYER_LIGHT_GREEN"] = color.RGBA{128, 255, 128, 255}
	colorMap["COLOR_PLAYER_LIGHT_GREEN_TEXT"] = color.RGBA{179, 255, 179, 255}
	colorMap["COLOR_PLAYER_LIGHT_BLUE"] = color.RGBA{128, 179, 255, 255}
	colorMap["COLOR_PLAYER_BLUE_TEXT"] = color.RGBA{128, 179, 255, 255}
	colorMap["COLOR_PLAYER_LIGHT_YELLOW"] = color.RGBA{255, 255, 128, 255}
	colorMap["COLOR_PLAYER_LIGHT_YELLOW_TEXT"] = color.RGBA{230, 255, 128, 255}
	colorMap["COLOR_PLAYER_LIGHT_PURPLE"] = color.RGBA{179, 153, 255, 255}
	colorMap["COLOR_PLAYER_LIGHT_PURPLE_TEXT"] = color.RGBA{179, 153, 255, 255}
	colorMap["COLOR_PLAYER_LIGHT_ORANGE"] = color.RGBA{230, 166, 82, 255}
	colorMap["COLOR_PLAYER_LIGHT_ORANGE_TEXT"] = color.RGBA{255, 191, 89, 255}
	colorMap["COLOR_PLAYER_MIDDLE_PURPLE"] = color.RGBA{172, 30, 185, 255}
	colorMap["COLOR_PLAYER_MIDDLE_PURPLE_TEXT"] = color.RGBA{206, 101, 161, 255}
	colorMap["COLOR_PLAYER_GOLDENROD"] = color.RGBA{222, 159, 0, 255}
	colorMap["COLOR_PLAYER_DARK_GRAY"] = color.RGBA{94, 94, 94, 255}
	colorMap["COLOR_PLAYER_DARK_GRAY_TEXT"] = color.RGBA{144, 144, 144, 255}
	colorMap["COLOR_PLAYER_MIDDLE_GREEN"] = color.RGBA{52, 147, 0, 255}
	colorMap["COLOR_PLAYER_MIDDLE_GREEN_TEXT"] = color.RGBA{62, 157, 0, 255}
	colorMap["COLOR_PLAYER_DARK_LEMON"] = color.RGBA{216, 202, 10, 255}
	colorMap["COLOR_PLAYER_DARK_LEMON_TEXT"] = color.RGBA{231, 217, 25, 255}
	colorMap["COLOR_PLAYER_MIDDLE_BLUE"] = color.RGBA{0, 56, 233, 255}
	colorMap["COLOR_PLAYER_MIDDLE_BLUE_TEXT"] = color.RGBA{129, 185, 181, 255}
	colorMap["COLOR_PLAYER_MIDDLE_CYAN"] = color.RGBA{0, 163, 181, 255}
	colorMap["COLOR_PLAYER_MIDDLE_CYAN_TEXT"] = color.RGBA{40, 203, 221, 255}
	colorMap["COLOR_PLAYER_MAROON"] = color.RGBA{131, 51, 40, 255}
	colorMap["COLOR_PLAYER_LIGHT_BROWN"] = color.RGBA{132, 88, 19, 255}
	colorMap["COLOR_PLAYER_LIGHT_BROWN_TEXT"] = color.RGBA{148, 110, 52, 255}
	colorMap["COLOR_PLAYER_DARK_ORANGE"] = color.RGBA{224, 60, 0, 255}
	colorMap["COLOR_PLAYER_DARK_ORANGE_TEXT"] = color.RGBA{242, 78, 18, 255}
	colorMap["COLOR_PLAYER_DARK_DARK_GREEN_TEXT"] = color.RGBA{91, 159, 91, 255}
	colorMap["COLOR_PLAYER_PALE_RED"] = color.RGBA{199, 72, 61, 255}
	colorMap["COLOR_PLAYER_DARK_INDIGO"] = color.RGBA{78, 5, 213, 255}
	colorMap["COLOR_PLAYER_DARK_INDIGO_TEXT"] = color.RGBA{136, 92, 219, 255}
	colorMap["COLOR_PLAYER_PALE_ORANGE"] = color.RGBA{220, 120, 38, 255}
	colorMap["COLOR_PLAYER_LIGHT_BLACK"] = color.RGBA{64, 64, 64, 255}
	colorMap["COLOR_PLAYER_LIGHT_BLACK_TEXT"] = color.RGBA{100, 100, 100, 255}
	colorMap["COLOR_PLAYER_MINOR_ICON"] = color.RGBA{0, 0, 0, 255}
	colorMap["COLOR_PLAYER_BARBARIAN_ICON"] = color.RGBA{184, 0, 0, 255}
	colorMap["COLOR_PLAYER_AMERICA_ICON"] = color.RGBA{255, 255, 255, 255}
	colorMap["COLOR_PLAYER_ARABIA_ICON"] = color.RGBA{146, 221, 10, 255}
	colorMap["COLOR_PLAYER_AZTEC_ICON"] = color.RGBA{137, 239, 213, 255}
	colorMap["COLOR_PLAYER_CHINA_ICON"] = color.RGBA{255, 255, 255, 255}
	colorMap["COLOR_PLAYER_EGYPT_ICON"] = color.RGBA{83, 0, 208, 255}
	colorMap["COLOR_PLAYER_ENGLAND_ICON"] = color.RGBA{255, 255, 255, 255}
	colorMap["COLOR_PLAYER_FRANCE_ICON"] = color.RGBA{235, 235, 139, 255}
	colorMap["COLOR_PLAYER_GERMANY_ICON"] = color.RGBA{37, 43, 33, 255}
	colorMap["COLOR_PLAYER_GREECE_ICON"] = color.RGBA{65, 141, 254, 255}
	colorMap["COLOR_PLAYER_INDIA_ICON"] = color.RGBA{255, 153, 50, 255}
	colorMap["COLOR_PLAYER_IROQUOIS_ICON"] = color.RGBA{252, 202, 129, 255}
	colorMap["COLOR_PLAYER_JAPAN_ICON"] = color.RGBA{184, 0, 0, 255}
	colorMap["COLOR_PLAYER_OTTOMAN_ICON"] = color.RGBA{18, 82, 30, 255}
	colorMap["COLOR_PLAYER_PERSIA_ICON"] = color.RGBA{245, 230, 55, 255}
	colorMap["COLOR_PLAYER_ROME_ICON"] = color.RGBA{240, 199, 0, 255}
	colorMap["COLOR_PLAYER_RUSSIA_ICON"] = color.RGBA{0, 0, 0, 255}
	colorMap["COLOR_PLAYER_SIAM_ICON"] = color.RGBA{177, 8, 3, 255}
	colorMap["COLOR_PLAYER_SONGHAI_ICON"] = color.RGBA{90, 0, 10, 255}
	colorMap["COLOR_PLAYER_BARBARIAN_BACKGROUND"] = color.RGBA{0, 0, 0, 255}
	colorMap["COLOR_PLAYER_AMERICA_BACKGROUND"] = color.RGBA{31, 51, 120, 255}
	colorMap["COLOR_PLAYER_ARABIA_BACKGROUND"] = color.RGBA{43, 88, 46, 255}
	colorMap["COLOR_PLAYER_AZTEC_BACKGROUND"] = color.RGBA{161, 57, 35, 255}
	colorMap["COLOR_PLAYER_CHINA_BACKGROUND"] = color.RGBA{0, 149, 82, 255}
	colorMap["COLOR_PLAYER_EGYPT_BACKGROUND"] = color.RGBA{255, 252, 3, 255}
	colorMap["COLOR_PLAYER_ENGLAND_BACKGROUND"] = color.RGBA{109, 2, 0, 255}
	colorMap["COLOR_PLAYER_FRANCE_BACKGROUND"] = color.RGBA{65, 141, 254, 255}
	colorMap["COLOR_PLAYER_GERMANY_BACKGROUND"] = color.RGBA{179, 178, 184, 255}
	colorMap["COLOR_PLAYER_GREECE_BACKGROUND"] = color.RGBA{255, 255, 255, 255}
	colorMap["COLOR_PLAYER_INDIA_BACKGROUND"] = color.RGBA{18, 136, 7, 255}
	colorMap["COLOR_PLAYER_IROQUOIS_BACKGROUND"] = color.RGBA{65, 87, 87, 255}
	colorMap["COLOR_PLAYER_JAPAN_BACKGROUND"] = color.RGBA{255, 255, 255, 255}
	colorMap["COLOR_PLAYER_OTTOMAN_BACKGROUND"] = color.RGBA{247, 249, 200, 255}
	colorMap["COLOR_PLAYER_PERSIA_BACKGROUND"] = color.RGBA{177, 8, 3, 255}
	colorMap["COLOR_PLAYER_ROME_BACKGROUND"] = color.RGBA{70, 0, 118, 255}
	colorMap["COLOR_PLAYER_RUSSIA_BACKGROUND"] = color.RGBA{239, 180, 0, 255}
	colorMap["COLOR_PLAYER_SIAM_BACKGROUND"] = color.RGBA{245, 230, 55, 255}
	colorMap["COLOR_PLAYER_SONGHAI_BACKGROUND"] = color.RGBA{214, 145, 19, 255}
	colorMap["COLOR_CLEAR"] = color.RGBA{255, 255, 255, 0}
	colorMap["COLOR_ALPHA_GREY"] = color.RGBA{25, 25, 25, 114}
	colorMap["COLOR_WHITE"] = color.RGBA{255, 255, 255, 255}
	colorMap["COLOR_BLACK"] = color.RGBA{0, 0, 0, 255}
	colorMap["COLOR_DARK_GREY"] = color.RGBA{63, 63, 63, 255}
	colorMap["COLOR_GREY"] = color.RGBA{127, 127, 127, 255}
	colorMap["COLOR_LIGHT_GREY"] = color.RGBA{191, 191, 191, 255}
	colorMap["COLOR_GREEN"] = color.RGBA{0, 255, 0, 255}
	colorMap["COLOR_BLUE"] = color.RGBA{0, 0, 255, 255}
	colorMap["COLOR_XP_BLUE"] = color.RGBA{0, 119, 252, 255}
	colorMap["COLOR_CYAN"] = color.RGBA{0, 255, 255, 255}
	colorMap["COLOR_YELLOW"] = color.RGBA{255, 255, 0, 255}
	colorMap["COLOR_MAGENTA"] = color.RGBA{255, 0, 255, 255}
	colorMap["COLOR_YIELD_FOOD"] = color.RGBA{252, 147, 40, 255}
	colorMap["COLOR_YIELD_PRODUCTION"] = color.RGBA{112, 142, 188, 255}
	colorMap["COLOR_YIELD_GOLD"] = color.RGBA{255, 239, 20, 255}
	colorMap["COLOR_CITY_BLUE"] = color.RGBA{17, 117, 204, 255}
	colorMap["COLOR_CITY_GREY"] = color.RGBA{89, 63, 63, 255}
	colorMap["COLOR_CITY_BROWN"] = color.RGBA{178, 114, 0, 255}
	colorMap["COLOR_CITY_GREEN"] = color.RGBA{45, 135, 107, 255}
	colorMap["COLOR_FONT_RED"] = color.RGBA{255, 76, 38, 255}
	colorMap["COLOR_FONT_GREEN"] = color.RGBA{25, 242, 0, 255}
	colorMap["COLOR_RESEARCH_STORED"] = color.RGBA{0, 255, 255, 255}
	colorMap["COLOR_RESEARCH_RATE"] = color.RGBA{0, 255, 255, 153}
	colorMap["COLOR_CULTURE_STORED"] = color.RGBA{153, 0, 255, 255}
	colorMap["COLOR_CULTURE_RATE"] = color.RGBA{153, 0, 255, 153}
	colorMap["COLOR_GREAT_PEOPLE_STORED"] = color.RGBA{255, 255, 0, 255}
	colorMap["COLOR_GREAT_PEOPLE_RATE"] = color.RGBA{255, 255, 0, 153}
	colorMap["COLOR_NEGATIVE_RATE"] = color.RGBA{255, 0, 0, 165}
	colorMap["COLOR_EMPTY"] = color.RGBA{0, 0, 0, 102}
	colorMap["COLOR_POPUP_TEXT"] = color.RGBA{255, 255, 255, 255}
	colorMap["COLOR_POPUP_SELECTED"] = color.RGBA{255, 255, 0, 191}
	colorMap["COLOR_TECH_TEXT"] = color.RGBA{127, 255, 25, 255}
	colorMap["COLOR_UNIT_TEXT"] = color.RGBA{255, 255, 0, 255}
	colorMap["COLOR_BUILDING_TEXT"] = color.RGBA{204, 204, 216, 255}
	colorMap["COLOR_PROJECT_TEXT"] = color.RGBA{204, 204, 216, 255}
	colorMap["COLOR_HIGHLIGHT_TEXT"] = color.RGBA{102, 229, 255, 255}
	colorMap["COLOR_ALT_HIGHLIGHT_TEXT"] = color.RGBA{127, 255, 25, 255}
	colorMap["COLOR_WARNING_TEXT"] = color.RGBA{255, 76, 76, 255}
	colorMap["COLOR_POSITIVE_TEXT"] = color.RGBA{127, 255, 25, 255}
	colorMap["COLOR_NEGATIVE_TEXT"] = color.RGBA{255, 76, 76, 255}
	colorMap["COLOR_BROWN_TEXT"] = color.RGBA{102, 61, 40, 255}
	colorMap["COLOR_SELECTED_TEXT"] = color.RGBA{255, 209, 124, 255}
	colorMap["COLOR_WATER_TEXT"] = color.RGBA{178, 178, 255, 255}
	colorMap["COLOR_MENU_BLUE"] = color.RGBA{71, 211, 242, 255}
	colorMap["COLOR_DAWN_OF_MAN_TEXT"] = color.RGBA{56, 22, 7, 255}
	colorMap["COLOR_ADVISOR_HIGHLIGHT_TEXT"] = color.RGBA{255, 255, 0, 255}
	colorMap["COLOR_TECH_GREEN"] = color.RGBA{40, 178, 68, 127}
	colorMap["COLOR_TECH_BLUE"] = color.RGBA{53, 58, 173, 127}
	colorMap["COLOR_TECH_WORKING"] = color.RGBA{53, 58, 173, 127}
	colorMap["COLOR_TECH_BLACK"] = color.RGBA{0, 0, 0, 127}
	colorMap["COLOR_TECH_RED"] = color.RGBA{255, 0, 0, 127}
	colorMap["COLOR_RED"] = color.RGBA{255, 0, 0, 255}
	colorMap["COLOR_PLAYER_BLACK"] = color.RGBA{33, 33, 33, 255}
	colorMap["COLOR_PLAYER_BLACK_TEXT"] = color.RGBA{204, 206, 217, 255}
	colorMap["COLOR_PLAYER_BLUE"] = color.RGBA{53, 102, 255, 255}
	colorMap["COLOR_PLAYER_LIGHT_BLUE_TEXT"] = color.RGBA{178, 204, 255, 255}
	colorMap["COLOR_PLAYER_BROWN"] = color.RGBA{99, 61, 0, 255}
	colorMap["COLOR_PLAYER_BROWN_TEXT"] = color.RGBA{229, 165, 76, 255}
	colorMap["COLOR_PLAYER_CYAN"] = color.RGBA{17, 204, 244, 255}
	colorMap["COLOR_PLAYER_CYAN_TEXT"] = color.RGBA{153, 255, 248, 255}
	colorMap["COLOR_PLAYER_DARK_BLUE"] = color.RGBA{40, 0, 163, 255}
	colorMap["COLOR_PLAYER_DARK_BLUE_TEXT"] = color.RGBA{165, 140, 229, 255}
	colorMap["COLOR_PLAYER_DARK_CYAN"] = color.RGBA{0, 137, 140, 255}
	colorMap["COLOR_PLAYER_DARK_CYAN_TEXT"] = color.RGBA{0, 211, 200, 255}
	colorMap["COLOR_PLAYER_DARK_GREEN"] = color.RGBA{0, 99, 0, 255}
	colorMap["COLOR_PLAYER_DARK_DARK_GREEN"] = color.RGBA{0, 68, 0, 255}
	colorMap["COLOR_PLAYER_DARK_GREEN_TEXT"] = color.RGBA{143, 204, 143, 255}
	colorMap["COLOR_PLAYER_DARK_PINK"] = color.RGBA{175, 0, 96, 255}
	colorMap["COLOR_PLAYER_DARK_PINK_TEXT"] = color.RGBA{255, 0, 255, 255}
	colorMap["COLOR_PLAYER_DARK_PURPLE"] = color.RGBA{114, 0, 124, 255}
	colorMap["COLOR_PLAYER_DARK_PURPLE_TEXT"] = color.RGBA{204, 89, 216, 255}
	colorMap["COLOR_PLAYER_DARK_RED"] = color.RGBA{158, 0, 0, 255}
	colorMap["COLOR_PLAYER_DARK_RED_TEXT"] = color.RGBA{255, 56, 56, 255}
	colorMap["COLOR_PLAYER_DARK_YELLOW"] = color.RGBA{247, 191, 0, 255}
	colorMap["COLOR_PLAYER_DARK_YELLOW_TEXT"] = color.RGBA{255, 204, 0, 255}
	colorMap["COLOR_PLAYER_GRAY"] = color.RGBA{178, 178, 178, 255}
	colorMap["COLOR_PLAYER_GRAY_TEXT"] = color.RGBA{204, 204, 204, 255}
	colorMap["COLOR_PLAYER_GREEN"] = color.RGBA{124, 224, 0, 255}
	colorMap["COLOR_PLAYER_GREEN_TEXT"] = color.RGBA{123, 224, 0, 255}
	colorMap["COLOR_PLAYER_ORANGE"] = color.RGBA{252, 89, 0, 255}
	colorMap["COLOR_PLAYER_ORANGE_TEXT"] = color.RGBA{253, 117, 0, 255}
	colorMap["COLOR_PLAYER_PEACH"] = color.RGBA{255, 216, 142, 255}
	colorMap["COLOR_PLAYER_PEACH_TEXT"] = color.RGBA{194, 177, 100, 255}
	colorMap["COLOR_PLAYER_PINK"] = color.RGBA{249, 170, 124, 255}
	colorMap["COLOR_PLAYER_PINK_TEXT"] = color.RGBA{249, 183, 145, 255}
	colorMap["COLOR_PLAYER_PURPLE"] = color.RGBA{196, 86, 255, 255}
	colorMap["COLOR_PLAYER_PURPLE_TEXT"] = color.RGBA{216, 165, 255, 255}
	colorMap["COLOR_PLAYER_RED"] = color.RGBA{219, 5, 5, 255}
	colorMap["COLOR_PLAYER_RED_TEXT"] = color.RGBA{255, 75, 106, 255}
	colorMap["COLOR_PLAYER_WHITE"] = color.RGBA{229, 229, 229, 255}
	colorMap["COLOR_PLAYER_WHITE_TEXT"] = color.RGBA{255, 242, 242, 255}
	colorMap["COLOR_PLAYER_YELLOW"] = color.RGBA{255, 255, 43, 255}
	colorMap["COLOR_PLAYER_YELLOW_TEXT"] = color.RGBA{253, 255, 44, 255}
	colorMap["COLOR_PLAYER_LIGHT_GREEN"] = color.RGBA{127, 255, 127, 255}
	colorMap["COLOR_PLAYER_LIGHT_GREEN_TEXT"] = color.RGBA{178, 255, 178, 255}
	colorMap["COLOR_PLAYER_LIGHT_BLUE"] = color.RGBA{127, 178, 255, 255}
	colorMap["COLOR_PLAYER_BLUE_TEXT"] = color.RGBA{127, 178, 255, 255}
	colorMap["COLOR_PLAYER_LIGHT_YELLOW"] = color.RGBA{255, 255, 127, 255}
	colorMap["COLOR_PLAYER_LIGHT_YELLOW_TEXT"] = color.RGBA{229, 255, 127, 255}
	colorMap["COLOR_PLAYER_LIGHT_PURPLE"] = color.RGBA{178, 153, 255, 255}
	colorMap["COLOR_PLAYER_LIGHT_PURPLE_TEXT"] = color.RGBA{178, 153, 255, 255}
	colorMap["COLOR_PLAYER_LIGHT_ORANGE"] = color.RGBA{229, 165, 81, 255}
	colorMap["COLOR_PLAYER_LIGHT_ORANGE_TEXT"] = color.RGBA{255, 191, 89, 255}
	colorMap["COLOR_PLAYER_MIDDLE_PURPLE"] = color.RGBA{172, 30, 184, 255}
	colorMap["COLOR_PLAYER_MIDDLE_PURPLE_TEXT"] = color.RGBA{206, 100, 160, 255}
	colorMap["COLOR_PLAYER_GOLDENROD"] = color.RGBA{222, 159, 0, 255}
	colorMap["COLOR_PLAYER_DARK_GRAY"] = color.RGBA{94, 94, 94, 255}
	colorMap["COLOR_PLAYER_DARK_GRAY_TEXT"] = color.RGBA{144, 144, 144, 255}
	colorMap["COLOR_PLAYER_MIDDLE_GREEN"] = color.RGBA{52, 146, 0, 255}
	colorMap["COLOR_PLAYER_MIDDLE_GREEN_TEXT"] = color.RGBA{61, 157, 0, 255}
	colorMap["COLOR_PLAYER_DARK_LEMON"] = color.RGBA{215, 201, 9, 255}
	colorMap["COLOR_PLAYER_DARK_LEMON_TEXT"] = color.RGBA{231, 217, 24, 255}
	colorMap["COLOR_PLAYER_MIDDLE_BLUE"] = color.RGBA{0, 56, 233, 255}
	colorMap["COLOR_PLAYER_MIDDLE_BLUE_TEXT"] = color.RGBA{128, 184, 181, 255}
	colorMap["COLOR_PLAYER_MIDDLE_CYAN"] = color.RGBA{0, 162, 181, 255}
	colorMap["COLOR_PLAYER_MIDDLE_CYAN_TEXT"] = color.RGBA{40, 202, 221, 255}
	colorMap["COLOR_PLAYER_MAROON"] = color.RGBA{131, 51, 40, 255}
	colorMap["COLOR_PLAYER_LIGHT_BROWN"] = color.RGBA{132, 87, 19, 255}
	colorMap["COLOR_PLAYER_LIGHT_BROWN_TEXT"] = color.RGBA{147, 109, 52, 255}
	colorMap["COLOR_PLAYER_DARK_ORANGE"] = color.RGBA{223, 59, 0, 255}
	colorMap["COLOR_PLAYER_DARK_ORANGE_TEXT"] = color.RGBA{241, 78, 18, 255}
	colorMap["COLOR_PLAYER_DARK_DARK_GREEN_TEXT"] = color.RGBA{90, 159, 90, 255}
	colorMap["COLOR_PLAYER_PALE_RED"] = color.RGBA{198, 71, 60, 255}
	colorMap["COLOR_PLAYER_DARK_INDIGO"] = color.RGBA{78, 5, 212, 255}
	colorMap["COLOR_PLAYER_DARK_INDIGO_TEXT"] = color.RGBA{135, 92, 219, 255}
	colorMap["COLOR_PLAYER_PALE_ORANGE"] = color.RGBA{220, 120, 37, 255}
	colorMap["COLOR_PLAYER_LIGHT_BLACK"] = color.RGBA{64, 64, 64, 255}
	colorMap["COLOR_PLAYER_LIGHT_BLACK_TEXT"] = color.RGBA{99, 99, 99, 255}
	colorMap["COLOR_PLAYER_MINOR_ICON"] = color.RGBA{0, 0, 0, 255}
	colorMap["COLOR_PLAYER_BARBARIAN_ICON"] = color.RGBA{183, 0, 0, 255}
	colorMap["COLOR_PLAYER_AMERICA_ICON"] = color.RGBA{255, 255, 255, 255}
	colorMap["COLOR_PLAYER_ARABIA_ICON"] = color.RGBA{146, 221, 9, 255}
	colorMap["COLOR_PLAYER_AZTEC_ICON"] = color.RGBA{136, 238, 212, 255}
	colorMap["COLOR_PLAYER_CHINA_ICON"] = color.RGBA{255, 255, 255, 255}
	colorMap["COLOR_PLAYER_EGYPT_ICON"] = color.RGBA{82, 0, 208, 255}
	colorMap["COLOR_PLAYER_ENGLAND_ICON"] = color.RGBA{255, 255, 255, 255}
	colorMap["COLOR_PLAYER_FRANCE_ICON"] = color.RGBA{235, 235, 138, 255}
	colorMap["COLOR_PLAYER_GERMANY_ICON"] = color.RGBA{36, 43, 32, 255}
	colorMap["COLOR_PLAYER_GREECE_ICON"] = color.RGBA{65, 141, 253, 255}
	colorMap["COLOR_PLAYER_INDIA_ICON"] = color.RGBA{255, 153, 49, 255}
	colorMap["COLOR_PLAYER_IROQUOIS_ICON"] = color.RGBA{251, 201, 129, 255}
	colorMap["COLOR_PLAYER_JAPAN_ICON"] = color.RGBA{184, 0, 0, 255}
	colorMap["COLOR_PLAYER_OTTOMAN_ICON"] = color.RGBA{18, 82, 30, 255}
	colorMap["COLOR_PLAYER_PERSIA_ICON"] = color.RGBA{245, 230, 55, 255}
	colorMap["COLOR_PLAYER_ROME_ICON"] = color.RGBA{239, 198, 0, 255}
	colorMap["COLOR_PLAYER_RUSSIA_ICON"] = color.RGBA{0, 0, 0, 255}
	colorMap["COLOR_PLAYER_SIAM_ICON"] = color.RGBA{176, 7, 3, 255}
	colorMap["COLOR_PLAYER_SONGHAI_ICON"] = color.RGBA{90, 0, 9, 255}
	colorMap["COLOR_PLAYER_BARBARIAN_BACKGROUND"] = color.RGBA{0, 0, 0, 255}
	colorMap["COLOR_PLAYER_AMERICA_BACKGROUND"] = color.RGBA{31, 51, 120, 255}
	colorMap["COLOR_PLAYER_ARABIA_BACKGROUND"] = color.RGBA{43, 87, 45, 255}
	colorMap["COLOR_PLAYER_AZTEC_BACKGROUND"] = color.RGBA{161, 57, 34, 255}
	colorMap["COLOR_PLAYER_CHINA_BACKGROUND"] = color.RGBA{0, 148, 82, 255}
	colorMap["COLOR_PLAYER_EGYPT_BACKGROUND"] = color.RGBA{255, 251, 3, 255}
	colorMap["COLOR_PLAYER_ENGLAND_BACKGROUND"] = color.RGBA{108, 2, 0, 255}
	colorMap["COLOR_PLAYER_FRANCE_BACKGROUND"] = color.RGBA{65, 141, 253, 255}
	colorMap["COLOR_PLAYER_GERMANY_BACKGROUND"] = color.RGBA{179, 177, 184, 255}
	colorMap["COLOR_PLAYER_GREECE_BACKGROUND"] = color.RGBA{255, 255, 255, 255}
	colorMap["COLOR_PLAYER_INDIA_BACKGROUND"] = color.RGBA{18, 135, 6, 255}
	colorMap["COLOR_PLAYER_IROQUOIS_BACKGROUND"] = color.RGBA{65, 86, 86, 255}
	colorMap["COLOR_PLAYER_JAPAN_BACKGROUND"] = color.RGBA{255, 255, 255, 255}
	colorMap["COLOR_PLAYER_OTTOMAN_BACKGROUND"] = color.RGBA{247, 248, 199, 255}
	colorMap["COLOR_PLAYER_PERSIA_BACKGROUND"] = color.RGBA{176, 7, 3, 255}
	colorMap["COLOR_PLAYER_ROME_BACKGROUND"] = color.RGBA{70, 0, 118, 255}
	colorMap["COLOR_PLAYER_RUSSIA_BACKGROUND"] = color.RGBA{238, 180, 0, 255}
	colorMap["COLOR_PLAYER_SIAM_BACKGROUND"] = color.RGBA{245, 230, 55, 255}
	colorMap["COLOR_PLAYER_SONGHAI_BACKGROUND"] = color.RGBA{213, 145, 19, 255}
	colorMap["COLOR_PLAYER_MONGOL_ICON"] = color.RGBA{255, 120, 0, 255}
	colorMap["COLOR_PLAYER_MONGOL_BACKGROUND"] = color.RGBA{81, 0, 8, 255}
	colorMap["COLOR_PLAYER_INCA_ICON"] = color.RGBA{6, 159, 119, 255}
	colorMap["COLOR_PLAYER_INCA_BACKGROUND"] = color.RGBA{255, 184, 33, 255}
	colorMap["COLOR_PLAYER_SPAIN_ICON"] = color.RGBA{255, 168, 168, 255}
	colorMap["COLOR_PLAYER_SPAIN_BACKGROUND"] = color.RGBA{83, 26, 26, 255}
	colorMap["COLOR_PLAYER_POLYNESIA_ICON"] = color.RGBA{255, 255, 74, 255}
	colorMap["COLOR_PLAYER_POLYNESIA_BACKGROUND"] = color.RGBA{217, 88, 0, 255}
	colorMap["COLOR_PLAYER_DENMARK_ICON"] = color.RGBA{239, 231, 179, 255}
	colorMap["COLOR_PLAYER_DENMARK_BACKGROUND"] = color.RGBA{108, 42, 20, 255}
	colorMap["COLOR_PLAYER_KOREA_ICON"] = color.RGBA{255, 0, 0, 255}
	colorMap["COLOR_PLAYER_KOREA_BACKGROUND"] = color.RGBA{26, 32, 96, 255}
	colorMap["COLOR_PLAYER_NEBUCHADNEZZAR_ICON"] = color.RGBA{200, 248, 255, 255}
	colorMap["COLOR_PLAYER_NEBUCHADNEZZAR_BACKGROUND"] = color.RGBA{43, 81, 97, 255}
	colorMap["COLOR_PLAYER_AUSTRIA_ICON"] = color.RGBA{255, 255, 255, 255}
	colorMap["COLOR_PLAYER_AUSTRIA_BACKGROUND"] = color.RGBA{234, 0, 0, 255}
	colorMap["COLOR_PLAYER_BYZANTIUM_ICON"] = color.RGBA{60, 0, 108, 255}
	colorMap["COLOR_PLAYER_BYZANTIUM_BACKGROUND"] = color.RGBA{113, 161, 232, 255}
	colorMap["COLOR_PLAYER_CARTHAGE_ICON"] = color.RGBA{80, 0, 136, 255}
	colorMap["COLOR_PLAYER_CARTHAGE_BACKGROUND"] = color.RGBA{204, 204, 204, 255}
	colorMap["COLOR_PLAYER_CELTS_ICON"] = color.RGBA{147, 169, 255, 255}
	colorMap["COLOR_PLAYER_CELTS_BACKGROUND"] = color.RGBA{21, 91, 62, 255}
	colorMap["COLOR_PLAYER_HUNS_ICON"] = color.RGBA{69, 0, 3, 255}
	colorMap["COLOR_PLAYER_HUNS_BACKGROUND"] = color.RGBA{179, 177, 163, 255}
	colorMap["COLOR_PLAYER_MAYA_ICON"] = color.RGBA{23, 62, 65, 255}
	colorMap["COLOR_PLAYER_MAYA_BACKGROUND"] = color.RGBA{197, 140, 98, 255}
	colorMap["COLOR_PLAYER_NETHERLANDS_ICON"] = color.RGBA{255, 255, 255, 255}
	colorMap["COLOR_PLAYER_NETHERLANDS_BACKGROUND"] = color.RGBA{255, 143, 0, 255}
	colorMap["COLOR_PLAYER_SWEDEN_ICON"] = color.RGBA{248, 246, 2, 255}
	colorMap["COLOR_PLAYER_SWEDEN_BACKGROUND"] = color.RGBA{7, 7, 165, 255}
	colorMap["COLOR_FADING_POSITIVE_TEXT"] = color.RGBA{55, 121, 3, 255}
	colorMap["COLOR_FADING_NEGATIVE_TEXT"] = color.RGBA{133, 15, 15, 255}
	colorMap["COLOR_PLAYER_ETHIOPIA_ICON"] = color.RGBA{255, 45, 45, 255}
	colorMap["COLOR_PLAYER_ETHIOPIA_BACKGROUND"] = color.RGBA{1, 39, 14, 255}
	colorMap["COLOR_PLAYER_ASSYRIA_ICON"] = color.RGBA{255, 169, 13, 255}
	colorMap["COLOR_PLAYER_ASSYRIA_BACKGROUND"] = color.RGBA{255, 244, 173, 255}
	colorMap["COLOR_PLAYER_BRAZIL_ICON"] = color.RGBA{41, 83, 44, 255}
	colorMap["COLOR_PLAYER_BRAZIL_BACKGROUND"] = color.RGBA{150, 222, 10, 255}
	colorMap["COLOR_PLAYER_INDONESIA_ICON"] = color.RGBA{159, 47, 29, 255}
	colorMap["COLOR_PLAYER_INDONESIA_BACKGROUND"] = color.RGBA{111, 211, 217, 255}
	colorMap["COLOR_PLAYER_MOROCCO_ICON"] = color.RGBA{40, 179, 80, 255}
	colorMap["COLOR_PLAYER_MOROCCO_BACKGROUND"] = color.RGBA{144, 2, 0, 255}
	colorMap["COLOR_PLAYER_POLAND_ICON"] = color.RGBA{57, 0, 0, 255}
	colorMap["COLOR_PLAYER_POLAND_BACKGROUND"] = color.RGBA{244, 5, 0, 255}
	colorMap["COLOR_PLAYER_PORTUGAL_ICON"] = color.RGBA{4, 20, 124, 255}
	colorMap["COLOR_PLAYER_PORTUGAL_BACKGROUND"] = color.RGBA{255, 255, 255, 255}
	colorMap["COLOR_PLAYER_SHOSHONE_ICON"] = color.RGBA{24, 240, 205, 255}
	colorMap["COLOR_PLAYER_SHOSHONE_BACKGROUND"] = color.RGBA{73, 59, 46, 255}
	colorMap["COLOR_PLAYER_VENICE_ICON"] = color.RGBA{255, 254, 215, 255}
	colorMap["COLOR_PLAYER_VENICE_BACKGROUND"] = color.RGBA{102, 33, 162, 255}
	colorMap["COLOR_PLAYER_ZULU_ICON"] = color.RGBA{106, 50, 24, 255}
	colorMap["COLOR_PLAYER_ZULU_BACKGROUND"] = color.RGBA{255, 232, 214, 255}
	return colorMap
}

func initCivColorMap() map[string]CivColor {
	civColorMap := make(map[string]CivColor)
	civColorMap["PLAYERCOLOR_BLACK"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_WHITE"],
		InnerColor: colorMap["COLOR_PLAYER_BLACK"],
		TextColor:  colorMap["COLOR_PLAYER_BLACK_TEXT"],
	}
	civColorMap["PLAYERCOLOR_BLUE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_WHITE"],
		InnerColor: colorMap["COLOR_PLAYER_BLUE"],
		TextColor:  colorMap["COLOR_PLAYER_BLUE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_BROWN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_YELLOW"],
		InnerColor: colorMap["COLOR_PLAYER_BROWN"],
		TextColor:  colorMap["COLOR_PLAYER_BROWN_TEXT"],
	}
	civColorMap["PLAYERCOLOR_CYAN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_BLACK"],
		InnerColor: colorMap["COLOR_PLAYER_CYAN"],
		TextColor:  colorMap["COLOR_PLAYER_CYAN_TEXT"],
	}
	civColorMap["PLAYERCOLOR_DARK_BLUE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_YELLOW"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_BLUE"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_BLUE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_DARK_CYAN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_WHITE"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_CYAN"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_CYAN_TEXT"],
	}
	civColorMap["PLAYERCOLOR_DARK_GREEN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_YELLOW"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_GREEN"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_GREEN_TEXT"],
	}
	civColorMap["PLAYERCOLOR_DARK_PINK"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_YELLOW"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_PINK"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_PINK_TEXT"],
	}
	civColorMap["PLAYERCOLOR_DARK_PURPLE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_YELLOW"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_PURPLE"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_PURPLE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_DARK_RED"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_YELLOW"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_RED"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_RED_TEXT"],
	}
	civColorMap["PLAYERCOLOR_DARK_YELLOW"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_RED"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_YELLOW"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_YELLOW_TEXT"],
	}
	civColorMap["PLAYERCOLOR_GRAY"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_BLACK"],
		InnerColor: colorMap["COLOR_PLAYER_GRAY"],
		TextColor:  colorMap["COLOR_PLAYER_GRAY_TEXT"],
	}
	civColorMap["PLAYERCOLOR_GREEN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_BLACK"],
		InnerColor: colorMap["COLOR_PLAYER_GREEN"],
		TextColor:  colorMap["COLOR_PLAYER_GREEN_TEXT"],
	}
	civColorMap["PLAYERCOLOR_ORANGE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_WHITE"],
		InnerColor: colorMap["COLOR_PLAYER_ORANGE"],
		TextColor:  colorMap["COLOR_PLAYER_ORANGE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_PEACH"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_BLACK"],
		InnerColor: colorMap["COLOR_PLAYER_PEACH"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}
	civColorMap["PLAYERCOLOR_PINK"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_RED"],
		InnerColor: colorMap["COLOR_PLAYER_PINK"],
		TextColor:  colorMap["COLOR_PLAYER_PINK_TEXT"],
	}
	civColorMap["PLAYERCOLOR_PURPLE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_BLACK"],
		InnerColor: colorMap["COLOR_PLAYER_PURPLE"],
		TextColor:  colorMap["COLOR_PLAYER_PURPLE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_RED"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_WHITE"],
		InnerColor: colorMap["COLOR_PLAYER_RED"],
		TextColor:  colorMap["COLOR_PLAYER_RED_TEXT"],
	}
	civColorMap["PLAYERCOLOR_WHITE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_RED"],
		InnerColor: colorMap["COLOR_PLAYER_WHITE"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_YELLOW"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_BLUE"],
		InnerColor: colorMap["COLOR_PLAYER_YELLOW"],
		TextColor:  colorMap["COLOR_PLAYER_YELLOW_TEXT"],
	}
	civColorMap["PLAYERCOLOR_LIGHT_GREEN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_BLUE"],
		InnerColor: colorMap["COLOR_PLAYER_LIGHT_GREEN"],
		TextColor:  colorMap["COLOR_PLAYER_LIGHT_GREEN_TEXT"],
	}
	civColorMap["PLAYERCOLOR_LIGHT_BLUE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_BLACK"],
		InnerColor: colorMap["COLOR_PLAYER_LIGHT_BLUE"],
		TextColor:  colorMap["COLOR_PLAYER_LIGHT_BLUE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_LIGHT_YELLOW"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_BLACK"],
		InnerColor: colorMap["COLOR_PLAYER_LIGHT_YELLOW"],
		TextColor:  colorMap["COLOR_PLAYER_LIGHT_YELLOW_TEXT"],
	}
	civColorMap["PLAYERCOLOR_LIGHT_PURPLE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_BLACK"],
		InnerColor: colorMap["COLOR_PLAYER_LIGHT_PURPLE"],
		TextColor:  colorMap["COLOR_PLAYER_LIGHT_PURPLE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_LIGHT_ORANGE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_DARK_GREEN"],
		InnerColor: colorMap["COLOR_PLAYER_LIGHT_ORANGE"],
		TextColor:  colorMap["COLOR_PLAYER_LIGHT_ORANGE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_MIDDLE_PURPLE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_GOLDENROD"],
		InnerColor: colorMap["COLOR_PLAYER_MIDDLE_PURPLE"],
		TextColor:  colorMap["COLOR_PLAYER_MIDDLE_PURPLE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_DARK_GRAY"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_YELLOW"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_GRAY"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_GRAY_TEXT"],
	}
	civColorMap["PLAYERCOLOR_MIDDLE_GREEN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_CYAN_TEXT"],
		InnerColor: colorMap["COLOR_PLAYER_MIDDLE_GREEN"],
		TextColor:  colorMap["COLOR_PLAYER_MIDDLE_GREEN_TEXT"],
	}
	civColorMap["PLAYERCOLOR_DARK_LEMON"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_BLACK"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_LEMON"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_LEMON_TEXT"],
	}
	civColorMap["PLAYERCOLOR_MIDDLE_BLUE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_RED_TEXT"],
		InnerColor: colorMap["COLOR_PLAYER_MIDDLE_BLUE"],
		TextColor:  colorMap["COLOR_PLAYER_MIDDLE_BLUE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_MIDDLE_CYAN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MAROON"],
		InnerColor: colorMap["COLOR_PLAYER_MIDDLE_CYAN"],
		TextColor:  colorMap["COLOR_PLAYER_MIDDLE_CYAN_TEXT"],
	}
	civColorMap["PLAYERCOLOR_LIGHT_BROWN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_BLACK"],
		InnerColor: colorMap["COLOR_PLAYER_LIGHT_BROWN"],
		TextColor:  colorMap["COLOR_PLAYER_LIGHT_BROWN_TEXT"],
	}
	civColorMap["PLAYERCOLOR_DARK_ORANGE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_BLACK"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_ORANGE"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_ORANGE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_DARK_DARK_GREEN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_PALE_RED"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_DARK_GREEN"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_DARK_GREEN_TEXT"],
	}
	civColorMap["PLAYERCOLOR_DARK_INDIGO"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_PALE_ORANGE"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_INDIGO"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_INDIGO_TEXT"],
	}
	civColorMap["PLAYERCOLOR_RED_AND_GOLD"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_YELLOW"],
		InnerColor: colorMap["COLOR_PLAYER_RED"],
		TextColor:  colorMap["COLOR_PLAYER_RED_TEXT"],
	}
	civColorMap["PLAYERCOLOR_GOLD_AND_BLACK"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_BLACK"],
		InnerColor: colorMap["COLOR_PLAYER_GOLDENROD"],
		TextColor:  colorMap["COLOR_PLAYER_GOLDENROD"],
	}
	civColorMap["PLAYERCOLOR_GREEN_AND_BLACK"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_BLACK"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_GREEN"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_GREEN_TEXT"],
	}
	civColorMap["PLAYERCOLOR_DARK_CYAN_AND_LEMON"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_LIGHT_YELLOW"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_CYAN"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_CYAN_TEXT"],
	}
	civColorMap["PLAYERCOLOR_BLACK_AND_GREEN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MIDDLE_GREEN"],
		InnerColor: colorMap["COLOR_PLAYER_LIGHT_BLACK"],
		TextColor:  colorMap["COLOR_PLAYER_LIGHT_BLACK_TEXT"],
	}
	civColorMap["PLAYERCOLOR_GREEN_AND_WHITE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_WHITE"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_GREEN"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_CYAN_AND_GRAY"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_GRAY"],
		InnerColor: colorMap["COLOR_PLAYER_CYAN"],
		TextColor:  colorMap["COLOR_PLAYER_CYAN_TEXT"],
	}
	civColorMap["PLAYERCOLOR_DARK_INDIGO_AND_WHITE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_WHITE"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_INDIGO"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_INDIGO_TEXT"],
	}
	civColorMap["PLAYERCOLOR_ORANGE_AND_GREEN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_GREEN"],
		InnerColor: colorMap["COLOR_PLAYER_ORANGE"],
		TextColor:  colorMap["COLOR_PLAYER_ORANGE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_BARBARIAN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_BARBARIAN_BACKGROUND"],
		InnerColor: colorMap["COLOR_PLAYER_BARBARIAN_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}
	civColorMap["PLAYERCOLOR_MINOR_WHITE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_WHITE"],
		InnerColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}
	civColorMap["PLAYERCOLOR_MINOR_GRAY"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_GRAY"],
		InnerColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}
	civColorMap["PLAYERCOLOR_MINOR_BLUE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_BLUE"],
		InnerColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}
	civColorMap["PLAYERCOLOR_MINOR_MIDDLE_BLUE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MIDDLE_BLUE"],
		InnerColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}
	civColorMap["PLAYERCOLOR_MINOR_CYAN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_CYAN"],
		InnerColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}
	civColorMap["PLAYERCOLOR_MINOR_MIDDLE_CYAN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MIDDLE_CYAN"],
		InnerColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}
	civColorMap["PLAYERCOLOR_MINOR_PEACH"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_PEACH"],
		InnerColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}
	civColorMap["PLAYERCOLOR_MINOR_GREEN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_GREEN"],
		InnerColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}
	civColorMap["PLAYERCOLOR_MINOR_LIGHT_GREEN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_LIGHT_GREEN"],
		InnerColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}
	civColorMap["PLAYERCOLOR_MINOR_LIGHT_BLUE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_LIGHT_BLUE"],
		InnerColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}
	civColorMap["PLAYERCOLOR_MINOR_PURPLE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_PURPLE"],
		InnerColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}
	civColorMap["PLAYERCOLOR_MINOR_MIDDLE_PURPLE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MIDDLE_PURPLE"],
		InnerColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}
	civColorMap["PLAYERCOLOR_MINOR_LIGHT_PURPLE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_LIGHT_PURPLE"],
		InnerColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}
	civColorMap["PLAYERCOLOR_MINOR_LIGHT_ORANGE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_LIGHT_ORANGE"],
		InnerColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}
	civColorMap["PLAYERCOLOR_MINOR_YELLOW"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_YELLOW"],
		InnerColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}
	civColorMap["PLAYERCOLOR_MINOR_LIGHT_YELLOW"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_LIGHT_YELLOW"],
		InnerColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}
	civColorMap["PLAYERCOLOR_MINOR_GOLDENROD"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_GOLDENROD"],
		InnerColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}
	civColorMap["PLAYERCOLOR_MINOR_DARK_LEMON"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_LEMON"],
		InnerColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}
	civColorMap["PLAYERCOLOR_AMERICA"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_AMERICA_BACKGROUND"],
		InnerColor: colorMap["COLOR_PLAYER_AMERICA_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_ARABIA"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_ARABIA_BACKGROUND"],
		InnerColor: colorMap["COLOR_PLAYER_ARABIA_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_AZTEC"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_AZTEC_BACKGROUND"],
		InnerColor: colorMap["COLOR_PLAYER_AZTEC_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_CHINA"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_CHINA_BACKGROUND"],
		InnerColor: colorMap["COLOR_PLAYER_CHINA_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_EGYPT"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_EGYPT_BACKGROUND"],
		InnerColor: colorMap["COLOR_PLAYER_EGYPT_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_ENGLAND"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_ENGLAND_BACKGROUND"],
		InnerColor: colorMap["COLOR_PLAYER_ENGLAND_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_FRANCE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_FRANCE_BACKGROUND"],
		InnerColor: colorMap["COLOR_PLAYER_FRANCE_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_GERMANY"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_GERMANY_BACKGROUND"],
		InnerColor: colorMap["COLOR_PLAYER_GERMANY_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_GREECE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_GREECE_BACKGROUND"],
		InnerColor: colorMap["COLOR_PLAYER_GREECE_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_INDIA"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_INDIA_BACKGROUND"],
		InnerColor: colorMap["COLOR_PLAYER_INDIA_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_IROQUOIS"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_IROQUOIS_BACKGROUND"],
		InnerColor: colorMap["COLOR_PLAYER_IROQUOIS_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_JAPAN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_JAPAN_BACKGROUND"],
		InnerColor: colorMap["COLOR_PLAYER_JAPAN_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_OTTOMAN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_OTTOMAN_BACKGROUND"],
		InnerColor: colorMap["COLOR_PLAYER_OTTOMAN_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_PERSIA"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_PERSIA_BACKGROUND"],
		InnerColor: colorMap["COLOR_PLAYER_PERSIA_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_ROME"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_ROME_BACKGROUND"],
		InnerColor: colorMap["COLOR_PLAYER_ROME_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_RUSSIA"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_RUSSIA_BACKGROUND"],
		InnerColor: colorMap["COLOR_PLAYER_RUSSIA_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_SIAM"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_SIAM_BACKGROUND"],
		InnerColor: colorMap["COLOR_PLAYER_SIAM_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}
	civColorMap["PLAYERCOLOR_SONGHAI"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_SONGHAI_BACKGROUND"],
		InnerColor: colorMap["COLOR_PLAYER_SONGHAI_ICON"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_ASSYRIA"] = CivColor{
		OuterColor: color.RGBA{255, 243, 173, 255}, // light yellow
		InnerColor: color.RGBA{255, 168, 12, 255},  // yellow
		TextColor:  color.RGBA{255, 168, 12, 255},  // yellow
	}
	civColorMap["PLAYERCOLOR_AUSTRIA"] = CivColor{
		OuterColor: color.RGBA{234, 0, 0, 255}, // red
		InnerColor: color.RGBA{255, 255, 255, 255},
		TextColor:  color.RGBA{255, 255, 255, 255},
	}
	civColorMap["PLAYERCOLOR_AZTECS"] = CivColor{
		OuterColor: color.RGBA{161, 57, 34, 255},   // red
		InnerColor: color.RGBA{136, 238, 212, 255}, // light blue,
		TextColor:  color.RGBA{136, 238, 212, 255}, // light blue,
	}
	civColorMap["PLAYERCOLOR_BABYLON"] = CivColor{
		OuterColor: color.RGBA{43, 81, 97, 255},    // dark blue
		InnerColor: color.RGBA{200, 248, 255, 255}, // light blue
		TextColor:  color.RGBA{200, 248, 255, 255}, // light blue
	}
	civColorMap["PLAYERCOLOR_BRAZIL"] = CivColor{
		OuterColor: color.RGBA{149, 221, 10, 255}, // light green
		InnerColor: color.RGBA{41, 83, 44, 255},   // dark green
		TextColor:  color.RGBA{41, 83, 44, 255},   // dark green
	}
	civColorMap["PLAYERCOLOR_CELTS"] = CivColor{
		OuterColor: color.RGBA{21, 91, 62, 255},    // dark green
		InnerColor: color.RGBA{147, 169, 255, 255}, // light blue
		TextColor:  color.RGBA{147, 169, 255, 255}, // light blue
	}
	civColorMap["PLAYERCOLOR_ETHIOPIA"] = CivColor{
		OuterColor: color.RGBA{1, 39, 14, 255},   // dark green
		InnerColor: color.RGBA{255, 45, 45, 255}, // red
		TextColor:  color.RGBA{255, 45, 45, 255}, // red
	}
	civColorMap["PLAYERCOLOR_HUNS"] = CivColor{
		OuterColor: color.RGBA{179, 177, 163, 255}, // gray
		InnerColor: color.RGBA{69, 0, 3, 255},      // dark red
		TextColor:  color.RGBA{69, 0, 3, 255},      // dark red
	}
	civColorMap["PLAYERCOLOR_INCA"] = CivColor{
		OuterColor: color.RGBA{255, 184, 33, 255}, // yellow
		InnerColor: color.RGBA{6, 159, 119, 255},  // green
		TextColor:  color.RGBA{6, 159, 119, 255},  // green
	}
	civColorMap["PLAYERCOLOR_INDONESIA"] = CivColor{
		OuterColor: color.RGBA{110, 210, 217, 255}, // cyan
		InnerColor: color.RGBA{158, 46, 28, 255},   // red
		TextColor:  color.RGBA{158, 46, 28, 255},
	}
	civColorMap["PLAYERCOLOR_MAYA"] = CivColor{
		OuterColor: color.RGBA{197, 140, 98, 255}, // yellow
		InnerColor: color.RGBA{23, 62, 65, 255},   // dark blue
		TextColor:  color.RGBA{23, 62, 65, 255},   // dark blue
	}
	civColorMap["PLAYERCOLOR_MONGOL"] = CivColor{
		OuterColor: color.RGBA{81, 0, 8, 255},    // dark red
		InnerColor: color.RGBA{255, 120, 0, 255}, // orange
		TextColor:  color.RGBA{255, 120, 0, 255}, // orange
	}
	civColorMap["PLAYERCOLOR_MOROCCO"] = CivColor{
		OuterColor: color.RGBA{144, 2, 0, 255},   // dark red
		InnerColor: color.RGBA{39, 178, 79, 255}, // green
		TextColor:  color.RGBA{39, 178, 79, 255}, // green
	}
	civColorMap["PLAYERCOLOR_NETHERLANDS"] = CivColor{
		OuterColor: color.RGBA{255, 143, 0, 255},   // orange
		InnerColor: color.RGBA{255, 255, 255, 255}, // white
		TextColor:  color.RGBA{255, 255, 255, 255}, // white
	}
	civColorMap["PLAYERCOLOR_POLAND"] = CivColor{
		OuterColor: color.RGBA{244, 5, 0, 255}, // light red
		InnerColor: color.RGBA{56, 0, 0, 255},  // dark red
		TextColor:  color.RGBA{56, 0, 0, 255},  // dark red
	}
	civColorMap["PLAYERCOLOR_PORTUGAL"] = CivColor{
		OuterColor: color.RGBA{255, 255, 255, 255}, // white
		InnerColor: color.RGBA{3, 20, 124, 255},    // dark blue
		TextColor:  color.RGBA{3, 20, 124, 255},    // dark blue
	}
	civColorMap["PLAYERCOLOR_SPAIN"] = CivColor{
		OuterColor: color.RGBA{83, 26, 26, 255},    // dark red
		InnerColor: color.RGBA{244, 168, 168, 255}, // pink
		TextColor:  color.RGBA{244, 168, 168, 255}, // pink
	}
	civColorMap["PLAYERCOLOR_SWEDEN"] = CivColor{
		OuterColor: color.RGBA{7, 7, 165, 255},   // dark blue
		InnerColor: color.RGBA{248, 246, 2, 255}, // yellow
		TextColor:  color.RGBA{248, 246, 2, 255}, // yellow
	}
	civColorMap["PLAYERCOLOR_ZULU"] = CivColor{
		OuterColor: color.RGBA{255, 231, 213, 255}, // beige
		InnerColor: color.RGBA{106, 49, 24, 255},   // dark red
		TextColor:  color.RGBA{106, 49, 24, 255},   // dark red
	}
	civColorMap["PLAYERCOLOR_BLACK"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_BLACK"],
		InnerColor: colorMap["COLOR_PLAYER_WHITE"],
		TextColor:  colorMap["COLOR_PLAYER_BLACK_TEXT"],
	}

	civColorMap["PLAYERCOLOR_BLUE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_BLUE"],
		InnerColor: colorMap["COLOR_PLAYER_WHITE"],
		TextColor:  colorMap["COLOR_PLAYER_BLUE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_BROWN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_BROWN"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_YELLOW"],
		TextColor:  colorMap["COLOR_PLAYER_BROWN_TEXT"],
	}

	civColorMap["PLAYERCOLOR_CYAN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_CYAN"],
		InnerColor: colorMap["COLOR_PLAYER_BLACK"],
		TextColor:  colorMap["COLOR_PLAYER_CYAN_TEXT"],
	}

	civColorMap["PLAYERCOLOR_DARK_BLUE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_BLUE"],
		InnerColor: colorMap["COLOR_PLAYER_YELLOW"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_BLUE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_DARK_CYAN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_CYAN"],
		InnerColor: colorMap["COLOR_PLAYER_WHITE"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_CYAN_TEXT"],
	}

	civColorMap["PLAYERCOLOR_DARK_GREEN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_GREEN"],
		InnerColor: colorMap["COLOR_PLAYER_YELLOW"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_GREEN_TEXT"],
	}

	civColorMap["PLAYERCOLOR_DARK_PINK"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_PINK"],
		InnerColor: colorMap["COLOR_PLAYER_YELLOW"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_PINK_TEXT"],
	}

	civColorMap["PLAYERCOLOR_DARK_PURPLE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_PURPLE"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_YELLOW"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_PURPLE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_DARK_RED"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_RED"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_YELLOW"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_RED_TEXT"],
	}

	civColorMap["PLAYERCOLOR_DARK_YELLOW"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_YELLOW"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_RED"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_YELLOW_TEXT"],
	}

	civColorMap["PLAYERCOLOR_GRAY"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_GRAY"],
		InnerColor: colorMap["COLOR_PLAYER_BLACK"],
		TextColor:  colorMap["COLOR_PLAYER_GRAY_TEXT"],
	}

	civColorMap["PLAYERCOLOR_GREEN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_GREEN"],
		InnerColor: colorMap["COLOR_PLAYER_BLACK"],
		TextColor:  colorMap["COLOR_PLAYER_GREEN_TEXT"],
	}

	civColorMap["PLAYERCOLOR_ORANGE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_ORANGE"],
		InnerColor: colorMap["COLOR_PLAYER_WHITE"],
		TextColor:  colorMap["COLOR_PLAYER_ORANGE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_PEACH"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_PEACH"],
		InnerColor: colorMap["COLOR_PLAYER_BLACK"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}

	civColorMap["PLAYERCOLOR_PINK"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_PINK"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_RED"],
		TextColor:  colorMap["COLOR_PLAYER_PINK_TEXT"],
	}

	civColorMap["PLAYERCOLOR_PURPLE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_PURPLE"],
		InnerColor: colorMap["COLOR_PLAYER_BLACK"],
		TextColor:  colorMap["COLOR_PLAYER_PURPLE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_RED"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_RED"],
		InnerColor: colorMap["COLOR_PLAYER_WHITE"],
		TextColor:  colorMap["COLOR_PLAYER_RED_TEXT"],
	}

	civColorMap["PLAYERCOLOR_WHITE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_WHITE"],
		InnerColor: colorMap["COLOR_PLAYER_RED"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_YELLOW"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_YELLOW"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_BLUE"],
		TextColor:  colorMap["COLOR_PLAYER_YELLOW_TEXT"],
	}

	civColorMap["PLAYERCOLOR_LIGHT_GREEN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_LIGHT_GREEN"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_BLUE"],
		TextColor:  colorMap["COLOR_PLAYER_LIGHT_GREEN_TEXT"],
	}

	civColorMap["PLAYERCOLOR_LIGHT_BLUE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_LIGHT_BLUE"],
		InnerColor: colorMap["COLOR_PLAYER_BLACK"],
		TextColor:  colorMap["COLOR_PLAYER_LIGHT_BLUE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_LIGHT_YELLOW"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_LIGHT_YELLOW"],
		InnerColor: colorMap["COLOR_PLAYER_BLACK"],
		TextColor:  colorMap["COLOR_PLAYER_LIGHT_YELLOW_TEXT"],
	}

	civColorMap["PLAYERCOLOR_LIGHT_PURPLE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_LIGHT_PURPLE"],
		InnerColor: colorMap["COLOR_PLAYER_BLACK"],
		TextColor:  colorMap["COLOR_PLAYER_LIGHT_PURPLE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_LIGHT_ORANGE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_LIGHT_ORANGE"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_DARK_GREEN"],
		TextColor:  colorMap["COLOR_PLAYER_LIGHT_ORANGE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_MIDDLE_PURPLE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MIDDLE_PURPLE"],
		InnerColor: colorMap["COLOR_PLAYER_GOLDENROD"],
		TextColor:  colorMap["COLOR_PLAYER_MIDDLE_PURPLE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_DARK_GRAY"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_GRAY"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_YELLOW"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_GRAY_TEXT"],
	}

	civColorMap["PLAYERCOLOR_MIDDLE_GREEN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MIDDLE_GREEN"],
		InnerColor: colorMap["COLOR_PLAYER_CYAN_TEXT"],
		TextColor:  colorMap["COLOR_PLAYER_MIDDLE_GREEN_TEXT"],
	}

	civColorMap["PLAYERCOLOR_DARK_LEMON"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_LEMON"],
		InnerColor: colorMap["COLOR_PLAYER_BLACK"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_LEMON_TEXT"],
	}

	civColorMap["PLAYERCOLOR_MIDDLE_BLUE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MIDDLE_BLUE"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_RED_TEXT"],
		TextColor:  colorMap["COLOR_PLAYER_MIDDLE_BLUE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_MIDDLE_CYAN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MIDDLE_CYAN"],
		InnerColor: colorMap["COLOR_PLAYER_MAROON"],
		TextColor:  colorMap["COLOR_PLAYER_MIDDLE_CYAN_TEXT"],
	}

	civColorMap["PLAYERCOLOR_LIGHT_BROWN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_LIGHT_BROWN"],
		InnerColor: colorMap["COLOR_PLAYER_BLACK"],
		TextColor:  colorMap["COLOR_PLAYER_LIGHT_BROWN_TEXT"],
	}

	civColorMap["PLAYERCOLOR_DARK_ORANGE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_ORANGE"],
		InnerColor: colorMap["COLOR_PLAYER_BLACK"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_ORANGE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_DARK_DARK_GREEN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_DARK_GREEN"],
		InnerColor: colorMap["COLOR_PLAYER_PALE_RED"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_DARK_GREEN_TEXT"],
	}

	civColorMap["PLAYERCOLOR_DARK_INDIGO"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_INDIGO"],
		InnerColor: colorMap["COLOR_PLAYER_PALE_ORANGE"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_INDIGO_TEXT"],
	}

	civColorMap["PLAYERCOLOR_RED_AND_GOLD"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_RED"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_YELLOW"],
		TextColor:  colorMap["COLOR_PLAYER_RED_TEXT"],
	}

	civColorMap["PLAYERCOLOR_GOLD_AND_BLACK"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_GOLDENROD"],
		InnerColor: colorMap["COLOR_PLAYER_BLACK"],
		TextColor:  colorMap["COLOR_PLAYER_GOLDENROD"],
	}

	civColorMap["PLAYERCOLOR_GREEN_AND_BLACK"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_GREEN"],
		InnerColor: colorMap["COLOR_PLAYER_BLACK"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_GREEN_TEXT"],
	}

	civColorMap["PLAYERCOLOR_DARK_CYAN_AND_LEMON"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_CYAN"],
		InnerColor: colorMap["COLOR_PLAYER_LIGHT_YELLOW"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_CYAN_TEXT"],
	}

	civColorMap["PLAYERCOLOR_BLACK_AND_GREEN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_LIGHT_BLACK"],
		InnerColor: colorMap["COLOR_PLAYER_MIDDLE_GREEN"],
		TextColor:  colorMap["COLOR_PLAYER_LIGHT_BLACK_TEXT"],
	}

	civColorMap["PLAYERCOLOR_GREEN_AND_WHITE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_GREEN"],
		InnerColor: colorMap["COLOR_PLAYER_WHITE"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_CYAN_AND_GRAY"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_CYAN"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_GRAY"],
		TextColor:  colorMap["COLOR_PLAYER_CYAN_TEXT"],
	}

	civColorMap["PLAYERCOLOR_DARK_INDIGO_AND_WHITE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DARK_INDIGO"],
		InnerColor: colorMap["COLOR_PLAYER_WHITE"],
		TextColor:  colorMap["COLOR_PLAYER_DARK_INDIGO_TEXT"],
	}

	civColorMap["PLAYERCOLOR_ORANGE_AND_GREEN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_ORANGE"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_GREEN"],
		TextColor:  colorMap["COLOR_PLAYER_ORANGE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_BARBARIAN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_BARBARIAN_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_BARBARIAN_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}

	civColorMap["PLAYERCOLOR_MINOR_WHITE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_WHITE"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}

	civColorMap["PLAYERCOLOR_MINOR_GRAY"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_GRAY"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}

	civColorMap["PLAYERCOLOR_MINOR_BLUE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_BLUE"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}

	civColorMap["PLAYERCOLOR_MINOR_MIDDLE_BLUE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_MIDDLE_BLUE"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}

	civColorMap["PLAYERCOLOR_MINOR_CYAN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_CYAN"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}

	civColorMap["PLAYERCOLOR_MINOR_MIDDLE_CYAN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_MIDDLE_CYAN"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}

	civColorMap["PLAYERCOLOR_MINOR_PEACH"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_PEACH"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}

	civColorMap["PLAYERCOLOR_MINOR_GREEN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_GREEN"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}

	civColorMap["PLAYERCOLOR_MINOR_LIGHT_GREEN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_LIGHT_GREEN"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}

	civColorMap["PLAYERCOLOR_MINOR_LIGHT_BLUE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_LIGHT_BLUE"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}

	civColorMap["PLAYERCOLOR_MINOR_PURPLE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_PURPLE"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}

	civColorMap["PLAYERCOLOR_MINOR_MIDDLE_PURPLE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_MIDDLE_PURPLE"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}

	civColorMap["PLAYERCOLOR_MINOR_LIGHT_PURPLE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_LIGHT_PURPLE"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}

	civColorMap["PLAYERCOLOR_MINOR_LIGHT_ORANGE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_LIGHT_ORANGE"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}

	civColorMap["PLAYERCOLOR_MINOR_YELLOW"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_YELLOW"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}

	civColorMap["PLAYERCOLOR_MINOR_LIGHT_YELLOW"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_LIGHT_YELLOW"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}

	civColorMap["PLAYERCOLOR_MINOR_GOLDENROD"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_GOLDENROD"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}

	civColorMap["PLAYERCOLOR_MINOR_DARK_LEMON"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MINOR_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_DARK_LEMON"],
		TextColor:  colorMap["COLOR_PLAYER_PEACH_TEXT"],
	}

	civColorMap["PLAYERCOLOR_AMERICA"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_AMERICA_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_AMERICA_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_ARABIA"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_ARABIA_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_ARABIA_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_AZTEC"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_AZTEC_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_AZTEC_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_CHINA"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_CHINA_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_CHINA_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_EGYPT"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_EGYPT_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_EGYPT_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_ENGLAND"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_ENGLAND_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_ENGLAND_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_FRANCE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_FRANCE_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_FRANCE_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_GERMANY"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_GERMANY_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_GERMANY_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_GREECE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_GREECE_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_GREECE_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_INDIA"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_INDIA_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_INDIA_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_IROQUOIS"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_IROQUOIS_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_IROQUOIS_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_JAPAN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_JAPAN_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_JAPAN_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_OTTOMAN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_OTTOMAN_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_OTTOMAN_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_PERSIA"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_PERSIA_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_PERSIA_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_ROME"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_ROME_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_ROME_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_RUSSIA"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_RUSSIA_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_RUSSIA_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_SIAM"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_SIAM_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_SIAM_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_SONGHAI"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_SONGHAI_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_SONGHAI_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_MONGOL"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MONGOL_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_MONGOL_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_INCA"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_INCA_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_INCA_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_SPAIN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_SPAIN_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_SPAIN_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_POLYNESIA"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_POLYNESIA_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_POLYNESIA_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_DENMARK"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_DENMARK_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_DENMARK_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_KOREA"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_KOREA_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_KOREA_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_BABYLON"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_NEBUCHADNEZZAR_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_NEBUCHADNEZZAR_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_AUSTRIA"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_AUSTRIA_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_AUSTRIA_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_BYZANTIUM"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_BYZANTIUM_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_BYZANTIUM_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_CARTHAGE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_CARTHAGE_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_CARTHAGE_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_CELTS"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_CELTS_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_CELTS_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_HUNS"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_HUNS_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_HUNS_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_MAYA"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MAYA_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_MAYA_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_NETHERLANDS"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_NETHERLANDS_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_NETHERLANDS_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_SWEDEN"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_SWEDEN_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_SWEDEN_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_ETHIOPIA"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_ETHIOPIA_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_ETHIOPIA_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_ASSYRIA"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_ASSYRIA_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_ASSYRIA_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_BRAZIL"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_BRAZIL_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_BRAZIL_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_INDONESIA"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_INDONESIA_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_INDONESIA_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_MOROCCO"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_MOROCCO_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_MOROCCO_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_POLAND"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_POLAND_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_POLAND_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_PORTUGAL"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_PORTUGAL_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_PORTUGAL_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_SHOSHONE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_SHOSHONE_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_SHOSHONE_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_VENICE"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_VENICE_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_VENICE_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}

	civColorMap["PLAYERCOLOR_ZULU"] = CivColor{
		OuterColor: colorMap["COLOR_PLAYER_ZULU_ICON"],
		InnerColor: colorMap["COLOR_PLAYER_ZULU_BACKGROUND"],
		TextColor:  colorMap["COLOR_PLAYER_WHITE_TEXT"],
	}
	return civColorMap
}

func overrideColorMap(civColorOverrides []fileio.CivColorOverride) {
	for _, override := range civColorOverrides {
		civKey := override.CivKey
		outerColor := convertCivColorInfoToRGBA(override.OuterColor)
		innerColor := convertCivColorInfoToRGBA(override.InnerColor)

		civColorMap[civKey] = CivColor{
			OuterColor: outerColor,
			InnerColor: innerColor,
			TextColor:  innerColor,
		}
	}
}

func convertFractionToRGBA(fractionRed float64, fractionGreen float64, fractionBlue float64) color.RGBA {
	return color.RGBA{
		uint8(math.Round(fractionRed * 255)),
		uint8(math.Round(fractionGreen * 255)),
		uint8(math.Round(fractionBlue * 255)),
		255,
	}
}

func convertCivColorInfoToRGBA(civColorInfo fileio.CivColorInfo) color.RGBA {
	if civColorInfo.Model == "constant" {
		return colorMap[civColorInfo.ColorConstant]
	} else if civColorInfo.Model == "fraction" {
		return convertFractionToRGBA(civColorInfo.Red, civColorInfo.Green, civColorInfo.Blue)
	} else if civColorInfo.Model == "rgb255" {
		return color.RGBA{uint8(civColorInfo.Red), uint8(civColorInfo.Green), uint8(civColorInfo.Blue), 255}
	} else {
		return color.RGBA{0, 0, 0, 255}
	}
}
