package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println(err.Error())
	}

	arguments := os.Args
	fmt.Println(arguments)
	fmt.Println(arguments[0])
	// fmt.Println(arguments[1])
	// fmt.Println(arguments[2])

	errChdir := os.Chdir("baren")

	if errChdir != nil {
		fmt.Println("error:", errChdir)
	}

	errChmod := os.Chmod("test-chmod.txt", 555)
	if errChmod != nil {
		fmt.Println("error chmod:", errChmod)
	} else {
		fmt.Println("permission untuk file test-chmod.txt berubah")
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")
	fmt.Println(username)
	fmt.Println(password)

	os.Clearenv()

	if pesan, err := os.Executable(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pesan)
	}

	// os.Exit(0)

	greeting := func(name string) string {
		switch name {
		case "WAKTU":
			return "Evening"
		case "NAME":
			return "Baren"
		}
		return ""
	}

	fmt.Println(os.Expand("Good ${WAKTU}, $NAME", greeting))

	os.Setenv("NAME", "baren")
	os.Setenv("PASSWORD", "new-password")

	fmt.Println("ENV NAME : ", os.ExpandEnv("$NAME, ENV PASS : ${PASSWORD}"))

	fmt.Println(os.Getenv("NAME"))
	fmt.Println(os.Getenv("PASSWORD"))

	fmt.Println("id dari function getId : ", os.Getegid())

	fmt.Println("uid dari function getUID :", os.Getuid())

	fmt.Println("id dari function getgid", os.Getgid())

	if getGroups, errGetGroups := os.Getgroups(); errGetGroups != nil {
		fmt.Println("error getGroups : ", errGetGroups)
	} else {
		fmt.Println(getGroups)
	}

	fmt.Println(os.Getpagesize())

	fmt.Println(os.Getpid())

	if path, errGetWD := os.Getwd(); errGetWD != nil {
		fmt.Println("error function getwd : ", errGetWD)
	} else {
		fmt.Println(path)
	}

	if _, errStat := os.Stat("main.go"); os.IsNotExist(errStat) {
		fmt.Println(errStat)
	} else {
		fmt.Println("file ditemukan")
	}

	fmt.Println(os.IsPathSeparator(2))

	uid := os.Getuid()
	id := os.Getegid()

	if errChown := os.Chown("test-chmod.txt", uid, id); errChown != nil {
		fmt.Println(errChown)
	} else {
		fmt.Println("sepertinya berhasil")
	}

	if errLink := os.Link("test-chmod.txt", "new-chmod"); errLink != nil {
		fmt.Println(errLink)
	} else {
		fmt.Println("function os link tidak error")
	}

	showEnv := func(key string) {
		value, ok := os.LookupEnv(key)
		if ok {
			fmt.Println("key :", key, "value :", value)
		} else {
			fmt.Println("key : ", key, "not defined")
		}
	}

	showEnv("NAME")
	showEnv("PASSWORD")
	showEnv("username")

	if errMkdir := os.Mkdir("main.go", 777); errMkdir != nil {
		fmt.Println("error mkdir :", errMkdir)
	}

	if errMkdirAll := os.MkdirAll("dibuat-dengan-mkdir-all", 777); errMkdirAll != nil {
		fmt.Println("error MkdirAll :", errMkdirAll)
	}

	buatError := errors.New("coba error")
	fmt.Println(buatError)

	if r, w, errPipe := os.Pipe(); errPipe != nil {
		fmt.Println(errPipe)
	} else {
		fmt.Println("r  = ", *r, "w = ", *w)
	}

	if readLink, errReadLink := os.Readlink("example"); errReadLink != nil {
		fmt.Println("error:", errReadLink)
	} else {
		fmt.Println(readLink)
	}

}
