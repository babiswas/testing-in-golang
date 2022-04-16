package myapp
import "testing"

func TestCalulate(t *testing.T){
   if calculate(2)!=4{
      t.Error("Expected 2+2 is equal to 4")
   }

   if sum(3,4)!=7{
      t.Error("Expected 3+4 is equal to 7")
   }
}