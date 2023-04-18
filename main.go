package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"math/rand"
	"strings"
	"time"
)

func geekSayWord(word string, translations map[string]string) string {
	for k, v := range translations {
		if word == k {
			return v
		}
	}
	return word
}

func geekSayPhrase(phrase interface{}, translations map[string]string) string {
	out := ""
	switch v := phrase.(type) {
	case string:
		for _, v := range strings.Split(v, " ") {
			out += geekSayWord(strings.ToLower(v), translations) + " "
		}
	case []string:
		for _, v := range v {
			out += geekSayWord(strings.ToLower(v), translations) + " "
		}
	}

	return out
}

func main() {
	translations := map[string]string{
		"hide":         "incognito",
		"hidden":       "incognito",
		"screenshot":   "start+prt sc",
		"input":        "<input>",
		"into":         "*",
		"emotion":      "emoticon",
		"repeat":       "loop",
		"style":        "css",
		"save":         "ctrl+s",
		"break":        "<br>",
		"control":      "ctrl",
		"rule":         "ctrl",
		"escape":       "esc",
		"black":        "#000",
		"white":        "#fff",
		"bang":         "!",
		"not":          "!",
		"new":          "ctrl+n",
		"open":         "ctrl+o",
		"anchor":       "<a>",
		"bold":         "<b>",
		"cut":          "ctrl+x",
		"copy":         "ctrl+c",
		"paste":        "ctrl+v",
		"quit":         "ctrl+q",
		"close":        "ctrl+w",
		"help":         "f1",
		"rename":       "f2",
		"undo":         "ctrl+z",
		"redo":         "ctrl+y",
		"search":       "ctrl+f",
		"find":         "ctrl+f",
		"replace":      "ctrl+r",
		"print":        "ctrl+p",
		"greet":        "ping",
		"get":          "git",
		"got":          "git",
		"home":         "127.0.0.1",
		"bed":          "127.0.0.1",
		"house":        "127.0.0.1",
		"local":        "127.0.0.1",
		"localhost":    "127.0.0.1",
		"fuck":         "fork",
		"hello":        "ping",
		"hi":           "ping",
		"hey":          "ping",
		"command":      "cmd",
		"forget":       "ctrl+z",
		"world":        "0.0.0.0/0",
		"want":         "require",
		"need":         "require",
		"name":         "alias",
		"address":      "url",
		"pie":          "π",
		"function":     "fn",
		"kill":         "alt+f4",
		"stop":         "abort",
		"refresh":      "f5",
		"slow":         "O(n^n)",
		"fast":         "O(1)",
		"leet":         "1337",
		"hacker":       "haxor",
		"geek":         "g33k",
		"plus":         "+",
		"minus":        "-",
		"and":          "&&",
		"or":           "||",
		"binary":       "01",
		"call":         "ping",
		"me":           "self",
		"owner":        "admin",
		"god":          "root",
		"please":       "sudo",
		"at":           "@",
		"fix":          "debug",
		"shit":         "dump",
		"bullshit":     "dump",
		"poop":         "dump",
		"delete":       "del",
		"force":        "ma",
		"coffee":       "covfefe",
		"number":       "#",
		"image":        "img",
		"photo":        "img",
		"pic":          "img",
		"picture":      "img",
		"snap":         "img",
		"snapshot":     "img",
		"bye":          "exit",
		"noob":         "n00b",
		"life":         "42",
		"universe":     "42",
		"everything":   "42",
		"lifetime":     "runtime",
		"sleep":        "shutdown",
		"red":          "#ff0000",
		"green":        "#00ff00",
		"blue":         "#0000ff",
		"marry":        "merge",
		"propose":      "commit",
		"notfound":     "404!",
		"link":         "url",
		"send":         "push",
		"alternate":    "alt",
		"alternative":  "alt",
		"inspect":      "ctrl+shift+i",
		"equal":        "=",
		"remove":       "rm",
		"move":         "mv",
		"random":       "rand",
		"mathematics":  "math",
		"love":         "<3",
		"smile":        ":)",
		"sad":          ":(",
		"list":         "ls",
		"ok":           "200",
		"okay":         "200",
		"question":     "?",
		"history":      "log",
		"on":           "1",
		"off":          "0",
		"pixel":        "px",
		"say":          "log",
		"body":         "<body>",
		"head":         "<head>",
		"title":        "<title>",
		"restart":      "ctrl+alt+del",
		"death":        "BSOD",
		"wait":         "load",
		"waiting":      "loading",
		"unauthorized": "401",
		"nothing":      "void",
		"all":          "*",
		"lineartime":   "o(N)",
		"change":       "alt",
		"pink":         "#ffc0cb",
		"developer":    "dev",
		"switch":       "alt+tab",
		"information":  "info",
		"space":        "' '",
		"full":         "overload",
		"modulo":       "%",
		"mind":         "database",
		"heart":        "SMPS",
		"semicolon":    ";",
		"future":       "ctrl+y",
		"past":         "ctrl+z",
		"missing":      "404!",
		"unavailable":  "503!",
		"overloaded":   "502!",
		"forbidden":    "403!",
		"timeout":      "408!",
		"repository":   "repo",
		"disapperar":   "heisenbug",
		"clueless":     "Jimmy",
		"imaginary":    "Unicorny",
		"catastrophic": "hidenbug",
		"useless":      "reality 101 faliure",
		"different":    "nested",
		"calling":      "recursion",
		"heartbeat":    "ping",
		"boobs":        "80085",
		"error":        "404",
		"top":          "</head>",
		"not found":    "404",
		"run":          "ctrl+F5",
		"comment":      "//comment",
		"table":        "<table>",
		"gray":         "#ccc",
		"grey":         "#ccc",
		"correct":      "legit",
		"teapot":       "418",
		"moved":        "301",
		"bad":          "400",
		"gone":         "410",
		"form":         "<form>",
		"select":       "<select>",
		"leave":        "alt+f4",
		"mistake":      "ERROR",
		"divide":       "/",
	}
	quotes := [...]string{
		"Everything is under control",                   // Everything is under ctrl
		"Give me a break",                               // Give me a <br>
		"Hello World",                                   // ping 0.0.0.0/0
		"Bye World",                                     // exit 0.0.0.0/0
		"Forget that ever happened",                     // ctrl+z that ever happened
		"Home is where the heart is",                    // 127.0.0.1 is where the SMPS is
		"Home is where ideas come to life",              // 127.0.0.1 is where ideas come to 42
		"Home is where the wifi connects automatically", // 127.0.0.1 is where the wifi connects automatically
		"fuck that shit",                                // fork that dump
		"go hard or go home",                            // go hard || go 127.0.0.1
		"I will call you",                               // I will ping you
		"I want the world to know me",                   // I require the 0.0.0.0/0 to know self
		"I am the owner",                                // I am the admin
		"I am God",                                      // I am root
		"please make me a sandwich",                     // sudo make self a sandwich
		"coffee is life",                                // covfefe is 42
		"what's your name",                              // what's your alias
		"I will find you and I will fix you!",           // I will ctrl+f you && I will debug you!
		"I need help",                                   // I require f1
		"geek alert",                                    // g33k alert
		"geek inside",                                   // g33k inside
		"copy that!",                                    // ctrl+c that!
		"get out of home",                               // git out of 127.0.0.1
		"not everything is black and white",             // ! 42 is #000 && #fff
		"to be or not to be, that is the question",      // to be || ! to be, that is the ?
		"i was 5 and he was 6",                          // i was 101 && he was 110
		"he wore black and i wore white",                // he wore #000 && i wore #fff
		"bang bang my baby shot me down",                // ! ! my baby shot self down
		"Today is the first day of the rest of your lifetime", // Today is the first day of the rest of your runtime
		"i need to sleep", // i require to shutdown
		"There are ten types of people in the world: those who understand binary and those who don't", // There are ten types of people in the world: those who understand 01 && those who don't
		"I love my life",                  // I <3 my 42
		"I love my house",                 // I <3 my 127.0.0.1
		"heart break",                     // SMPS <br>
		"I like to fix shit",              // I like to debug dump
		"eat sleep code repeat",           // eat shutdown code repeat
		"roses are red, violets are blue", // roses are ##ff0000, violets are ##0000ff
		"The greatest trick the devil ever pulled was convincing the world he didn't exist", // The greatest trick the devil ever pulled was convincing the 0.0.0.0/0 he didn't exist
		"home sweet home",                          // 127.0.0.1 sweet 127.0.0.1
		"history repeats itself",                   // log repeats itself
		"My dog likes to greet me at the door.",    // My dog likes to ping me at the door.
		"keep calm and geek on",                    // keep calm && g33k 1
		"somebody save me",                         // somebody ctrl+s me
		"Easy as pie",                              // easy as π
		"Live, laugh, love",                        // live, laugh, <3
		"Green with envy",                          // #00ff00 with envy
		"There's no place like home",               //There's no place like 127.0.0.1
		"Green with envy",                          // #00ff00 with envy
		"Once in a blue moon",                      //Once in a #0000ff moon
		"Hello All",                                // Hello *
		"It is lineartime",                         // It is o(N)
		"you yourself are the anchor of your life", //you yourself are the <a> of your life
		"you are the love of my life",              //you are the <3 of my 42
		"pink panther",                             //  #ffc0cb panther
		"Get busy living or get busy dying",        // git busy living || git busy dying
		"out of the blue",                          // out of the #0000ff
		"caught red-handed",                        // caught #ff0000 handed
		"I love you",                               //I <3 you
		"He has a style",                           //He has a css
		"hide it!",                                 //incognito it!
		"99 little bugs in the code. 99 little bugs in the code. Take one down, patch it around. 127 little bugs in the code", //cb - so many bugs
		"Hey Google! Find me the best repo.", //ping Google! ctrl+f self the best repo
		"I need to restart my life",          // I require to ctrl+alt+del my 42
		"Paste my Code",                      // ctrl+v my Code
		"Paste my Copy",                      // ctrl+v my ctrl+c
		"What's new?",                        // What's ctrl+n
		"You're done coding! Go to sleep.",   // You're done coding! Go to shutdown
		"Life is a big question that even Google can't find answer.",     //42 is a big ? that even Google can't ctrl+f answer
		"Have no friends not equal to yourself.",                         // Have no friends ! = to yourself
		"The past does not equal the future.",                            //The past does ! = the future
		"I just switch tabs, not the love ones",                          // I just alt+tab tabs, not the <3 ones
		"I have not failed. I’ve just found 10,000 ways that won’t work", // I have ! failed. I’ve just found 0 ways that won’t work
		"Life is a big question that even Google can't find answer.",     //42 is a big ? that even Google can't ctrl+f answer
		"Have no friends not equal to yourself.",                         // Have no friends ! = to yourself
		"The past does not equal the future.",                            //The past does ! = the future
		"I just switch tabs, not the love ones",                          // I just alt+tab tabs, not the <3 ones
		"to be or not to be",                                             // 2b || !2b
		"programming joke",                                               // My </code>
		"home sweet home",                                                // 127.0.0.1 sweet 127.0.0.1
		"politics is bullshit",                                           //politics is dump
		"I want the address of my love",                                  // I want the url of my <3
		"Never get to bed mad, stay up and snap",                         // Never git to 127.0.0.1 mad,stay up && img
		"Heart is missing",                                               //SMPS is 404!
		"Information is lost from the deleted repository",                //Info is lost from the del repo
		"Be different",                                                   // Be !=
		"Don't leave blank space in your code!",                          // Don't alt+f4 blank ' ' in your code!
		"Sometimes it's okay to run",                                     // Sometimes it's okay to ctrl+F5
		"Divide and conquer",                                             // / and conquer
	}

	var rootCmd = &cobra.Command{
		Use:  "geeksay",
		Args: cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			out := ""
			if args[0] == "r" {
				rand.Seed(time.Now().UnixMilli())
				randomInt := rand.Intn(len(quotes))
				out += geekSayPhrase(quotes[randomInt], translations)
			} else {
				out += geekSayPhrase(args, translations)
			}
			fmt.Println(out)
		},
	}
	rootCmd.Execute()
}
