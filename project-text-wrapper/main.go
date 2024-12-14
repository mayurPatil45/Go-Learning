package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println("Text Wrapper Challenge")
	const text = `Her mom had warned her. She had been warned time and again, but she had refused to believe her. She had done everything right and she knew she would be rewarded for doing so with the promotion. So when the promotion was given to her main rival, it not only stung, it threw her belief system into disarray. It was her first big lesson in life, but not the last.
Betty was a creature of habit and she thought she liked it that way. That was until Dave showed up in her life. She now had a choice to make and it would determine whether her lie remained the same or if it would change forever.
The time had come for Nancy to say goodbye. She had been dreading this moment for a good six months, and it had finally arrived despite her best efforts to forestall it. No matter how hard she tried, she couldn't keep the inevitable from happening. So the time had come for a normal person to say goodbye and move on. It was at this moment that Nancy decided not to be a normal person. After all the time and effort she had expended, she couldn't bring herself to do it.`

	const maxWidth = 40

	var lw int // Line Width
	for _, r := range text {
		fmt.Printf("%c", r)
		switch lw++; {
		case lw > maxWidth && r != '\n' && unicode.IsSpace(r):
			fmt.Println()
			fallthrough
		case r == '\n':
			lw = 0
		}
	}
	fmt.Println("")
}
