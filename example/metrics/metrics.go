package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

//定义
var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:      "request_total",
			Help:      "Number of request processed by this service.",
		}, []string{},
	)

	requestLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:      "request_latency_seconds",
			Help:      "Time spent in this service.",
			Buckets:   []float64{0.01, 0.02, 0.05, 0.1, 0.2, 0.5, 1.0, 2.0, 5.0, 10.0, 20.0, 30.0, 60.0, 120.0, 300.0},
		}, []string{},
	)
        //cpu温度
    	cpuTemp = prometheus.NewGauge(
           	prometheus.GaugeOpts{
            	Name: "cpu_Temperature", 
            	Help: "Cpu's temperature.", 
    	})
)

type RequestLatency struct {
	histo *prometheus.HistogramVec
	start time.Time
}

func Register() {
	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(requestLatency)
}

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

//获取cpu温度
func  MeasureTemperature() {
    	cpuTemp.Set(float64(rand.Int31n(30)+45))
}
