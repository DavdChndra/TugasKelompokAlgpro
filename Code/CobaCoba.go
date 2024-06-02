package main

import (
    "fmt"
)

type Subject struct {
    Name  string
    Score int
}

type Student struct {
    Name     string
    Subjects [3]Subject
}

func main() {
    var students [2]Student

    for i := range students {
        // Input untuk nama pelajar
        fmt.Printf("Masukkan nama pelajar %d: ", i+1)
        fmt.Scanln(&students[i].Name)

        for j := range students[i].Subjects {
            // Input untuk nama subjek
            fmt.Printf("  Masukkan nama subjek %d untuk pelajar %s: ", j+1, students[i].Name)
            fmt.Scanln(&students[i].Subjects[j].Name)

            // Input untuk nilai subjek
            fmt.Printf("  Masukkan nilai untuk subjek %s: ", students[i].Subjects[j].Name)
            fmt.Scanln(&students[i].Subjects[j].Score)
        }
    }

    // Menampilkan data yang telah dimasukkan
    for i, student := range students {
        fmt.Printf("Student %d: %s\n", i+1, student.Name)
        for j, subject := range student.Subjects {
            fmt.Printf("  Subject %d: %s, Score: %d\n", j+1, subject.Name, subject.Score)
        }
    }
}
