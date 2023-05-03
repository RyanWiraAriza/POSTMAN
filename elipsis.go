package main
import "fmt"

func main(){
	arr:=[]int{1,2}
	fmt.Println("Sebelum:",arr)
	swap(arr...)
	fmt.Println("Setelah:",arr)
}
func swap(nums ...int){
	for i:=0;i<len(nums)-1;i+=2{
		nums[i],nums[i+1]=nums[i+1],nums[i]
	}

}