package main

import "net/http"

func sign(w http.ResponseWriter, req *http.Request) {
	//make sure we got post
	if req.Method != "POST" {
		http.NotFound(w, req)
		return
	}

	entry := NewEntry()
	entry.Name = req.FormValue("name")
	entry.Message = req.FormValue("message")

	if entry.Name == "" {
		entry.Name = "Some dummy who forgot a name"
	}
	if entry.Message == "" {
		entry.Message = "Some dummy who forgot a message."
	}

	s := session.Clone()
	defer s.Close()

	coll := s.DB("gostbook").C("entries")
	if err := coll.Insert(entry); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
}
