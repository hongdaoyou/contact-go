
ALL: contact_go_arr contact_go_link


contact_go_arr:contact_go_arr.go
	go build -o $@ $^

contact_go_link: contact_go_link.go
	go build -o $@ $^

clean:
	rm -rf contact_go_arr contact_go_link
