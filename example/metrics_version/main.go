package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"strconv"
	"example/metrics"
)

// 设置两个处理函数，前者为index，后者为Handler函数（来自"github.com/prometheus/client_golang/prometheus/promhttp"，将所有metrics中的指标都响应给请求）
func main(){
	http.HandleFunc("/abc", index)
	http.Handle("/metrics", promhttp.Handler())
	metrics.Register()	
	err := http.ListenAndServe(":5565", nil) // 设置监听，绑定在5565端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	timer:=metrics.NewAdmissionLatency()	// 记录从访问路由到响应的延迟
	metrics.RequestIncrease()	//请求次数加一
	metrics.getCpuTemperature()	//调用metrics中的cpu温度获取函数
	num:=os.Getenv("Num")
	if num==""{
		Fibonacci(10)	//延时
		_,err:=w.Write([]byte("there is no env Num. Computation successed\n"))
		if err!=nil{
			log.Println("err:"+err.Error()+" No\n")
		}
	}else{
		numInt,_:=strconv.Atoi(num)
		Fibonacci(numInt)
		_,err:=w.Write([]byte("there is env Num. Computation successed\n"))
		if err!=nil{
			log.Println("err:"+err.Error()+" Yes\n")
		}
	}
	timer.Observe()
}

func Fibonacci(n int)int{
	if n<=2{
		return 1
	}else{
		return Fibonacci(n-1)+Fibonacci(n-2)
	}
}
