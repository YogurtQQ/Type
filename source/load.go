package main

import (
	"fmt"
	"math"
	"time"

	//"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Load() {
	fileBytes, err := ioutil.ReadFile("./database/saved_results.csv")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	DB = strings.Split(string(fileBytes), "\n")
}

var changing_saved_results bool = false

func Save_result(m *discordgo.MessageCreate, Text_ID int, wpm float64) {
	var WPM_rounded = (math.Round(wpm*10) / 10)
	var wpm_str = fmt.Sprint(WPM_rounded)
	var wpm_str_save = fmt.Sprint(wpm)

	for true {
		if !changing_saved_results {
			changing_saved_results = true
			f, err := os.OpenFile("./database/saved_results.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}
			Text_ID_str := strconv.Itoa(Text_ID)
			if _, err := f.Write([]byte(Text_ID_str + " # " + m.Author.ID + " # " + m.Author.Username + " # " + wpm_str + " # " + time.Now().Format("02/01/2006 15:04") + " # " + wpm_str_save + " # " + m.Message.ID + "\n")); err != nil {
				f.Close()
				log.Fatal(err)
			}
			if err := f.Close(); err != nil {
				log.Fatal(err)
			}
			changing_saved_results = false
			break
		}
	}
}

func Update() {
	for true {
		if !changing_saved_results {
			changing_saved_results = true

			f, err := os.Create("./database/saved_results.csv")
			if err != nil {
				log.Fatal(err)
			}

			defer f.Close()

			for _, line := range DB {
				_, err := fmt.Fprintln(f, line)
				if err != nil {
					log.Fatal(err)
				}
			}
			changing_saved_results = false
			break
		}
	}
}

func Log(m *discordgo.MessageCreate) {
	f, err := os.OpenFile("./database/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := f.Write([]byte("[" + time.Now().Format("02/01/2006 15:04:05") + "] <#" + m.ChannelID + "> " + m.Author.ID + ", " + m.Author.Username + "> " + m.Content + "\n")); err != nil {
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func Load_texts() {
	fileBytes, err := ioutil.ReadFile("./database/texts.csv")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	Texts = strings.Split(string(fileBytes), "\n")
}
