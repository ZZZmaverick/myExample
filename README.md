# myExample
 ├─ example
 │  ├── Dockerfile                   制作镜像所使用
 │  ├── README
 │  ├── deploy                       部署资源对象时使用的配置文件
 │  │   ├── deployment.yaml          云服务的Deployment配置文件
 │  │   ├── metrics_service.yaml     
 │  │   ├── prometheus.config.yml    prometheus抓取目标配置文件
 │  │   ├── prometheus.deploy.yml    prometheus部署所使用Deployment
 │  │   ├── prometheus.rbac.yml      prometheus权限配置文件
 │  │   └── service.yaml
 │  ├── go.mod                       依赖管理
 │  ├── go.sum                       依赖管理
 │  ├── metrics                      Exporter
 │  │   └── metrics.go
 │  ├── metrics_version              
 │  │   └── main.go
 │  └── without_metrics             
 │      └── main.go
 └── nginx   
    └── nginx-deployment.yaml
    └── nginx-svc.yaml
    
    
    
主要是example/metrics/metrics.go和example/metrics_version/main.go，二者共同构成一个简单的exporter

example/metrics/metrics.go中：
//初始化计时器
func NewAdmissionLatency() *RequestLatency {
	return &RequestLatency{
		histo: requestLatency,
		start: time.Now(),
	}
}
//测量执行时间
func (t *RequestLatency) Observe() {
	(*t.histo).WithLabelValues().Observe(time.Now().Sub(t.start).Seconds())
}
//增加请求计数
func RequestIncrease() {
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

