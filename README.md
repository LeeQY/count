# count
Simple counts storage for postgresql.

####Install:  
`go get github.com/LeeQY/count`

####Get count
```Go
//pass in the database, table, field and id
n := count.Get(db, "table", "value", 1)
```

####Change count
```Go
//pass in with an increment
count.Add(db, "table", "value", 1, -3)
```