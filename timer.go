package zwheeltimer
import(
	"time"
)

type Timer struct {
	Id int64 					//用来管理不同的环形定时器
	CurrentIndex int 			//当前指示的位置
	WheelSize int 				//一共有多少格
}

//创建一个环形定时器
func NewTimer(tick time.Duration, wheelSize int64){

}