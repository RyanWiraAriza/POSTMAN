package main
import "fmt"

func main(){
var n int
fmt.Scan(&n)

var array1[]string
var array2[]string
var data string
for index:=0;index<n;index++{
fmt.Scan(&data)
array1=append(array1,data)
}
for index:=0;index<n;index++{
fmt.Scan(&data)
array2=append(array2,data)
}

//equalizing
for i:=0;i<n;i++{
if array1[i]==array2[i]{
continue;
}
fmt.Printf("Index ke %d berbeda",i+1)
fmt.Println()
}

}