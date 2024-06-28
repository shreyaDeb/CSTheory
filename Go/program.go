//Shreya Deb

func readinput()string{

        file, err := os.Open(os.Args[1])

        if err != nil {
                log.Fatalf("Could not open file: %s", err)
        }

	scanner := bufio.NewScanner(file)
        scanner.Split(bufio.ScanLines)

        var inputlines []string

        for scanner.Scan() {
                inputlines = append(inputlines, scanner.Text())
        }

	str1 := strings.Join(inputlines, " ")

        file.Close()

        return(str1)
}


func readinputdigits (str1 string)string{

        var result []int

        re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
        str2 := re.FindAllString(str1, -1)

        str3 := strings.Join(str2, ",")
        s := strings.Split(str3, ",")

        for i := 0; i < len(s); i++ {
                jbb := s[i]
                jbb2, err := strconv.Atoi(jbb)
                if err != nil {
                log.Fatal(err)
        }
                result = append(result, jbb2)
        }

	fmt.Println(" ")

        var b bytes.Buffer

        for i := 0; i < len(result); i++ {
                jbb1 := result[i]
                jbb3 := string(jbb1)
                b.WriteString(jbb3)
                fmt.Printf(jbb3)
        }

	fmt.Println(" ")

        return(b.String())

}

func output(str3 string){

        f, err := os.Create(os.Args[2])

        if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

        _, err2 := f.WriteString(str3)

    if err2 != nil {
        log.Fatal(err2)
    }

     	fmt.Println(" ")

    fmt.Println("Copied to Output File")

    return

}



func main() {

	str1 := readinput()

        str4 := readinputdigits(str1)

        output(str4)

        fmt.Println("End of program")

}
