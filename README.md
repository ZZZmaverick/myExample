# myExample
   
主要是example/metrics/metrics.go和example/metrics_version/main.go，二者共同构成一个简单的exporter。

example/metrics/metrics.go中：
NewAdmissionLatency() *RequestLatency用来初始化计时器，(t *RequestLatency) Observe()用于测量执行时间，RequestIncrease()用来增加请求计数。
 {
	requestCount.WithLabelValues().Add(1)
}


example/metrics_version/main.go中：
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
//index函数，main中的第一个调用
func index(w http.ResponseWriter, r *http.Request) {
	timer:=metrics.NewAdmissionLatency()	// 记录从访问路由到响应的延迟
	metrics.RequestIncrease()	//请求次数加一
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


