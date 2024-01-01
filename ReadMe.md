#First_Main_Server

This is my first Go server. This server serves two static pages on two different routes and also sends email to the user as well as the server owner.

#Routes
http://localhost:8080/ :- This route contains a welcome page with a button which directs you to form page.
http://localhost:8080/form :- This route serves a form page that allows users to fill a form and submit it. On submit a page with user details is serves and the data along with a message is sent to the server admin and on successful delivery of mail, user is notified through an email.