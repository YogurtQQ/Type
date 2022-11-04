package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	"strconv"
	"strings"
	"github.com/bwmarrin/discordgo"
)

var DB []string

var Date string

var Current_text string = "Escucha, las reglas propias... se tratan de decidir conseguir algo usando medios y maneras propias para conseguirlo. Por eso decimos que son nuestras reglas. Precisamente por eso podemos afrontar sinceramente los desafíos y darlo todo. Y si fracasamos, hay que retomar la práctica y soportar duros entrenamientos para lograrlo. Y así, dedicándote a ello, creas tus propias reglas."
var Is_started bool
var Started_when = time.Now().UnixMilli()

var Text_message_ID string

var WPM float64
var WPM_str string
var WPM_str_save string

var Random int

func Is_illegal(ζ string) bool {
	var π = strings.Contains(ζ, "​")
	return π
}

func υ() string {

	var Texts_arr = strings.Split(Current_text, " ")
	var λ string
	for i := 0; i < len(Texts_arr); i++ {
		if i != len(Texts_arr)-1 {
			λ = λ + Texts_arr[i] + "​ "
		} else {
			λ = λ + Texts_arr[i]
		}
	}
	return λ
}

func Typing_test(s *discordgo.Session, m *discordgo.MessageCreate) {
	Is_started = true

	rand.Seed(time.Now().UnixNano())
	Random = rand.Intn(How_many_texts())
	Random = rand.Intn(400)
	//Random = 76

	Current_text = Texts[Random]

	var Test_message, _ = s.ChannelMessageSend(m.ChannelID, "```🔴 Preparados...```")
	Text_message_ID = Test_message.ID
	time.Sleep(1 * time.Second)
	s.ChannelMessageEdit(m.ChannelID, Test_message.ID, "```🟡 Listos...```")
	time.Sleep(1 * time.Second)
	var Text_shown, _ = s.ChannelMessageEdit(m.ChannelID, Test_message.ID, "**" + υ() + "**")
	_ = Text_shown
	var Started_when_time = Text_shown.EditedTimestamp
	Started_when = Started_when_time.UnixMilli()
}

func Judge(m *discordgo.MessageCreate, S string) int8 {
	if S == Current_text {
		return 1
	}

	var Content_arr = strings.Split(m.Content, " ")
	var Current_text_arr = strings.Split(Current_text, " ")

	if Content_arr[0] == Current_text_arr[0] {
		if (len(m.Content) > len(Current_text) - 20) {
				return 2
		}

		if (len(m.Content) > len(Current_text) - 60) {
				return 4
		}
	}
	return 3
}

func Calculate(m *discordgo.MessageCreate) {
	var sent_when, _ = SnowflakeTimestamp(m.Message.ID)
	Date = sent_when.Format("02/01/2006 15:04")
	var sent_when_unixmilli = sent_when.UnixMilli()

	var sent_when_unixmilli_float64 float64 = float64(sent_when_unixmilli)

	var length = len([]rune(Current_text))
	var length_as_a_float float64 = float64(length)

	var Started_when_float float64 = float64(Started_when)
	WPM = length_as_a_float / 5 / ((sent_when_unixmilli_float64-Started_when_float)-1000) * 60000
}

var Error_list string
var Errors int
var Errors_str string

func Errors_calculate(sent string, current string) {
	
	/* reseting */
	Errors = 0
	Error_list = ""

	A := sent
	sent_arrayed := strings.Split(A, " ")

	B := current
	text_arrayed := strings.Split(B, " ")

	if len(text_arrayed) == len(sent_arrayed) {
		for i := 0; i < len(text_arrayed); i++ {
				if text_arrayed[i] != sent_arrayed[i] {
					if Error_list != "" {
						Error_list = Error_list + ", " + sent_arrayed[i]
						Errors++
					} else {
						Error_list = sent_arrayed[i]
						Errors++
					}
				}
			}
	}

	if  len(text_arrayed) > len(sent_arrayed) {

	}


	Errors_str = strconv.FormatInt(int64(Errors), 10)
}

func Show_result(s *discordgo.Session, m *discordgo.MessageCreate) {
	var WPM_rounded = (math.Round(WPM*10)/10)
	WPM_str = fmt.Sprint(WPM_rounded)

	WPM_str_save = fmt.Sprint(WPM)

	s.ChannelMessageSend(m.ChannelID, "```diff\n+ " + m.Author.Username + ", has terminado.\nTu resultado es: " + WPM_str + " WPM```")
}

func Show_result_not_improved(s *discordgo.Session, m *discordgo.MessageCreate) {
	var WPM_rounded = (math.Round(WPM*10)/10)
	WPM_str = fmt.Sprint(WPM_rounded)

	s.ChannelMessageSend(m.ChannelID, "```diff\n- No has superado tu anterior marca, " + m.Author.Username + ".\nTu resultado es: " + WPM_str + " WPM```")
}

func Show_result_with_errors(s *discordgo.Session, m *discordgo.MessageCreate) {
	var WPM_rounded = (math.Round(WPM*10)/10)
	if Errors < 1 {
		WPM_str = fmt.Sprint(WPM_rounded)
		s.ChannelMessageSend(m.ChannelID, "```diff\n- " + m.Author.Username + ", no has terminado correctamente.\nPusiste una palabra de más o un doble espacio, así que no se pudieron calcular tus errores.\nTu resultado hubiera sido: " + WPM_str + " WPM```")

	} else {
		WPM_str = fmt.Sprint(WPM_rounded)
		s.ChannelMessageSend(m.ChannelID, "```diff\n- " + m.Author.Username + ", no has terminado correctamente.\nHas cometido " + Errors_str + " errores: " + Error_list + "\nTu resultado hubiera sido: " + WPM_str + " WPM```")
	}
	
}

func TT_short(s *discordgo.Session, m *discordgo.MessageCreate) {
	Is_started = true

	rand.Seed(time.Now().UnixNano())
	Random = rand.Intn(435 - 400) + 400
	//Random = 76
	Current_text = Texts[Random]

	var Test_message, _ = s.ChannelMessageSend(m.ChannelID, "```🔴 (Textos cortos) Preparados...```")
	time.Sleep(1 * time.Second)
	s.ChannelMessageEdit(m.ChannelID, Test_message.ID, "```🟡 (Textos cortos) Listos...```")
	time.Sleep(1 * time.Second)
	var Text_shown, _ = s.ChannelMessageEdit(m.ChannelID, Test_message.ID, "**" + υ() + "**")
	_ = Text_shown
	var Started_when_time = Text_shown.EditedTimestamp
	Started_when = Started_when_time.UnixMilli()
}