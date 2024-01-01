package main

import (
	"fmt"
	"net/http"
	"net/smtp"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
	}

	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "./static/form.html")
	} else {
		name := r.FormValue("name")
		email := r.FormValue("email")
		message := r.FormValue("message")
		fmt.Fprintf(w, "Name: %s \n", name)
		fmt.Fprintf(w, "Email: %s \n", email)
		fmt.Fprintf(w, "Message: %s \n", message)
		auth := smtp.PlainAuth(
			"",
			"sharmapriyanshu594@gmail.com",
			"badr fffo uchn irvs",
			"smtp.gmail.com",
		)

		msg := "Subject: Inquiry Mail" + "\n" + message
		err := smtp.SendMail(
			"smtp.gmail.com:587",
			auth,
			"sharmapriyanshu594@gmail.com",
			[]string{"sharmapriyanshu594@gmail.com"},
			[]byte(msg),
		)
		if err != nil {
			fmt.Println(err)
		}
		msg1 := "Subject: Success Msg" + "\n" + "Your response has been delivered Successfully"
		err1 := smtp.SendMail(
			"smtp.gmail.com:587",
			auth,
			"sharmapriyanshu594@gmail.com",
			[]string{email},
			[]byte(msg1),
		)
		if err1 != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/", fs)
	http.HandleFunc("/form", formHandler)
	fmt.Println("Server Running At Port: 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error: ", err)
	}

}
