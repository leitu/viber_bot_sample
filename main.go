package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"io/ioutil"

	"github.com/mileusna/viber"
)

const NewID = "xxxddfdf"

func main() {
	v := &viber.Viber{
		AppKey: "MY-TOKEN",
		Sender: viber.Sender{
			Name:   "Sample_bot",
			Avatar: "https://mysite.com/img/avatar.jpg",
		},
		Message:   myMsgReceivedFunc, // your function for handling messages
		Delivered: myDeliveredFunc,   // your function for delivery report
	}
	v.Seen = mySeenFunc // or assign events after declaration

	// this have to be your webhook, pass it your viber app as http handler
	http.Handle("/viber/webhook/", v)
	http.ListenAndServe(":80", nil)
}

// myMsgReceivedFunc will be called everytime when user send us a message
func myMsgReceivedFunc(v *viber.Viber, u viber.User, m viber.Message, token uint64, t time.Time) {
	switch m.(type) {

	case *viber.TextMessage:
		v.SendTextMessage(u.ID, "Thank you for your message")
		log.Println(u.ID)
		txt := m.(*viber.TextMessage).Text

		if txt == "test" {
			log.Println("This is new one")
			d1 := []byte("hello\ngo\n")
			err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
			check(err)

			content, err := ioutil.ReadFile("/tmp/dat1")
			data := fmt.Sprintf("%s", content)
			v.SendTextMessage(u.ID, data)
		} else if txt == "all" {
			v.SendTextMessage(u.ID, "ahaha")
			v.SendTextMessage(NewID, "ahahaha")
		}

	case *viber.URLMessage:
		url := m.(*viber.URLMessage).Media
		v.SendTextMessage(u.ID, "You have sent me an interesting link "+url)

	case *viber.PictureMessage:
		v.SendTextMessage(u.ID, "Nice pic!")

	}
}

func myDeliveredFunc(v *viber.Viber, userID string, token uint64, t time.Time) {
	log.Println("Message ID", token, "delivered to user ID", userID)
}

func mySeenFunc(v *viber.Viber, userID string, token uint64, t time.Time) {
	log.Println("Message ID", token, "seen by user ID", userID)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// All events that you can assign your function, declarations must match
// ConversationStarted func(v *Viber, u User, conversationType, context string, subscribed bool, token uint64, t time.Time) Message
// Message             func(v *Viber, u User, m Message, token uint64, t time.Time)
// Subscribed          func(v *Viber, u User, token uint64, t time.Time)
// Unsubscribed        func(v *Viber, userID string, token uint64, t time.Time)
// Delivered           func(v *Viber, userID string, token uint64, t time.Time)
// Seen                func(v *Viber, userID string, token uint64, t time.Time)
// Failed              func(v *Viber, userID string, token uint64, descr string, t time.Time)
