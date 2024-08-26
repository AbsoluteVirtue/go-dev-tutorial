if time.Now().Hour() < 12 {
	fmt.Println("Good morning.")
} else {
	fmt.Println("Good afternoon (or evening).")
}

birthday, _ := time.Parse("Jan 2 2006", "Nov 10 2009") // time.Time
age := time.Since(birthday)                            // time.Duration
fmt.Printf("Go is %d days old\n", age/(time.Hour*24))

t := time.Now()
fmt.Println(t.In(time.UTC))
home, _ := time.LoadLocation("Australia/Sydney")
fmt.Println(t.In(home))
