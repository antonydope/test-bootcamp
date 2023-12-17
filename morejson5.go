package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

type Friends struct {
    ID         int    `json:"FriendsID"`
    FirstName  string `json:"Firstname"`
    SecondName string `json:"Secondname"`
    Age        int    `json:"age"`
    Hobby      string `json:"Hobby"`
}

func main() {
    // Open a MySQL database connection
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/myfriends")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Create the Friends table if it doesn't exist
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS friends (
            id INT AUTO_INCREMENT PRIMARY KEY,
            first_name VARCHAR(50) NOT NULL,
            second_name VARCHAR(50) NOT NULL,
            age INT NOT NULL,
            hobby VARCHAR(50) NOT NULL
        );
    `)
    if err != nil {
        log.Fatal(err)
    }

    // Clear existing data (optional)
    _, err = db.Exec("DELETE FROM friends;")
    if err != nil {
        log.Fatal(err)
    }

    firstFriend := Friends{12, "Antony", "Musumba", 16, "Coding"}
    secondFriend := Friends{13, "John", "Kiptoo", 34, "Driving"}
    thirdFriend := Friends{14, "Kevin", "Kipchumba", 23, "Drinking"}

    // Insert data into the Friends table if it doesn't exist
    var count int
    err = db.QueryRow("SELECT COUNT(*) FROM friends").Scan(&count)
    if err != nil {
        log.Fatal(err)
    }

    if count == 0 {
        _, err = db.Exec(`
            INSERT INTO friends (first_name, second_name, age, hobby) VALUES
            (?, ?, ?, ?), (?, ?, ?, ?), (?, ?, ?, ?);
        `, firstFriend.FirstName, firstFriend.SecondName, firstFriend.Age, firstFriend.Hobby,
            secondFriend.FirstName, secondFriend.SecondName, secondFriend.Age, secondFriend.Hobby,
            thirdFriend.FirstName, thirdFriend.SecondName, thirdFriend.Age, thirdFriend.Hobby)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println("Data inserted successfully.")
    } else {
        fmt.Println("Data already exists. Skipping insertion.")
    }

    // Query the Friends table and print the results
    rows, err := db.Query(`SELECT * FROM friends`)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var friendsInfo []Friends
    for rows.Next() {
        var friend Friends
        err := rows.Scan(&friend.ID, &friend.FirstName, &friend.SecondName, &friend.Age, &friend.Hobby)
        if err != nil {
            log.Fatal(err)
        }
        friendsInfo = append(friendsInfo, friend)
    }

    // Print the unmarshalled friends data
    fmt.Println("THE UNMARSHALLED FRIENDS DATA")
    for _, y := range friendsInfo {
        fmt.Printf("ID         : %d,\nFirstName  : %s,\nSecondName : %s,\nAge        : %d,\nHobby      : %s\n  \n", y.ID, y.FirstName, y.SecondName, y.Age, y.Hobby)
    }
}
