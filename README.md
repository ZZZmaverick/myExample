# myExample
   
主要是example/metrics/metrics.go和example/metrics_version/main.go，二者共同构成一个简单的exporter。

example/metrics/metrics.go中：
NewAdmissionLatency() *RequestLatency用来初始化计时器，(t *RequestLatency) Observe()用于测量执行时间，RequestIncrease()用来增加请求计数。另外加入了表示cpu温度的变量cpuTemp，并设置了getCpuTemperature()来获取cpu温度。


example/metrics_version/main.go中：
main()先设置两个处理函数，前者为index()函数（本文件中），后者为Handler函数（来自"github.com/prometheus/client_golang/prometheus/promhttp" ，将所有metrics中的指标都响应给请求），调用metrics.Register()做指标注册，然后判断条件，设置监听端口，满足则绑定在5565端口。
index()函数是main中的第一个调用，它先调用metrics.NewAdmissionLatency()初始化计时器，记录从访问路由到响应的延迟，进行请求次数累加，然后调用metrics中的getCpuTemperature()函数来获取cpu温度，再进行Num值的获取和判断，经Fibonacci()函数调用延时后，最后响应请求。

详细代码注释等见文件内...

