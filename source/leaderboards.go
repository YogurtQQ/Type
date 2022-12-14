package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Two_Slices struct {
	WPM      float64
	Username string
	Date     string
}

func Tops(s *discordgo.Session, m *discordgo.MessageCreate) {
	var What = strings.Replace(m.Content, ".tops ", "", -1)

	var WHERE int
	for i := 0; i < len(Texts); i++ {
		if strings.HasPrefix(Texts[i], What) {
			WHERE = i
		}
	}

	var WHERE_str = strconv.Itoa(WHERE)
	var FOUND bool = false
	var FOUND_how_many_times int = 0

	var ts []Two_Slices

	var slice Two_Slices

	Load()
	for i := 0; i < len(DB); i++ {
		if strings.HasPrefix(DB[i], WHERE_str+" #") {
			if i == 0 {

			} else {
				var CL = strings.Split(DB[i], " # ")
				WPM_f64, _ := strconv.ParseFloat(CL[5], 8)

				slice = Two_Slices{
					WPM_f64, CL[2], CL[4],
				}
				ts = append(ts, slice)
			}
			FOUND = true
			FOUND_how_many_times++
		}
	}

	sort.Slice(ts, func(i, j int) bool {
		return ts[i].WPM > ts[j].WPM
	})

	var DISPLAY string
	var C int = 1
	for i := 0; i < 10; i++ {
		if i != 0 {
			if i > len(ts)-1 {
				var C_str = strconv.Itoa(C)
				DISPLAY = DISPLAY + C_str + "." + "\n"
			} else {
				var C_str = strconv.Itoa(C)
				var WPM_str = fmt.Sprintf("%.1f", ts[i].WPM)
				DISPLAY = DISPLAY + C_str + ". " + ts[i].Username + " (" + WPM_str + ") " + ts[i].Date + "\n"
			}
		} else {
			if len(ts) > 0 {
				var C_str = strconv.Itoa(C)
				var WPM_str = fmt.Sprintf("%.1f", ts[i].WPM)
				DISPLAY = DISPLAY + C_str + ". " + ts[i].Username + " (" + WPM_str + ") " + ts[i].Date + "\n"
			}
		}
		C++
	}

	var FOUND_how_many_times_str = strconv.Itoa(FOUND_how_many_times)
	if FOUND {
		var Texts_arr = strings.Split(Texts[WHERE], " ")
		var λ string
		for i := 0; i < len(Texts_arr); i++ {
			if i != len(Texts_arr)-1 {
				λ = λ + Texts_arr[i] + "\u200b "
			} else {
				λ = λ + Texts_arr[i]
			}
		}
		s.ChannelMessageSend(m.ChannelID, "```"+λ+"```")
		s.ChannelMessageSend(m.ChannelID, "```diff\n+ Se encontraron "+FOUND_how_many_times_str+" marcas del texto ID:"+WHERE_str+"\n"+DISPLAY+"```")
	} else {
		s.ChannelMessageSend(m.ChannelID, "```diff\n- No se encontró o no existe.```")
	}

}

func Top(s *discordgo.Session, m *discordgo.MessageCreate, Text_ID int) {
	var What = Texts[Text_ID]

	var WHERE int
	for i := 0; i < len(Texts); i++ {
		if strings.HasPrefix(Texts[i], What) {
			WHERE = i
		}
	}

	var WHERE_str = strconv.Itoa(WHERE)

	var ts []Two_Slices

	var slice Two_Slices

	Load()
	for i := 0; i < len(DB); i++ {
		if strings.HasPrefix(DB[i], WHERE_str+" #") {
			if i == 0 {

			} else {
				var CL = strings.Split(DB[i], " # ")

				var WPM_f64 float64

				if len(CL) > 5 {
					WPM_f64, _ = strconv.ParseFloat(CL[5], 8)
				} else {
					WPM_f64, _ = strconv.ParseFloat(CL[3], 8)
				}

				slice = Two_Slices{
					WPM_f64, CL[2], CL[4],
				}
				ts = append(ts, slice)
			}
		}
	}

	sort.Slice(ts, func(i, j int) bool {
		return ts[i].WPM > ts[j].WPM
	})

	var DISPLAY string
	var C int = 1
	for i := 0; i < 5; i++ {
		if i != 0 {
			if i > len(ts)-1 {
				var C_str = strconv.Itoa(C)
				DISPLAY = DISPLAY + C_str + "." + "\n"
			} else {
				var C_str = strconv.Itoa(C)
				var WPM_str = fmt.Sprintf("%.1f", ts[i].WPM)
				DISPLAY = DISPLAY + C_str + ". " + ts[i].Username + " (" + WPM_str + ") " + ts[i].Date + "\n"
			}
		} else {
			if len(ts) > 0 {
				var C_str = strconv.Itoa(C)
				var WPM_str = fmt.Sprintf("%.1f", ts[i].WPM)
				DISPLAY = DISPLAY + C_str + ". " + ts[i].Username + " (" + WPM_str + ") " + ts[i].Date + "\n"
			}
		}
		C++
	}

	s.ChannelMessageSend(m.ChannelID, "```🏆 LEADERBOARD 🏆\n"+DISPLAY+"```")
}

func TopsID(s *discordgo.Session, m *discordgo.MessageCreate) {
	var What = strings.Replace(m.Content, ".topsID ", "", -1)

	var What_int, _ = strconv.Atoi(What)

	var FOUND bool = false
	var FOUND_how_many_times int = 0

	var ts []Two_Slices

	var slice Two_Slices

	Load()
	for i := 0; i < len(DB); i++ {
		if strings.HasPrefix(DB[i], What+" #") {
			if i == 0 {

			} else {
				var WPM_f64 float64

				var CL = strings.Split(DB[i], " # ")
				if len(CL) > 5 {
					WPM_f64, _ = strconv.ParseFloat(CL[5], 8)
				} else {
					WPM_f64, _ = strconv.ParseFloat(CL[3], 8)
				}

				slice = Two_Slices{
					WPM_f64, CL[2], CL[4],
				}
				ts = append(ts, slice)
			}
			FOUND = true
			FOUND_how_many_times++
		}
	}

	sort.Slice(ts, func(i, j int) bool {
		return ts[i].WPM > ts[j].WPM
	})

	var DISPLAY string
	var C int = 1
	for i := 0; i < 5; i++ {
		if i != 0 {
			if i > len(ts)-1 {
				var C_str = strconv.Itoa(C)
				DISPLAY = DISPLAY + C_str + "." + "\n"
			} else {
				var C_str = strconv.Itoa(C)
				var WPM_str = fmt.Sprintf("%.1f", ts[i].WPM)
				DISPLAY = DISPLAY + C_str + ". " + ts[i].Username + " (" + WPM_str + ") " + ts[i].Date + "\n"
			}
		} else {
			if len(ts) > 0 {
				var C_str = strconv.Itoa(C)
				var WPM_str = fmt.Sprintf("%.1f", ts[i].WPM)
				DISPLAY = DISPLAY + C_str + ". " + ts[i].Username + " (" + WPM_str + ") " + ts[i].Date + "\n"
			}
		}
		C++
	}

	var FOUND_how_many_times_str = strconv.Itoa(FOUND_how_many_times)
	if FOUND {
		var Texts_arr = strings.Split(Texts[What_int], " ")
		var λ string
		for i := 0; i < len(Texts_arr); i++ {
			if i != len(Texts_arr)-1 {
				λ = λ + Texts_arr[i] + "\u200b "
			} else {
				λ = λ + Texts_arr[i]
			}
		}
		s.ChannelMessageSend(m.ChannelID, "```"+λ+"```")
		s.ChannelMessageSend(m.ChannelID, "```diff\n+ Se encontraron "+FOUND_how_many_times_str+" marcas del texto ID:"+What+"\n"+DISPLAY+"```")
	} else {
		s.ChannelMessageSend(m.ChannelID, "```diff\n- No se encontró```")
	}

}

var Date_temp string
var WPM_temp string

func Is_already_in_top(m *discordgo.MessageCreate, Text_ID int, wpm float64) bool {
	Load()

	var Random_str = strconv.Itoa(Text_ID)
	for i := 0; i < len(DB); i++ {
		if strings.HasPrefix(DB[i], Random_str) {
			var CL = strings.Split(DB[i], " # ")
			if CL[1] == m.Author.ID {
				CL_f64, _ := strconv.ParseFloat(CL[5], 8)
				if CL_f64 > wpm {
					Date_temp = CL[4]
					WPM_temp = CL[3]
					return true
				}
			}
		}
	}
	return false
}

func Is_already_in_top_LOWER(s *discordgo.Session, m *discordgo.MessageCreate, Text_ID int, wpm float64) {
	Load()

	var Random_str = strconv.Itoa(Text_ID)
	for i := 0; i < len(DB); i++ {
		if strings.HasPrefix(DB[i], Random_str) {
			var CL = strings.Split(DB[i], " # ")
			if CL[1] == m.Author.ID {
				CL_f64, _ := strconv.ParseFloat(CL[5], 8)
				if CL_f64 < wpm {
					Delete_last_score_because_improved = true
					Date_temp = CL[4]
					WPM_temp = CL[3]
					DB[i] = ""
					Update()
				}
			}
		}
	}
}

func Stat_list(s *discordgo.Session, m *discordgo.MessageCreate) string {
	var stats string
	var C int

	Load()
	for i := 0; i < len(DB); i++ {
		var CL = strings.Split(DB[i], " # ")
		/* Check if is not empty line */
		if len(CL) > 1 {
			if CL[1] == m.Author.ID {
				C++
				stats = stats + DB[i] + "\n"
			}
		}
	}

	C_str := strconv.Itoa(C)
	stats = C_str

	return stats
}

func Stats(s *discordgo.Session, m *discordgo.MessageCreate) {
	var Texts_where_user_is_in []string

	var How_many_texts_in int
	var Tops_1 int

	Load()
	for i := 0; i < len(DB); i++ {
		var CL = strings.Split(DB[i], " # ")
		/* Check if is not empty line */
		if len(CL) > 1 {
			if CL[1] == m.Author.ID {
				if !(Slice_contains(Texts_where_user_is_in, CL[0])) {
					Texts_where_user_is_in = append(Texts_where_user_is_in, CL[0])
					How_many_texts_in++
					var what_text = CL[0]
					var what_WPM, _ = strconv.ParseFloat(CL[5], 64)

					var Not_top bool = false
					for k := 0; k < len(DB); k++ {
						var CK = strings.Split(DB[k], " # ")
						if CK[0] == what_text && CK[1] != m.Author.ID {
							if len(CK) > 5 {
								CK_float, _ := strconv.ParseFloat(CK[5], 64)

								if CK_float > what_WPM {
									Not_top = true
								}
							}
						}
					}
					if Not_top == false {
						Tops_1++
					}
				}
			}
		}
	}
	var TOPS_str = strconv.Itoa(Tops_1)
	s.ChannelMessageSend(m.ChannelID, "```css\nHas participado en "+Stat_list(s, m)+" textos\n"+m.Author.Username+", tienes "+TOPS_str+" tops 1.```")
}

type Two_Slices2 struct {
	User_list              string
	User_ID_list           string
	User_tops              int
	User_how_many_texts_in int
}

func Leaderboards(s *discordgo.Session, m *discordgo.MessageCreate) {
	// 👷🚧

	var User_list []string
	var User_ID_list []string
	var User_tops []int
	var User_how_many_texts_in []int

	Load()

	for i := 0; i < len(DB); i++ {
		var CL = strings.Split(DB[i], " # ")
		if len(CL) > 1 {
			if !Slice_contains(User_list, CL[2]) && CL[1] != "0" {
				User_list = append(User_list, CL[2])
				User_ID_list = append(User_ID_list, CL[1])
			}
		}
	}

	for u := 0; u < len(User_list); u++ {
		var Texts_where_user_is_in []string
		var How_many_texts_in int
		var Tops_1 int

		Load()
		for i := 0; i < len(DB); i++ {
			var CL = strings.Split(DB[i], " # ")
			if len(CL) > 5 {
				if CL[1] == User_ID_list[u] {
					if !(Slice_contains(Texts_where_user_is_in, CL[0])) {
						Texts_where_user_is_in = append(Texts_where_user_is_in, CL[0])
						How_many_texts_in++
						var what_text = CL[0]
						var what_WPM, _ = strconv.ParseFloat(CL[5], 64)

						var Not_top bool = false
						for k := 0; k < len(DB); k++ {
							var CK = strings.Split(DB[k], " # ")
							if CK[0] == what_text && CK[1] != User_ID_list[u] {
								if len(CK) > 5 {
									CK_float, _ := strconv.ParseFloat(CK[5], 64)

									if CK_float > what_WPM {
										Not_top = true
									}
								}
							}
						}
						if Not_top == false {
							Tops_1++
						}
					}
				}
			}
		}
		if false {
			User_tops = append(User_tops, 0)

		} else {
			User_tops = append(User_tops, Tops_1)
			User_how_many_texts_in = append(User_how_many_texts_in, How_many_texts_in)
		}
	}

	var ts2 []Two_Slices2

	var slice2 Two_Slices2
	/*
		var User_list []string
		var User_ID_list []string
		var User_tops []int
	*/
	for i := range User_list {
		slice2 = Two_Slices2{
			User_list[i], User_ID_list[i], User_tops[i], User_how_many_texts_in[i],
		}
		ts2 = append(ts2, slice2)
	}

	//fmt.Println(ts2)
	sort.Slice(ts2, func(i, j int) bool {
		return ts2[i].User_tops > ts2[j].User_tops
	})
	//fmt.Println(ts2)
	var result string

	var c int
	for i := range ts2 {
		c++
		if c > 10 {
			break
		}
		var c_str = strconv.Itoa(c)
		var ts2i_User_tops_str = strconv.Itoa(ts2[i].User_tops)
		var ts2i_User_how_many_texts_in_str = strconv.Itoa(ts2[i].User_how_many_texts_in)
		result = result + "\n" + c_str + ". " + fmt.Sprint(ts2[i].User_list+" con "+ts2i_User_tops_str+" tops 1 en "+ts2i_User_how_many_texts_in_str+" textos.")
	}

	s.ChannelMessageSend(m.ChannelID, "```css\n📈 Leaderboards\n"+result+"```")
}
